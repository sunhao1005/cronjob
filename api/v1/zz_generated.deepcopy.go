//go:build !ignore_autogenerated
// +build !ignore_autogenerated

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

// Code generated by controller-gen. DO NOT EDIT.

package v1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ContainersInfo) DeepCopyInto(out *ContainersInfo) {
	*out = *in
	out.Resources = in.Resources
	out.Ports = in.Ports
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ContainersInfo.
func (in *ContainersInfo) DeepCopy() *ContainersInfo {
	if in == nil {
		return nil
	}
	out := new(ContainersInfo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CronJob) DeepCopyInto(out *CronJob) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CronJob.
func (in *CronJob) DeepCopy() *CronJob {
	if in == nil {
		return nil
	}
	out := new(CronJob)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CronJob) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CronJobList) DeepCopyInto(out *CronJobList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]CronJob, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CronJobList.
func (in *CronJobList) DeepCopy() *CronJobList {
	if in == nil {
		return nil
	}
	out := new(CronJobList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CronJobList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CronJobSpec) DeepCopyInto(out *CronJobSpec) {
	*out = *in
	out.Selector = in.Selector
	out.Containers = in.Containers
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CronJobSpec.
func (in *CronJobSpec) DeepCopy() *CronJobSpec {
	if in == nil {
		return nil
	}
	out := new(CronJobSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CronJobStatus) DeepCopyInto(out *CronJobStatus) {
	*out = *in
	if in.PodNames != nil {
		in, out := &in.PodNames, &out.PodNames
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CronJobStatus.
func (in *CronJobStatus) DeepCopy() *CronJobStatus {
	if in == nil {
		return nil
	}
	out := new(CronJobStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LabelsInfo) DeepCopyInto(out *LabelsInfo) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LabelsInfo.
func (in *LabelsInfo) DeepCopy() *LabelsInfo {
	if in == nil {
		return nil
	}
	out := new(LabelsInfo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LimitsInfo) DeepCopyInto(out *LimitsInfo) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LimitsInfo.
func (in *LimitsInfo) DeepCopy() *LimitsInfo {
	if in == nil {
		return nil
	}
	out := new(LimitsInfo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MatchLabelsInfo) DeepCopyInto(out *MatchLabelsInfo) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MatchLabelsInfo.
func (in *MatchLabelsInfo) DeepCopy() *MatchLabelsInfo {
	if in == nil {
		return nil
	}
	out := new(MatchLabelsInfo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MetadataInfo) DeepCopyInto(out *MetadataInfo) {
	*out = *in
	out.Labels = in.Labels
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MetadataInfo.
func (in *MetadataInfo) DeepCopy() *MetadataInfo {
	if in == nil {
		return nil
	}
	out := new(MetadataInfo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PortsInfo) DeepCopyInto(out *PortsInfo) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PortsInfo.
func (in *PortsInfo) DeepCopy() *PortsInfo {
	if in == nil {
		return nil
	}
	out := new(PortsInfo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourcesInfo) DeepCopyInto(out *ResourcesInfo) {
	*out = *in
	out.Limits = in.Limits
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourcesInfo.
func (in *ResourcesInfo) DeepCopy() *ResourcesInfo {
	if in == nil {
		return nil
	}
	out := new(ResourcesInfo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SelectorInfo) DeepCopyInto(out *SelectorInfo) {
	*out = *in
	out.MatchLabels = in.MatchLabels
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SelectorInfo.
func (in *SelectorInfo) DeepCopy() *SelectorInfo {
	if in == nil {
		return nil
	}
	out := new(SelectorInfo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SpecInfo) DeepCopyInto(out *SpecInfo) {
	*out = *in
	out.Containers = in.Containers
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SpecInfo.
func (in *SpecInfo) DeepCopy() *SpecInfo {
	if in == nil {
		return nil
	}
	out := new(SpecInfo)
	in.DeepCopyInto(out)
	return out
}
