package constants

import (
	"github.com/solo-io/supergloo/pkg/install/istio"
)

var SupportedIstioVersions = []string{
	istio.IstioVersion103,
	istio.IstioVersion105,
	istio.IstioVersion106,
}
