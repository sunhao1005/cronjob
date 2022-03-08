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

package controllers

import (
	"context"
	"fmt"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/intstr"
	"k8s.io/utils/pointer"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
	"sigs.k8s.io/controller-runtime/pkg/log"

	demotestv1 "cronjob/api/v1"
)

// CronJobReconciler reconciles a CronJob object
type CronJobReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=demotest.daocloud.io,resources=cronjobs,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=demotest.daocloud.io,resources=cronjobs/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=demotest.daocloud.io,resources=cronjobs/finalizers,verbs=update
// +kubebuilder:rbac:groups=core,resources=pods,verbs=get;list;watch
// +kubebuilder:rbac:groups=core,resources=pods/exec,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=core,resources=services,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=apps,resources=deployments,verbs=get;list;watch;create;update;patch;delete

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the CronJob object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.11.0/pkg/reconcile
func (r *CronJobReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	logger := log.FromContext(ctx)
	logger.Info("start reconcile")
	fmt.Printf("*********| Name %s | Namespace %s | NamespacedName %s |********", req.Name, req.Namespace, req.NamespacedName)
	jobObject := demotestv1.CronJob{}
	err := r.Get(ctx, req.NamespacedName, &jobObject)
	if err != nil {
		logger.Error(err, "Reconcile", "Get error", req.NamespacedName)
		if errors.IsNotFound(err) {
			return ctrl.Result{}, nil
		}
		return ctrl.Result{}, err
	}
	logger.Info("Reconcile", "副本", jobObject.Spec.Replicas)
	logger.Info("Reconcile", "选择器", jobObject.Spec.Selector)
	logger.Info("Reconcile", "容器", jobObject.Spec.Containers)

	//创建service
	service := &corev1.Service{}
	err = r.Get(ctx, req.NamespacedName, service)
	if err != nil {
		if errors.IsNotFound(err) {
			//构造service对象
			service = &corev1.Service{
				ObjectMeta: metav1.ObjectMeta{
					Name:      jobObject.Name + "-service",
					Namespace: jobObject.Namespace,
					Labels:    jobObject.Labels},
				Spec: corev1.ServiceSpec{
					Selector: map[string]string{"app": jobObject.Name},
					Type:     corev1.ServiceTypeNodePort,
					Ports: []corev1.ServicePort{
						{
							Name:       jobObject.Name,
							Port:       21960,
							TargetPort: intstr.FromInt(9600),
							Protocol:   corev1.ProtocolTCP,
							NodePort:   31960,
						},
					}},
			}
			logger.Info("Reconcile", "构造service成功", service)
			//建立关联
			err = controllerutil.SetControllerReference(&jobObject, service, r.Scheme)
			if err != nil {
				logger.Error(err, "Reconcile", "SetControllerReference", "error")
				return ctrl.Result{}, err
			}
			logger.Info("Reconcile", "建立连接成功", "SetControllerReference")
			// create service
			err = r.Create(ctx, service)
			if err != nil {
				logger.Error(err, "Reconcile", "Create service", "error")
				return ctrl.Result{}, err
			}
			logger.Info("Reconcile", "创建service成功", "Create service success")
		}
		logger.Error(err, "Reconcile", "Get", "service")
		return ctrl.Result{}, err
	}
	//创建deployment
	deployment := &appsv1.Deployment{}
	err = r.Get(ctx, req.NamespacedName, deployment)
	if err != nil {
		if errors.IsNotFound(err) {
			//构造deployment对象
			deployment = &appsv1.Deployment{
				ObjectMeta: metav1.ObjectMeta{
					Name:      jobObject.Name,
					Namespace: jobObject.Namespace,
					Labels:    jobObject.Labels},
				Spec: appsv1.DeploymentSpec{
					Selector: &metav1.LabelSelector{MatchLabels: map[string]string{"app": jobObject.Name}},
					Replicas: pointer.Int32(int32(jobObject.Spec.Replicas)),
					Template: corev1.PodTemplateSpec{
						ObjectMeta: metav1.ObjectMeta{Labels: map[string]string{"app": jobObject.Name}},
						Spec: corev1.PodSpec{
							Containers: []corev1.Container{{
								Name:            jobObject.Name,
								Image:           jobObject.Spec.Containers.Image,
								ImagePullPolicy: corev1.PullPolicy(jobObject.Spec.Containers.ImagePullPolicy),
								Ports:           []corev1.ContainerPort{{ContainerPort: 9600}},
								Resources: corev1.ResourceRequirements{
									//Requests: corev1.ResourceList{},
									Limits: corev1.ResourceList{
										"cpu":    resource.MustParse(jobObject.Spec.Containers.Resources.Limits.Cpu),
										"memory": resource.MustParse(jobObject.Spec.Containers.Resources.Limits.Cpu)},
								},
							}},
						},
					},
				},
			}
			logger.Info("Reconcile", "构造deployment成功", deployment)
			//建立关联
			//create deployment
			//建立关联
			err = controllerutil.SetControllerReference(&jobObject, deployment, r.Scheme)
			if err != nil {
				logger.Error(err, "Reconcile", "SetControllerReference", "deployment")
				return ctrl.Result{}, err
			}
			logger.Info("Reconcile", "建立连接成功", "SetControllerReference deployment")
			// create service
			err = r.Create(ctx, deployment)
			if err != nil {
				logger.Error(err, "Reconcile", "Create", "deployment")
				return ctrl.Result{}, err
			}
			logger.Info("Reconcile", "创建成功", "deployment")
		}
		logger.Error(err, "Reconcile", "Get", "Deployment")
		return ctrl.Result{}, nil
	}
	//更新status
	jobObject.Status.Replicas = 1
	/*var joblist v1.JobList
	err = r.List(ctx, &joblist, client.InNamespace(req.Namespace))
	if err != nil {
		logger.Error(err, "Reconcile", "List", "pod")
		return ctrl.Result{}, nil
	}
	var podsName []string
	for _, item := range joblist.Items {
		podsName = append(podsName, item.Name)
	}
	jobObject.Status.PodNames = podsName
	logger.Info("Reconcile", "pods信息", joblist)*/

	err = r.Status().Update(ctx, &jobObject)
	if err != nil {
		logger.Error(err, "Reconcile", "Status", "Update")
		return ctrl.Result{}, err
	}
	logger.Info("Reconcile", "状态更新完成", jobObject.Status)
	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *CronJobReconciler) SetupWithManager(mgr ctrl.Manager) error {
	fmt.Println("********* Scheme: ", r.Scheme, " **********")
	fmt.Println("********* Client.Status: ", r.Client.Status(), " **********")
	return ctrl.NewControllerManagedBy(mgr).
		For(&demotestv1.CronJob{}).
		Complete(r)
}
