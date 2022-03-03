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
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// CronJobSpec defines the desired state of CronJob
type CronJobSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Foo is an example field of CronJob. Edit cronjob_types.go to remove/update
	//Foo string `json:"foo,omitempty"`
	Selector SelectorInfo `json:"selector"`
}
type SelectorInfo struct {
	MatchLabels MatchLabelsInfo `json:"matchLabels"`
	Replicas    int64           `json:"replicas"`
	Template    TemplateInfo    `json:"template"`
}
type TemplateInfo struct {
	Metadata MatchLabelsInfo `json:"metadata"`
	Spec     SpecInfo        `json:"spec"`
}

type SpecInfo struct {
	Containers ContainersInfo `json:"containers"`
}
type ContainersInfo struct {
	Name            string        `json:"name"`
	Image           string        `json:"image"`
	ImagePullPolicy string        `json:"imagePullPolicy"`
	Resources       ResourcesInfo `json:"resources"`
	Ports           PortsInfo     `json:"ports"`
}

type ResourcesInfo struct {
	Limits LimitsInfo `json:"limits"`
}
type LimitsInfo struct {
	Cpu    string `json:"cpu"`
	Memory string `json:"memory"`
}

type PortsInfo struct {
	ContainerPort string `json:"containerPort"`
}

type MetadataInfo struct {
	Labels LabelsInfo `json:"labels"`
}
type LabelsInfo struct {
	App string `json:"app"`
}

type MatchLabelsInfo struct {
	App string `json:"app"`
}

// CronJobStatus defines the observed state of CronJob
type CronJobStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	Replicas int      `json:"replicas"`
	PodNames []string `json:"podNames"` //这里二次开发加入计数和pod列表
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// CronJob is the Schema for the cronjobs API
type CronJob struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   CronJobSpec   `json:"spec,omitempty"`
	Status CronJobStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// CronJobList contains a list of CronJob
type CronJobList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CronJob `json:"items"`
}

func init() {
	SchemeBuilder.Register(&CronJob{}, &CronJobList{})
}
