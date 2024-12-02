//go:build !ignore_autogenerated

/*
Copyright 2024 Feast Community.

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
	"k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AuthzConfig) DeepCopyInto(out *AuthzConfig) {
	*out = *in
	if in.KubernetesAuthz != nil {
		in, out := &in.KubernetesAuthz, &out.KubernetesAuthz
		*out = new(KubernetesAuthz)
		(*in).DeepCopyInto(*out)
	}
<<<<<<< HEAD
	if in.OidcAuthz != nil {
		in, out := &in.OidcAuthz, &out.OidcAuthz
		*out = new(OidcAuthz)
		**out = **in
	}
=======
>>>>>>> 39eb4d80c (feat: RBAC Authorization in Feast Operator (#4786))
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AuthzConfig.
func (in *AuthzConfig) DeepCopy() *AuthzConfig {
	if in == nil {
		return nil
	}
	out := new(AuthzConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DefaultConfigs) DeepCopyInto(out *DefaultConfigs) {
	*out = *in
	if in.Image != nil {
		in, out := &in.Image, &out.Image
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DefaultConfigs.
func (in *DefaultConfigs) DeepCopy() *DefaultConfigs {
	if in == nil {
		return nil
	}
	out := new(DefaultConfigs)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FeatureStore) DeepCopyInto(out *FeatureStore) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FeatureStore.
func (in *FeatureStore) DeepCopy() *FeatureStore {
	if in == nil {
		return nil
	}
	out := new(FeatureStore)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *FeatureStore) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FeatureStoreList) DeepCopyInto(out *FeatureStoreList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]FeatureStore, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FeatureStoreList.
func (in *FeatureStoreList) DeepCopy() *FeatureStoreList {
	if in == nil {
		return nil
	}
	out := new(FeatureStoreList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *FeatureStoreList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FeatureStoreRef) DeepCopyInto(out *FeatureStoreRef) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FeatureStoreRef.
func (in *FeatureStoreRef) DeepCopy() *FeatureStoreRef {
	if in == nil {
		return nil
	}
	out := new(FeatureStoreRef)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FeatureStoreServices) DeepCopyInto(out *FeatureStoreServices) {
	*out = *in
	if in.OfflineStore != nil {
		in, out := &in.OfflineStore, &out.OfflineStore
		*out = new(OfflineStore)
		(*in).DeepCopyInto(*out)
	}
	if in.OnlineStore != nil {
		in, out := &in.OnlineStore, &out.OnlineStore
		*out = new(OnlineStore)
		(*in).DeepCopyInto(*out)
	}
	if in.Registry != nil {
		in, out := &in.Registry, &out.Registry
		*out = new(Registry)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FeatureStoreServices.
func (in *FeatureStoreServices) DeepCopy() *FeatureStoreServices {
	if in == nil {
		return nil
	}
	out := new(FeatureStoreServices)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FeatureStoreSpec) DeepCopyInto(out *FeatureStoreSpec) {
	*out = *in
	if in.Services != nil {
		in, out := &in.Services, &out.Services
		*out = new(FeatureStoreServices)
		(*in).DeepCopyInto(*out)
	}
	if in.AuthzConfig != nil {
		in, out := &in.AuthzConfig, &out.AuthzConfig
		*out = new(AuthzConfig)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FeatureStoreSpec.
func (in *FeatureStoreSpec) DeepCopy() *FeatureStoreSpec {
	if in == nil {
		return nil
	}
	out := new(FeatureStoreSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FeatureStoreStatus) DeepCopyInto(out *FeatureStoreStatus) {
	*out = *in
	in.Applied.DeepCopyInto(&out.Applied)
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]metav1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	out.ServiceHostnames = in.ServiceHostnames
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FeatureStoreStatus.
func (in *FeatureStoreStatus) DeepCopy() *FeatureStoreStatus {
	if in == nil {
		return nil
	}
	out := new(FeatureStoreStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubernetesAuthz) DeepCopyInto(out *KubernetesAuthz) {
	*out = *in
	if in.Roles != nil {
		in, out := &in.Roles, &out.Roles
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubernetesAuthz.
func (in *KubernetesAuthz) DeepCopy() *KubernetesAuthz {
	if in == nil {
		return nil
	}
	out := new(KubernetesAuthz)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LocalRegistryConfig) DeepCopyInto(out *LocalRegistryConfig) {
	*out = *in
	in.ServiceConfigs.DeepCopyInto(&out.ServiceConfigs)
	if in.Persistence != nil {
		in, out := &in.Persistence, &out.Persistence
		*out = new(RegistryPersistence)
		(*in).DeepCopyInto(*out)
	}
	if in.TLS != nil {
		in, out := &in.TLS, &out.TLS
		*out = new(TlsConfigs)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LocalRegistryConfig.
func (in *LocalRegistryConfig) DeepCopy() *LocalRegistryConfig {
	if in == nil {
		return nil
	}
	out := new(LocalRegistryConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OfflineStore) DeepCopyInto(out *OfflineStore) {
	*out = *in
	in.ServiceConfigs.DeepCopyInto(&out.ServiceConfigs)
	if in.Persistence != nil {
		in, out := &in.Persistence, &out.Persistence
		*out = new(OfflineStorePersistence)
		(*in).DeepCopyInto(*out)
	}
	if in.TLS != nil {
		in, out := &in.TLS, &out.TLS
<<<<<<< HEAD
		*out = new(TlsConfigs)
=======
		*out = new(OfflineTlsConfigs)
>>>>>>> 668d47b8e (feat: Add TLS support to the Operator (#4796))
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OfflineStore.
func (in *OfflineStore) DeepCopy() *OfflineStore {
	if in == nil {
		return nil
	}
	out := new(OfflineStore)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OfflineStoreDBStorePersistence) DeepCopyInto(out *OfflineStoreDBStorePersistence) {
	*out = *in
<<<<<<< HEAD
<<<<<<< HEAD
	out.SecretRef = in.SecretRef
=======
	if in.SecretRef != nil {
		in, out := &in.SecretRef, &out.SecretRef
		*out = new(v1.LocalObjectReference)
		**out = **in
	}
>>>>>>> 863a82cb7 (feat: Added feast Go operator db stores support (#4771))
=======
	out.SecretRef = in.SecretRef
>>>>>>> cac619cfa (fix: Fix db store types in Operator CRD (#4798))
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OfflineStoreDBStorePersistence.
func (in *OfflineStoreDBStorePersistence) DeepCopy() *OfflineStoreDBStorePersistence {
	if in == nil {
		return nil
	}
	out := new(OfflineStoreDBStorePersistence)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OfflineStoreFilePersistence) DeepCopyInto(out *OfflineStoreFilePersistence) {
	*out = *in
	if in.PvcConfig != nil {
		in, out := &in.PvcConfig, &out.PvcConfig
		*out = new(PvcConfig)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OfflineStoreFilePersistence.
func (in *OfflineStoreFilePersistence) DeepCopy() *OfflineStoreFilePersistence {
	if in == nil {
		return nil
	}
	out := new(OfflineStoreFilePersistence)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OfflineStorePersistence) DeepCopyInto(out *OfflineStorePersistence) {
	*out = *in
	if in.FilePersistence != nil {
		in, out := &in.FilePersistence, &out.FilePersistence
		*out = new(OfflineStoreFilePersistence)
		(*in).DeepCopyInto(*out)
<<<<<<< HEAD
	}
	if in.DBPersistence != nil {
		in, out := &in.DBPersistence, &out.DBPersistence
		*out = new(OfflineStoreDBStorePersistence)
		**out = **in
=======
>>>>>>> 6c1a66ea8 (feat: PVC configuration and impl (#4750))
	}
	if in.DBPersistence != nil {
		in, out := &in.DBPersistence, &out.DBPersistence
		*out = new(OfflineStoreDBStorePersistence)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OfflineStorePersistence.
func (in *OfflineStorePersistence) DeepCopy() *OfflineStorePersistence {
	if in == nil {
		return nil
	}
	out := new(OfflineStorePersistence)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
<<<<<<< HEAD
func (in *OidcAuthz) DeepCopyInto(out *OidcAuthz) {
	*out = *in
	out.SecretRef = in.SecretRef
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OidcAuthz.
func (in *OidcAuthz) DeepCopy() *OidcAuthz {
	if in == nil {
		return nil
	}
	out := new(OidcAuthz)
=======
func (in *OfflineTlsConfigs) DeepCopyInto(out *OfflineTlsConfigs) {
	*out = *in
	in.TlsConfigs.DeepCopyInto(&out.TlsConfigs)
	if in.VerifyClient != nil {
		in, out := &in.VerifyClient, &out.VerifyClient
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OfflineTlsConfigs.
func (in *OfflineTlsConfigs) DeepCopy() *OfflineTlsConfigs {
	if in == nil {
		return nil
	}
	out := new(OfflineTlsConfigs)
>>>>>>> 668d47b8e (feat: Add TLS support to the Operator (#4796))
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OnlineStore) DeepCopyInto(out *OnlineStore) {
	*out = *in
	in.ServiceConfigs.DeepCopyInto(&out.ServiceConfigs)
	if in.Persistence != nil {
		in, out := &in.Persistence, &out.Persistence
		*out = new(OnlineStorePersistence)
		(*in).DeepCopyInto(*out)
	}
	if in.TLS != nil {
		in, out := &in.TLS, &out.TLS
		*out = new(TlsConfigs)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OnlineStore.
func (in *OnlineStore) DeepCopy() *OnlineStore {
	if in == nil {
		return nil
	}
	out := new(OnlineStore)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OnlineStoreDBStorePersistence) DeepCopyInto(out *OnlineStoreDBStorePersistence) {
	*out = *in
<<<<<<< HEAD
<<<<<<< HEAD
	out.SecretRef = in.SecretRef
=======
	if in.SecretRef != nil {
		in, out := &in.SecretRef, &out.SecretRef
		*out = new(v1.LocalObjectReference)
		**out = **in
	}
>>>>>>> 863a82cb7 (feat: Added feast Go operator db stores support (#4771))
=======
	out.SecretRef = in.SecretRef
>>>>>>> cac619cfa (fix: Fix db store types in Operator CRD (#4798))
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OnlineStoreDBStorePersistence.
func (in *OnlineStoreDBStorePersistence) DeepCopy() *OnlineStoreDBStorePersistence {
	if in == nil {
		return nil
	}
	out := new(OnlineStoreDBStorePersistence)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OnlineStoreFilePersistence) DeepCopyInto(out *OnlineStoreFilePersistence) {
	*out = *in
	if in.PvcConfig != nil {
		in, out := &in.PvcConfig, &out.PvcConfig
		*out = new(PvcConfig)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OnlineStoreFilePersistence.
func (in *OnlineStoreFilePersistence) DeepCopy() *OnlineStoreFilePersistence {
	if in == nil {
		return nil
	}
	out := new(OnlineStoreFilePersistence)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OnlineStorePersistence) DeepCopyInto(out *OnlineStorePersistence) {
	*out = *in
	if in.FilePersistence != nil {
		in, out := &in.FilePersistence, &out.FilePersistence
		*out = new(OnlineStoreFilePersistence)
		(*in).DeepCopyInto(*out)
<<<<<<< HEAD
	}
	if in.DBPersistence != nil {
		in, out := &in.DBPersistence, &out.DBPersistence
		*out = new(OnlineStoreDBStorePersistence)
		**out = **in
=======
>>>>>>> 6c1a66ea8 (feat: PVC configuration and impl (#4750))
	}
	if in.DBPersistence != nil {
		in, out := &in.DBPersistence, &out.DBPersistence
		*out = new(OnlineStoreDBStorePersistence)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OnlineStorePersistence.
func (in *OnlineStorePersistence) DeepCopy() *OnlineStorePersistence {
	if in == nil {
		return nil
	}
	out := new(OnlineStorePersistence)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OptionalConfigs) DeepCopyInto(out *OptionalConfigs) {
	*out = *in
	if in.Env != nil {
		in, out := &in.Env, &out.Env
		*out = new([]v1.EnvVar)
		if **in != nil {
			in, out := *in, *out
			*out = make([]v1.EnvVar, len(*in))
			for i := range *in {
				(*in)[i].DeepCopyInto(&(*out)[i])
			}
		}
	}
	if in.EnvFrom != nil {
		in, out := &in.EnvFrom, &out.EnvFrom
		*out = new([]v1.EnvFromSource)
		if **in != nil {
			in, out := *in, *out
			*out = make([]v1.EnvFromSource, len(*in))
			for i := range *in {
				(*in)[i].DeepCopyInto(&(*out)[i])
			}
		}
	}
	if in.ImagePullPolicy != nil {
		in, out := &in.ImagePullPolicy, &out.ImagePullPolicy
		*out = new(v1.PullPolicy)
		**out = **in
	}
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = new(v1.ResourceRequirements)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OptionalConfigs.
func (in *OptionalConfigs) DeepCopy() *OptionalConfigs {
	if in == nil {
		return nil
	}
	out := new(OptionalConfigs)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PvcConfig) DeepCopyInto(out *PvcConfig) {
	*out = *in
	if in.Ref != nil {
		in, out := &in.Ref, &out.Ref
		*out = new(v1.LocalObjectReference)
		**out = **in
	}
	if in.Create != nil {
		in, out := &in.Create, &out.Create
		*out = new(PvcCreate)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PvcConfig.
func (in *PvcConfig) DeepCopy() *PvcConfig {
	if in == nil {
		return nil
	}
	out := new(PvcConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PvcCreate) DeepCopyInto(out *PvcCreate) {
	*out = *in
<<<<<<< HEAD
	if in.AccessModes != nil {
		in, out := &in.AccessModes, &out.AccessModes
		*out = make([]v1.PersistentVolumeAccessMode, len(*in))
		copy(*out, *in)
	}
=======
>>>>>>> 6c1a66ea8 (feat: PVC configuration and impl (#4750))
	if in.StorageClassName != nil {
		in, out := &in.StorageClassName, &out.StorageClassName
		*out = new(string)
		**out = **in
	}
	in.Resources.DeepCopyInto(&out.Resources)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PvcCreate.
func (in *PvcCreate) DeepCopy() *PvcCreate {
	if in == nil {
		return nil
	}
	out := new(PvcCreate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Registry) DeepCopyInto(out *Registry) {
	*out = *in
	if in.Local != nil {
		in, out := &in.Local, &out.Local
		*out = new(LocalRegistryConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.Remote != nil {
		in, out := &in.Remote, &out.Remote
		*out = new(RemoteRegistryConfig)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Registry.
func (in *Registry) DeepCopy() *Registry {
	if in == nil {
		return nil
	}
	out := new(Registry)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RegistryDBStorePersistence) DeepCopyInto(out *RegistryDBStorePersistence) {
	*out = *in
<<<<<<< HEAD
<<<<<<< HEAD
	out.SecretRef = in.SecretRef
=======
	if in.SecretRef != nil {
		in, out := &in.SecretRef, &out.SecretRef
		*out = new(v1.LocalObjectReference)
		**out = **in
	}
>>>>>>> 863a82cb7 (feat: Added feast Go operator db stores support (#4771))
=======
	out.SecretRef = in.SecretRef
>>>>>>> cac619cfa (fix: Fix db store types in Operator CRD (#4798))
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RegistryDBStorePersistence.
func (in *RegistryDBStorePersistence) DeepCopy() *RegistryDBStorePersistence {
	if in == nil {
		return nil
	}
	out := new(RegistryDBStorePersistence)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RegistryFilePersistence) DeepCopyInto(out *RegistryFilePersistence) {
	*out = *in
	if in.PvcConfig != nil {
		in, out := &in.PvcConfig, &out.PvcConfig
		*out = new(PvcConfig)
		(*in).DeepCopyInto(*out)
	}
<<<<<<< HEAD
<<<<<<< HEAD
=======
>>>>>>> bc64ddfac (feat: Object store persistence in operator (#4758))
	if in.S3AdditionalKwargs != nil {
		in, out := &in.S3AdditionalKwargs, &out.S3AdditionalKwargs
		*out = new(map[string]string)
		if **in != nil {
			in, out := *in, *out
			*out = make(map[string]string, len(*in))
			for key, val := range *in {
				(*out)[key] = val
			}
		}
	}
<<<<<<< HEAD
=======
>>>>>>> 6c1a66ea8 (feat: PVC configuration and impl (#4750))
=======
>>>>>>> bc64ddfac (feat: Object store persistence in operator (#4758))
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RegistryFilePersistence.
func (in *RegistryFilePersistence) DeepCopy() *RegistryFilePersistence {
	if in == nil {
		return nil
	}
	out := new(RegistryFilePersistence)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RegistryPersistence) DeepCopyInto(out *RegistryPersistence) {
	*out = *in
	if in.FilePersistence != nil {
		in, out := &in.FilePersistence, &out.FilePersistence
		*out = new(RegistryFilePersistence)
		(*in).DeepCopyInto(*out)
<<<<<<< HEAD
	}
	if in.DBPersistence != nil {
		in, out := &in.DBPersistence, &out.DBPersistence
		*out = new(RegistryDBStorePersistence)
		**out = **in
=======
>>>>>>> 6c1a66ea8 (feat: PVC configuration and impl (#4750))
	}
	if in.DBPersistence != nil {
		in, out := &in.DBPersistence, &out.DBPersistence
		*out = new(RegistryDBStorePersistence)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RegistryPersistence.
func (in *RegistryPersistence) DeepCopy() *RegistryPersistence {
	if in == nil {
		return nil
	}
	out := new(RegistryPersistence)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RemoteRegistryConfig) DeepCopyInto(out *RemoteRegistryConfig) {
	*out = *in
	if in.Hostname != nil {
		in, out := &in.Hostname, &out.Hostname
		*out = new(string)
		**out = **in
	}
	if in.FeastRef != nil {
		in, out := &in.FeastRef, &out.FeastRef
		*out = new(FeatureStoreRef)
		**out = **in
	}
	if in.TLS != nil {
		in, out := &in.TLS, &out.TLS
		*out = new(TlsRemoteRegistryConfigs)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RemoteRegistryConfig.
func (in *RemoteRegistryConfig) DeepCopy() *RemoteRegistryConfig {
	if in == nil {
		return nil
	}
	out := new(RemoteRegistryConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecretKeyNames) DeepCopyInto(out *SecretKeyNames) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecretKeyNames.
func (in *SecretKeyNames) DeepCopy() *SecretKeyNames {
	if in == nil {
		return nil
	}
	out := new(SecretKeyNames)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceConfigs) DeepCopyInto(out *ServiceConfigs) {
	*out = *in
	in.DefaultConfigs.DeepCopyInto(&out.DefaultConfigs)
	in.OptionalConfigs.DeepCopyInto(&out.OptionalConfigs)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceConfigs.
func (in *ServiceConfigs) DeepCopy() *ServiceConfigs {
	if in == nil {
		return nil
	}
	out := new(ServiceConfigs)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceHostnames) DeepCopyInto(out *ServiceHostnames) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceHostnames.
func (in *ServiceHostnames) DeepCopy() *ServiceHostnames {
	if in == nil {
		return nil
	}
	out := new(ServiceHostnames)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TlsConfigs) DeepCopyInto(out *TlsConfigs) {
	*out = *in
	if in.SecretRef != nil {
		in, out := &in.SecretRef, &out.SecretRef
		*out = new(v1.LocalObjectReference)
		**out = **in
	}
	out.SecretKeyNames = in.SecretKeyNames
	if in.Disable != nil {
		in, out := &in.Disable, &out.Disable
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TlsConfigs.
func (in *TlsConfigs) DeepCopy() *TlsConfigs {
	if in == nil {
		return nil
	}
	out := new(TlsConfigs)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TlsRemoteRegistryConfigs) DeepCopyInto(out *TlsRemoteRegistryConfigs) {
	*out = *in
	out.ConfigMapRef = in.ConfigMapRef
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TlsRemoteRegistryConfigs.
func (in *TlsRemoteRegistryConfigs) DeepCopy() *TlsRemoteRegistryConfigs {
	if in == nil {
		return nil
	}
	out := new(TlsRemoteRegistryConfigs)
	in.DeepCopyInto(out)
	return out
}
