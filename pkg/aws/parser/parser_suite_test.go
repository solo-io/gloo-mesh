package aws_utils_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestAwsUtils(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "AWS Parser Suite")
}
