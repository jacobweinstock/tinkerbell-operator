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

	"github.com/go-logr/logr"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	tinkerbellv1 "github.com/jacobweinstock/tinkerbell-operator/api/v1"
	v1 "github.com/jacobweinstock/tinkerbell-operator/api/v1"
)

// HardwareReconciler reconciles a Hardware object
type HardwareReconciler struct {
	client.Client
	Log    logr.Logger
	Scheme *runtime.Scheme
}

// +kubebuilder:rbac:groups=tinkerbell.tinkerbell.org,resources=hardwares,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=tinkerbell.tinkerbell.org,resources=hardwares/status,verbs=get;update;patch

func (r *HardwareReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	ctx := context.Background()
	log := r.Log.WithValues("hardware", req.NamespacedName)
	log.V(0).Info("debugging", "msg", "doing reconcile", "req", req)
	instance := &v1.Hardware{}
	err := r.Get(ctx, req.NamespacedName, instance)
	if err != nil {
		if errors.IsNotFound(err) {
			// object not found, could have been deleted after
			// reconcile request, hence don't requeue
			return ctrl.Result{}, nil
		}

		// error reading the object, requeue the request
		return ctrl.Result{}, err
	}
	log.V(0).Info("debugging", "instance", fmt.Sprintf("%+v", instance.Spec))

	// your logic here

	return ctrl.Result{}, nil
}

func (r *HardwareReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&tinkerbellv1.Hardware{}).
		Complete(r)
}
