// +build !ignore_autogenerated

/*
 * Copyright 2018-2019, EnMasse authors.
 * License: Apache License 2.0 (see the file LICENSE or http://apache.org/licenses/LICENSE-2.0.html).
 */

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1beta2

import (
	v1beta1 "github.com/enmasseproject/enmasse/pkg/apis/enmasse/v1beta1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MessagingInfra) DeepCopyInto(out *MessagingInfra) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MessagingInfra.
func (in *MessagingInfra) DeepCopy() *MessagingInfra {
	if in == nil {
		return nil
	}
	out := new(MessagingInfra)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MessagingInfra) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MessagingInfraCondition) DeepCopyInto(out *MessagingInfraCondition) {
	*out = *in
	in.LastTransitionTime.DeepCopyInto(&out.LastTransitionTime)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MessagingInfraCondition.
func (in *MessagingInfraCondition) DeepCopy() *MessagingInfraCondition {
	if in == nil {
		return nil
	}
	out := new(MessagingInfraCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MessagingInfraList) DeepCopyInto(out *MessagingInfraList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]MessagingInfra, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MessagingInfraList.
func (in *MessagingInfraList) DeepCopy() *MessagingInfraList {
	if in == nil {
		return nil
	}
	out := new(MessagingInfraList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MessagingInfraList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MessagingInfraReference) DeepCopyInto(out *MessagingInfraReference) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MessagingInfraReference.
func (in *MessagingInfraReference) DeepCopy() *MessagingInfraReference {
	if in == nil {
		return nil
	}
	out := new(MessagingInfraReference)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MessagingInfraSpec) DeepCopyInto(out *MessagingInfraSpec) {
	*out = *in
	if in.Selector != nil {
		in, out := &in.Selector, &out.Selector
		*out = new(Selector)
		(*in).DeepCopyInto(*out)
	}
	in.Router.DeepCopyInto(&out.Router)
	in.Broker.DeepCopyInto(&out.Broker)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MessagingInfraSpec.
func (in *MessagingInfraSpec) DeepCopy() *MessagingInfraSpec {
	if in == nil {
		return nil
	}
	out := new(MessagingInfraSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MessagingInfraSpecBroker) DeepCopyInto(out *MessagingInfraSpecBroker) {
	*out = *in
	if in.InitImage != nil {
		in, out := &in.InitImage, &out.InitImage
		*out = new(v1beta1.ImageOverride)
		**out = **in
	}
	if in.Image != nil {
		in, out := &in.Image, &out.Image
		*out = new(v1beta1.ImageOverride)
		**out = **in
	}
	if in.ScalingStrategy != nil {
		in, out := &in.ScalingStrategy, &out.ScalingStrategy
		*out = new(MessagingInfraSpecBrokerScalingStrategy)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MessagingInfraSpecBroker.
func (in *MessagingInfraSpecBroker) DeepCopy() *MessagingInfraSpecBroker {
	if in == nil {
		return nil
	}
	out := new(MessagingInfraSpecBroker)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MessagingInfraSpecBrokerScalingStrategy) DeepCopyInto(out *MessagingInfraSpecBrokerScalingStrategy) {
	*out = *in
	if in.Static != nil {
		in, out := &in.Static, &out.Static
		*out = new(MessagingInfraSpecBrokerScalingStrategyStatic)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MessagingInfraSpecBrokerScalingStrategy.
func (in *MessagingInfraSpecBrokerScalingStrategy) DeepCopy() *MessagingInfraSpecBrokerScalingStrategy {
	if in == nil {
		return nil
	}
	out := new(MessagingInfraSpecBrokerScalingStrategy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MessagingInfraSpecBrokerScalingStrategyStatic) DeepCopyInto(out *MessagingInfraSpecBrokerScalingStrategyStatic) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MessagingInfraSpecBrokerScalingStrategyStatic.
func (in *MessagingInfraSpecBrokerScalingStrategyStatic) DeepCopy() *MessagingInfraSpecBrokerScalingStrategyStatic {
	if in == nil {
		return nil
	}
	out := new(MessagingInfraSpecBrokerScalingStrategyStatic)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MessagingInfraSpecRouter) DeepCopyInto(out *MessagingInfraSpecRouter) {
	*out = *in
	if in.Image != nil {
		in, out := &in.Image, &out.Image
		*out = new(v1beta1.ImageOverride)
		**out = **in
	}
	if in.ScalingStrategy != nil {
		in, out := &in.ScalingStrategy, &out.ScalingStrategy
		*out = new(MessagingInfraSpecRouterScalingStrategy)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MessagingInfraSpecRouter.
func (in *MessagingInfraSpecRouter) DeepCopy() *MessagingInfraSpecRouter {
	if in == nil {
		return nil
	}
	out := new(MessagingInfraSpecRouter)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MessagingInfraSpecRouterScalingStrategy) DeepCopyInto(out *MessagingInfraSpecRouterScalingStrategy) {
	*out = *in
	if in.Static != nil {
		in, out := &in.Static, &out.Static
		*out = new(MessagingInfraSpecRouterScalingStrategyStatic)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MessagingInfraSpecRouterScalingStrategy.
func (in *MessagingInfraSpecRouterScalingStrategy) DeepCopy() *MessagingInfraSpecRouterScalingStrategy {
	if in == nil {
		return nil
	}
	out := new(MessagingInfraSpecRouterScalingStrategy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MessagingInfraSpecRouterScalingStrategyStatic) DeepCopyInto(out *MessagingInfraSpecRouterScalingStrategyStatic) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MessagingInfraSpecRouterScalingStrategyStatic.
func (in *MessagingInfraSpecRouterScalingStrategyStatic) DeepCopy() *MessagingInfraSpecRouterScalingStrategyStatic {
	if in == nil {
		return nil
	}
	out := new(MessagingInfraSpecRouterScalingStrategyStatic)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MessagingInfraStatus) DeepCopyInto(out *MessagingInfraStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]MessagingInfraCondition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MessagingInfraStatus.
func (in *MessagingInfraStatus) DeepCopy() *MessagingInfraStatus {
	if in == nil {
		return nil
	}
	out := new(MessagingInfraStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MessagingTenant) DeepCopyInto(out *MessagingTenant) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MessagingTenant.
func (in *MessagingTenant) DeepCopy() *MessagingTenant {
	if in == nil {
		return nil
	}
	out := new(MessagingTenant)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MessagingTenant) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MessagingTenantCondition) DeepCopyInto(out *MessagingTenantCondition) {
	*out = *in
	in.LastTransitionTime.DeepCopyInto(&out.LastTransitionTime)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MessagingTenantCondition.
func (in *MessagingTenantCondition) DeepCopy() *MessagingTenantCondition {
	if in == nil {
		return nil
	}
	out := new(MessagingTenantCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MessagingTenantList) DeepCopyInto(out *MessagingTenantList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]MessagingTenant, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MessagingTenantList.
func (in *MessagingTenantList) DeepCopy() *MessagingTenantList {
	if in == nil {
		return nil
	}
	out := new(MessagingTenantList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MessagingTenantList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MessagingTenantSpec) DeepCopyInto(out *MessagingTenantSpec) {
	*out = *in
	if in.MessagingInfraRef != nil {
		in, out := &in.MessagingInfraRef, &out.MessagingInfraRef
		*out = new(MessagingInfraReference)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MessagingTenantSpec.
func (in *MessagingTenantSpec) DeepCopy() *MessagingTenantSpec {
	if in == nil {
		return nil
	}
	out := new(MessagingTenantSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MessagingTenantStatus) DeepCopyInto(out *MessagingTenantStatus) {
	*out = *in
	if in.MessagingInfraRef != nil {
		in, out := &in.MessagingInfraRef, &out.MessagingInfraRef
		*out = new(MessagingInfraReference)
		**out = **in
	}
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]MessagingTenantCondition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MessagingTenantStatus.
func (in *MessagingTenantStatus) DeepCopy() *MessagingTenantStatus {
	if in == nil {
		return nil
	}
	out := new(MessagingTenantStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Selector) DeepCopyInto(out *Selector) {
	*out = *in
	if in.Namespaces != nil {
		in, out := &in.Namespaces, &out.Namespaces
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Selector.
func (in *Selector) DeepCopy() *Selector {
	if in == nil {
		return nil
	}
	out := new(Selector)
	in.DeepCopyInto(out)
	return out
}
