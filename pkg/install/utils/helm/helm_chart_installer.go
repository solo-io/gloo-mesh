package helm

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"

	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/api/meta"
	"k8s.io/cli-runtime/pkg/genericclioptions/resource"

	"k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"

	"github.com/avast/retry-go"
	kubecrds "github.com/solo-io/supergloo/pkg2/kube"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/solo-io/go-utils/contextutils"
	"github.com/solo-io/go-utils/errors"
	apiexts "k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset"
	kubeerrs "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/helm/pkg/chartutil"
	"k8s.io/helm/pkg/kube"
	"k8s.io/helm/pkg/manifest"
	"k8s.io/helm/pkg/proto/hapi/chart"
	"k8s.io/helm/pkg/renderutil"
	"k8s.io/helm/pkg/tiller"
	"k8s.io/helm/pkg/timeconv"
)

// an interface allowing these methods to be mocked
type Installer interface {
	// create the resources described in the manifest
	CreateFromManifests(ctx context.Context, namespace string, manifests Manifests) error
	// delete the resources described in the manifest
	DeleteFromManifests(ctx context.Context, namespace string, manifests Manifests) error
	// perform a diff and apply patches to migrate from original to updated manifests
	UpdateFromManifests(ctx context.Context, namespace string, original, updated Manifests, recreatePods bool) error
}

type helmInstaller struct{}

func NewHelmInstaller() Installer {
	return &helmInstaller{}
}

func (*helmInstaller) CreateFromManifests(ctx context.Context, namespace string, manifests Manifests) error {
	return createFromManifests(ctx, namespace, manifests)
}

func (*helmInstaller) DeleteFromManifests(ctx context.Context, namespace string, manifests Manifests) error {
	return deleteFromManifests(ctx, namespace, manifests)
}

func (*helmInstaller) UpdateFromManifests(ctx context.Context, namespace string, original, updated Manifests, recreatePods bool) error {
	return updateFromManifests(ctx, namespace, original, updated, recreatePods)
}

var defaultKubeVersion = fmt.Sprintf("%s.%s", chartutil.DefaultKubeVersion.Major, chartutil.DefaultKubeVersion.Minor)

func RenderManifests(ctx context.Context, chartUri, values, releaseName, namespace, kubeVersion string, releaseIsInstall bool) (Manifests, error) {
	var file io.Reader
	if strings.HasPrefix(chartUri, "http://") || strings.HasPrefix(chartUri, "https://") {
		resp, err := http.Get(chartUri)
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			return nil, errors.Errorf("http GET returned status %d", resp.StatusCode)
		}

		file = resp.Body
	} else {
		path, err := filepath.Abs(chartUri)
		if err != nil {
			return nil, errors.Wrapf(err, "getting absolute path for %v", chartUri)
		}

		f, err := os.Open(path)
		if err != nil {
			return nil, errors.Wrapf(err, "opening file %v", path)
		}
		file = f
	}

	if kubeVersion == "" {
		kubeVersion = defaultKubeVersion
	}
	renderOpts := renderutil.Options{
		ReleaseOptions: chartutil.ReleaseOptions{
			Name:      releaseName,
			IsInstall: releaseIsInstall,
			IsUpgrade: !releaseIsInstall,
			Time:      timeconv.Now(),
			Namespace: namespace,
		},
		KubeVersion: kubeVersion,
	}

	// Check chart requirements to make sure all dependencies are present in /charts
	c, err := chartutil.LoadArchive(file)
	if err != nil {
		return nil, errors.Wrapf(err, "loading chart")
	}

	config := &chart.Config{Raw: values, Values: map[string]*chart.Value{}}
	renderedTemplates, err := renderutil.Render(c, config, renderOpts)
	if err != nil {
		return nil, err
	}

	for file, man := range renderedTemplates {
		if isEmptyManifest(man) {
			contextutils.LoggerFrom(ctx).Warnf("is an empty manifest, removing %v", file)
			delete(renderedTemplates, file)
		}
	}
	manifests := manifest.SplitManifests(renderedTemplates)
	return tiller.SortByKind(manifests), nil
}

