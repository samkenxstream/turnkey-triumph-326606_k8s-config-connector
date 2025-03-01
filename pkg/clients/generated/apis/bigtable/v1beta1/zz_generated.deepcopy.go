//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Copyright 2020 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// *** DISCLAIMER ***
// Config Connector's go-client for CRDs is currently in ALPHA, which means
// that future versions of the go-client may include breaking changes.
// Please try it out and give us feedback!

// Code generated by main. DO NOT EDIT.

package v1beta1

import (
	v1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/k8s/v1alpha1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppprofileSingleClusterRouting) DeepCopyInto(out *AppprofileSingleClusterRouting) {
	*out = *in
	if in.AllowTransactionalWrites != nil {
		in, out := &in.AllowTransactionalWrites, &out.AllowTransactionalWrites
		*out = new(bool)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppprofileSingleClusterRouting.
func (in *AppprofileSingleClusterRouting) DeepCopy() *AppprofileSingleClusterRouting {
	if in == nil {
		return nil
	}
	out := new(AppprofileSingleClusterRouting)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BigtableAppProfile) DeepCopyInto(out *BigtableAppProfile) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BigtableAppProfile.
func (in *BigtableAppProfile) DeepCopy() *BigtableAppProfile {
	if in == nil {
		return nil
	}
	out := new(BigtableAppProfile)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BigtableAppProfile) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BigtableAppProfileList) DeepCopyInto(out *BigtableAppProfileList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]BigtableAppProfile, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BigtableAppProfileList.
func (in *BigtableAppProfileList) DeepCopy() *BigtableAppProfileList {
	if in == nil {
		return nil
	}
	out := new(BigtableAppProfileList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BigtableAppProfileList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BigtableAppProfileSpec) DeepCopyInto(out *BigtableAppProfileSpec) {
	*out = *in
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.InstanceRef != nil {
		in, out := &in.InstanceRef, &out.InstanceRef
		*out = new(v1alpha1.ResourceRef)
		**out = **in
	}
	if in.MultiClusterRoutingClusterIds != nil {
		in, out := &in.MultiClusterRoutingClusterIds, &out.MultiClusterRoutingClusterIds
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.MultiClusterRoutingUseAny != nil {
		in, out := &in.MultiClusterRoutingUseAny, &out.MultiClusterRoutingUseAny
		*out = new(bool)
		**out = **in
	}
	if in.ResourceID != nil {
		in, out := &in.ResourceID, &out.ResourceID
		*out = new(string)
		**out = **in
	}
	if in.SingleClusterRouting != nil {
		in, out := &in.SingleClusterRouting, &out.SingleClusterRouting
		*out = new(AppprofileSingleClusterRouting)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BigtableAppProfileSpec.
func (in *BigtableAppProfileSpec) DeepCopy() *BigtableAppProfileSpec {
	if in == nil {
		return nil
	}
	out := new(BigtableAppProfileSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BigtableAppProfileStatus) DeepCopyInto(out *BigtableAppProfileStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1alpha1.Condition, len(*in))
		copy(*out, *in)
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.ObservedGeneration != nil {
		in, out := &in.ObservedGeneration, &out.ObservedGeneration
		*out = new(int)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BigtableAppProfileStatus.
func (in *BigtableAppProfileStatus) DeepCopy() *BigtableAppProfileStatus {
	if in == nil {
		return nil
	}
	out := new(BigtableAppProfileStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BigtableGCPolicy) DeepCopyInto(out *BigtableGCPolicy) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BigtableGCPolicy.
func (in *BigtableGCPolicy) DeepCopy() *BigtableGCPolicy {
	if in == nil {
		return nil
	}
	out := new(BigtableGCPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BigtableGCPolicy) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BigtableGCPolicyList) DeepCopyInto(out *BigtableGCPolicyList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]BigtableGCPolicy, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BigtableGCPolicyList.
func (in *BigtableGCPolicyList) DeepCopy() *BigtableGCPolicyList {
	if in == nil {
		return nil
	}
	out := new(BigtableGCPolicyList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BigtableGCPolicyList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BigtableGCPolicySpec) DeepCopyInto(out *BigtableGCPolicySpec) {
	*out = *in
	if in.DeletionPolicy != nil {
		in, out := &in.DeletionPolicy, &out.DeletionPolicy
		*out = new(string)
		**out = **in
	}
	if in.GcRules != nil {
		in, out := &in.GcRules, &out.GcRules
		*out = new(string)
		**out = **in
	}
	out.InstanceRef = in.InstanceRef
	if in.MaxAge != nil {
		in, out := &in.MaxAge, &out.MaxAge
		*out = make([]GcpolicyMaxAge, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.MaxVersion != nil {
		in, out := &in.MaxVersion, &out.MaxVersion
		*out = make([]GcpolicyMaxVersion, len(*in))
		copy(*out, *in)
	}
	if in.Mode != nil {
		in, out := &in.Mode, &out.Mode
		*out = new(string)
		**out = **in
	}
	out.TableRef = in.TableRef
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BigtableGCPolicySpec.
func (in *BigtableGCPolicySpec) DeepCopy() *BigtableGCPolicySpec {
	if in == nil {
		return nil
	}
	out := new(BigtableGCPolicySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BigtableGCPolicyStatus) DeepCopyInto(out *BigtableGCPolicyStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1alpha1.Condition, len(*in))
		copy(*out, *in)
	}
	if in.ObservedGeneration != nil {
		in, out := &in.ObservedGeneration, &out.ObservedGeneration
		*out = new(int)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BigtableGCPolicyStatus.
func (in *BigtableGCPolicyStatus) DeepCopy() *BigtableGCPolicyStatus {
	if in == nil {
		return nil
	}
	out := new(BigtableGCPolicyStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BigtableInstance) DeepCopyInto(out *BigtableInstance) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BigtableInstance.
func (in *BigtableInstance) DeepCopy() *BigtableInstance {
	if in == nil {
		return nil
	}
	out := new(BigtableInstance)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BigtableInstance) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BigtableInstanceList) DeepCopyInto(out *BigtableInstanceList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]BigtableInstance, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BigtableInstanceList.
func (in *BigtableInstanceList) DeepCopy() *BigtableInstanceList {
	if in == nil {
		return nil
	}
	out := new(BigtableInstanceList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BigtableInstanceList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BigtableInstanceSpec) DeepCopyInto(out *BigtableInstanceSpec) {
	*out = *in
	if in.Cluster != nil {
		in, out := &in.Cluster, &out.Cluster
		*out = make([]InstanceCluster, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.DeletionProtection != nil {
		in, out := &in.DeletionProtection, &out.DeletionProtection
		*out = new(bool)
		**out = **in
	}
	if in.DisplayName != nil {
		in, out := &in.DisplayName, &out.DisplayName
		*out = new(string)
		**out = **in
	}
	if in.InstanceType != nil {
		in, out := &in.InstanceType, &out.InstanceType
		*out = new(string)
		**out = **in
	}
	if in.ResourceID != nil {
		in, out := &in.ResourceID, &out.ResourceID
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BigtableInstanceSpec.
func (in *BigtableInstanceSpec) DeepCopy() *BigtableInstanceSpec {
	if in == nil {
		return nil
	}
	out := new(BigtableInstanceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BigtableInstanceStatus) DeepCopyInto(out *BigtableInstanceStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1alpha1.Condition, len(*in))
		copy(*out, *in)
	}
	if in.ObservedGeneration != nil {
		in, out := &in.ObservedGeneration, &out.ObservedGeneration
		*out = new(int)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BigtableInstanceStatus.
func (in *BigtableInstanceStatus) DeepCopy() *BigtableInstanceStatus {
	if in == nil {
		return nil
	}
	out := new(BigtableInstanceStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BigtableTable) DeepCopyInto(out *BigtableTable) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BigtableTable.
func (in *BigtableTable) DeepCopy() *BigtableTable {
	if in == nil {
		return nil
	}
	out := new(BigtableTable)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BigtableTable) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BigtableTableList) DeepCopyInto(out *BigtableTableList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]BigtableTable, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BigtableTableList.
func (in *BigtableTableList) DeepCopy() *BigtableTableList {
	if in == nil {
		return nil
	}
	out := new(BigtableTableList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BigtableTableList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BigtableTableSpec) DeepCopyInto(out *BigtableTableSpec) {
	*out = *in
	if in.ColumnFamily != nil {
		in, out := &in.ColumnFamily, &out.ColumnFamily
		*out = make([]TableColumnFamily, len(*in))
		copy(*out, *in)
	}
	if in.DeletionProtection != nil {
		in, out := &in.DeletionProtection, &out.DeletionProtection
		*out = new(string)
		**out = **in
	}
	out.InstanceRef = in.InstanceRef
	if in.ResourceID != nil {
		in, out := &in.ResourceID, &out.ResourceID
		*out = new(string)
		**out = **in
	}
	if in.SplitKeys != nil {
		in, out := &in.SplitKeys, &out.SplitKeys
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BigtableTableSpec.
func (in *BigtableTableSpec) DeepCopy() *BigtableTableSpec {
	if in == nil {
		return nil
	}
	out := new(BigtableTableSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BigtableTableStatus) DeepCopyInto(out *BigtableTableStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1alpha1.Condition, len(*in))
		copy(*out, *in)
	}
	if in.ObservedGeneration != nil {
		in, out := &in.ObservedGeneration, &out.ObservedGeneration
		*out = new(int)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BigtableTableStatus.
func (in *BigtableTableStatus) DeepCopy() *BigtableTableStatus {
	if in == nil {
		return nil
	}
	out := new(BigtableTableStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GcpolicyMaxAge) DeepCopyInto(out *GcpolicyMaxAge) {
	*out = *in
	if in.Days != nil {
		in, out := &in.Days, &out.Days
		*out = new(int)
		**out = **in
	}
	if in.Duration != nil {
		in, out := &in.Duration, &out.Duration
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GcpolicyMaxAge.
func (in *GcpolicyMaxAge) DeepCopy() *GcpolicyMaxAge {
	if in == nil {
		return nil
	}
	out := new(GcpolicyMaxAge)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GcpolicyMaxVersion) DeepCopyInto(out *GcpolicyMaxVersion) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GcpolicyMaxVersion.
func (in *GcpolicyMaxVersion) DeepCopy() *GcpolicyMaxVersion {
	if in == nil {
		return nil
	}
	out := new(GcpolicyMaxVersion)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InstanceAutoscalingConfig) DeepCopyInto(out *InstanceAutoscalingConfig) {
	*out = *in
	if in.StorageTarget != nil {
		in, out := &in.StorageTarget, &out.StorageTarget
		*out = new(int)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InstanceAutoscalingConfig.
func (in *InstanceAutoscalingConfig) DeepCopy() *InstanceAutoscalingConfig {
	if in == nil {
		return nil
	}
	out := new(InstanceAutoscalingConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InstanceCluster) DeepCopyInto(out *InstanceCluster) {
	*out = *in
	if in.AutoscalingConfig != nil {
		in, out := &in.AutoscalingConfig, &out.AutoscalingConfig
		*out = new(InstanceAutoscalingConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.KmsKeyRef != nil {
		in, out := &in.KmsKeyRef, &out.KmsKeyRef
		*out = new(v1alpha1.ResourceRef)
		**out = **in
	}
	if in.NumNodes != nil {
		in, out := &in.NumNodes, &out.NumNodes
		*out = new(int)
		**out = **in
	}
	if in.StorageType != nil {
		in, out := &in.StorageType, &out.StorageType
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InstanceCluster.
func (in *InstanceCluster) DeepCopy() *InstanceCluster {
	if in == nil {
		return nil
	}
	out := new(InstanceCluster)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TableColumnFamily) DeepCopyInto(out *TableColumnFamily) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TableColumnFamily.
func (in *TableColumnFamily) DeepCopy() *TableColumnFamily {
	if in == nil {
		return nil
	}
	out := new(TableColumnFamily)
	in.DeepCopyInto(out)
	return out
}
