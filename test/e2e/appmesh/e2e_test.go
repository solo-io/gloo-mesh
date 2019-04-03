package appmesh_test

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/solo-io/supergloo/install/helm/supergloo/generate"
	sgtestutils "github.com/solo-io/supergloo/test/testutils"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/solo-io/supergloo/cli/test/utils"
)

const superglooNamespace = "supergloo-system"

var _ = Describe("E2e", func() {
	It("registers and tests appmesh", func() {
		// install discovery via cli
		// start discovery
		var superglooErr error
		projectRoot := filepath.Join(os.Getenv("GOPATH"), "src", os.Getenv("PROJECT_ROOT"))
		err := generate.Run("dev", "Always", projectRoot)
		if err == nil {
			superglooErr = utils.Supergloo(fmt.Sprintf("init --release latest --values %s", filepath.Join(projectRoot, generate.ValuesOutput)))
		} else {
			superglooErr = utils.Supergloo("init --release latest")
		}
		Expect(superglooErr).NotTo(HaveOccurred())

		// TODO (ilackarms): add a flag to switch between starting supergloo locally and deploying via cli
		sgtestutils.DeleteSuperglooPods(kube, superglooNamespace)
		appmeshName := "appmesh"

		testRegisterAppmesh(appmeshName)

		testUnregisterAppmesh(appmeshName)
	})
})

/*
   tests
*/
func testRegisterAppmesh(meshName string) {

}

func testUnregisterAppmesh(meshName string) {

}
