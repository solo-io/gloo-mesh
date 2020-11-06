package trafficpolicyutils

import discoveryv1alpha2 "github.com/solo-io/gloo-mesh/pkg/api/discovery.mesh.gloo.solo.io/v1alpha2"

func ContainsPort(
	destKubeServicePorts []*discoveryv1alpha2.TrafficTargetSpec_KubeService_KubeServicePort,
	port uint32,
) bool {
	for _, destPort := range destKubeServicePorts {
		if destPort.Port == port {
			return true
		}
	}
	return false
}
