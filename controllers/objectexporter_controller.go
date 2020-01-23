/*

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

package controllers

import (
	"context"
	"fmt"

	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"

	"k8s.io/apimachinery/pkg/types"

	"sigs.k8s.io/controller-runtime/pkg/builder"

	"github.com/go-logr/logr"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	exporterv1alpha1 "github.com/jrthrawny/kube-object-exporter/api/v1alpha1"
)

// ObjectExporterReconciler reconciles a ObjectExporter object
type ObjectExporterReconciler struct {
	client.Client
	Log      logr.Logger
	exporter types.NamespacedName
	schema   runtime.Scheme
	gvk      schema.GroupVersionKind
}

// +kubebuilder:rbac:groups=exporter.thrawny.com,resources=objectexporters,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=exporter.thrawny.com,resources=objectexporters/status,verbs=get;update;patch

func (r *ObjectExporterReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	_ = context.Background()
	_ = r.Log.WithValues("objectexporter", req.NamespacedName)

	// your logic here

	return ctrl.Result{}, nil
}

// ObjectExporterReconciler reconciles a ObjectExporter object
type ObjectExportManagerReconciler struct {
	client.Client
	Log logr.Logger
	mgr ctrl.Manager
}

// +kubebuilder:rbac:groups=exporter.thrawny.com,resources=objectexporters,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=exporter.thrawny.com,resources=objectexporters/status,verbs=get;update;patch
func (r *ObjectExportManagerReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	ctx := context.Background()
	_ = r.Log.WithValues("objectexporter", req.NamespacedName)

	exporter := &exporterv1alpha1.ObjectExporter{}
	err := r.Get(ctx, req.NamespacedName, exporter)
	if client.IgnoreNotFound(err) != nil {
		return ctrl.Result{}, err
	}

	gvk := schema.GroupVersionKind{
		Group:   exporter.Spec.Identifier.Group,
		Version: exporter.Spec.Identifier.Version,
		Kind:    exporter.Spec.Identifier.Kind,
	}

	sch := r.mgr.GetScheme()

	if sch.Recognizes(gvk) {
		return ctrl.Result{}, fmt.Errorf("%v is not recognized as a valid type", exporter.Spec.Identifier.GroupVersionKind)
	}

	rec := ObjectExportManagerReconciler{
		Client:   r.mgr.GetClient(),
		Log:      ctrl.Log.WithName("controllers").WithName("ObjectExporter"),
		exporter: req.NamespacedName,
		gvk:      gvk,
		schema:   sch,
	}

	err = builder.ControllerManagedBy(r.mgr).For(&exporterv1alpha1.ObjectExporter{}).Complete(r)
	if err != nil {
		return ctrl.Result{}, err
	}

	return ctrl.Result{}, nil
}

func (r *ObjectExportManagerReconciler) SetupWithManager(mgr ctrl.Manager) error {
	//controller, err := builder.ControllerManagedBy(r.mgr).For(&exporterv1alpha1.ObjectExporter{}).Complete
	//builder := ctrl.NewControllerManagedBy(mgr)
	r.mgr = mgr
	return builder.ControllerManagedBy(r.mgr).For(&exporterv1alpha1.ObjectExporter{}).Complete(r)
}

// AdmissionEnforcer
// RetroactiveEnforcer

// ** EnforcementStrategy:

/*
policy.site.ddev.live/ResourceQuota
Scope: Application | Organization | Deployment
rules:
- metric: External
  resource: ddevlive_site_info
  LabelSelector
*/
