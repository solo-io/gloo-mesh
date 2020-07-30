package io

import (
	"github.com/solo-io/service-mesh-hub/codegen/constants"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

var (
	CertificateIssuerInputTypes = Snapshot{
		schema.GroupVersion{
			Group:   "certificates." + constants.ServiceMeshHubApiGroupSuffix,
			Version: "v1alpha2",
		}: {
			"CertificateRequest",
		},
	}

	CertificateIssuerOutputTypes = Snapshot{
		corev1.SchemeGroupVersion: {
			"Secret",
		},
	}

	CertificateAgentInputTypes = Snapshot{
		schema.GroupVersion{
			Group:   "certificates." + constants.ServiceMeshHubApiGroupSuffix,
			Version: "v1alpha2",
		}: {
			"IssuedCertificate",
		},
	}

	CertificateAgentOutputTypes = Snapshot{
		schema.GroupVersion{
			Group:   "certificates." + constants.ServiceMeshHubApiGroupSuffix,
			Version: "v1alpha2",
		}: {
			"CertificateRequest",
		},
	}
)
