package settings

import (
	"context"

	smh_settings_types "github.com/solo-io/service-mesh-hub/pkg/api/core.smh.solo.io/v1alpha1/types"
)

//go:generate mockgen -source ./interfaces.go -destination ./mocks/mock_interfaces.go

// Convenience wrapper around fetching the global Settings object.
type SettingsHelperClient interface {
	// Fetch Settings for the specified AWS account ID, returns (nil, nil) if no settings found.
	GetAWSSettingsForAccount(ctx context.Context, accountId string) (*smh_settings_types.SettingsSpec_AwsAccount, error)
}
