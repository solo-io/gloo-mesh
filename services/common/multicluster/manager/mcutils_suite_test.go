package mc_manager_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestMcutils(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Multi Cluster Manager Suite")
}
