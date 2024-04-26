/*
Copyright 2024.

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

package controller

import (
	"context"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/types"

	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
	"sigs.k8s.io/controller-runtime/pkg/log"

	"github.com/john/api/v1beta1"
	johnv1beta1 "github.com/john/api/v1beta1"
	"github.com/john/internal/controller/utils"
	appv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	netv1 "k8s.io/api/networking/v1"
)

// AppReconciler reconciles a App object
type AppReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=john.john.tech,resources=apps,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=john.john.tech,resources=apps/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=john.john.tech,resources=apps/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the App object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.16.3/pkg/reconcile
func (r *AppReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	logger := log.FromContext(ctx)
	app := &v1beta1.App{}

	if err := r.Get(ctx, req.NamespacedName, app); err != nil {
		logger.Error(err, "not found")
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}
	deployment := utils.NewDeployment(app)
	// set owner
	if err := controllerutil.SetControllerReference(app, deployment, r.Scheme); err != nil {
		logger.Error(err, "failed to set deployment object owner")
		return ctrl.Result{}, err
	}
	// find the deployment which has the same name
	d := &appv1.Deployment{}
	if err := r.Get(ctx,req.NamespacedName,d); err != nil {
		if errors.IsNotFound(err) {
			if err := r.Create(ctx,deployment); err != nil {
				logger.Error(err, "failed to create deployment")
				return ctrl.Result{}, err
			}
		}
	}else {
		if err := r.Update(ctx,d); err != nil {
			logger.Error(err, "failed to update deployment")
			return ctrl.Result{}, err
		}
	}

	// service
	service := utils.NewService(app)
	if err := controllerutil.SetControllerReference(app,service,r.Scheme); err != nil{
		logger.Error(err,"failed to set owner to the service")
		return ctrl.Result{}, err
	}
	s := &corev1.Service{}
	if err := r.Get(ctx,types.NamespacedName{Name:app.Name,Namespace:app.Namespace}, s); err != nil {

		if app.Spec.EnableService {
			if errors.IsNotFound(err) {
				if err := r.Create(ctx,service); err != nil {
					logger.Error(err, "failed to create service")
					return ctrl.Result{}, err
				}
			}else {
				return ctrl.Result{}, err
			}
		}
	}else {
		if app.Spec.EnableService{
			logger.Info("skip update service")
		}else {
			if err := r.Delete(ctx,s); err != nil {
				logger.Error(err, "failed to delete service")
				return ctrl.Result{}, err
			}
		}
	}

	// Ingress

	ingress := utils.NewIngress(app)
	if err := controllerutil.SetControllerReference(app,ingress,r.Scheme); err != nil {
		logger.Error(err,"failed to set owner to the ingress")
	}

	i := &netv1.Ingress{}
	if err := r.Get(ctx,types.NamespacedName{Name: app.Name, Namespace: app.Namespace},i); err !=nil{
		if app.Spec.EnableIngress {
			if errors.IsNotFound(err) {
				if err := r.Create(ctx,ingress); err != nil {
					logger.Error(err, "failed to create ingress in current namespace")
					return ctrl.Result{}, err
				}
			}else {
				logger.Error(err, "failed to find ingress in current namespace")
				return ctrl.Result{}, err
			}
		}
		
	}else {
		if app.Spec.EnableIngress{
			logger.Info("skip to update the ingress")
		}else {
			if err := r.Delete(ctx,i); err != nil {
				logger.Error(err, "failed to delete ingress in current namespace")
				return ctrl.Result{}, err
			}
		}
	}


	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *AppReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&johnv1beta1.App{}).
		Owns(&appv1.Deployment{}).
		Owns(&corev1.Service{}).
		Owns(&netv1.Ingress{}).
		Complete(r)
}
