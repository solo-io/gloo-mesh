package docsgen

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	plugin_go "github.com/golang/protobuf/protoc-gen-go/plugin"
	"github.com/iancoleman/strcase"
	gendoc "github.com/pseudomuto/protoc-gen-doc"
	"github.com/pseudomuto/protokit"
	"github.com/solo-io/gloo-mesh/codegen/constants"
	"github.com/solo-io/solo-kit/pkg/code-generator/collector"
)

var (
	links map[string]string

	protoDocTemplate = filepath.Join(moduleRoot, "docs", "docsgen", "proto_docs_template.tmpl")

	protoIndexTemplate = filepath.Join(moduleRoot, "docs", "docsgen", "proto_index_template.tmpl")
)

func generateApiDocs(root string, opts ProtoOptions) error {
	// flush directory for idempotence
	apiDocsDir := filepath.Join(root, opts.OutputDir)
	os.RemoveAll(apiDocsDir)
	os.MkdirAll(apiDocsDir, 0755)

	if opts.ProtoRoot == "" {
		opts.ProtoRoot = filepath.Join(moduleRoot, "vendor_any")
	}
	return generateProtoDocs(opts.ProtoRoot, protoDocTemplate, apiDocsDir)
}

func buildCompleteFilename(destDir string, file *gendoc.File) string {
	// qualify file name with proto package to avoid collisions
	// locate all generated files in same directory for easier linking
	filename := filepath.Join(destDir, strings.ReplaceAll(file.Name, "/", "."))
	return strings.TrimSuffix(filename, ".proto") + ".md"
}

func generateProtoDocs(protoDir, templateFile, destDir string) error {
	tmpDir, err := ioutil.TempDir("", "proto-docs")
	if err != nil {
		return err
	}

	defer os.RemoveAll(tmpDir)

	docsTemplate, err := collectDescriptors(protoDir, tmpDir)
	if err != nil {
		return err
	}

	links = collectLinks(destDir, docsTemplate)

	templateContents, err := ioutil.ReadFile(templateFile)

	tmpl, err := template.New(templateFile).Funcs(templateFuncs(links)).Parse(string(templateContents))
	if err != nil {
		return err
	}

	// generate API docs
	for _, file := range docsTemplate.Files {
		filename := buildCompleteFilename(destDir, file)
		destFile, err := os.Create(filename)
		if err != nil {
			return err
		}
		defer destFile.Close()
		removeDescriptions(file)
		if err := tmpl.Execute(destFile, file); err != nil {
			return err
		}
	}

	// generate index page
	return generateProtoDocsIndex(docsTemplate, links, destDir)
}

// Remove description from imported protos to avoid incorrectly rendering non-solo Hugo themes
func removeDescriptions(file *gendoc.File) {
	if strings.Contains(file.Package, constants.GlooMeshApiGroupSuffix) {
		return
	}
	for _, msg := range file.Messages {
		msg.Description = ""
	}
	for _, enum := range file.Enums {
		enum.Description = ""
	}
	for _, service := range file.Services {
		service.Description = ""
	}
	for _, extension := range file.Extensions {
		extension.Description = ""
	}
}

// generate proto docs index page
func generateProtoDocsIndex(descriptors *gendoc.Template, links map[string]string, destDir string) error {

	templateContents, err := ioutil.ReadFile(protoIndexTemplate)

	tmpl, err := template.New(protoIndexTemplate).Funcs(templateFuncs(links)).Parse(string(templateContents))
	if err != nil {
		return err
	}

	destFile, err := os.Create(filepath.Join(destDir, "_index.md"))
	if err != nil {
		return err
	}

	indexData := collectIndexData(descriptors, links)

	if err := tmpl.Execute(destFile, indexData); err != nil {
		return err
	}
	return nil
}

type IndexEntry struct {
	Name string
	Link string
}

type PackageIndex struct {
	// maps CRD name to CrdIndex
	Crds     map[string]*IndexEntry
	Services []*IndexEntry
}

