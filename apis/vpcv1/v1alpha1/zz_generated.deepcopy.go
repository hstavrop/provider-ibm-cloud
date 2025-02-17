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
func (in *IP) DeepCopyInto(out *IP) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IP.
func (in *IP) DeepCopy() *IP {
	if in == nil {
		return nil
	}
	out := new(IP)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkACLIdentity) DeepCopyInto(out *NetworkACLIdentity) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkACLIdentity.
func (in *NetworkACLIdentity) DeepCopy() *NetworkACLIdentity {
	if in == nil {
		return nil
	}
	out := new(NetworkACLIdentity)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkACLReference) DeepCopyInto(out *NetworkACLReference) {
	*out = *in
	if in.Deleted != nil {
		in, out := &in.Deleted, &out.Deleted
		*out = new(NetworkACLReferenceDeleted)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkACLReference.
func (in *NetworkACLReference) DeepCopy() *NetworkACLReference {
	if in == nil {
		return nil
	}
	out := new(NetworkACLReference)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkACLReferenceDeleted) DeepCopyInto(out *NetworkACLReferenceDeleted) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkACLReferenceDeleted.
func (in *NetworkACLReferenceDeleted) DeepCopy() *NetworkACLReferenceDeleted {
	if in == nil {
		return nil
	}
	out := new(NetworkACLReferenceDeleted)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PublicGatewayIdentity) DeepCopyInto(out *PublicGatewayIdentity) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PublicGatewayIdentity.
func (in *PublicGatewayIdentity) DeepCopy() *PublicGatewayIdentity {
	if in == nil {
		return nil
	}
	out := new(PublicGatewayIdentity)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PublicGatewayReference) DeepCopyInto(out *PublicGatewayReference) {
	*out = *in
	if in.Deleted != nil {
		in, out := &in.Deleted, &out.Deleted
		*out = new(PublicGatewayReferenceDeleted)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PublicGatewayReference.
func (in *PublicGatewayReference) DeepCopy() *PublicGatewayReference {
	if in == nil {
		return nil
	}
	out := new(PublicGatewayReference)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PublicGatewayReferenceDeleted) DeepCopyInto(out *PublicGatewayReferenceDeleted) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PublicGatewayReferenceDeleted.
func (in *PublicGatewayReferenceDeleted) DeepCopy() *PublicGatewayReferenceDeleted {
	if in == nil {
		return nil
	}
	out := new(PublicGatewayReferenceDeleted)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceGroupIdentity) DeepCopyInto(out *ResourceGroupIdentity) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceGroupIdentity.
func (in *ResourceGroupIdentity) DeepCopy() *ResourceGroupIdentity {
	if in == nil {
		return nil
	}
	out := new(ResourceGroupIdentity)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceGroupReference) DeepCopyInto(out *ResourceGroupReference) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceGroupReference.
func (in *ResourceGroupReference) DeepCopy() *ResourceGroupReference {
	if in == nil {
		return nil
	}
	out := new(ResourceGroupReference)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RoutingTableIdentity) DeepCopyInto(out *RoutingTableIdentity) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RoutingTableIdentity.
func (in *RoutingTableIdentity) DeepCopy() *RoutingTableIdentity {
	if in == nil {
		return nil
	}
	out := new(RoutingTableIdentity)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RoutingTableReference) DeepCopyInto(out *RoutingTableReference) {
	*out = *in
	if in.Deleted != nil {
		in, out := &in.Deleted, &out.Deleted
		*out = new(RoutingTableReferenceDeleted)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RoutingTableReference.
func (in *RoutingTableReference) DeepCopy() *RoutingTableReference {
	if in == nil {
		return nil
	}
	out := new(RoutingTableReference)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RoutingTableReferenceDeleted) DeepCopyInto(out *RoutingTableReferenceDeleted) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RoutingTableReferenceDeleted.
func (in *RoutingTableReferenceDeleted) DeepCopy() *RoutingTableReferenceDeleted {
	if in == nil {
		return nil
	}
	out := new(RoutingTableReferenceDeleted)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecurityGroupReference) DeepCopyInto(out *SecurityGroupReference) {
	*out = *in
	if in.Deleted != nil {
		in, out := &in.Deleted, &out.Deleted
		*out = new(SecurityGroupReferenceDeleted)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecurityGroupReference.
func (in *SecurityGroupReference) DeepCopy() *SecurityGroupReference {
	if in == nil {
		return nil
	}
	out := new(SecurityGroupReference)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecurityGroupReferenceDeleted) DeepCopyInto(out *SecurityGroupReferenceDeleted) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecurityGroupReferenceDeleted.
func (in *SecurityGroupReferenceDeleted) DeepCopy() *SecurityGroupReferenceDeleted {
	if in == nil {
		return nil
	}
	out := new(SecurityGroupReferenceDeleted)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Subnet) DeepCopyInto(out *Subnet) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Subnet.
func (in *Subnet) DeepCopy() *Subnet {
	if in == nil {
		return nil
	}
	out := new(Subnet)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Subnet) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SubnetList) DeepCopyInto(out *SubnetList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Subnet, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SubnetList.
func (in *SubnetList) DeepCopy() *SubnetList {
	if in == nil {
		return nil
	}
	out := new(SubnetList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SubnetList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SubnetObservation) DeepCopyInto(out *SubnetObservation) {
	*out = *in
	if in.CreatedAt != nil {
		in, out := &in.CreatedAt, &out.CreatedAt
		*out = (*in).DeepCopy()
	}
	in.NetworkACL.DeepCopyInto(&out.NetworkACL)
	in.PublicGateway.DeepCopyInto(&out.PublicGateway)
	out.ResourceGroup = in.ResourceGroup
	in.RoutingTable.DeepCopyInto(&out.RoutingTable)
	in.VPC.DeepCopyInto(&out.VPC)
	out.Zone = in.Zone
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SubnetObservation.
func (in *SubnetObservation) DeepCopy() *SubnetObservation {
	if in == nil {
		return nil
	}
	out := new(SubnetObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SubnetParameters) DeepCopyInto(out *SubnetParameters) {
	*out = *in
	if in.ByTocalCount != nil {
		in, out := &in.ByTocalCount, &out.ByTocalCount
		*out = new(SubnetPrototypeSubnetByTotalCount)
		(*in).DeepCopyInto(*out)
	}
	if in.ByCIDR != nil {
		in, out := &in.ByCIDR, &out.ByCIDR
		*out = new(SubnetPrototypeSubnetByCIDR)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SubnetParameters.
func (in *SubnetParameters) DeepCopy() *SubnetParameters {
	if in == nil {
		return nil
	}
	out := new(SubnetParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SubnetPrototypeSubnetByCIDR) DeepCopyInto(out *SubnetPrototypeSubnetByCIDR) {
	*out = *in
	if in.IPVersion != nil {
		in, out := &in.IPVersion, &out.IPVersion
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.NetworkACL != nil {
		in, out := &in.NetworkACL, &out.NetworkACL
		*out = new(NetworkACLIdentity)
		**out = **in
	}
	if in.PublicGateway != nil {
		in, out := &in.PublicGateway, &out.PublicGateway
		*out = new(PublicGatewayIdentity)
		**out = **in
	}
	if in.ResourceGroup != nil {
		in, out := &in.ResourceGroup, &out.ResourceGroup
		*out = new(ResourceGroupIdentity)
		**out = **in
	}
	if in.RoutingTable != nil {
		in, out := &in.RoutingTable, &out.RoutingTable
		*out = new(RoutingTableIdentity)
		**out = **in
	}
	in.VPC.DeepCopyInto(&out.VPC)
	if in.Zone != nil {
		in, out := &in.Zone, &out.Zone
		*out = new(ZoneIdentity)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SubnetPrototypeSubnetByCIDR.
func (in *SubnetPrototypeSubnetByCIDR) DeepCopy() *SubnetPrototypeSubnetByCIDR {
	if in == nil {
		return nil
	}
	out := new(SubnetPrototypeSubnetByCIDR)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SubnetPrototypeSubnetByTotalCount) DeepCopyInto(out *SubnetPrototypeSubnetByTotalCount) {
	*out = *in
	if in.IPVersion != nil {
		in, out := &in.IPVersion, &out.IPVersion
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.NetworkACL != nil {
		in, out := &in.NetworkACL, &out.NetworkACL
		*out = new(NetworkACLIdentity)
		**out = **in
	}
	if in.PublicGateway != nil {
		in, out := &in.PublicGateway, &out.PublicGateway
		*out = new(PublicGatewayIdentity)
		**out = **in
	}
	if in.ResourceGroup != nil {
		in, out := &in.ResourceGroup, &out.ResourceGroup
		*out = new(ResourceGroupIdentity)
		**out = **in
	}
	if in.RoutingTable != nil {
		in, out := &in.RoutingTable, &out.RoutingTable
		*out = new(RoutingTableIdentity)
		**out = **in
	}
	in.VPC.DeepCopyInto(&out.VPC)
	out.Zone = in.Zone
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SubnetPrototypeSubnetByTotalCount.
func (in *SubnetPrototypeSubnetByTotalCount) DeepCopy() *SubnetPrototypeSubnetByTotalCount {
	if in == nil {
		return nil
	}
	out := new(SubnetPrototypeSubnetByTotalCount)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SubnetSpec) DeepCopyInto(out *SubnetSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SubnetSpec.
func (in *SubnetSpec) DeepCopy() *SubnetSpec {
	if in == nil {
		return nil
	}
	out := new(SubnetSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SubnetStatus) DeepCopyInto(out *SubnetStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SubnetStatus.
func (in *SubnetStatus) DeepCopy() *SubnetStatus {
	if in == nil {
		return nil
	}
	out := new(SubnetStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VPC) DeepCopyInto(out *VPC) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VPC.
func (in *VPC) DeepCopy() *VPC {
	if in == nil {
		return nil
	}
	out := new(VPC)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VPC) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VPCIdentity) DeepCopyInto(out *VPCIdentity) {
	*out = *in
	if in.VPCRef != nil {
		in, out := &in.VPCRef, &out.VPCRef
		*out = new(corev1alpha1.Reference)
		**out = **in
	}
	if in.VPCSelector != nil {
		in, out := &in.VPCSelector, &out.VPCSelector
		*out = new(corev1alpha1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VPCIdentity.
func (in *VPCIdentity) DeepCopy() *VPCIdentity {
	if in == nil {
		return nil
	}
	out := new(VPCIdentity)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VPCList) DeepCopyInto(out *VPCList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]VPC, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VPCList.
func (in *VPCList) DeepCopy() *VPCList {
	if in == nil {
		return nil
	}
	out := new(VPCList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VPCList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VPCObservation) DeepCopyInto(out *VPCObservation) {
	*out = *in
	if in.CreatedAt != nil {
		in, out := &in.CreatedAt, &out.CreatedAt
		*out = (*in).DeepCopy()
	}
	if in.CseSourceIps != nil {
		in, out := &in.CseSourceIps, &out.CseSourceIps
		*out = make([]VpccseSourceIP, len(*in))
		copy(*out, *in)
	}
	in.DefaultNetworkACL.DeepCopyInto(&out.DefaultNetworkACL)
	in.DefaultRoutingTable.DeepCopyInto(&out.DefaultRoutingTable)
	in.DefaultSecurityGroup.DeepCopyInto(&out.DefaultSecurityGroup)
	out.ResourceGroup = in.ResourceGroup
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VPCObservation.
func (in *VPCObservation) DeepCopy() *VPCObservation {
	if in == nil {
		return nil
	}
	out := new(VPCObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VPCParameters) DeepCopyInto(out *VPCParameters) {
	*out = *in
	if in.AddressPrefixManagement != nil {
		in, out := &in.AddressPrefixManagement, &out.AddressPrefixManagement
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.ResourceGroup != nil {
		in, out := &in.ResourceGroup, &out.ResourceGroup
		*out = new(ResourceGroupIdentity)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VPCParameters.
func (in *VPCParameters) DeepCopy() *VPCParameters {
	if in == nil {
		return nil
	}
	out := new(VPCParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VPCReference) DeepCopyInto(out *VPCReference) {
	*out = *in
	if in.Deleted != nil {
		in, out := &in.Deleted, &out.Deleted
		*out = new(VPCReferenceDeleted)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VPCReference.
func (in *VPCReference) DeepCopy() *VPCReference {
	if in == nil {
		return nil
	}
	out := new(VPCReference)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VPCReferenceDeleted) DeepCopyInto(out *VPCReferenceDeleted) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VPCReferenceDeleted.
func (in *VPCReferenceDeleted) DeepCopy() *VPCReferenceDeleted {
	if in == nil {
		return nil
	}
	out := new(VPCReferenceDeleted)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VPCSpec) DeepCopyInto(out *VPCSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VPCSpec.
func (in *VPCSpec) DeepCopy() *VPCSpec {
	if in == nil {
		return nil
	}
	out := new(VPCSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VPCStatus) DeepCopyInto(out *VPCStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VPCStatus.
func (in *VPCStatus) DeepCopy() *VPCStatus {
	if in == nil {
		return nil
	}
	out := new(VPCStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VpccseSourceIP) DeepCopyInto(out *VpccseSourceIP) {
	*out = *in
	out.IP = in.IP
	out.Zone = in.Zone
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VpccseSourceIP.
func (in *VpccseSourceIP) DeepCopy() *VpccseSourceIP {
	if in == nil {
		return nil
	}
	out := new(VpccseSourceIP)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ZoneIdentity) DeepCopyInto(out *ZoneIdentity) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ZoneIdentity.
func (in *ZoneIdentity) DeepCopy() *ZoneIdentity {
	if in == nil {
		return nil
	}
	out := new(ZoneIdentity)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ZoneReference) DeepCopyInto(out *ZoneReference) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ZoneReference.
func (in *ZoneReference) DeepCopy() *ZoneReference {
	if in == nil {
		return nil
	}
	out := new(ZoneReference)
	in.DeepCopyInto(out)
	return out
}
