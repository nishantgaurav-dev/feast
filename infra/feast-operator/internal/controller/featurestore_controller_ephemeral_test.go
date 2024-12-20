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

package controller

import (
	"context"
	"encoding/base64"
	"fmt"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"gopkg.in/yaml.v3"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	apimeta "k8s.io/apimachinery/pkg/api/meta"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/selection"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/util/intstr"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"

	"github.com/feast-dev/feast/infra/feast-operator/api/feastversion"
	feastdevv1alpha1 "github.com/feast-dev/feast/infra/feast-operator/api/v1alpha1"
<<<<<<< HEAD
<<<<<<< HEAD
	"github.com/feast-dev/feast/infra/feast-operator/internal/controller/handler"
=======
>>>>>>> 6c1a66ea8 (feat: PVC configuration and impl (#4750))
=======
	"github.com/feast-dev/feast/infra/feast-operator/internal/controller/handler"
>>>>>>> 39eb4d80c (feat: RBAC Authorization in Feast Operator (#4786))
	"github.com/feast-dev/feast/infra/feast-operator/internal/controller/services"
)

var _ = Describe("FeatureStore Controller-Ephemeral services", func() {
	Context("When deploying a resource with all ephemeral services", func() {
		const resourceName = "services-ephemeral"
<<<<<<< HEAD
<<<<<<< HEAD
		const offlineType = "duckdb"
=======
>>>>>>> 6c1a66ea8 (feat: PVC configuration and impl (#4750))
=======
		const offlineType = "duckdb"
>>>>>>> 668d47b8e (feat: Add TLS support to the Operator (#4796))
		var pullPolicy = corev1.PullAlways
		var testEnvVarName = "testEnvVarName"
		var testEnvVarValue = "testEnvVarValue"

		ctx := context.Background()

		typeNamespacedName := types.NamespacedName{
			Name:      resourceName,
			Namespace: "default",
		}
		featurestore := &feastdevv1alpha1.FeatureStore{}
		onlineStorePath := "/data/online.db"
		registryPath := "/data/registry.db"
<<<<<<< HEAD
<<<<<<< HEAD

		BeforeEach(func() {
			createEnvFromSecretAndConfigMap()
=======
		offlineType := "duckdb"
=======
>>>>>>> 668d47b8e (feat: Add TLS support to the Operator (#4796))

		BeforeEach(func() {
>>>>>>> 6c1a66ea8 (feat: PVC configuration and impl (#4750))
			By("creating the custom resource for the Kind FeatureStore")
			err := k8sClient.Get(ctx, typeNamespacedName, featurestore)
			if err != nil && errors.IsNotFound(err) {
<<<<<<< HEAD
<<<<<<< HEAD
				resource := createFeatureStoreResource(resourceName, image, pullPolicy, &[]corev1.EnvVar{{Name: testEnvVarName, Value: testEnvVarValue},
<<<<<<< HEAD
					{Name: "fieldRefName", ValueFrom: &corev1.EnvVarSource{FieldRef: &corev1.ObjectFieldSelector{APIVersion: "v1", FieldPath: "metadata.namespace"}}}}, withEnvFrom())
=======
=======
				resource := createFeatureStoreResource(resourceName, image, pullPolicy, replicas, &[]corev1.EnvVar{{Name: testEnvVarName, Value: testEnvVarValue},
>>>>>>> 47204bcaf (feat: Add online/offline replica support (#4812))
=======
				resource := createFeatureStoreResource(resourceName, image, pullPolicy, &[]corev1.EnvVar{{Name: testEnvVarName, Value: testEnvVarValue},
>>>>>>> b0a04af1d (fix: Refactor Operator to deploy all feast services to the same Deployment/Pod (#4863))
					{Name: "fieldRefName", ValueFrom: &corev1.EnvVarSource{FieldRef: &corev1.ObjectFieldSelector{APIVersion: "v1", FieldPath: "metadata.namespace"}}}})
>>>>>>> 6c1a66ea8 (feat: PVC configuration and impl (#4750))
				resource.Spec.Services.OfflineStore.Persistence = &feastdevv1alpha1.OfflineStorePersistence{
					FilePersistence: &feastdevv1alpha1.OfflineStoreFilePersistence{
						Type: offlineType,
					},
				}
				resource.Spec.Services.OnlineStore.Persistence = &feastdevv1alpha1.OnlineStorePersistence{
					FilePersistence: &feastdevv1alpha1.OnlineStoreFilePersistence{
						Path: onlineStorePath,
					},
				}
				resource.Spec.Services.Registry = &feastdevv1alpha1.Registry{
					Local: &feastdevv1alpha1.LocalRegistryConfig{
						Persistence: &feastdevv1alpha1.RegistryPersistence{
							FilePersistence: &feastdevv1alpha1.RegistryFilePersistence{
								Path: registryPath,
							},
						},
					},
				}

				Expect(k8sClient.Create(ctx, resource)).To(Succeed())
			}
		})
		AfterEach(func() {
			resource := &feastdevv1alpha1.FeatureStore{}
			err := k8sClient.Get(ctx, typeNamespacedName, resource)
			Expect(err).NotTo(HaveOccurred())

			By("Cleanup the specific resource instance FeatureStore")
			Expect(k8sClient.Delete(ctx, resource)).To(Succeed())
<<<<<<< HEAD

			deleteEnvFromSecretAndConfigMap()
=======
>>>>>>> 6c1a66ea8 (feat: PVC configuration and impl (#4750))
		})

		It("should successfully reconcile the resource", func() {
			By("Reconciling the created resource")
			controllerReconciler := &FeatureStoreReconciler{
				Client: k8sClient,
				Scheme: k8sClient.Scheme(),
			}

			_, err := controllerReconciler.Reconcile(ctx, reconcile.Request{
				NamespacedName: typeNamespacedName,
			})
			Expect(err).NotTo(HaveOccurred())

			resource := &feastdevv1alpha1.FeatureStore{}
			err = k8sClient.Get(ctx, typeNamespacedName, resource)
			Expect(err).NotTo(HaveOccurred())

			feast := services.FeastServices{
<<<<<<< HEAD
<<<<<<< HEAD
=======
>>>>>>> 39eb4d80c (feat: RBAC Authorization in Feast Operator (#4786))
				Handler: handler.FeastHandler{
					Client:       controllerReconciler.Client,
					Context:      ctx,
					Scheme:       controllerReconciler.Scheme,
					FeatureStore: resource,
				},
<<<<<<< HEAD
=======
				Client:       controllerReconciler.Client,
				Context:      ctx,
				Scheme:       controllerReconciler.Scheme,
				FeatureStore: resource,
>>>>>>> 6c1a66ea8 (feat: PVC configuration and impl (#4750))
=======
>>>>>>> 39eb4d80c (feat: RBAC Authorization in Feast Operator (#4786))
			}
			Expect(resource.Status).NotTo(BeNil())
			Expect(resource.Status.FeastVersion).To(Equal(feastversion.FeastVersion))
			Expect(resource.Status.ClientConfigMap).To(Equal(feast.GetFeastServiceName(services.ClientFeastType)))
			Expect(resource.Status.Applied.FeastProject).To(Equal(resource.Spec.FeastProject))
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
			Expect(resource.Status.Applied.AuthzConfig).To(BeNil())
=======
>>>>>>> 6c1a66ea8 (feat: PVC configuration and impl (#4750))
=======
			Expect(resource.Status.Applied.AuthzConfig).To(Equal(&feastdevv1alpha1.AuthzConfig{}))
>>>>>>> 39eb4d80c (feat: RBAC Authorization in Feast Operator (#4786))
=======
			Expect(resource.Status.Applied.AuthzConfig).To(BeNil())
>>>>>>> 668d47b8e (feat: Add TLS support to the Operator (#4796))
			Expect(resource.Status.Applied.Services).NotTo(BeNil())
			Expect(resource.Status.Applied.Services.OfflineStore).NotTo(BeNil())
			Expect(resource.Status.Applied.Services.OfflineStore.Persistence).NotTo(BeNil())
			Expect(resource.Status.Applied.Services.OfflineStore.Persistence.FilePersistence).NotTo(BeNil())
			Expect(resource.Status.Applied.Services.OfflineStore.Persistence.FilePersistence.Type).To(Equal(offlineType))
			Expect(resource.Status.Applied.Services.OfflineStore.ImagePullPolicy).To(BeNil())
			Expect(resource.Status.Applied.Services.OfflineStore.Resources).To(BeNil())
			Expect(resource.Status.Applied.Services.OfflineStore.Image).To(Equal(&services.DefaultImage))
			Expect(resource.Status.Applied.Services.OnlineStore).NotTo(BeNil())
			Expect(resource.Status.Applied.Services.OnlineStore.Persistence).NotTo(BeNil())
			Expect(resource.Status.Applied.Services.OnlineStore.Persistence.FilePersistence).NotTo(BeNil())
			Expect(resource.Status.Applied.Services.OnlineStore.Persistence.FilePersistence.Path).To(Equal(onlineStorePath))
			Expect(resource.Status.Applied.Services.OnlineStore.Env).To(Equal(&[]corev1.EnvVar{{Name: testEnvVarName, Value: testEnvVarValue}, {Name: "fieldRefName", ValueFrom: &corev1.EnvVarSource{FieldRef: &corev1.ObjectFieldSelector{APIVersion: "v1", FieldPath: "metadata.namespace"}}}}))
<<<<<<< HEAD
			Expect(resource.Status.Applied.Services.OnlineStore.EnvFrom).To(Equal(withEnvFrom()))
=======
>>>>>>> 6c1a66ea8 (feat: PVC configuration and impl (#4750))
			Expect(resource.Status.Applied.Services.OnlineStore.ImagePullPolicy).To(Equal(&pullPolicy))
			Expect(resource.Status.Applied.Services.OnlineStore.Resources).NotTo(BeNil())
			Expect(resource.Status.Applied.Services.OnlineStore.Image).To(Equal(&image))
			Expect(resource.Status.Applied.Services.Registry).NotTo(BeNil())
			Expect(resource.Status.Applied.Services.Registry.Local).NotTo(BeNil())
			Expect(resource.Status.Applied.Services.Registry.Local.Persistence).NotTo(BeNil())
			Expect(resource.Status.Applied.Services.Registry.Local.Persistence.FilePersistence).NotTo(BeNil())
			Expect(resource.Status.Applied.Services.Registry.Local.Persistence.FilePersistence.Path).To(Equal(registryPath))
			Expect(resource.Status.Applied.Services.Registry.Local.ImagePullPolicy).To(BeNil())
			Expect(resource.Status.Applied.Services.Registry.Local.Resources).To(BeNil())
			Expect(resource.Status.Applied.Services.Registry.Local.Image).To(Equal(&services.DefaultImage))

			Expect(resource.Status.ServiceHostnames.OfflineStore).To(Equal(feast.GetFeastServiceName(services.OfflineFeastType) + "." + resource.Namespace + domain))
			Expect(resource.Status.ServiceHostnames.OnlineStore).To(Equal(feast.GetFeastServiceName(services.OnlineFeastType) + "." + resource.Namespace + domain))
			Expect(resource.Status.ServiceHostnames.Registry).To(Equal(feast.GetFeastServiceName(services.RegistryFeastType) + "." + resource.Namespace + domain))

			Expect(resource.Status.Conditions).NotTo(BeEmpty())
			cond := apimeta.FindStatusCondition(resource.Status.Conditions, feastdevv1alpha1.ReadyType)
			Expect(cond).ToNot(BeNil())
			Expect(cond.Status).To(Equal(metav1.ConditionTrue))
			Expect(cond.Reason).To(Equal(feastdevv1alpha1.ReadyReason))
			Expect(cond.Type).To(Equal(feastdevv1alpha1.ReadyType))
			Expect(cond.Message).To(Equal(feastdevv1alpha1.ReadyMessage))

<<<<<<< HEAD
<<<<<<< HEAD
			cond = apimeta.FindStatusCondition(resource.Status.Conditions, feastdevv1alpha1.AuthorizationReadyType)
			Expect(cond).To(BeNil())

=======
>>>>>>> 6c1a66ea8 (feat: PVC configuration and impl (#4750))
=======
			cond = apimeta.FindStatusCondition(resource.Status.Conditions, feastdevv1alpha1.AuthorizationReadyType)
			Expect(cond).To(BeNil())

>>>>>>> 39eb4d80c (feat: RBAC Authorization in Feast Operator (#4786))
			cond = apimeta.FindStatusCondition(resource.Status.Conditions, feastdevv1alpha1.RegistryReadyType)
			Expect(cond).ToNot(BeNil())
			Expect(cond.Status).To(Equal(metav1.ConditionTrue))
			Expect(cond.Reason).To(Equal(feastdevv1alpha1.ReadyReason))
			Expect(cond.Type).To(Equal(feastdevv1alpha1.RegistryReadyType))
			Expect(cond.Message).To(Equal(feastdevv1alpha1.RegistryReadyMessage))

			cond = apimeta.FindStatusCondition(resource.Status.Conditions, feastdevv1alpha1.ClientReadyType)
			Expect(cond).ToNot(BeNil())
			Expect(cond.Status).To(Equal(metav1.ConditionTrue))
			Expect(cond.Reason).To(Equal(feastdevv1alpha1.ReadyReason))
			Expect(cond.Type).To(Equal(feastdevv1alpha1.ClientReadyType))
			Expect(cond.Message).To(Equal(feastdevv1alpha1.ClientReadyMessage))

			cond = apimeta.FindStatusCondition(resource.Status.Conditions, feastdevv1alpha1.OfflineStoreReadyType)
			Expect(cond).ToNot(BeNil())
			Expect(cond.Status).To(Equal(metav1.ConditionTrue))
			Expect(cond.Reason).To(Equal(feastdevv1alpha1.ReadyReason))
			Expect(cond.Type).To(Equal(feastdevv1alpha1.OfflineStoreReadyType))
			Expect(cond.Message).To(Equal(feastdevv1alpha1.OfflineStoreReadyMessage))

			cond = apimeta.FindStatusCondition(resource.Status.Conditions, feastdevv1alpha1.OnlineStoreReadyType)
			Expect(cond).ToNot(BeNil())
			Expect(cond.Status).To(Equal(metav1.ConditionTrue))
			Expect(cond.Reason).To(Equal(feastdevv1alpha1.ReadyReason))
			Expect(cond.Type).To(Equal(feastdevv1alpha1.OnlineStoreReadyType))
			Expect(cond.Message).To(Equal(feastdevv1alpha1.OnlineStoreReadyMessage))

			Expect(resource.Status.Phase).To(Equal(feastdevv1alpha1.ReadyPhase))

<<<<<<< HEAD
<<<<<<< HEAD
			// check deployment
			deploy := &appsv1.Deployment{}
			objMeta := feast.GetObjectMeta()
			err = k8sClient.Get(ctx, types.NamespacedName{
				Name:      objMeta.Name,
				Namespace: objMeta.Namespace,
			}, deploy)
			Expect(err).NotTo(HaveOccurred())
			Expect(deploy.Spec.Replicas).To(Equal(&services.DefaultReplicas))
			Expect(controllerutil.HasControllerReference(deploy)).To(BeTrue())
			Expect(deploy.Spec.Template.Spec.Containers).To(HaveLen(3))
=======
=======
			// check deployment
>>>>>>> b0a04af1d (fix: Refactor Operator to deploy all feast services to the same Deployment/Pod (#4863))
			deploy := &appsv1.Deployment{}
			objMeta := feast.GetObjectMeta()
			err = k8sClient.Get(ctx, types.NamespacedName{
				Name:      objMeta.Name,
				Namespace: objMeta.Namespace,
			}, deploy)
			Expect(err).NotTo(HaveOccurred())
			Expect(deploy.Spec.Replicas).To(Equal(&services.DefaultReplicas))
			Expect(controllerutil.HasControllerReference(deploy)).To(BeTrue())
<<<<<<< HEAD
			Expect(deploy.Spec.Template.Spec.Containers).To(HaveLen(1))
>>>>>>> 6c1a66ea8 (feat: PVC configuration and impl (#4750))
=======
			Expect(deploy.Spec.Template.Spec.Containers).To(HaveLen(3))
>>>>>>> b0a04af1d (fix: Refactor Operator to deploy all feast services to the same Deployment/Pod (#4863))

			svc := &corev1.Service{}
			err = k8sClient.Get(ctx, types.NamespacedName{
				Name:      feast.GetFeastServiceName(services.RegistryFeastType),
				Namespace: resource.Namespace,
			},
				svc)
			Expect(err).NotTo(HaveOccurred())
			Expect(controllerutil.HasControllerReference(svc)).To(BeTrue())
<<<<<<< HEAD
<<<<<<< HEAD
			Expect(svc.Spec.Ports[0].TargetPort).To(Equal(intstr.FromInt(int(services.FeastServiceConstants[services.RegistryFeastType].TargetHttpPort))))
=======
			Expect(svc.Spec.Ports[0].TargetPort).To(Equal(intstr.FromInt(int(services.FeastServiceConstants[services.RegistryFeastType].TargetPort))))
>>>>>>> 6c1a66ea8 (feat: PVC configuration and impl (#4750))
=======
			Expect(svc.Spec.Ports[0].TargetPort).To(Equal(intstr.FromInt(int(services.FeastServiceConstants[services.RegistryFeastType].TargetHttpPort))))
>>>>>>> 668d47b8e (feat: Add TLS support to the Operator (#4796))
		})

		It("should properly encode a feature_store.yaml config", func() {
			By("Reconciling the created resource")
			controllerReconciler := &FeatureStoreReconciler{
				Client: k8sClient,
				Scheme: k8sClient.Scheme(),
			}

			_, err := controllerReconciler.Reconcile(ctx, reconcile.Request{
				NamespacedName: typeNamespacedName,
			})
			Expect(err).NotTo(HaveOccurred())

			resource := &feastdevv1alpha1.FeatureStore{}
			err = k8sClient.Get(ctx, typeNamespacedName, resource)
			Expect(err).NotTo(HaveOccurred())

			req, err := labels.NewRequirement(services.NameLabelKey, selection.Equals, []string{resource.Name})
			Expect(err).NotTo(HaveOccurred())
			labelSelector := labels.NewSelector().Add(*req)
			listOpts := &client.ListOptions{Namespace: resource.Namespace, LabelSelector: labelSelector}
			deployList := appsv1.DeploymentList{}
			err = k8sClient.List(ctx, &deployList, listOpts)
			Expect(err).NotTo(HaveOccurred())
<<<<<<< HEAD
<<<<<<< HEAD
			Expect(deployList.Items).To(HaveLen(1))
=======
			Expect(deployList.Items).To(HaveLen(3))
>>>>>>> 6c1a66ea8 (feat: PVC configuration and impl (#4750))
=======
			Expect(deployList.Items).To(HaveLen(1))
>>>>>>> b0a04af1d (fix: Refactor Operator to deploy all feast services to the same Deployment/Pod (#4863))

			svcList := corev1.ServiceList{}
			err = k8sClient.List(ctx, &svcList, listOpts)
			Expect(err).NotTo(HaveOccurred())
			Expect(svcList.Items).To(HaveLen(3))

			cmList := corev1.ConfigMapList{}
			err = k8sClient.List(ctx, &cmList, listOpts)
			Expect(err).NotTo(HaveOccurred())
			Expect(cmList.Items).To(HaveLen(1))

			feast := services.FeastServices{
<<<<<<< HEAD
<<<<<<< HEAD
=======
>>>>>>> 39eb4d80c (feat: RBAC Authorization in Feast Operator (#4786))
				Handler: handler.FeastHandler{
					Client:       controllerReconciler.Client,
					Context:      ctx,
					Scheme:       controllerReconciler.Scheme,
					FeatureStore: resource,
				},
<<<<<<< HEAD
			}

			// check deployment
			deploy := &appsv1.Deployment{}
			objMeta := feast.GetObjectMeta()
			err = k8sClient.Get(ctx, types.NamespacedName{
				Name:      objMeta.Name,
				Namespace: objMeta.Namespace,
			}, deploy)
			Expect(err).NotTo(HaveOccurred())
			Expect(deploy.Spec.Template.Spec.Containers).To(HaveLen(3))
			registryContainer := services.GetRegistryContainer(deploy.Spec.Template.Spec.Containers)
			Expect(registryContainer.Env).To(HaveLen(1))
			env := getFeatureStoreYamlEnvVar(registryContainer.Env)
			Expect(env).NotTo(BeNil())

			fsYamlStr, err := feast.GetServiceFeatureStoreYamlBase64()
=======
				Client:       controllerReconciler.Client,
				Context:      ctx,
				Scheme:       controllerReconciler.Scheme,
				FeatureStore: resource,
=======
>>>>>>> 39eb4d80c (feat: RBAC Authorization in Feast Operator (#4786))
			}

			// check deployment
			deploy := &appsv1.Deployment{}
			objMeta := feast.GetObjectMeta()
			err = k8sClient.Get(ctx, types.NamespacedName{
				Name:      objMeta.Name,
				Namespace: objMeta.Namespace,
			}, deploy)
			Expect(err).NotTo(HaveOccurred())
			Expect(deploy.Spec.Template.Spec.Containers).To(HaveLen(3))
			registryContainer := services.GetRegistryContainer(deploy.Spec.Template.Spec.Containers)
			Expect(registryContainer.Env).To(HaveLen(1))
			env := getFeatureStoreYamlEnvVar(registryContainer.Env)
			Expect(env).NotTo(BeNil())

<<<<<<< HEAD
			fsYamlStr, err := feast.GetServiceFeatureStoreYamlBase64(services.RegistryFeastType)
>>>>>>> 6c1a66ea8 (feat: PVC configuration and impl (#4750))
=======
			fsYamlStr, err := feast.GetServiceFeatureStoreYamlBase64()
>>>>>>> b0a04af1d (fix: Refactor Operator to deploy all feast services to the same Deployment/Pod (#4863))
			Expect(err).NotTo(HaveOccurred())
			Expect(fsYamlStr).To(Equal(env.Value))

			envByte, err := base64.StdEncoding.DecodeString(env.Value)
			Expect(err).NotTo(HaveOccurred())
			repoConfig := &services.RepoConfig{}
			err = yaml.Unmarshal(envByte, repoConfig)
			Expect(err).NotTo(HaveOccurred())
			testConfig := &services.RepoConfig{
				Project:                       feastProject,
				Provider:                      services.LocalProviderType,
				EntityKeySerializationVersion: feastdevv1alpha1.SerializationVersion,
<<<<<<< HEAD
<<<<<<< HEAD
				OfflineStore: services.OfflineStoreConfig{
					Type: services.OfflineFilePersistenceDuckDbConfigType,
				},
=======
>>>>>>> 6c1a66ea8 (feat: PVC configuration and impl (#4750))
=======
				OfflineStore: services.OfflineStoreConfig{
					Type: services.OfflineFilePersistenceDuckDbConfigType,
				},
>>>>>>> b0a04af1d (fix: Refactor Operator to deploy all feast services to the same Deployment/Pod (#4863))
				Registry: services.RegistryConfig{
					RegistryType: services.RegistryFileConfigType,
					Path:         registryPath,
				},
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
=======
>>>>>>> b0a04af1d (fix: Refactor Operator to deploy all feast services to the same Deployment/Pod (#4863))
				OnlineStore: services.OnlineStoreConfig{
					Path: onlineStorePath,
					Type: services.OnlineSqliteConfigType,
				},
				AuthzConfig: noAuthzConfig(),
			}
			Expect(repoConfig).To(Equal(testConfig))

			offlineContainer := services.GetOfflineContainer(deploy.Spec.Template.Spec.Containers)
			Expect(offlineContainer.Env).To(HaveLen(1))
<<<<<<< HEAD
			assertEnvFrom(*offlineContainer)
			env = getFeatureStoreYamlEnvVar(offlineContainer.Env)
			Expect(env).NotTo(BeNil())

			//check envFrom for offlineContainer
			assertEnvFrom(*offlineContainer)

			fsYamlStr, err = feast.GetServiceFeatureStoreYamlBase64()
=======
=======
				AuthzConfig: noAuthzConfig(),
>>>>>>> 39eb4d80c (feat: RBAC Authorization in Feast Operator (#4786))
			}
			Expect(repoConfig).To(Equal(testConfig))

			// check offline config
			deploy = &appsv1.Deployment{}
			err = k8sClient.Get(ctx, types.NamespacedName{
				Name:      feast.GetFeastServiceName(services.OfflineFeastType),
				Namespace: resource.Namespace,
			},
				deploy)
			Expect(err).NotTo(HaveOccurred())
			Expect(deploy.Spec.Template.Spec.Containers).To(HaveLen(1))
			Expect(deploy.Spec.Template.Spec.Containers[0].Env).To(HaveLen(1))
			env = getFeatureStoreYamlEnvVar(deploy.Spec.Template.Spec.Containers[0].Env)
			Expect(env).NotTo(BeNil())

			fsYamlStr, err = feast.GetServiceFeatureStoreYamlBase64(services.OfflineFeastType)
>>>>>>> 6c1a66ea8 (feat: PVC configuration and impl (#4750))
=======
			env = getFeatureStoreYamlEnvVar(offlineContainer.Env)
			Expect(env).NotTo(BeNil())

			fsYamlStr, err = feast.GetServiceFeatureStoreYamlBase64()
>>>>>>> b0a04af1d (fix: Refactor Operator to deploy all feast services to the same Deployment/Pod (#4863))
			Expect(err).NotTo(HaveOccurred())
			Expect(fsYamlStr).To(Equal(env.Value))

			envByte, err = base64.StdEncoding.DecodeString(env.Value)
			Expect(err).NotTo(HaveOccurred())
			repoConfigOffline := &services.RepoConfig{}
			err = yaml.Unmarshal(envByte, repoConfigOffline)
			Expect(err).NotTo(HaveOccurred())
<<<<<<< HEAD
<<<<<<< HEAD
			Expect(repoConfigOffline).To(Equal(testConfig))

			onlineContainer := services.GetOnlineContainer(deploy.Spec.Template.Spec.Containers)
			Expect(onlineContainer.Env).To(HaveLen(3))
			Expect(onlineContainer.ImagePullPolicy).To(Equal(corev1.PullAlways))
			env = getFeatureStoreYamlEnvVar(onlineContainer.Env)
			Expect(env).NotTo(BeNil())

			fsYamlStr, err = feast.GetServiceFeatureStoreYamlBase64()
=======
			regRemote := services.RegistryConfig{
				RegistryType: services.RegistryRemoteConfigType,
				Path:         fmt.Sprintf("feast-%s-registry.default.svc.cluster.local:80", resourceName),
			}
			offlineConfig := &services.RepoConfig{
				Project:                       feastProject,
				Provider:                      services.LocalProviderType,
				EntityKeySerializationVersion: feastdevv1alpha1.SerializationVersion,
				OfflineStore: services.OfflineStoreConfig{
					Type: services.OfflineFilePersistenceDuckDbConfigType,
				},
				Registry:    regRemote,
				AuthzConfig: noAuthzConfig(),
			}
			Expect(repoConfigOffline).To(Equal(offlineConfig))
=======
			Expect(repoConfigOffline).To(Equal(testConfig))
>>>>>>> b0a04af1d (fix: Refactor Operator to deploy all feast services to the same Deployment/Pod (#4863))

			onlineContainer := services.GetOnlineContainer(deploy.Spec.Template.Spec.Containers)
			Expect(onlineContainer.Env).To(HaveLen(3))
			Expect(onlineContainer.ImagePullPolicy).To(Equal(corev1.PullAlways))
			env = getFeatureStoreYamlEnvVar(onlineContainer.Env)
			Expect(env).NotTo(BeNil())

<<<<<<< HEAD
			fsYamlStr, err = feast.GetServiceFeatureStoreYamlBase64(services.OnlineFeastType)
>>>>>>> 6c1a66ea8 (feat: PVC configuration and impl (#4750))
=======
			fsYamlStr, err = feast.GetServiceFeatureStoreYamlBase64()
>>>>>>> b0a04af1d (fix: Refactor Operator to deploy all feast services to the same Deployment/Pod (#4863))
			Expect(err).NotTo(HaveOccurred())
			Expect(fsYamlStr).To(Equal(env.Value))

			envByte, err = base64.StdEncoding.DecodeString(env.Value)
			Expect(err).NotTo(HaveOccurred())
			repoConfigOnline := &services.RepoConfig{}
			err = yaml.Unmarshal(envByte, repoConfigOnline)
			Expect(err).NotTo(HaveOccurred())
<<<<<<< HEAD
<<<<<<< HEAD
			Expect(repoConfigOnline).To(Equal(testConfig))
=======
			offlineRemote := services.OfflineStoreConfig{
				Host: fmt.Sprintf("feast-%s-offline.default.svc.cluster.local", resourceName),
				Type: services.OfflineRemoteConfigType,
				Port: services.HttpPort,
			}
			onlineConfig := &services.RepoConfig{
				Project:                       feastProject,
				Provider:                      services.LocalProviderType,
				EntityKeySerializationVersion: feastdevv1alpha1.SerializationVersion,
				OfflineStore:                  offlineRemote,
				OnlineStore: services.OnlineStoreConfig{
					Path: onlineStorePath,
					Type: services.OnlineSqliteConfigType,
				},
				Registry:    regRemote,
				AuthzConfig: noAuthzConfig(),
			}
			Expect(repoConfigOnline).To(Equal(onlineConfig))
			Expect(deploy.Spec.Template.Spec.Containers[0].Env).To(HaveLen(3))
>>>>>>> 6c1a66ea8 (feat: PVC configuration and impl (#4750))
=======
			Expect(repoConfigOnline).To(Equal(testConfig))
>>>>>>> b0a04af1d (fix: Refactor Operator to deploy all feast services to the same Deployment/Pod (#4863))

			// check client config
			cm := &corev1.ConfigMap{}
			name := feast.GetFeastServiceName(services.ClientFeastType)
			err = k8sClient.Get(ctx, types.NamespacedName{
				Name:      name,
				Namespace: resource.Namespace,
			},
				cm)
			Expect(err).NotTo(HaveOccurred())
			repoConfigClient := &services.RepoConfig{}
			err = yaml.Unmarshal([]byte(cm.Data[services.FeatureStoreYamlCmKey]), repoConfigClient)
			Expect(err).NotTo(HaveOccurred())
			clientConfig := &services.RepoConfig{
				Project:                       feastProject,
				Provider:                      services.LocalProviderType,
				EntityKeySerializationVersion: feastdevv1alpha1.SerializationVersion,
<<<<<<< HEAD
<<<<<<< HEAD
=======
>>>>>>> b0a04af1d (fix: Refactor Operator to deploy all feast services to the same Deployment/Pod (#4863))
				OfflineStore: services.OfflineStoreConfig{
					Host: fmt.Sprintf("feast-%s-offline.default.svc.cluster.local", resourceName),
					Type: services.OfflineRemoteConfigType,
					Port: services.HttpPort,
				},
<<<<<<< HEAD
=======
				OfflineStore:                  offlineRemote,
>>>>>>> 6c1a66ea8 (feat: PVC configuration and impl (#4750))
=======
>>>>>>> b0a04af1d (fix: Refactor Operator to deploy all feast services to the same Deployment/Pod (#4863))
				OnlineStore: services.OnlineStoreConfig{
					Path: fmt.Sprintf("http://feast-%s-online.default.svc.cluster.local:80", resourceName),
					Type: services.OnlineRemoteConfigType,
				},
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
=======
>>>>>>> b0a04af1d (fix: Refactor Operator to deploy all feast services to the same Deployment/Pod (#4863))
				Registry: services.RegistryConfig{
					RegistryType: services.RegistryRemoteConfigType,
					Path:         fmt.Sprintf("feast-%s-registry.default.svc.cluster.local:80", resourceName),
				},
<<<<<<< HEAD
				AuthzConfig: noAuthzConfig(),
=======
				Registry: regRemote,
>>>>>>> 6c1a66ea8 (feat: PVC configuration and impl (#4750))
=======
				Registry:    regRemote,
=======
>>>>>>> b0a04af1d (fix: Refactor Operator to deploy all feast services to the same Deployment/Pod (#4863))
				AuthzConfig: noAuthzConfig(),
>>>>>>> 39eb4d80c (feat: RBAC Authorization in Feast Operator (#4786))
			}
			Expect(repoConfigClient).To(Equal(clientConfig))

			// change paths and reconcile
			resourceNew := resource.DeepCopy()
			newOnlineStorePath := "/data/new_online.db"
			newRegistryPath := "/data/new_registry.db"
			resourceNew.Spec.Services.OnlineStore.Persistence.FilePersistence.Path = newOnlineStorePath
			resourceNew.Spec.Services.Registry.Local.Persistence.FilePersistence.Path = newRegistryPath
			err = k8sClient.Update(ctx, resourceNew)
			Expect(err).NotTo(HaveOccurred())
			_, err = controllerReconciler.Reconcile(ctx, reconcile.Request{
				NamespacedName: typeNamespacedName,
			})
			Expect(err).NotTo(HaveOccurred())

			resource = &feastdevv1alpha1.FeatureStore{}
			err = k8sClient.Get(ctx, typeNamespacedName, resource)
			Expect(err).NotTo(HaveOccurred())
<<<<<<< HEAD
<<<<<<< HEAD
			feast.Handler.FeatureStore = resource

			// check registry
<<<<<<< HEAD
			err = k8sClient.Get(ctx, types.NamespacedName{
				Name:      objMeta.Name,
				Namespace: objMeta.Namespace,
			}, deploy)
			Expect(err).NotTo(HaveOccurred())
			registryContainer = services.GetRegistryContainer(deploy.Spec.Template.Spec.Containers)
			env = getFeatureStoreYamlEnvVar(registryContainer.Env)
			Expect(env).NotTo(BeNil())
			fsYamlStr, err = feast.GetServiceFeatureStoreYamlBase64()
=======
			feast.FeatureStore = resource
=======
			feast.Handler.FeatureStore = resource
>>>>>>> 39eb4d80c (feat: RBAC Authorization in Feast Operator (#4786))

			// check registry config
			deploy = &appsv1.Deployment{}
=======
>>>>>>> b0a04af1d (fix: Refactor Operator to deploy all feast services to the same Deployment/Pod (#4863))
			err = k8sClient.Get(ctx, types.NamespacedName{
				Name:      objMeta.Name,
				Namespace: objMeta.Namespace,
			}, deploy)
			Expect(err).NotTo(HaveOccurred())
			registryContainer = services.GetRegistryContainer(deploy.Spec.Template.Spec.Containers)
			env = getFeatureStoreYamlEnvVar(registryContainer.Env)
			Expect(env).NotTo(BeNil())
<<<<<<< HEAD
			fsYamlStr, err = feast.GetServiceFeatureStoreYamlBase64(services.RegistryFeastType)
>>>>>>> 6c1a66ea8 (feat: PVC configuration and impl (#4750))
=======
			fsYamlStr, err = feast.GetServiceFeatureStoreYamlBase64()
>>>>>>> b0a04af1d (fix: Refactor Operator to deploy all feast services to the same Deployment/Pod (#4863))
			Expect(err).NotTo(HaveOccurred())
			Expect(fsYamlStr).To(Equal(env.Value))

			envByte, err = base64.StdEncoding.DecodeString(env.Value)
			Expect(err).NotTo(HaveOccurred())
			repoConfig = &services.RepoConfig{}
			err = yaml.Unmarshal(envByte, repoConfig)
			Expect(err).NotTo(HaveOccurred())
<<<<<<< HEAD
<<<<<<< HEAD
			testConfig.OnlineStore.Path = newOnlineStorePath
=======
>>>>>>> 6c1a66ea8 (feat: PVC configuration and impl (#4750))
=======
			testConfig.OnlineStore.Path = newOnlineStorePath
>>>>>>> b0a04af1d (fix: Refactor Operator to deploy all feast services to the same Deployment/Pod (#4863))
			testConfig.Registry.Path = newRegistryPath
			Expect(repoConfig).To(Equal(testConfig))

			// check offline config
<<<<<<< HEAD
<<<<<<< HEAD
			offlineContainer = services.GetRegistryContainer(deploy.Spec.Template.Spec.Containers)
			env = getFeatureStoreYamlEnvVar(offlineContainer.Env)
			Expect(env).NotTo(BeNil())

			fsYamlStr, err = feast.GetServiceFeatureStoreYamlBase64()
=======
			deploy = &appsv1.Deployment{}
			err = k8sClient.Get(ctx, types.NamespacedName{
				Name:      feast.GetFeastServiceName(services.OfflineFeastType),
				Namespace: resource.Namespace,
			},
				deploy)
			Expect(err).NotTo(HaveOccurred())
			env = getFeatureStoreYamlEnvVar(deploy.Spec.Template.Spec.Containers[0].Env)
			Expect(env).NotTo(BeNil())

			fsYamlStr, err = feast.GetServiceFeatureStoreYamlBase64(services.OfflineFeastType)
>>>>>>> 6c1a66ea8 (feat: PVC configuration and impl (#4750))
=======
			offlineContainer = services.GetRegistryContainer(deploy.Spec.Template.Spec.Containers)
			env = getFeatureStoreYamlEnvVar(offlineContainer.Env)
			Expect(env).NotTo(BeNil())

			fsYamlStr, err = feast.GetServiceFeatureStoreYamlBase64()
>>>>>>> b0a04af1d (fix: Refactor Operator to deploy all feast services to the same Deployment/Pod (#4863))
			Expect(err).NotTo(HaveOccurred())
			Expect(fsYamlStr).To(Equal(env.Value))

			envByte, err = base64.StdEncoding.DecodeString(env.Value)
			Expect(err).NotTo(HaveOccurred())
			repoConfigOffline = &services.RepoConfig{}
			err = yaml.Unmarshal(envByte, repoConfigOffline)
			Expect(err).NotTo(HaveOccurred())
<<<<<<< HEAD
<<<<<<< HEAD
			Expect(repoConfigOffline).To(Equal(testConfig))

			// check online config
			onlineContainer = services.GetOnlineContainer(deploy.Spec.Template.Spec.Containers)
			env = getFeatureStoreYamlEnvVar(onlineContainer.Env)
			Expect(env).NotTo(BeNil())

			//check envFrom
			// Validate `envFrom` for ConfigMap and Secret
			assertEnvFrom(*onlineContainer)

			fsYamlStr, err = feast.GetServiceFeatureStoreYamlBase64()
=======
			Expect(repoConfigOffline).To(Equal(offlineConfig))
=======
			Expect(repoConfigOffline).To(Equal(testConfig))
>>>>>>> b0a04af1d (fix: Refactor Operator to deploy all feast services to the same Deployment/Pod (#4863))

			// check online config
			onlineContainer = services.GetOnlineContainer(deploy.Spec.Template.Spec.Containers)
			env = getFeatureStoreYamlEnvVar(onlineContainer.Env)
			Expect(env).NotTo(BeNil())

<<<<<<< HEAD
			fsYamlStr, err = feast.GetServiceFeatureStoreYamlBase64(services.OnlineFeastType)
>>>>>>> 6c1a66ea8 (feat: PVC configuration and impl (#4750))
=======
			fsYamlStr, err = feast.GetServiceFeatureStoreYamlBase64()
>>>>>>> b0a04af1d (fix: Refactor Operator to deploy all feast services to the same Deployment/Pod (#4863))
			Expect(err).NotTo(HaveOccurred())
			Expect(fsYamlStr).To(Equal(env.Value))

			envByte, err = base64.StdEncoding.DecodeString(env.Value)
			Expect(err).NotTo(HaveOccurred())

			repoConfigOnline = &services.RepoConfig{}
			err = yaml.Unmarshal(envByte, repoConfigOnline)
			Expect(err).NotTo(HaveOccurred())
<<<<<<< HEAD
<<<<<<< HEAD
			testConfig.OnlineStore.Path = newOnlineStorePath
			Expect(repoConfigOnline).To(Equal(testConfig))
=======
			onlineConfig.OnlineStore.Path = newOnlineStorePath
			Expect(repoConfigOnline).To(Equal(onlineConfig))
>>>>>>> 6c1a66ea8 (feat: PVC configuration and impl (#4750))
=======
			testConfig.OnlineStore.Path = newOnlineStorePath
			Expect(repoConfigOnline).To(Equal(testConfig))
>>>>>>> b0a04af1d (fix: Refactor Operator to deploy all feast services to the same Deployment/Pod (#4863))
		})
	})
})
