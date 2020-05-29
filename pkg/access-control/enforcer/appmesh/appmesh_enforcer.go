package appmesh

import (
	"context"

	access_control_enforcer "github.com/solo-io/service-mesh-hub/pkg/access-control/enforcer"
	zephyr_discovery "github.com/solo-io/service-mesh-hub/pkg/api/discovery.zephyr.solo.io/v1alpha1"
	"github.com/solo-io/service-mesh-hub/pkg/aws/appmesh/translation"
)

const (
	EnforcerId = "appmesh_enforcer"
)

type appmeshEnforcer struct {
	appmeshTranslationReconciler translation.AppmeshTranslationReconciler
}

type AppmeshEnforcer access_control_enforcer.AccessPolicyMeshEnforcer

func NewAppmeshEnforcer(
	appmeshTranslationReconciler translation.AppmeshTranslationReconciler,
) AppmeshEnforcer {
	return &appmeshEnforcer{appmeshTranslationReconciler: appmeshTranslationReconciler}
}

func (a *appmeshEnforcer) Name() string {
	return EnforcerId
}

func (a *appmeshEnforcer) StartEnforcing(ctx context.Context, mesh *zephyr_discovery.Mesh) error {
	return a.appmeshTranslationReconciler.Reconcile(ctx, mesh)
}

func (a *appmeshEnforcer) StopEnforcing(ctx context.Context, mesh *zephyr_discovery.Mesh) error {
	return nil
}
