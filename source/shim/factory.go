/*
Copyright 2023 Hidde Beydals <yelling@hhh.computer>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package shim

import (
	"context"
	"time"

	helper "github.com/fluxcd/pkg/runtime/controller"
	helmgetter "helm.sh/helm/v3/pkg/getter"
	kuberecorder "k8s.io/client-go/tools/record"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/fluxcd/source-controller/internal/controller"
)

// HelmRepositoryReconcilerFactory is a factory for a HelmRepositoryReconciler.
//
// It allows you to theoretically create a HelmRepositoryReconciler without
// having to deploy the source-controller by registering it within your own
// controller-runtime manager.
//
// However, usage of this is NOT RECOMMENDED, as it is NOT OFFICIALLY SUPPORTED
// due to the fact that it is not a public API and may cause issues for users
// running the helm-controller in their cluster. In more detail, it can cause
// issues with:
//
//   - The life-cycle of the HelmRepository Custom Resource Definition, as the
//     user of your project may have a different version of the HelmRepository
//     CRD than the one you are using.
//   - The reconciliation of HelmRepository resources, as your project may
//     attempt to reconcile a HelmRepository resource which is actually supposed
//     to be reconciled by the source-controller running in the user's cluster.
//   - Any other hidden surprises that may be lurking when using this factory
//     in combination with the source-controller.
//
// In addition, this factory is not guaranteed to be stable, and may change
// (or break) without notice.
//
// The recommended way to use the source-controller is to deploy it in the
// cluster, while creating HelmRepository resources in the cluster using the
// Kubernetes API (in combination with owner references).
//
// If you still want to use this factory, may God be with you (and your users),
// and please ensure to allow the user to disable this functionality within
// your project.
type HelmRepositoryReconcilerFactory struct {
	Client        client.Client
	EventRecorder kuberecorder.EventRecorder
	Metrics       helper.Metrics

	Getters        helmgetter.Providers
	Storage        Storage
	ControllerName string

	Cache         Cache
	CacheRecorder CacheRecorder
	TTL           time.Duration
}

// HelmRepositoryReconcilerOptions is shim for the internal
// controller.HelmRepositoryReconcilerOptions.
type HelmRepositoryReconcilerOptions controller.HelmRepositoryReconcilerOptions

// SetupWithManager sets up the internal HelmRepositoryReconciler with a
// manager of choice. Before using this, please read the documentation of
// HelmRepositoryReconcilerFactory.
func (f *HelmRepositoryReconcilerFactory) SetupWithManager(_ context.Context, mgr ctrl.Manager, opts HelmRepositoryReconcilerOptions) error {
	r := &controller.HelmRepositoryReconciler{
		Client:         f.Client,
		EventRecorder:  f.EventRecorder,
		Metrics:        f.Metrics,
		Getters:        f.Getters,
		Storage:        f.Storage,
		ControllerName: f.ControllerName,
		Cache:          f.Cache,
		CacheRecorder:  f.CacheRecorder,
		TTL:            f.TTL,
	}
	return r.SetupWithManagerAndOptions(mgr, controller.HelmRepositoryReconcilerOptions(opts))
}

// RegistryClientGeneratorFunc is a shim for the internal
// controller.RegistryClientGeneratorFunc.
type RegistryClientGeneratorFunc controller.RegistryClientGeneratorFunc

// HelmChartReconcilerFactory is a factory for a HelmChartReconciler.
//
// It allows you to theoretically create a HelmChartReconciler without
// having to deploy the source-controller by registering it within your own
// controller-runtime manager.
//
// However, usage of this is NOT RECOMMENDED, as it is NOT OFFICIALLY SUPPORTED
// due to the fact that it is not a public API and may cause issues for users
// running the helm-controller in their cluster. In more detail, it can cause
// issues with:
//
//   - The life-cycle of the HelmChart Custom Resource Definition, as the
//     user of your project may have a different version of the HelmChart
//     CRD than the one you are using.
//   - The reconciliation of HelmChart resources, as your project may
//     attempt to reconcile a HelmChart resource which is actually supposed
//     to be reconciled by the source-controller running in the user's cluster.
//   - Any other hidden surprises that may be lurking when using this factory
//     in combination with the source-controller.
//
// In addition, this factory is not guaranteed to be stable, and may change
// (or break) without notice.
//
// The recommended way to use the source-controller is to deploy it in the
// cluster, while creating HelmChart resources in the cluster using the
// Kubernetes API (in combination with owner references).
//
// If you still want to use this factory, may God be with you (and your users),
// and please ensure to allow the user to disable this functionality within
// your project.
type HelmChartReconcilerFactory struct {
	Client        client.Client
	EventRecorder kuberecorder.EventRecorder
	Metrics       helper.Metrics

	RegistryClientGenerator RegistryClientGeneratorFunc
	Storage                 Storage
	Getters                 helmgetter.Providers
	ControllerName          string

	Cache         Cache
	CacheRecorder CacheRecorder
	TTL           time.Duration
}

// HelmChartReconcilerOptions is shim for the internal
// controller.HelmChartReconcilerOptions.
type HelmChartReconcilerOptions controller.HelmChartReconcilerOptions

// SetupWithManager sets up the internal HelmRepositoryReconciler with a
// manager of choice. Before using this, please read the documentation of
// HelmRepositoryReconcilerFactory.
func (f *HelmChartReconcilerFactory) SetupWithManager(ctx context.Context, mgr ctrl.Manager, opts HelmRepositoryReconcilerOptions) error {
	r := &controller.HelmChartReconciler{
		Client:                  f.Client,
		EventRecorder:           f.EventRecorder,
		Metrics:                 f.Metrics,
		RegistryClientGenerator: controller.RegistryClientGeneratorFunc(f.RegistryClientGenerator),
		Storage:                 f.Storage,
		Getters:                 f.Getters,
		ControllerName:          f.ControllerName,
		Cache:                   f.Cache,
		CacheRecorder:           f.CacheRecorder,
		TTL:                     f.TTL,
	}
	return r.SetupWithManagerAndOptions(ctx, mgr, controller.HelmChartReconcilerOptions(opts))
}
