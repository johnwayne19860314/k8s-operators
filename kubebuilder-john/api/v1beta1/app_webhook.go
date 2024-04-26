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

package v1beta1

import (
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/validation/field"
	ctrl "sigs.k8s.io/controller-runtime"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"
)

// log is for logging in this package.
var applog = logf.Log.WithName("app-resource")

// SetupWebhookWithManager will setup the manager to manage the webhooks
func (r *App) SetupWebhookWithManager(mgr ctrl.Manager) error {
	return ctrl.NewWebhookManagedBy(mgr).
		For(r).
		Complete()
}

// TODO(user): EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!

//+kubebuilder:webhook:path=/mutate-john-john-tech-v1beta1-app,mutating=true,failurePolicy=fail,sideEffects=None,groups=john.john.tech,resources=apps,verbs=create;update,versions=v1beta1,name=mapp.kb.io,admissionReviewVersions=v1

var _ webhook.Defaulter = &App{}

// Default implements webhook.Defaulter so a webhook will be registered for the type
func (r *App) Default() {
	applog.Info("default", "name", r.Name)

	// TODO(user): fill in your defaulting logic.
	r.Spec.EnableIngress = !r.Spec.EnableIngress
}

// TODO(user): change verbs to "verbs=create;update;delete" if you want to enable deletion validation.
//+kubebuilder:webhook:path=/validate-john-john-tech-v1beta1-app,mutating=false,failurePolicy=fail,sideEffects=None,groups=john.john.tech,resources=apps,verbs=create;update,versions=v1beta1,name=vapp.kb.io,admissionReviewVersions=v1

var _ webhook.Validator = &App{}

// ValidateCreate implements webhook.Validator so a webhook will be registered for the type
func (r *App) ValidateCreate() (admission.Warnings, error) {
	applog.Info("validate create", "name", r.Name)

	// TODO(user): fill in your validation logic upon object creation.
	err := r.validateApp()
	return nil, err
}

// ValidateUpdate implements webhook.Validator so a webhook will be registered for the type
func (r *App) ValidateUpdate(old runtime.Object) (admission.Warnings, error) {
	applog.Info("validate update", "name", r.Name)

	// TODO(user): fill in your validation logic upon object update.
	err := r.validateApp()
	return nil, err

}

func (r *App) validateApp() error {
	if !r.Spec.EnableService && r.Spec.EnableIngress {
		return apierrors.NewInvalid(GroupVersion.WithKind("App").GroupKind(), r.Name,
			field.ErrorList{
				field.Invalid(field.NewPath("enable_service"),
					r.Spec.EnableService,
					"enable_service should be true when enable_ingress is true"),
			},
		)
	}
	return nil
}

// ValidateDelete implements webhook.Validator so a webhook will be registered for the type
func (r *App) ValidateDelete() (admission.Warnings, error) {
	applog.Info("validate delete", "name", r.Name)

	// TODO(user): fill in your validation logic upon object deletion.
	return nil, nil
}