func collectIndexData(template *gendoc.Template, links map[string]string) map[string]*PackageIndex {
	indexData := map[string]*PackageIndex{}
	for _, file := range template.Files {
		if !strings.Contains(file.Package, constants.GlooMeshApiGroupSuffix) {
			continue
		}
		packageIndex, ok := indexData[file.Package]
		if !ok {
			indexData[file.Package] = &PackageIndex{
				Crds:     map[string]*IndexEntry{},
				Services: []*IndexEntry{},
			}
			packageIndex = indexData[file.Package]
		}

		for _, msg := range file.Messages {
			if strings.HasSuffix(msg.FullName, "Spec") {
				crdName := strings.Replace(msg.FullName, "Spec", "", 1)
				crdIndex, ok := packageIndex.Crds[crdName]
				if !ok {
					packageIndex.Crds[crdName] = &IndexEntry{}
					crdIndex = packageIndex.Crds[crdName]
				}
				crdIndex.Name = strings.Split(crdName, ".")[len(strings.Split(crdName, "."))-1]
				crdIndex.Link = links[msg.FullName]
			}
		}
		for _, service := range file.Services {
			packageIndex.Services = append(packageIndex.Services, &IndexEntry{
				Name: service.FullName,
				Link: links[service.FullName],
			})
		}
	}
	return indexData
}

// iterate gendoc template files and construct mapping from proto message name to relative link
func collectLinks(destDir string, template *gendoc.Template) map[string]string {
	links := map[string]string{}
	for _, file := range template.Files {
		filename := filepath.Base(buildCompleteFilename(destDir, file))
		// links consists of "<filename.md>#<message/enumName>"
		for _, msg := range file.Messages {
			if a, ok := links[msg.FullName]; ok && a != msg.FullName {
				log.Printf("warning: found multiple definitions of proto msg %s: %+v", msg.FullName, []string{a, filepath.Base(filename) + "#" + msg.FullName})
			}
			links[msg.FullName] = buildLink(filename, msg.FullName)
		}
		for _, enum := range file.Enums {
			if a, ok := links[enum.FullName]; ok && a != enum.FullName {
				log.Printf("warning: found multiple definitions of proto enum %s: %+v", enum.FullName, []string{a, filepath.Base(filename) + "#" + enum.FullName})
			}
			links[enum.FullName] = buildLink(filename, enum.FullName)
		}
		for _, service := range file.Services {
			if a, ok := links[service.FullName]; ok && a != service.FullName {
				log.Printf("warning: found multiple definitions of proto service %s: %+v", service.FullName, []string{a, filepath.Base(filename) + "#" + service.FullName})
			}
			links[service.FullName] = buildLink(filename, service.FullName)
		}
		for _, extension := range file.Extensions {
			if a, ok := links[extension.FullName]; ok && a != extension.FullName {
				log.Printf("warning: found multiple definitions of proto extension %s: %+v", extension.FullName, []string{a, filepath.Base(filename) + "#" + extension.FullName})
			}
			links[extension.FullName] = buildLink(filename, extension.FullName)
		}
	}
	return links
}

func buildLink(filename, protoTypeFullName string) string {
	return fmt.Sprintf("{{< versioned_link_path fromRoot=\"/reference/api/%s\" >}}", strings.ReplaceAll(filepath.Base(filename), ".md", "")+"#"+protoTypeFullName)
}

func collectDescriptors(protoDir, outDir string, customImports ...string) (*gendoc.Template, error) {
	descriptors, err := collector.NewCollector(
		customImports,
		[]string{protoDir},
		nil,
		[]string{},
		outDir,
		func(file string) bool {
			return true
		}).CollectDescriptorsFromRoot(protoDir, nil)
	if err != nil {
		return nil, err
	}

	req := &plugin_go.CodeGeneratorRequest{}

	for _, file := range descriptors {
		req.FileToGenerate = append(req.FileToGenerate, file.GetName())
		req.ProtoFile = append(req.ProtoFile, file.FileDescriptorProto)
	}

	return gendoc.NewTemplate(protokit.ParseCodeGenRequest(req)), nil
}

func templateFuncs(links map[string]string) template.FuncMap {
	return template.FuncMap{
		"lowerCamel": strcase.ToLowerCamel,
		"replaceNewLine": func(str string) string {
			str = strings.ReplaceAll(str, "\n\n", "<br>")
			return strings.ReplaceAll(str, "\n", " ")
		},
		"cleanFileName": func(str string) string {
			return filepath.Base(str)
		},
		"link_to_type": func(v interface{}) string {
			switch fieldType := v.(type) {
			case *gendoc.MessageField:
				link, ok := links[fieldType.FullType]
				if ok {
					return link
				} else if strings.Contains(fieldType.FullType, ".") {
					log.Printf("link not found for %v", fieldType)
				}
			}
			return ""
		},
		"include_in_menu": func(pkg interface{}) bool {
			if strings.Contains(pkg.(string), constants.GlooMeshApiGroupSuffix) {
				return true
			}
			return false
		},
	}
}
