// +build !ignore_autogenerated

/*
Copyright The Kubernetes Authors.

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

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1alpha2

import (
	v1alpha1 "github.com/vmware/global-load-balancing-services-for-kubernetes/internal/apis/amko/v1alpha1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppSelector) DeepCopyInto(out *AppSelector) {
	*out = *in
	if in.Label != nil {
		in, out := &in.Label, &out.Label
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppSelector.
func (in *AppSelector) DeepCopy() *AppSelector {
	if in == nil {
		return nil
	}
	out := new(AppSelector)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterProperty) DeepCopyInto(out *ClusterProperty) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterProperty.
func (in *ClusterProperty) DeepCopy() *ClusterProperty {
	if in == nil {
		return nil
	}
	out := new(ClusterProperty)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GDPSpec) DeepCopyInto(out *GDPSpec) {
	*out = *in
	in.MatchRules.DeepCopyInto(&out.MatchRules)
	if in.MatchClusters != nil {
		in, out := &in.MatchClusters, &out.MatchClusters
		*out = make([]ClusterProperty, len(*in))
		copy(*out, *in)
	}
	if in.TrafficSplit != nil {
		in, out := &in.TrafficSplit, &out.TrafficSplit
		*out = make([]TrafficSplitElem, len(*in))
		copy(*out, *in)
	}
	if in.HealthMonitorRefs != nil {
		in, out := &in.HealthMonitorRefs, &out.HealthMonitorRefs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.TTL != nil {
		in, out := &in.TTL, &out.TTL
		*out = new(int)
		**out = **in
	}
	if in.SitePersistenceRef != nil {
		in, out := &in.SitePersistenceRef, &out.SitePersistenceRef
		*out = new(string)
		**out = **in
	}
	if in.PoolAlgorithmSettings != nil {
		in, out := &in.PoolAlgorithmSettings, &out.PoolAlgorithmSettings
		*out = new(v1alpha1.PoolAlgorithmSettings)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GDPSpec.
func (in *GDPSpec) DeepCopy() *GDPSpec {
	if in == nil {
		return nil
	}
	out := new(GDPSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GDPStatus) DeepCopyInto(out *GDPStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GDPStatus.
func (in *GDPStatus) DeepCopy() *GDPStatus {
	if in == nil {
		return nil
	}
	out := new(GDPStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GlobalDeploymentPolicy) DeepCopyInto(out *GlobalDeploymentPolicy) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GlobalDeploymentPolicy.
func (in *GlobalDeploymentPolicy) DeepCopy() *GlobalDeploymentPolicy {
	if in == nil {
		return nil
	}
	out := new(GlobalDeploymentPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *GlobalDeploymentPolicy) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GlobalDeploymentPolicyList) DeepCopyInto(out *GlobalDeploymentPolicyList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]GlobalDeploymentPolicy, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GlobalDeploymentPolicyList.
func (in *GlobalDeploymentPolicyList) DeepCopy() *GlobalDeploymentPolicyList {
	if in == nil {
		return nil
	}
	out := new(GlobalDeploymentPolicyList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *GlobalDeploymentPolicyList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MatchRules) DeepCopyInto(out *MatchRules) {
	*out = *in
	in.AppSelector.DeepCopyInto(&out.AppSelector)
	in.NamespaceSelector.DeepCopyInto(&out.NamespaceSelector)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MatchRules.
func (in *MatchRules) DeepCopy() *MatchRules {
	if in == nil {
		return nil
	}
	out := new(MatchRules)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NamespaceSelector) DeepCopyInto(out *NamespaceSelector) {
	*out = *in
	if in.Label != nil {
		in, out := &in.Label, &out.Label
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NamespaceSelector.
func (in *NamespaceSelector) DeepCopy() *NamespaceSelector {
	if in == nil {
		return nil
	}
	out := new(NamespaceSelector)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TrafficSplitElem) DeepCopyInto(out *TrafficSplitElem) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TrafficSplitElem.
func (in *TrafficSplitElem) DeepCopy() *TrafficSplitElem {
	if in == nil {
		return nil
	}
	out := new(TrafficSplitElem)
	in.DeepCopyInto(out)
	return out
}
