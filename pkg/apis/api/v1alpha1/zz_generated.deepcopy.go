// +build !ignore_autogenerated

// Code generated by operator-sdk. DO NOT EDIT.

package v1alpha1

import (
	v1beta1 "github.com/astarte-platform/astarte-kubernetes-operator/external/voyager/v1beta1"
	appsv1 "k8s.io/api/apps/v1"
	v1 "k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Astarte) DeepCopyInto(out *Astarte) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Astarte.
func (in *Astarte) DeepCopy() *Astarte {
	if in == nil {
		return nil
	}
	out := new(Astarte)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Astarte) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AstarteAPISpec) DeepCopyInto(out *AstarteAPISpec) {
	*out = *in
	if in.SSL != nil {
		in, out := &in.SSL, &out.SSL
		*out = new(bool)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AstarteAPISpec.
func (in *AstarteAPISpec) DeepCopy() *AstarteAPISpec {
	if in == nil {
		return nil
	}
	out := new(AstarteAPISpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AstarteAppengineAPISpec) DeepCopyInto(out *AstarteAppengineAPISpec) {
	*out = *in
	in.AstarteGenericAPISpec.DeepCopyInto(&out.AstarteGenericAPISpec)
	if in.MaxResultsLimit != nil {
		in, out := &in.MaxResultsLimit, &out.MaxResultsLimit
		*out = new(int)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AstarteAppengineAPISpec.
func (in *AstarteAppengineAPISpec) DeepCopy() *AstarteAppengineAPISpec {
	if in == nil {
		return nil
	}
	out := new(AstarteAppengineAPISpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AstarteCFSSLCARootConfigSigningCAConstraintSpec) DeepCopyInto(out *AstarteCFSSLCARootConfigSigningCAConstraintSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AstarteCFSSLCARootConfigSigningCAConstraintSpec.
func (in *AstarteCFSSLCARootConfigSigningCAConstraintSpec) DeepCopy() *AstarteCFSSLCARootConfigSigningCAConstraintSpec {
	if in == nil {
		return nil
	}
	out := new(AstarteCFSSLCARootConfigSigningCAConstraintSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AstarteCFSSLCARootConfigSigningDefaultSpec) DeepCopyInto(out *AstarteCFSSLCARootConfigSigningDefaultSpec) {
	*out = *in
	if in.Usages != nil {
		in, out := &in.Usages, &out.Usages
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.CAConstraint != nil {
		in, out := &in.CAConstraint, &out.CAConstraint
		*out = new(AstarteCFSSLCARootConfigSigningCAConstraintSpec)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AstarteCFSSLCARootConfigSigningDefaultSpec.
func (in *AstarteCFSSLCARootConfigSigningDefaultSpec) DeepCopy() *AstarteCFSSLCARootConfigSigningDefaultSpec {
	if in == nil {
		return nil
	}
	out := new(AstarteCFSSLCARootConfigSigningDefaultSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AstarteCFSSLCARootConfigSigningSpec) DeepCopyInto(out *AstarteCFSSLCARootConfigSigningSpec) {
	*out = *in
	if in.Default != nil {
		in, out := &in.Default, &out.Default
		*out = new(AstarteCFSSLCARootConfigSigningDefaultSpec)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AstarteCFSSLCARootConfigSigningSpec.
func (in *AstarteCFSSLCARootConfigSigningSpec) DeepCopy() *AstarteCFSSLCARootConfigSigningSpec {
	if in == nil {
		return nil
	}
	out := new(AstarteCFSSLCARootConfigSigningSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AstarteCFSSLCARootConfigSpec) DeepCopyInto(out *AstarteCFSSLCARootConfigSpec) {
	*out = *in
	if in.Signing != nil {
		in, out := &in.Signing, &out.Signing
		*out = new(AstarteCFSSLCARootConfigSigningSpec)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AstarteCFSSLCARootConfigSpec.
func (in *AstarteCFSSLCARootConfigSpec) DeepCopy() *AstarteCFSSLCARootConfigSpec {
	if in == nil {
		return nil
	}
	out := new(AstarteCFSSLCARootConfigSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AstarteCFSSLCSRRootCACASpec) DeepCopyInto(out *AstarteCFSSLCSRRootCACASpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AstarteCFSSLCSRRootCACASpec.
func (in *AstarteCFSSLCSRRootCACASpec) DeepCopy() *AstarteCFSSLCSRRootCACASpec {
	if in == nil {
		return nil
	}
	out := new(AstarteCFSSLCSRRootCACASpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AstarteCFSSLCSRRootCAKeySpec) DeepCopyInto(out *AstarteCFSSLCSRRootCAKeySpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AstarteCFSSLCSRRootCAKeySpec.
func (in *AstarteCFSSLCSRRootCAKeySpec) DeepCopy() *AstarteCFSSLCSRRootCAKeySpec {
	if in == nil {
		return nil
	}
	out := new(AstarteCFSSLCSRRootCAKeySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AstarteCFSSLCSRRootCANamesSpec) DeepCopyInto(out *AstarteCFSSLCSRRootCANamesSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AstarteCFSSLCSRRootCANamesSpec.
func (in *AstarteCFSSLCSRRootCANamesSpec) DeepCopy() *AstarteCFSSLCSRRootCANamesSpec {
	if in == nil {
		return nil
	}
	out := new(AstarteCFSSLCSRRootCANamesSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AstarteCFSSLCSRRootCASpec) DeepCopyInto(out *AstarteCFSSLCSRRootCASpec) {
	*out = *in
	if in.Key != nil {
		in, out := &in.Key, &out.Key
		*out = new(AstarteCFSSLCSRRootCAKeySpec)
		**out = **in
	}
	if in.Names != nil {
		in, out := &in.Names, &out.Names
		*out = make([]AstarteCFSSLCSRRootCANamesSpec, len(*in))
		copy(*out, *in)
	}
	if in.CA != nil {
		in, out := &in.CA, &out.CA
		*out = new(AstarteCFSSLCSRRootCACASpec)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AstarteCFSSLCSRRootCASpec.
func (in *AstarteCFSSLCSRRootCASpec) DeepCopy() *AstarteCFSSLCSRRootCASpec {
	if in == nil {
		return nil
	}
	out := new(AstarteCFSSLCSRRootCASpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AstarteCFSSLDBConfigSpec) DeepCopyInto(out *AstarteCFSSLDBConfigSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AstarteCFSSLDBConfigSpec.
func (in *AstarteCFSSLDBConfigSpec) DeepCopy() *AstarteCFSSLDBConfigSpec {
	if in == nil {
		return nil
	}
	out := new(AstarteCFSSLDBConfigSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AstarteCFSSLSpec) DeepCopyInto(out *AstarteCFSSLSpec) {
	*out = *in
	if in.Deploy != nil {
		in, out := &in.Deploy, &out.Deploy
		*out = new(bool)
		**out = **in
	}
	if in.DBConfig != nil {
		in, out := &in.DBConfig, &out.DBConfig
		*out = new(AstarteCFSSLDBConfigSpec)
		**out = **in
	}
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = new(v1.ResourceRequirements)
		(*in).DeepCopyInto(*out)
	}
	if in.Storage != nil {
		in, out := &in.Storage, &out.Storage
		*out = new(AstartePersistentStorageSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.CSRRootCa != nil {
		in, out := &in.CSRRootCa, &out.CSRRootCa
		*out = new(AstarteCFSSLCSRRootCASpec)
		(*in).DeepCopyInto(*out)
	}
	if in.CARootConfig != nil {
		in, out := &in.CARootConfig, &out.CARootConfig
		*out = new(AstarteCFSSLCARootConfigSpec)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AstarteCFSSLSpec.
func (in *AstarteCFSSLSpec) DeepCopy() *AstarteCFSSLSpec {
	if in == nil {
		return nil
	}
	out := new(AstarteCFSSLSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AstarteCassandraSpec) DeepCopyInto(out *AstarteCassandraSpec) {
	*out = *in
	in.AstarteGenericClusteredResource.DeepCopyInto(&out.AstarteGenericClusteredResource)
	if in.Storage != nil {
		in, out := &in.Storage, &out.Storage
		*out = new(AstartePersistentStorageSpec)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AstarteCassandraSpec.
func (in *AstarteCassandraSpec) DeepCopy() *AstarteCassandraSpec {
	if in == nil {
		return nil
	}
	out := new(AstarteCassandraSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AstarteComponentsSpec) DeepCopyInto(out *AstarteComponentsSpec) {
	*out = *in
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = new(v1.ResourceRequirements)
		(*in).DeepCopyInto(*out)
	}
	in.Housekeeping.DeepCopyInto(&out.Housekeeping)
	in.RealmManagement.DeepCopyInto(&out.RealmManagement)
	in.Pairing.DeepCopyInto(&out.Pairing)
	in.DataUpdaterPlant.DeepCopyInto(&out.DataUpdaterPlant)
	in.AppengineAPI.DeepCopyInto(&out.AppengineAPI)
	in.TriggerEngine.DeepCopyInto(&out.TriggerEngine)
	in.Dashboard.DeepCopyInto(&out.Dashboard)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AstarteComponentsSpec.
func (in *AstarteComponentsSpec) DeepCopy() *AstarteComponentsSpec {
	if in == nil {
		return nil
	}
	out := new(AstarteComponentsSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AstarteDashboardConfigAuthSpec) DeepCopyInto(out *AstarteDashboardConfigAuthSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AstarteDashboardConfigAuthSpec.
func (in *AstarteDashboardConfigAuthSpec) DeepCopy() *AstarteDashboardConfigAuthSpec {
	if in == nil {
		return nil
	}
	out := new(AstarteDashboardConfigAuthSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AstarteDashboardConfigSpec) DeepCopyInto(out *AstarteDashboardConfigSpec) {
	*out = *in
	if in.Auth != nil {
		in, out := &in.Auth, &out.Auth
		*out = make([]AstarteDashboardConfigAuthSpec, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AstarteDashboardConfigSpec.
func (in *AstarteDashboardConfigSpec) DeepCopy() *AstarteDashboardConfigSpec {
	if in == nil {
		return nil
	}
	out := new(AstarteDashboardConfigSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AstarteDashboardSpec) DeepCopyInto(out *AstarteDashboardSpec) {
	*out = *in
	in.AstarteGenericClusteredResource.DeepCopyInto(&out.AstarteGenericClusteredResource)
	in.Config.DeepCopyInto(&out.Config)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AstarteDashboardSpec.
func (in *AstarteDashboardSpec) DeepCopy() *AstarteDashboardSpec {
	if in == nil {
		return nil
	}
	out := new(AstarteDashboardSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AstarteDataUpdaterPlantSpec) DeepCopyInto(out *AstarteDataUpdaterPlantSpec) {
	*out = *in
	in.AstarteGenericClusteredResource.DeepCopyInto(&out.AstarteGenericClusteredResource)
	if in.DataQueueCount != nil {
		in, out := &in.DataQueueCount, &out.DataQueueCount
		*out = new(int)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AstarteDataUpdaterPlantSpec.
func (in *AstarteDataUpdaterPlantSpec) DeepCopy() *AstarteDataUpdaterPlantSpec {
	if in == nil {
		return nil
	}
	out := new(AstarteDataUpdaterPlantSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AstarteGenericAPISpec) DeepCopyInto(out *AstarteGenericAPISpec) {
	*out = *in
	in.AstarteGenericClusteredResource.DeepCopyInto(&out.AstarteGenericClusteredResource)
	if in.DisableAuthentication != nil {
		in, out := &in.DisableAuthentication, &out.DisableAuthentication
		*out = new(bool)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AstarteGenericAPISpec.
func (in *AstarteGenericAPISpec) DeepCopy() *AstarteGenericAPISpec {
	if in == nil {
		return nil
	}
	out := new(AstarteGenericAPISpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AstarteGenericClusteredResource) DeepCopyInto(out *AstarteGenericClusteredResource) {
	*out = *in
	if in.Deploy != nil {
		in, out := &in.Deploy, &out.Deploy
		*out = new(bool)
		**out = **in
	}
	if in.Replicas != nil {
		in, out := &in.Replicas, &out.Replicas
		*out = new(int32)
		**out = **in
	}
	if in.AntiAffinity != nil {
		in, out := &in.AntiAffinity, &out.AntiAffinity
		*out = new(bool)
		**out = **in
	}
	if in.CustomAffinity != nil {
		in, out := &in.CustomAffinity, &out.CustomAffinity
		*out = new(v1.Affinity)
		(*in).DeepCopyInto(*out)
	}
	if in.DeploymentStrategy != nil {
		in, out := &in.DeploymentStrategy, &out.DeploymentStrategy
		*out = new(appsv1.DeploymentStrategy)
		(*in).DeepCopyInto(*out)
	}
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = new(v1.ResourceRequirements)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AstarteGenericClusteredResource.
func (in *AstarteGenericClusteredResource) DeepCopy() *AstarteGenericClusteredResource {
	if in == nil {
		return nil
	}
	out := new(AstarteGenericClusteredResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AstarteGenericComponentSpec) DeepCopyInto(out *AstarteGenericComponentSpec) {
	*out = *in
	in.API.DeepCopyInto(&out.API)
	in.Backend.DeepCopyInto(&out.Backend)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AstarteGenericComponentSpec.
func (in *AstarteGenericComponentSpec) DeepCopy() *AstarteGenericComponentSpec {
	if in == nil {
		return nil
	}
	out := new(AstarteGenericComponentSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AstarteGenericIngressSpec) DeepCopyInto(out *AstarteGenericIngressSpec) {
	*out = *in
	if in.Deploy != nil {
		in, out := &in.Deploy, &out.Deploy
		*out = new(bool)
		**out = **in
	}
	if in.Replicas != nil {
		in, out := &in.Replicas, &out.Replicas
		*out = new(int32)
		**out = **in
	}
	if in.TLSRef != nil {
		in, out := &in.TLSRef, &out.TLSRef
		*out = new(v1beta1.LocalTypedReference)
		**out = **in
	}
	if in.AnnotationsService != nil {
		in, out := &in.AnnotationsService, &out.AnnotationsService
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AstarteGenericIngressSpec.
func (in *AstarteGenericIngressSpec) DeepCopy() *AstarteGenericIngressSpec {
	if in == nil {
		return nil
	}
	out := new(AstarteGenericIngressSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AstarteList) DeepCopyInto(out *AstarteList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Astarte, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AstarteList.
func (in *AstarteList) DeepCopy() *AstarteList {
	if in == nil {
		return nil
	}
	out := new(AstarteList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AstarteList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AstartePersistentStorageSpec) DeepCopyInto(out *AstartePersistentStorageSpec) {
	*out = *in
	if in.Size != nil {
		in, out := &in.Size, &out.Size
		x := (*in).DeepCopy()
		*out = &x
	}
	if in.VolumeDefinition != nil {
		in, out := &in.VolumeDefinition, &out.VolumeDefinition
		*out = new(v1.Volume)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AstartePersistentStorageSpec.
func (in *AstartePersistentStorageSpec) DeepCopy() *AstartePersistentStorageSpec {
	if in == nil {
		return nil
	}
	out := new(AstartePersistentStorageSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AstarteRabbitMQConnectionSecretSpec) DeepCopyInto(out *AstarteRabbitMQConnectionSecretSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AstarteRabbitMQConnectionSecretSpec.
func (in *AstarteRabbitMQConnectionSecretSpec) DeepCopy() *AstarteRabbitMQConnectionSecretSpec {
	if in == nil {
		return nil
	}
	out := new(AstarteRabbitMQConnectionSecretSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AstarteRabbitMQConnectionSpec) DeepCopyInto(out *AstarteRabbitMQConnectionSpec) {
	*out = *in
	if in.Port != nil {
		in, out := &in.Port, &out.Port
		*out = new(int16)
		**out = **in
	}
	if in.Secret != nil {
		in, out := &in.Secret, &out.Secret
		*out = new(AstarteRabbitMQConnectionSecretSpec)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AstarteRabbitMQConnectionSpec.
func (in *AstarteRabbitMQConnectionSpec) DeepCopy() *AstarteRabbitMQConnectionSpec {
	if in == nil {
		return nil
	}
	out := new(AstarteRabbitMQConnectionSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AstarteRabbitMQSpec) DeepCopyInto(out *AstarteRabbitMQSpec) {
	*out = *in
	in.AstarteGenericClusteredResource.DeepCopyInto(&out.AstarteGenericClusteredResource)
	if in.Connection != nil {
		in, out := &in.Connection, &out.Connection
		*out = new(AstarteRabbitMQConnectionSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.Storage != nil {
		in, out := &in.Storage, &out.Storage
		*out = new(AstartePersistentStorageSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.AdditionalPlugins != nil {
		in, out := &in.AdditionalPlugins, &out.AdditionalPlugins
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AstarteRabbitMQSpec.
func (in *AstarteRabbitMQSpec) DeepCopy() *AstarteRabbitMQSpec {
	if in == nil {
		return nil
	}
	out := new(AstarteRabbitMQSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AstarteSpec) DeepCopyInto(out *AstarteSpec) {
	*out = *in
	if in.ImagePullPolicy != nil {
		in, out := &in.ImagePullPolicy, &out.ImagePullPolicy
		*out = new(v1.PullPolicy)
		**out = **in
	}
	if in.ImagePullSecrets != nil {
		in, out := &in.ImagePullSecrets, &out.ImagePullSecrets
		*out = make([]v1.LocalObjectReference, len(*in))
		copy(*out, *in)
	}
	in.DeploymentStrategy.DeepCopyInto(&out.DeploymentStrategy)
	if in.RBAC != nil {
		in, out := &in.RBAC, &out.RBAC
		*out = new(bool)
		**out = **in
	}
	in.API.DeepCopyInto(&out.API)
	in.RabbitMQ.DeepCopyInto(&out.RabbitMQ)
	in.Cassandra.DeepCopyInto(&out.Cassandra)
	in.VerneMQ.DeepCopyInto(&out.VerneMQ)
	in.CFSSL.DeepCopyInto(&out.CFSSL)
	in.Components.DeepCopyInto(&out.Components)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AstarteSpec.
func (in *AstarteSpec) DeepCopy() *AstarteSpec {
	if in == nil {
		return nil
	}
	out := new(AstarteSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AstarteStatus) DeepCopyInto(out *AstarteStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AstarteStatus.
func (in *AstarteStatus) DeepCopy() *AstarteStatus {
	if in == nil {
		return nil
	}
	out := new(AstarteStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AstarteVerneMQSpec) DeepCopyInto(out *AstarteVerneMQSpec) {
	*out = *in
	in.AstarteGenericClusteredResource.DeepCopyInto(&out.AstarteGenericClusteredResource)
	if in.Port != nil {
		in, out := &in.Port, &out.Port
		*out = new(int16)
		**out = **in
	}
	if in.Storage != nil {
		in, out := &in.Storage, &out.Storage
		*out = new(AstartePersistentStorageSpec)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AstarteVerneMQSpec.
func (in *AstarteVerneMQSpec) DeepCopy() *AstarteVerneMQSpec {
	if in == nil {
		return nil
	}
	out := new(AstarteVerneMQSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AstarteVoyagerIngress) DeepCopyInto(out *AstarteVoyagerIngress) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AstarteVoyagerIngress.
func (in *AstarteVoyagerIngress) DeepCopy() *AstarteVoyagerIngress {
	if in == nil {
		return nil
	}
	out := new(AstarteVoyagerIngress)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AstarteVoyagerIngress) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AstarteVoyagerIngressAPISpec) DeepCopyInto(out *AstarteVoyagerIngressAPISpec) {
	*out = *in
	in.AstarteGenericIngressSpec.DeepCopyInto(&out.AstarteGenericIngressSpec)
	if in.Cors != nil {
		in, out := &in.Cors, &out.Cors
		*out = new(bool)
		**out = **in
	}
	if in.ExposeHousekeeping != nil {
		in, out := &in.ExposeHousekeeping, &out.ExposeHousekeeping
		*out = new(bool)
		**out = **in
	}
	if in.ServeMetrics != nil {
		in, out := &in.ServeMetrics, &out.ServeMetrics
		*out = new(bool)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AstarteVoyagerIngressAPISpec.
func (in *AstarteVoyagerIngressAPISpec) DeepCopy() *AstarteVoyagerIngressAPISpec {
	if in == nil {
		return nil
	}
	out := new(AstarteVoyagerIngressAPISpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AstarteVoyagerIngressBrokerSpec) DeepCopyInto(out *AstarteVoyagerIngressBrokerSpec) {
	*out = *in
	in.AstarteGenericIngressSpec.DeepCopyInto(&out.AstarteGenericIngressSpec)
	if in.MaxConnections != nil {
		in, out := &in.MaxConnections, &out.MaxConnections
		*out = new(int)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AstarteVoyagerIngressBrokerSpec.
func (in *AstarteVoyagerIngressBrokerSpec) DeepCopy() *AstarteVoyagerIngressBrokerSpec {
	if in == nil {
		return nil
	}
	out := new(AstarteVoyagerIngressBrokerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AstarteVoyagerIngressDashboardSpec) DeepCopyInto(out *AstarteVoyagerIngressDashboardSpec) {
	*out = *in
	if in.SSL != nil {
		in, out := &in.SSL, &out.SSL
		*out = new(bool)
		**out = **in
	}
	if in.TLSRef != nil {
		in, out := &in.TLSRef, &out.TLSRef
		*out = new(v1beta1.LocalTypedReference)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AstarteVoyagerIngressDashboardSpec.
func (in *AstarteVoyagerIngressDashboardSpec) DeepCopy() *AstarteVoyagerIngressDashboardSpec {
	if in == nil {
		return nil
	}
	out := new(AstarteVoyagerIngressDashboardSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AstarteVoyagerIngressLetsEncryptSpec) DeepCopyInto(out *AstarteVoyagerIngressLetsEncryptSpec) {
	*out = *in
	if in.Use != nil {
		in, out := &in.Use, &out.Use
		*out = new(bool)
		**out = **in
	}
	if in.Staging != nil {
		in, out := &in.Staging, &out.Staging
		*out = new(bool)
		**out = **in
	}
	if in.Domains != nil {
		in, out := &in.Domains, &out.Domains
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.AutoHTTPChallenge != nil {
		in, out := &in.AutoHTTPChallenge, &out.AutoHTTPChallenge
		*out = new(bool)
		**out = **in
	}
	in.ChallengeProvider.DeepCopyInto(&out.ChallengeProvider)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AstarteVoyagerIngressLetsEncryptSpec.
func (in *AstarteVoyagerIngressLetsEncryptSpec) DeepCopy() *AstarteVoyagerIngressLetsEncryptSpec {
	if in == nil {
		return nil
	}
	out := new(AstarteVoyagerIngressLetsEncryptSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AstarteVoyagerIngressList) DeepCopyInto(out *AstarteVoyagerIngressList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]AstarteVoyagerIngress, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AstarteVoyagerIngressList.
func (in *AstarteVoyagerIngressList) DeepCopy() *AstarteVoyagerIngressList {
	if in == nil {
		return nil
	}
	out := new(AstarteVoyagerIngressList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AstarteVoyagerIngressList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AstarteVoyagerIngressSpec) DeepCopyInto(out *AstarteVoyagerIngressSpec) {
	*out = *in
	if in.ImagePullPolicy != nil {
		in, out := &in.ImagePullPolicy, &out.ImagePullPolicy
		*out = new(v1.PullPolicy)
		**out = **in
	}
	in.API.DeepCopyInto(&out.API)
	in.Dashboard.DeepCopyInto(&out.Dashboard)
	in.Broker.DeepCopyInto(&out.Broker)
	in.Letsencrypt.DeepCopyInto(&out.Letsencrypt)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AstarteVoyagerIngressSpec.
func (in *AstarteVoyagerIngressSpec) DeepCopy() *AstarteVoyagerIngressSpec {
	if in == nil {
		return nil
	}
	out := new(AstarteVoyagerIngressSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AstarteVoyagerIngressStatus) DeepCopyInto(out *AstarteVoyagerIngressStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AstarteVoyagerIngressStatus.
func (in *AstarteVoyagerIngressStatus) DeepCopy() *AstarteVoyagerIngressStatus {
	if in == nil {
		return nil
	}
	out := new(AstarteVoyagerIngressStatus)
	in.DeepCopyInto(out)
	return out
}
