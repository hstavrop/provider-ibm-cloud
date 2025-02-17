// +build !ignore_autogenerated

/*
Copyright 2020 The Crossplane Authors.

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

package v1alpha1

import (
	corev1alpha1 "github.com/crossplane/crossplane-runtime/apis/core/v1alpha1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ActivityTracking) DeepCopyInto(out *ActivityTracking) {
	*out = *in
	if in.ReadDataEvents != nil {
		in, out := &in.ReadDataEvents, &out.ReadDataEvents
		*out = new(bool)
		**out = **in
	}
	if in.WriteDataEvents != nil {
		in, out := &in.WriteDataEvents, &out.WriteDataEvents
		*out = new(bool)
		**out = **in
	}
	if in.ActivityTrackerCRN != nil {
		in, out := &in.ActivityTrackerCRN, &out.ActivityTrackerCRN
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ActivityTracking.
func (in *ActivityTracking) DeepCopy() *ActivityTracking {
	if in == nil {
		return nil
	}
	out := new(ActivityTracking)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Bucket) DeepCopyInto(out *Bucket) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Bucket.
func (in *Bucket) DeepCopy() *Bucket {
	if in == nil {
		return nil
	}
	out := new(Bucket)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Bucket) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BucketConfig) DeepCopyInto(out *BucketConfig) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BucketConfig.
func (in *BucketConfig) DeepCopy() *BucketConfig {
	if in == nil {
		return nil
	}
	out := new(BucketConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BucketConfig) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BucketConfigList) DeepCopyInto(out *BucketConfigList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]BucketConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BucketConfigList.
func (in *BucketConfigList) DeepCopy() *BucketConfigList {
	if in == nil {
		return nil
	}
	out := new(BucketConfigList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BucketConfigList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BucketConfigObservation) DeepCopyInto(out *BucketConfigObservation) {
	*out = *in
	in.TimeCreated.DeepCopyInto(&out.TimeCreated)
	in.TimeUpdated.DeepCopyInto(&out.TimeUpdated)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BucketConfigObservation.
func (in *BucketConfigObservation) DeepCopy() *BucketConfigObservation {
	if in == nil {
		return nil
	}
	out := new(BucketConfigObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BucketConfigParams) DeepCopyInto(out *BucketConfigParams) {
	*out = *in
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.NameRef != nil {
		in, out := &in.NameRef, &out.NameRef
		*out = new(corev1alpha1.Reference)
		**out = **in
	}
	if in.NameSelector != nil {
		in, out := &in.NameSelector, &out.NameSelector
		*out = new(corev1alpha1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.HardQuota != nil {
		in, out := &in.HardQuota, &out.HardQuota
		*out = new(int64)
		**out = **in
	}
	if in.Firewall != nil {
		in, out := &in.Firewall, &out.Firewall
		*out = new(Firewall)
		(*in).DeepCopyInto(*out)
	}
	if in.ActivityTracking != nil {
		in, out := &in.ActivityTracking, &out.ActivityTracking
		*out = new(ActivityTracking)
		(*in).DeepCopyInto(*out)
	}
	if in.MetricsMonitoring != nil {
		in, out := &in.MetricsMonitoring, &out.MetricsMonitoring
		*out = new(MetricsMonitoring)
		(*in).DeepCopyInto(*out)
	}
	if in.Headers != nil {
		in, out := &in.Headers, &out.Headers
		*out = new(map[string]string)
		if **in != nil {
			in, out := *in, *out
			*out = make(map[string]string, len(*in))
			for key, val := range *in {
				(*out)[key] = val
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BucketConfigParams.
func (in *BucketConfigParams) DeepCopy() *BucketConfigParams {
	if in == nil {
		return nil
	}
	out := new(BucketConfigParams)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BucketConfigSpec) DeepCopyInto(out *BucketConfigSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BucketConfigSpec.
func (in *BucketConfigSpec) DeepCopy() *BucketConfigSpec {
	if in == nil {
		return nil
	}
	out := new(BucketConfigSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BucketConfigStatus) DeepCopyInto(out *BucketConfigStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BucketConfigStatus.
func (in *BucketConfigStatus) DeepCopy() *BucketConfigStatus {
	if in == nil {
		return nil
	}
	out := new(BucketConfigStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BucketList) DeepCopyInto(out *BucketList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Bucket, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BucketList.
func (in *BucketList) DeepCopy() *BucketList {
	if in == nil {
		return nil
	}
	out := new(BucketList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BucketList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BucketObservation) DeepCopyInto(out *BucketObservation) {
	*out = *in
	if in.CreationDate != nil {
		in, out := &in.CreationDate, &out.CreationDate
		*out = (*in).DeepCopy()
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BucketObservation.
func (in *BucketObservation) DeepCopy() *BucketObservation {
	if in == nil {
		return nil
	}
	out := new(BucketObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BucketPararams) DeepCopyInto(out *BucketPararams) {
	*out = *in
	if in.IbmServiceInstanceID != nil {
		in, out := &in.IbmServiceInstanceID, &out.IbmServiceInstanceID
		*out = new(string)
		**out = **in
	}
	if in.IbmServiceInstanceIDRef != nil {
		in, out := &in.IbmServiceInstanceIDRef, &out.IbmServiceInstanceIDRef
		*out = new(corev1alpha1.Reference)
		**out = **in
	}
	if in.IbmServiceInstanceIDSelector != nil {
		in, out := &in.IbmServiceInstanceIDSelector, &out.IbmServiceInstanceIDSelector
		*out = new(corev1alpha1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.IbmSSEKpEncryptionAlgorithm != nil {
		in, out := &in.IbmSSEKpEncryptionAlgorithm, &out.IbmSSEKpEncryptionAlgorithm
		*out = new(string)
		**out = **in
	}
	if in.IbmSSEKpCustomerRootKeyCrn != nil {
		in, out := &in.IbmSSEKpCustomerRootKeyCrn, &out.IbmSSEKpCustomerRootKeyCrn
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BucketPararams.
func (in *BucketPararams) DeepCopy() *BucketPararams {
	if in == nil {
		return nil
	}
	out := new(BucketPararams)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BucketSpec) DeepCopyInto(out *BucketSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BucketSpec.
func (in *BucketSpec) DeepCopy() *BucketSpec {
	if in == nil {
		return nil
	}
	out := new(BucketSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BucketStatus) DeepCopyInto(out *BucketStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BucketStatus.
func (in *BucketStatus) DeepCopy() *BucketStatus {
	if in == nil {
		return nil
	}
	out := new(BucketStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Firewall) DeepCopyInto(out *Firewall) {
	*out = *in
	if in.AllowedIP != nil {
		in, out := &in.AllowedIP, &out.AllowedIP
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Firewall.
func (in *Firewall) DeepCopy() *Firewall {
	if in == nil {
		return nil
	}
	out := new(Firewall)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MetricsMonitoring) DeepCopyInto(out *MetricsMonitoring) {
	*out = *in
	if in.UsageMetricsEnabled != nil {
		in, out := &in.UsageMetricsEnabled, &out.UsageMetricsEnabled
		*out = new(bool)
		**out = **in
	}
	if in.RequestMetricsEnabled != nil {
		in, out := &in.RequestMetricsEnabled, &out.RequestMetricsEnabled
		*out = new(bool)
		**out = **in
	}
	if in.MetricsMonitoringCRN != nil {
		in, out := &in.MetricsMonitoringCRN, &out.MetricsMonitoringCRN
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MetricsMonitoring.
func (in *MetricsMonitoring) DeepCopy() *MetricsMonitoring {
	if in == nil {
		return nil
	}
	out := new(MetricsMonitoring)
	in.DeepCopyInto(out)
	return out
}
