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

	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/rest"
	kuberecorder "k8s.io/client-go/tools/record"
	"sigs.k8s.io/cli-utils/pkg/kstatus/polling"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	runtimeClient "github.com/fluxcd/pkg/runtime/client"
	helper "github.com/fluxcd/pkg/runtime/controller"

	"github.com/fluxcd/helm-controller/internal/controller"
)

// HelmReleaseReconcilerFactory is a factory for a HelmReleaseReconciler.
//
// It allows you to theoretically create a HelmReleaseReconciler without having
// to deploy the helm-controller by registering it within your own
// controller-runtime manager.
//
// However, usage of this is NOT RECOMMENDED, as it is NOT OFFICIALLY SUPPORTED
// due to the fact that it is not a public API and may cause issues for users
// running the helm-controller in their cluster. In more detail, it can cause
// issues with:
//
//   - The life-cycle of the HelmRelease Custom Resource Definition, as the user
//     of your project may have a different version of the HelmRelease CRD than
//     the one you are using.
//   - The reconciliation of HelmRelease resources, as your project may attempt
//     to reconcile a HelmRelease resource which is actually supposed to be
//     reconciled by the helm-controller running in the user's cluster.
//   - Any other hidden surprises that may be lurking when using this factory
//     in combination with the helm-controller.
//
// In addition, this factory is not guaranteed to be stable, and may change
// (or break) without notice.
//
// The recommended way to use the helm-controller is to deploy it in the
// cluster, while creating HelmRelease resources in the cluster using the
// Kubernetes API (in combination with owner references).
//
// If you still want to use this factory, may God be with you (and your users),
// and please ensure to allow the user to disable this functionality within
// your project.
type HelmReleaseReconcilerFactory struct {
	client.Client
	helper.Metrics

	Config                *rest.Config
	Scheme                *runtime.Scheme
	EventRecorder         kuberecorder.EventRecorder
	DefaultServiceAccount string
	NoCrossNamespaceRef   bool
	ClientOpts            runtimeClient.Options
	KubeConfigOpts        runtimeClient.KubeConfigOptions
	StatusPoller          *polling.StatusPoller
	PollingOpts           polling.Options
	ControllerName        string
}

// HelmReleaseReconcilerOptions is a shim for the internal
// controller.HelmReleaseReconcilerOptions.
type HelmReleaseReconcilerOptions controller.HelmReleaseReconcilerOptions

// SetupWithManager sets up the internal HelmReleaseReconciler with a
// manager of choice. Before using this, please read the documentation of
// HelmReleaseReconcilerFactory.
func (f *HelmReleaseReconcilerFactory) SetupWithManager(ctx context.Context, mgr ctrl.Manager, opts HelmReleaseReconcilerOptions) error {
	r := &controller.HelmReleaseReconciler{
		Client:                f.Client,
		Metrics:               f.Metrics,
		Config:                f.Config,
		Scheme:                f.Scheme,
		EventRecorder:         f.EventRecorder,
		DefaultServiceAccount: f.DefaultServiceAccount,
		NoCrossNamespaceRef:   f.NoCrossNamespaceRef,
		ClientOpts:            f.ClientOpts,
		KubeConfigOpts:        f.KubeConfigOpts,
		StatusPoller:          f.StatusPoller,
		PollingOpts:           f.PollingOpts,
		ControllerName:        f.ControllerName,
	}
	return r.SetupWithManager(ctx, mgr, controller.HelmReleaseReconcilerOptions(opts))
}
