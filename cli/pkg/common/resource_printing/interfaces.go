package resource_printing

import (
	"io"

	"k8s.io/apimachinery/pkg/runtime"
)

//go:generate mockgen -source ./interfaces.go -destination ./mocks/mock_interfaces.go

type ResourcePrinter interface {
	Print(out io.Writer, object runtime.Object, format OutputFormat) error
}
