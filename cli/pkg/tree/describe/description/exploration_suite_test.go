package description_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestExploration(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Exploration Suite")
}
