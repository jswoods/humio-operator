// +build !ignore_autogenerated

// Code generated by operator-sdk. DO NOT EDIT.

package v1alpha1

import (
	v1 "k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HumioCluster) DeepCopyInto(out *HumioCluster) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HumioCluster.
func (in *HumioCluster) DeepCopy() *HumioCluster {
	if in == nil {
		return nil
	}
	out := new(HumioCluster)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *HumioCluster) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HumioClusterIngressSpec) DeepCopyInto(out *HumioClusterIngressSpec) {
	*out = *in
	if in.Annotations != nil {
		in, out := &in.Annotations, &out.Annotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HumioClusterIngressSpec.
func (in *HumioClusterIngressSpec) DeepCopy() *HumioClusterIngressSpec {
	if in == nil {
		return nil
	}
	out := new(HumioClusterIngressSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HumioClusterList) DeepCopyInto(out *HumioClusterList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]HumioCluster, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HumioClusterList.
func (in *HumioClusterList) DeepCopy() *HumioClusterList {
	if in == nil {
		return nil
	}
	out := new(HumioClusterList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *HumioClusterList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HumioClusterSpec) DeepCopyInto(out *HumioClusterSpec) {
	*out = *in
	if in.EnvironmentVariables != nil {
		in, out := &in.EnvironmentVariables, &out.EnvironmentVariables
		*out = make([]v1.EnvVar, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	in.DataVolumeSource.DeepCopyInto(&out.DataVolumeSource)
	if in.ImagePullSecrets != nil {
		in, out := &in.ImagePullSecrets, &out.ImagePullSecrets
		*out = make([]v1.LocalObjectReference, len(*in))
		copy(*out, *in)
	}
	in.Affinity.DeepCopyInto(&out.Affinity)
	in.Resources.DeepCopyInto(&out.Resources)
	if in.ContainerSecurityContext != nil {
		in, out := &in.ContainerSecurityContext, &out.ContainerSecurityContext
		*out = new(v1.SecurityContext)
		(*in).DeepCopyInto(*out)
	}
	if in.PodSecurityContext != nil {
		in, out := &in.PodSecurityContext, &out.PodSecurityContext
		*out = new(v1.PodSecurityContext)
		(*in).DeepCopyInto(*out)
	}
	in.Ingress.DeepCopyInto(&out.Ingress)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HumioClusterSpec.
func (in *HumioClusterSpec) DeepCopy() *HumioClusterSpec {
	if in == nil {
		return nil
	}
	out := new(HumioClusterSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HumioClusterStatus) DeepCopyInto(out *HumioClusterStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HumioClusterStatus.
func (in *HumioClusterStatus) DeepCopy() *HumioClusterStatus {
	if in == nil {
		return nil
	}
	out := new(HumioClusterStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HumioExternalCluster) DeepCopyInto(out *HumioExternalCluster) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HumioExternalCluster.
func (in *HumioExternalCluster) DeepCopy() *HumioExternalCluster {
	if in == nil {
		return nil
	}
	out := new(HumioExternalCluster)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *HumioExternalCluster) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HumioExternalClusterList) DeepCopyInto(out *HumioExternalClusterList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]HumioExternalCluster, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HumioExternalClusterList.
func (in *HumioExternalClusterList) DeepCopy() *HumioExternalClusterList {
	if in == nil {
		return nil
	}
	out := new(HumioExternalClusterList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *HumioExternalClusterList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HumioExternalClusterSpec) DeepCopyInto(out *HumioExternalClusterSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HumioExternalClusterSpec.
func (in *HumioExternalClusterSpec) DeepCopy() *HumioExternalClusterSpec {
	if in == nil {
		return nil
	}
	out := new(HumioExternalClusterSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HumioExternalClusterStatus) DeepCopyInto(out *HumioExternalClusterStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HumioExternalClusterStatus.
func (in *HumioExternalClusterStatus) DeepCopy() *HumioExternalClusterStatus {
	if in == nil {
		return nil
	}
	out := new(HumioExternalClusterStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HumioIngestToken) DeepCopyInto(out *HumioIngestToken) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HumioIngestToken.
func (in *HumioIngestToken) DeepCopy() *HumioIngestToken {
	if in == nil {
		return nil
	}
	out := new(HumioIngestToken)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *HumioIngestToken) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HumioIngestTokenList) DeepCopyInto(out *HumioIngestTokenList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]HumioIngestToken, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HumioIngestTokenList.
func (in *HumioIngestTokenList) DeepCopy() *HumioIngestTokenList {
	if in == nil {
		return nil
	}
	out := new(HumioIngestTokenList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *HumioIngestTokenList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HumioIngestTokenSpec) DeepCopyInto(out *HumioIngestTokenSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HumioIngestTokenSpec.
func (in *HumioIngestTokenSpec) DeepCopy() *HumioIngestTokenSpec {
	if in == nil {
		return nil
	}
	out := new(HumioIngestTokenSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HumioIngestTokenStatus) DeepCopyInto(out *HumioIngestTokenStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HumioIngestTokenStatus.
func (in *HumioIngestTokenStatus) DeepCopy() *HumioIngestTokenStatus {
	if in == nil {
		return nil
	}
	out := new(HumioIngestTokenStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HumioParser) DeepCopyInto(out *HumioParser) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HumioParser.
func (in *HumioParser) DeepCopy() *HumioParser {
	if in == nil {
		return nil
	}
	out := new(HumioParser)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *HumioParser) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HumioParserList) DeepCopyInto(out *HumioParserList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]HumioParser, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HumioParserList.
func (in *HumioParserList) DeepCopy() *HumioParserList {
	if in == nil {
		return nil
	}
	out := new(HumioParserList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *HumioParserList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HumioParserSpec) DeepCopyInto(out *HumioParserSpec) {
	*out = *in
	if in.TagFields != nil {
		in, out := &in.TagFields, &out.TagFields
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.TestData != nil {
		in, out := &in.TestData, &out.TestData
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HumioParserSpec.
func (in *HumioParserSpec) DeepCopy() *HumioParserSpec {
	if in == nil {
		return nil
	}
	out := new(HumioParserSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HumioParserStatus) DeepCopyInto(out *HumioParserStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HumioParserStatus.
func (in *HumioParserStatus) DeepCopy() *HumioParserStatus {
	if in == nil {
		return nil
	}
	out := new(HumioParserStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HumioRepository) DeepCopyInto(out *HumioRepository) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HumioRepository.
func (in *HumioRepository) DeepCopy() *HumioRepository {
	if in == nil {
		return nil
	}
	out := new(HumioRepository)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *HumioRepository) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HumioRepositoryList) DeepCopyInto(out *HumioRepositoryList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]HumioRepository, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HumioRepositoryList.
func (in *HumioRepositoryList) DeepCopy() *HumioRepositoryList {
	if in == nil {
		return nil
	}
	out := new(HumioRepositoryList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *HumioRepositoryList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HumioRepositorySpec) DeepCopyInto(out *HumioRepositorySpec) {
	*out = *in
	out.Retention = in.Retention
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HumioRepositorySpec.
func (in *HumioRepositorySpec) DeepCopy() *HumioRepositorySpec {
	if in == nil {
		return nil
	}
	out := new(HumioRepositorySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HumioRepositoryStatus) DeepCopyInto(out *HumioRepositoryStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HumioRepositoryStatus.
func (in *HumioRepositoryStatus) DeepCopy() *HumioRepositoryStatus {
	if in == nil {
		return nil
	}
	out := new(HumioRepositoryStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HumioRetention) DeepCopyInto(out *HumioRetention) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HumioRetention.
func (in *HumioRetention) DeepCopy() *HumioRetention {
	if in == nil {
		return nil
	}
	out := new(HumioRetention)
	in.DeepCopyInto(out)
	return out
}
