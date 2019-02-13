package secret

import (
	"context"

	"github.com/solo-io/solo-kit/pkg/api/v1/resources"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources/core"

	istiov1 "github.com/solo-io/supergloo/pkg2/api/external/istio/encryption/v1"

	"github.com/solo-io/solo-kit/pkg/api/v1/clients"
	"github.com/solo-io/solo-kit/pkg/errors"
	v1 "github.com/solo-io/supergloo/pkg2/api/v1"
	"github.com/solo-io/supergloo/pkg2/kube"
)

// If you change this interface, you have to rerun mockgen
type SecretSyncer interface {
	SyncSecret(ctx context.Context, installNamespace string, encryption *v1.Encryption, secretList istiov1.IstioCacertsSecretList, preinstall bool) error
}

type KubeSecretSyncer struct {
	PodClient    kube.PodClient
	SecretClient kube.SecretClient

	IstioSecretClient istiov1.IstioCacertsSecretClient
	installNamespace  string
}

const (
	CustomRootCertificateSecretName  = "cacerts"
	DefaultRootCertificateSecretName = "istio.default"
	istioLabelKey                    = "istio"
	citadelLabelValue                = "citadel"
)

func (s *KubeSecretSyncer) SyncSecret(ctx context.Context, installNamespace string, encryption *v1.Encryption, secretList istiov1.IstioCacertsSecretList, preinstall bool) error {
	s.installNamespace = installNamespace
	if encryption == nil {
		return nil
	}
	if !encryption.TlsEnabled {
		return nil
	}
	encryptionSecret := encryption.Secret
	if encryptionSecret == nil {
		return nil
	}
	sourceSecret, err := secretList.Find(encryptionSecret.Namespace, encryptionSecret.Name)
	if err != nil {
		return errors.Wrapf(err, "Error finding secret referenced in mesh config (%s:%s)",
			encryptionSecret.Namespace, encryptionSecret.Name)
	}
	// this is where custom root certs will live once configured, if not found existingSecret will be nil
	existingSecret, _ := secretList.Find(s.installNamespace, CustomRootCertificateSecretName)
	return s.syncSecret(ctx, sourceSecret, existingSecret, preinstall)
}

func (s *KubeSecretSyncer) syncSecret(ctx context.Context, sourceSecret, existingSecret *istiov1.IstioCacertsSecret, preinstall bool) error {
	if err := validateTlsSecret(sourceSecret); err != nil {
		return errors.Wrapf(err, "invalid secret %v", sourceSecret.Metadata.Ref())
	}
	istioSecret := resources.Clone(sourceSecret).(*istiov1.IstioCacertsSecret)
	if existingSecret == nil {
		istioSecret.Metadata = core.Metadata{
			Namespace: s.installNamespace,
			Name:      CustomRootCertificateSecretName,
		}
		if _, err := s.IstioSecretClient.Write(istioSecret, clients.WriteOpts{
			Ctx: ctx,
		}); err != nil {
			return errors.Wrapf(err, "creating tool tls secret %v for istio", istioSecret.Metadata.Ref())
		}
		return nil
	}

	// move secret over to destination name/namespace
	istioSecret.SetMetadata(existingSecret.Metadata)
	istioSecret.Metadata.Annotations["created_by"] = "supergloo"
	// nothing to do
	if istioSecret.Equal(existingSecret) {
		return nil
	}
	if _, err := s.IstioSecretClient.Write(istioSecret, clients.WriteOpts{
		Ctx: ctx,
	}); err != nil {
		return errors.Wrapf(err, "updating tool tls secret %v for istio", istioSecret.Metadata.Ref())
	}

	if !preinstall {
		if err := s.restartCitadel(); err != nil {
			return errors.Wrapf(err, "Error restarting citadel")
		}
		if err := s.deleteIstioDefaultSecret(); err != nil {
			return errors.Wrapf(err, "Error removing existing default cert")
		}
	}

	return nil
}

func validateTlsSecret(secret *istiov1.IstioCacertsSecret) error {
	if secret.RootCert == "" {
		return errors.Errorf("Root cert is missing.")
	}
	if secret.CaKey == "" {
		return errors.Errorf("Private key is missing.")
	}
	return nil
}

func (s *KubeSecretSyncer) deleteIstioDefaultSecret() error {
	// Using Kube API directly cause we don't expect this secret to be tagged and it should be mostly a one-time op
	return s.SecretClient.Delete(s.installNamespace, DefaultRootCertificateSecretName)
}

func (s *KubeSecretSyncer) restartCitadel() error {
	selector := make(map[string]string)
	selector[istioLabelKey] = citadelLabelValue
	return s.PodClient.RestartPods(s.installNamespace, selector)
}
