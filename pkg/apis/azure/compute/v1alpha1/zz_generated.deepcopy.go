// +build !ignore_autogenerated

/*
Copyright 2018 The Crossplane Authors.

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
// Code generated by main. DO NOT EDIT.

package v1alpha1

import (
	v1 "k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AKSCluster) DeepCopyInto(out *AKSCluster) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AKSCluster.
func (in *AKSCluster) DeepCopy() *AKSCluster {
	if in == nil {
		return nil
	}
	out := new(AKSCluster)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AKSCluster) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AKSClusterList) DeepCopyInto(out *AKSClusterList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]AKSCluster, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AKSClusterList.
func (in *AKSClusterList) DeepCopy() *AKSClusterList {
	if in == nil {
		return nil
	}
	out := new(AKSClusterList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AKSClusterList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AKSClusterSpec) DeepCopyInto(out *AKSClusterSpec) {
	*out = *in
	if in.NodeCount != nil {
		in, out := &in.NodeCount, &out.NodeCount
		*out = new(int)
		**out = **in
	}
	if in.ClaimRef != nil {
		in, out := &in.ClaimRef, &out.ClaimRef
		*out = new(v1.ObjectReference)
		**out = **in
	}
	if in.ClassRef != nil {
		in, out := &in.ClassRef, &out.ClassRef
		*out = new(v1.ObjectReference)
		**out = **in
	}
	if in.ConnectionSecretRef != nil {
		in, out := &in.ConnectionSecretRef, &out.ConnectionSecretRef
		*out = new(v1.LocalObjectReference)
		**out = **in
	}
	out.ProviderRef = in.ProviderRef
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AKSClusterSpec.
func (in *AKSClusterSpec) DeepCopy() *AKSClusterSpec {
	if in == nil {
		return nil
	}
	out := new(AKSClusterSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AKSClusterStatus) DeepCopyInto(out *AKSClusterStatus) {
	*out = *in
	in.DeprecatedConditionedStatus.DeepCopyInto(&out.DeprecatedConditionedStatus)
	out.BindingStatusPhase = in.BindingStatusPhase
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AKSClusterStatus.
func (in *AKSClusterStatus) DeepCopy() *AKSClusterStatus {
	if in == nil {
		return nil
	}
	out := new(AKSClusterStatus)
	in.DeepCopyInto(out)
	return out
}
