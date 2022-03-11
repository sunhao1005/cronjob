/*
Copyright 2022.

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

package v1

import (
	"errors"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
)

// log is for logging in this package.
var cronjoblog = logf.Log.WithName("cronjob-resource")

func (r *CronJob) SetupWebhookWithManager(mgr ctrl.Manager) error {
	return ctrl.NewWebhookManagedBy(mgr).
		For(r).
		Complete()
}

// TODO(user): EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!

//+kubebuilder:webhook:path=/mutate-demotest-daocloud-io-v1-cronjob,mutating=true,failurePolicy=fail,sideEffects=None,groups=demotest.daocloud.io,resources=cronjobs,verbs=create;update,versions=v1,name=mcronjob.kb.io,admissionReviewVersions=v1

var _ webhook.Defaulter = &CronJob{}

// Default implements webhook.Defaulter so a webhook will be registered for the type
//设置默认配置
func (r *CronJob) Default() {
	cronjoblog.Info("default", "name", r.Name)
	// TODO(user): fill in your defaulting logic.
	if r.Spec.Replicas < 1 {
		r.Spec.Replicas = 1
	}
	cronjoblog.Info("default", "Replicas", r.Spec.Replicas)
}

// TODO(user): change verbs to "verbs=create;update;delete" if you want to enable deletion validation.
//+kubebuilder:webhook:path=/validate-demotest-daocloud-io-v1-cronjob,mutating=false,failurePolicy=fail,sideEffects=None,groups=demotest.daocloud.io,resources=cronjobs,verbs=create;update;delete,versions=v1,name=vcronjob.kb.io,admissionReviewVersions=v1

var _ webhook.Validator = &CronJob{}

// ValidateCreate implements webhook.Validator so a webhook will be registered for the type
//创建配置
func (r *CronJob) ValidateCreate() error {
	cronjoblog.Info("validate create", "name", r.Name)
	if r.Spec.Replicas > 3 {
		cronjoblog.Error(errors.New("replicas is too large"), "replicas", r.Spec.Replicas)
		return errors.New("replicas is too large")
	}
	// TODO(user): fill in your validation logic upon object creation.
	return nil
}

// ValidateUpdate implements webhook.Validator so a webhook will be registered for the type
//更新配置
func (r *CronJob) ValidateUpdate(old runtime.Object) error {
	cronjoblog.Info("validate update", "name", r.Name)

	// TODO(user): fill in your validation logic upon object update.
	return nil
}

// ValidateDelete implements webhook.Validator so a webhook will be registered for the type
//删除配置
func (r *CronJob) ValidateDelete() error {
	cronjoblog.Info("validate delete", "name", r.Name)

	// TODO(user): fill in your validation logic upon object deletion.
	return nil
}
