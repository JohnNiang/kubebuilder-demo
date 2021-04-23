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
	"encoding/base64"
	"time"

	"github.com/go-logr/logr"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	demov1 "demo/api/v1"
)

// DemoReconciler reconciles a Demo object
type DemoReconciler struct {
	client.Client
	Log    logr.Logger
	Scheme *runtime.Scheme
}

// +kubebuilder:rbac:groups=demo.devops.kubesphere,resources=demoes,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=demo.devops.kubesphere,resources=demoes/status,verbs=get;update;patch

func (r *DemoReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	ctx := context.Background()
	log := r.Log.WithValues("demo", req.NamespacedName)

	// your logic here
	var demo demov1.Demo
	// get demo data from api server
	err := r.Get(ctx, req.NamespacedName, &demo)
	if err != nil {
		log.Error(err, "Demo not found!")
		return ctrl.Result{}, nil
	}

	// get demo spec name and update demo status encrypted name
	encryptedName := base64.StdEncoding.EncodeToString([]byte(demo.Spec.Name))
	demo.Status.EncryptedName = encryptedName

	// update demo status
	err = r.Update(ctx, &demo)
	if err != nil {
		log.Error(err, "Failed to update demo")
		// returning err will let controller manager reconcile again
		return ctrl.Result{RequeueAfter: 1 * time.Minute}, err
	}

	log.Info("Successfully update demo status.")
	return ctrl.Result{}, nil
}

func (r *DemoReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&demov1.Demo{}).
		Complete(r)
}
