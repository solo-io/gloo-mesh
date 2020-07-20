package fieldutils_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	v1 "github.com/solo-io/skv2/pkg/api/core.skv2.solo.io/v1"
	. "github.com/solo-io/smh/pkg/mesh-networking/translation/utils/fieldutils"
	istiov1alpha3spec "istio.io/api/networking/v1alpha3"
)

var _ = Describe("FieldOwnership", func() {
	It("registers an owner for a field, and errors when another owner with lower or equal priority attempts to register", func() {
		fieldRegistry := NewOwnershipRegistry()

		// test with the real object on we use track ownership
		istioRoute := &istiov1alpha3spec.HTTPRoute{
			CorsPolicy: &istiov1alpha3spec.CorsPolicy{},
		}

		owner1 := &v1.ObjectRef{Name: "1"}
		owner2 := &v1.ObjectRef{Name: "2"}

		corsPolicyField := &istioRoute.CorsPolicy
		err := fieldRegistry.RegisterFieldOwnership(corsPolicyField, owner1, 1)
		Expect(err).NotTo(HaveOccurred())

		err = fieldRegistry.RegisterFieldOwnership(corsPolicyField, owner2, 0)
		Expect(err).To(HaveOccurred())
		Expect(err).To(Equal(FieldConflictError{
			Field:    corsPolicyField,
			Owners:   owner1,
			Priority: 1,
		}))

	})
})