func createFromManifests(ctx context.Context, namespace string, manifests Manifests) error {
	kc := kube.New(nil)

	crdManifests, nonCrdManifests := manifests.SplitByCrds()

	//crds come first
	if len(crdManifests) > 0 {
		crdInput := crdManifests.CombinedString()
		infos, err := kc.BuildUnstructured(namespace, bytes.NewBufferString(crdInput))
		if err != nil {
			return err
		}
		for _, info := range infos {
			helper := resource.NewHelper(info.Client, info.Mapping)
			if _, err := helper.Create(info.Namespace, true, info.Object, nil); err != nil {
				if !apierrors.IsAlreadyExists(err) {
					return errors.Wrapf(err, "creating %v", info.Object)
				}
			}
		}

		if err := waitForCrds(ctx, kc, crdInput); err != nil {
			return err
		}
	}

	if len(nonCrdManifests) > 0 {
		nonCrdInput := nonCrdManifests.CombinedString()
		if err := kc.Create(namespace, bytes.NewBufferString(nonCrdInput), 0, false); err != nil {
			return err
		}
	}

	return nil
}

func deleteFromManifests(ctx context.Context, namespace string, manifests Manifests) error {
	kc := kube.New(nil)

	for _, man := range manifests {
		contextutils.LoggerFrom(ctx).Infof("deleting manifest %v: %v", man.Name, man.Head)

		if err := kc.Delete(namespace, bytes.NewBufferString(man.Content)); err != nil {
			if kubeerrs.IsNotFound(err) || IsNoKindMatch(err) {
				contextutils.LoggerFrom(ctx).Warnf("not found, skipping %v", man.Name)
				continue
			}
			return err
		}
	}

	return nil
}

func updateFromManifests(ctx context.Context, namespace string, original, updated Manifests, recreatePods bool) error {
	kc := kube.New(nil)

	originalCrdManifests, originalNonCrdManifests := original.SplitByCrds()
	updatedCrdManifests, updatedNonCrdManifests := updated.SplitByCrds()

	//crds come first
	if len(originalCrdManifests) > 0 || len(updatedCrdManifests) > 0 {
		originalCrdInput := originalCrdManifests.CombinedString()
		updatedCrdInput := updatedCrdManifests.CombinedString()
		if err := kc.Update(
			namespace,
			bytes.NewBufferString(originalCrdInput),
			bytes.NewBufferString(updatedCrdInput),
			false, false, 0, false); err != nil {
			return err
		}
		if err := waitForCrds(ctx, kc, updatedCrdInput); err != nil {
			return err
		}
	}

	if len(originalNonCrdManifests) > 0 || len(updatedNonCrdManifests) > 0 {
		originalNonCrdInput := originalNonCrdManifests.CombinedString()
		updatedNonCrdInput := updatedNonCrdManifests.CombinedString()
		if err := kc.Update(
			namespace,
			bytes.NewBufferString(originalNonCrdInput),
			bytes.NewBufferString(updatedNonCrdInput),
			true, recreatePods, 0, false); err != nil {
			return err
		}
	}

	return nil
}

func waitForCrds(ctx context.Context, kc *kube.Client, manifestContent string) error {
	crds, err := kubecrds.CrdsFromManifest(manifestContent)
	if err != nil {
		return errors.Wrapf(err, "failed parsing crds from manifest")
	}

	restCfg, err := kc.ToRESTConfig()
	if err != nil {
		return errors.Wrapf(err, "getting kube rest cfg")
	}
	crdClientset, err := apiexts.NewForConfig(restCfg)
	if err != nil {
		return errors.Wrapf(err, "creating apiexts client")
	}
	for _, crd := range crds {
		crdName := crd.Name
		err = retry.Do(func() error {
			crd, err := crdClientset.ApiextensionsV1beta1().CustomResourceDefinitions().Get(crdName, v1.GetOptions{})
			if err != nil {
				return errors.Wrapf(err, "lookup crd %v", crdName)
			}

			var established bool
			for _, status := range crd.Status.Conditions {
				if status.Type == v1beta1.Established {
					established = true
					break
				}
			}

			if !established {
				return errors.Errorf("crd %v exists but not yet established by kube", crdName)
			}

			contextutils.LoggerFrom(ctx).Infof("registered crd %v", crd.ObjectMeta)
			return nil
		},
			retry.Delay(time.Millisecond*500),
			retry.DelayType(retry.FixedDelay),
		)
	}
	return nil
}

var commentRegex = regexp.MustCompile("#.*")

func isEmptyManifest(manifest string) bool {
	removeComments := commentRegex.ReplaceAllString(manifest, "")
	removeNewlines := strings.Replace(removeComments, "\n", "", -1)
	removeDashes := strings.Replace(removeNewlines, "---", "", -1)
	return removeDashes == ""
}

// consider moving to kube utils/errs package?

func IsNoKindMatch(err error) bool {
	_, ok := err.(*meta.NoKindMatchError)
	return ok
}
