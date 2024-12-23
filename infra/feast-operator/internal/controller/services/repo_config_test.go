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

package services

import (
	"fmt"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"gopkg.in/yaml.v3"
	corev1 "k8s.io/api/core/v1"
<<<<<<< HEAD
<<<<<<< HEAD
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
=======
>>>>>>> 863a82cb7 (feat: Added feast Go operator db stores support (#4771))
=======
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
>>>>>>> 33db9cabb (fix: Operator envVar positioning & tls.SecretRef.Name (#4806))

	feastdevv1alpha1 "github.com/feast-dev/feast/infra/feast-operator/api/v1alpha1"
)

var projectName = "test-project"

var _ = Describe("Repo Config", func() {
	Context("When creating the RepoConfig of a FeatureStore", func() {
		It("should successfully create the repo configs", func() {
			By("Having the minimal created resource")
			featureStore := minimalFeatureStore()
			ApplyDefaultsToStatus(featureStore)
<<<<<<< HEAD
<<<<<<< HEAD

=======
			var repoConfig RepoConfig
			repoConfig, err := getServiceRepoConfig(OfflineFeastType, featureStore, emptyMockExtractConfigFromSecret)
			Expect(err).NotTo(HaveOccurred())
			Expect(repoConfig.AuthzConfig.Type).To(Equal(NoAuthAuthType))
			Expect(repoConfig.OfflineStore).To(Equal(emptyOfflineStoreConfig()))
			Expect(repoConfig.OnlineStore).To(Equal(emptyOnlineStoreConfig()))
			Expect(repoConfig.Registry).To(Equal(emptyRegistryConfig()))

			repoConfig, err = getServiceRepoConfig(OnlineFeastType, featureStore, emptyMockExtractConfigFromSecret)
			Expect(err).NotTo(HaveOccurred())
			Expect(repoConfig.AuthzConfig.Type).To(Equal(NoAuthAuthType))
			Expect(repoConfig.OfflineStore).To(Equal(emptyOfflineStoreConfig()))
			Expect(repoConfig.OnlineStore).To(Equal(emptyOnlineStoreConfig()))
			Expect(repoConfig.Registry).To(Equal(emptyRegistryConfig()))

			repoConfig, err = getServiceRepoConfig(RegistryFeastType, featureStore, emptyMockExtractConfigFromSecret)
			Expect(err).NotTo(HaveOccurred())
			Expect(repoConfig.AuthzConfig.Type).To(Equal(NoAuthAuthType))
			Expect(repoConfig.OfflineStore).To(Equal(emptyOfflineStoreConfig()))
			Expect(repoConfig.OnlineStore).To(Equal(emptyOnlineStoreConfig()))
>>>>>>> 863a82cb7 (feat: Added feast Go operator db stores support (#4771))
			expectedRegistryConfig := RegistryConfig{
				RegistryType: "file",
<<<<<<< HEAD
				Path:         EphemeralPath + "/" + DefaultRegistryPath,
=======
				Path:         DefaultRegistryEphemeralPath,
>>>>>>> 6c1a66ea8 (feat: PVC configuration and impl (#4750))
=======

			expectedRegistryConfig := RegistryConfig{
				RegistryType: "file",
				Path:         EphemeralPath + "/" + DefaultRegistryPath,
>>>>>>> b0a04af1d (fix: Refactor Operator to deploy all feast services to the same Deployment/Pod (#4863))
			}
			expectedOnlineConfig := OnlineStoreConfig{
				Type: "sqlite",
				Path: EphemeralPath + "/" + DefaultOnlineStorePath,
			}

<<<<<<< HEAD
<<<<<<< HEAD
=======
			var repoConfig RepoConfig
>>>>>>> b0a04af1d (fix: Refactor Operator to deploy all feast services to the same Deployment/Pod (#4863))
=======
>>>>>>> e664d4a1f (fix: Improve status.applied updates & add offline pvc unit test (#4871))
			repoConfig, err := getServiceRepoConfig(featureStore, emptyMockExtractConfigFromSecret)
			Expect(err).NotTo(HaveOccurred())
			Expect(repoConfig.AuthzConfig.Type).To(Equal(NoAuthAuthType))
			Expect(repoConfig.OfflineStore).To(Equal(emptyOfflineStoreConfig))
			Expect(repoConfig.OnlineStore).To(Equal(expectedOnlineConfig))
			Expect(repoConfig.Registry).To(Equal(expectedRegistryConfig))

			By("Having the local registry resource")
			featureStore = minimalFeatureStore()
			featureStore.Spec.Services = &feastdevv1alpha1.FeatureStoreServices{
				Registry: &feastdevv1alpha1.Registry{
					Local: &feastdevv1alpha1.LocalRegistryConfig{
						Persistence: &feastdevv1alpha1.RegistryPersistence{
							FilePersistence: &feastdevv1alpha1.RegistryFilePersistence{
								Path: "file.db",
							},
						},
					},
				},
			}
			ApplyDefaultsToStatus(featureStore)
<<<<<<< HEAD
<<<<<<< HEAD

=======
			repoConfig, err = getServiceRepoConfig(OfflineFeastType, featureStore, emptyMockExtractConfigFromSecret)
			Expect(err).NotTo(HaveOccurred())
			Expect(repoConfig.AuthzConfig.Type).To(Equal(NoAuthAuthType))
			Expect(repoConfig.OfflineStore).To(Equal(emptyOfflineStoreConfig()))
			Expect(repoConfig.OnlineStore).To(Equal(emptyOnlineStoreConfig()))
			Expect(repoConfig.Registry).To(Equal(emptyRegistryConfig()))

			repoConfig, err = getServiceRepoConfig(OnlineFeastType, featureStore, emptyMockExtractConfigFromSecret)
			Expect(err).NotTo(HaveOccurred())
			Expect(repoConfig.AuthzConfig.Type).To(Equal(NoAuthAuthType))
			Expect(repoConfig.OfflineStore).To(Equal(emptyOfflineStoreConfig()))
			Expect(repoConfig.OnlineStore).To(Equal(emptyOnlineStoreConfig()))
			Expect(repoConfig.Registry).To(Equal(emptyRegistryConfig()))

			repoConfig, err = getServiceRepoConfig(RegistryFeastType, featureStore, emptyMockExtractConfigFromSecret)
			Expect(err).NotTo(HaveOccurred())
			Expect(repoConfig.AuthzConfig.Type).To(Equal(NoAuthAuthType))
			Expect(repoConfig.OfflineStore).To(Equal(emptyOfflineStoreConfig()))
			Expect(repoConfig.OnlineStore).To(Equal(emptyOnlineStoreConfig()))
>>>>>>> 863a82cb7 (feat: Added feast Go operator db stores support (#4771))
=======

>>>>>>> b0a04af1d (fix: Refactor Operator to deploy all feast services to the same Deployment/Pod (#4863))
			expectedRegistryConfig = RegistryConfig{
				RegistryType: "file",
				Path:         "file.db",
			}

			repoConfig, err = getServiceRepoConfig(featureStore, emptyMockExtractConfigFromSecret)
			Expect(err).NotTo(HaveOccurred())
			Expect(repoConfig.AuthzConfig.Type).To(Equal(NoAuthAuthType))
			Expect(repoConfig.OfflineStore).To(Equal(emptyOfflineStoreConfig))
			Expect(repoConfig.OnlineStore).To(Equal(expectedOnlineConfig))
			Expect(repoConfig.Registry).To(Equal(expectedRegistryConfig))

			By("Having the remote registry resource")
			featureStore = minimalFeatureStore()
			featureStore.Spec.Services = &feastdevv1alpha1.FeatureStoreServices{
				Registry: &feastdevv1alpha1.Registry{
					Remote: &feastdevv1alpha1.RemoteRegistryConfig{
						FeastRef: &feastdevv1alpha1.FeatureStoreRef{
							Name: "registry",
						},
					},
				},
			}
			ApplyDefaultsToStatus(featureStore)
<<<<<<< HEAD
<<<<<<< HEAD
			repoConfig, err = getServiceRepoConfig(featureStore, emptyMockExtractConfigFromSecret)
=======
			repoConfig, err = getServiceRepoConfig(OfflineFeastType, featureStore, emptyMockExtractConfigFromSecret)
>>>>>>> 863a82cb7 (feat: Added feast Go operator db stores support (#4771))
			Expect(err).NotTo(HaveOccurred())
			Expect(repoConfig.AuthzConfig.Type).To(Equal(NoAuthAuthType))
<<<<<<< HEAD
			Expect(repoConfig.OfflineStore).To(Equal(emptyOfflineStoreConfig))
			Expect(repoConfig.OnlineStore).To(Equal(expectedOnlineConfig))
			Expect(repoConfig.Registry).To(Equal(emptyRegistryConfig))
=======
			Expect(repoConfig.OfflineStore).To(Equal(emptyOfflineStoreConfig()))
			Expect(repoConfig.OnlineStore).To(Equal(emptyOnlineStoreConfig()))
			Expect(repoConfig.Registry).To(Equal(emptyRegistryConfig()))
>>>>>>> 39eb4d80c (feat: RBAC Authorization in Feast Operator (#4786))

<<<<<<< HEAD
			By("Having an offlineStore with PVC")
			mountPath := "/testing"
			expectedOnlineConfig.Path = mountPath + "/" + DefaultOnlineStorePath
			expectedRegistryConfig.Path = mountPath + "/" + DefaultRegistryPath

			featureStore = minimalFeatureStore()
			featureStore.Spec.Services = &feastdevv1alpha1.FeatureStoreServices{
				OfflineStore: &feastdevv1alpha1.OfflineStore{
					Persistence: &feastdevv1alpha1.OfflineStorePersistence{
						FilePersistence: &feastdevv1alpha1.OfflineStoreFilePersistence{
							PvcConfig: &feastdevv1alpha1.PvcConfig{
								MountPath: mountPath,
							},
						},
					},
				},
			}
			ApplyDefaultsToStatus(featureStore)
			appliedServices := featureStore.Status.Applied.Services
			Expect(appliedServices.OnlineStore).To(BeNil())
			Expect(appliedServices.Registry.Local.Persistence.FilePersistence.Path).To(Equal(expectedRegistryConfig.Path))

			repoConfig, err = getServiceRepoConfig(featureStore, emptyMockExtractConfigFromSecret)
=======
			repoConfig, err = getServiceRepoConfig(OnlineFeastType, featureStore, emptyMockExtractConfigFromSecret)
>>>>>>> 863a82cb7 (feat: Added feast Go operator db stores support (#4771))
			Expect(err).NotTo(HaveOccurred())
<<<<<<< HEAD
			Expect(repoConfig.OfflineStore).To(Equal(defaultOfflineStoreConfig))
			Expect(repoConfig.AuthzConfig.Type).To(Equal(NoAuthAuthType))
			Expect(repoConfig.Registry).To(Equal(expectedRegistryConfig))
			Expect(repoConfig.OnlineStore).To(Equal(expectedOnlineConfig))
=======
			Expect(repoConfig.AuthzConfig.Type).To(Equal(NoAuthAuthType))
			Expect(repoConfig.OfflineStore).To(Equal(emptyOfflineStoreConfig()))
			Expect(repoConfig.OnlineStore).To(Equal(emptyOnlineStoreConfig()))
			Expect(repoConfig.Registry).To(Equal(emptyRegistryConfig()))
>>>>>>> 39eb4d80c (feat: RBAC Authorization in Feast Operator (#4786))

<<<<<<< HEAD
=======
			repoConfig, err = getServiceRepoConfig(RegistryFeastType, featureStore, emptyMockExtractConfigFromSecret)
			Expect(err).NotTo(HaveOccurred())
			Expect(repoConfig.AuthzConfig.Type).To(Equal(NoAuthAuthType))
			Expect(repoConfig.OfflineStore).To(Equal(emptyOfflineStoreConfig()))
			Expect(repoConfig.OnlineStore).To(Equal(emptyOnlineStoreConfig()))
			Expect(repoConfig.Registry).To(Equal(emptyRegistryConfig()))
=======
			repoConfig, err = getServiceRepoConfig(featureStore, emptyMockExtractConfigFromSecret)
			Expect(err).NotTo(HaveOccurred())
			Expect(repoConfig.AuthzConfig.Type).To(Equal(NoAuthAuthType))
			Expect(repoConfig.OfflineStore).To(Equal(emptyOfflineStoreConfig))
			Expect(repoConfig.OnlineStore).To(Equal(expectedOnlineConfig))
			Expect(repoConfig.Registry).To(Equal(emptyRegistryConfig))
>>>>>>> b0a04af1d (fix: Refactor Operator to deploy all feast services to the same Deployment/Pod (#4863))

<<<<<<< HEAD
>>>>>>> 863a82cb7 (feat: Added feast Go operator db stores support (#4771))
=======
			By("Having an offlineStore with PVC")
			mountPath := "/testing"
			expectedOnlineConfig.Path = mountPath + "/" + DefaultOnlineStorePath
			expectedRegistryConfig.Path = mountPath + "/" + DefaultRegistryPath

			featureStore = minimalFeatureStore()
			featureStore.Spec.Services = &feastdevv1alpha1.FeatureStoreServices{
				OfflineStore: &feastdevv1alpha1.OfflineStore{
					Persistence: &feastdevv1alpha1.OfflineStorePersistence{
						FilePersistence: &feastdevv1alpha1.OfflineStoreFilePersistence{
							PvcConfig: &feastdevv1alpha1.PvcConfig{
								MountPath: mountPath,
							},
						},
					},
				},
			}
			ApplyDefaultsToStatus(featureStore)
			appliedServices := featureStore.Status.Applied.Services
			Expect(appliedServices.OnlineStore).To(BeNil())
			Expect(appliedServices.Registry.Local.Persistence.FilePersistence.Path).To(Equal(expectedRegistryConfig.Path))

			repoConfig, err = getServiceRepoConfig(featureStore, emptyMockExtractConfigFromSecret)
			Expect(err).NotTo(HaveOccurred())
			Expect(repoConfig.OfflineStore).To(Equal(defaultOfflineStoreConfig))
			Expect(repoConfig.AuthzConfig.Type).To(Equal(NoAuthAuthType))
			Expect(repoConfig.Registry).To(Equal(expectedRegistryConfig))
			Expect(repoConfig.OnlineStore).To(Equal(expectedOnlineConfig))

>>>>>>> e664d4a1f (fix: Improve status.applied updates & add offline pvc unit test (#4871))
			By("Having the all the file services")
			featureStore = minimalFeatureStore()
			featureStore.Spec.Services = &feastdevv1alpha1.FeatureStoreServices{
				OfflineStore: &feastdevv1alpha1.OfflineStore{
					Persistence: &feastdevv1alpha1.OfflineStorePersistence{
						FilePersistence: &feastdevv1alpha1.OfflineStoreFilePersistence{
							Type: "duckdb",
						},
					},
				},
				OnlineStore: &feastdevv1alpha1.OnlineStore{
					Persistence: &feastdevv1alpha1.OnlineStorePersistence{
						FilePersistence: &feastdevv1alpha1.OnlineStoreFilePersistence{
							Path: "/data/online.db",
						},
					},
				},
				Registry: &feastdevv1alpha1.Registry{
					Local: &feastdevv1alpha1.LocalRegistryConfig{
						Persistence: &feastdevv1alpha1.RegistryPersistence{
							FilePersistence: &feastdevv1alpha1.RegistryFilePersistence{
								Path: "/data/registry.db",
							},
						},
					},
				},
			}
			ApplyDefaultsToStatus(featureStore)
<<<<<<< HEAD
<<<<<<< HEAD

			expectedOfflineConfig := OfflineStoreConfig{
				Type: "duckdb",
			}
=======
			repoConfig, err = getServiceRepoConfig(OfflineFeastType, featureStore, emptyMockExtractConfigFromSecret)
			Expect(err).NotTo(HaveOccurred())
			Expect(repoConfig.AuthzConfig.Type).To(Equal(NoAuthAuthType))
			expectedOfflineConfig := OfflineStoreConfig{
				Type: "duckdb",
			}
			Expect(repoConfig.OfflineStore).To(Equal(expectedOfflineConfig))
			Expect(repoConfig.OnlineStore).To(Equal(emptyOnlineStoreConfig()))
			Expect(repoConfig.Registry).To(Equal(emptyRegistryConfig()))

			repoConfig, err = getServiceRepoConfig(OnlineFeastType, featureStore, emptyMockExtractConfigFromSecret)
			Expect(err).NotTo(HaveOccurred())
			Expect(repoConfig.AuthzConfig.Type).To(Equal(NoAuthAuthType))
			Expect(repoConfig.OfflineStore).To(Equal(emptyOfflineStoreConfig()))
			expectedOnlineConfig := OnlineStoreConfig{
				Type: "sqlite",
				Path: "/data/online.db",
			}
			Expect(repoConfig.OnlineStore).To(Equal(expectedOnlineConfig))
			Expect(repoConfig.Registry).To(Equal(emptyRegistryConfig()))

			repoConfig, err = getServiceRepoConfig(RegistryFeastType, featureStore, emptyMockExtractConfigFromSecret)
			Expect(err).NotTo(HaveOccurred())
			Expect(repoConfig.AuthzConfig.Type).To(Equal(NoAuthAuthType))
			Expect(repoConfig.OfflineStore).To(Equal(emptyOfflineStoreConfig()))
			Expect(repoConfig.OnlineStore).To(Equal(emptyOnlineStoreConfig()))
>>>>>>> 863a82cb7 (feat: Added feast Go operator db stores support (#4771))
=======

			expectedOfflineConfig := OfflineStoreConfig{
				Type: "duckdb",
			}
>>>>>>> b0a04af1d (fix: Refactor Operator to deploy all feast services to the same Deployment/Pod (#4863))
			expectedRegistryConfig = RegistryConfig{
				RegistryType: "file",
				Path:         "/data/registry.db",
			}
			expectedOnlineConfig = OnlineStoreConfig{
				Type: "sqlite",
				Path: "/data/online.db",
			}

			repoConfig, err = getServiceRepoConfig(featureStore, emptyMockExtractConfigFromSecret)
			Expect(err).NotTo(HaveOccurred())
			Expect(repoConfig.AuthzConfig.Type).To(Equal(NoAuthAuthType))
			Expect(repoConfig.OfflineStore).To(Equal(expectedOfflineConfig))
			Expect(repoConfig.OnlineStore).To(Equal(expectedOnlineConfig))
<<<<<<< HEAD
			Expect(repoConfig.Registry).To(Equal(expectedRegistryConfig))

			By("Having kubernetes authorization")
			featureStore = minimalFeatureStore()
			featureStore.Spec.AuthzConfig = &feastdevv1alpha1.AuthzConfig{
				KubernetesAuthz: &feastdevv1alpha1.KubernetesAuthz{},
			}
			featureStore.Spec.Services = &feastdevv1alpha1.FeatureStoreServices{
				OfflineStore: &feastdevv1alpha1.OfflineStore{},
				OnlineStore:  &feastdevv1alpha1.OnlineStore{},
				Registry: &feastdevv1alpha1.Registry{
					Local: &feastdevv1alpha1.LocalRegistryConfig{},
				},
			}
			ApplyDefaultsToStatus(featureStore)

			expectedOfflineConfig = OfflineStoreConfig{
				Type: "dask",
			}

			repoConfig, err = getServiceRepoConfig(featureStore, mockExtractConfigFromSecret)
			Expect(err).NotTo(HaveOccurred())
			Expect(repoConfig.AuthzConfig.Type).To(Equal(KubernetesAuthType))
			Expect(repoConfig.OfflineStore).To(Equal(expectedOfflineConfig))
			Expect(repoConfig.OnlineStore).To(Equal(defaultOnlineStoreConfig(featureStore)))
			Expect(repoConfig.Registry).To(Equal(defaultRegistryConfig(featureStore)))

			By("Having oidc authorization")
			featureStore.Spec.AuthzConfig = &feastdevv1alpha1.AuthzConfig{
				OidcAuthz: &feastdevv1alpha1.OidcAuthz{
					SecretRef: corev1.LocalObjectReference{
						Name: "oidc-secret",
					},
				},
			}
			ApplyDefaultsToStatus(featureStore)

			secretExtractionFunc := mockOidcConfigFromSecret(map[string]interface{}{
				string(OidcAuthDiscoveryUrl): "discovery-url",
				string(OidcClientId):         "client-id",
				string(OidcClientSecret):     "client-secret",
				string(OidcUsername):         "username",
				string(OidcPassword):         "password"})
			repoConfig, err = getServiceRepoConfig(featureStore, secretExtractionFunc)
			Expect(err).NotTo(HaveOccurred())
			Expect(repoConfig.AuthzConfig.Type).To(Equal(OidcAuthType))
			Expect(repoConfig.AuthzConfig.OidcParameters).To(HaveLen(2))
			Expect(repoConfig.AuthzConfig.OidcParameters).To(HaveKey(string(OidcClientId)))
			Expect(repoConfig.AuthzConfig.OidcParameters).To(HaveKey(string(OidcAuthDiscoveryUrl)))
			Expect(repoConfig.OfflineStore).To(Equal(expectedOfflineConfig))
			Expect(repoConfig.OnlineStore).To(Equal(defaultOnlineStoreConfig(featureStore)))
			Expect(repoConfig.Registry).To(Equal(defaultRegistryConfig(featureStore)))

			repoConfig, err = getClientRepoConfig(featureStore, secretExtractionFunc)
			Expect(err).NotTo(HaveOccurred())
			Expect(repoConfig.AuthzConfig.Type).To(Equal(OidcAuthType))
			Expect(repoConfig.AuthzConfig.OidcParameters).To(HaveLen(3))
			Expect(repoConfig.AuthzConfig.OidcParameters).To(HaveKey(string(OidcClientSecret)))
			Expect(repoConfig.AuthzConfig.OidcParameters).To(HaveKey(string(OidcUsername)))
			Expect(repoConfig.AuthzConfig.OidcParameters).To(HaveKey(string(OidcPassword)))

			By("Having the all the db services")
			featureStore = minimalFeatureStore()
			featureStore.Spec.Services = &feastdevv1alpha1.FeatureStoreServices{
				OfflineStore: &feastdevv1alpha1.OfflineStore{
					Persistence: &feastdevv1alpha1.OfflineStorePersistence{
						DBPersistence: &feastdevv1alpha1.OfflineStoreDBStorePersistence{
							Type: string(OfflineDBPersistenceSnowflakeConfigType),
							SecretRef: corev1.LocalObjectReference{
								Name: "offline-test-secret",
							},
						},
					},
				},
				OnlineStore: &feastdevv1alpha1.OnlineStore{
					Persistence: &feastdevv1alpha1.OnlineStorePersistence{
						DBPersistence: &feastdevv1alpha1.OnlineStoreDBStorePersistence{
							Type: string(OnlineDBPersistenceSnowflakeConfigType),
							SecretRef: corev1.LocalObjectReference{
								Name: "online-test-secret",
							},
						},
					},
				},
				Registry: &feastdevv1alpha1.Registry{
					Local: &feastdevv1alpha1.LocalRegistryConfig{
						Persistence: &feastdevv1alpha1.RegistryPersistence{
							DBPersistence: &feastdevv1alpha1.RegistryDBStorePersistence{
								Type: string(RegistryDBPersistenceSnowflakeConfigType),
								SecretRef: corev1.LocalObjectReference{
									Name: "registry-test-secret",
								},
							},
						},
					},
				},
			}
			parameterMap := createParameterMap()
			ApplyDefaultsToStatus(featureStore)
			featureStore.Spec.Services.OfflineStore.Persistence.FilePersistence = nil
			featureStore.Spec.Services.OnlineStore.Persistence.FilePersistence = nil
			featureStore.Spec.Services.Registry.Local.Persistence.FilePersistence = nil
			repoConfig, err = getServiceRepoConfig(featureStore, mockExtractConfigFromSecret)
			Expect(err).NotTo(HaveOccurred())
			newMap := CopyMap(parameterMap)
			port := parameterMap["port"].(int)
			delete(newMap, "port")
			expectedOfflineConfig = OfflineStoreConfig{
				Type:         OfflineDBPersistenceSnowflakeConfigType,
				Port:         port,
				DBParameters: newMap,
			}
			expectedOnlineConfig = OnlineStoreConfig{
				Type:         OnlineDBPersistenceSnowflakeConfigType,
				DBParameters: CopyMap(parameterMap),
			}
			expectedRegistryConfig = RegistryConfig{
				RegistryType: RegistryDBPersistenceSnowflakeConfigType,
				DBParameters: parameterMap,
			}
			Expect(repoConfig.OfflineStore).To(Equal(expectedOfflineConfig))
			Expect(repoConfig.OnlineStore).To(Equal(expectedOnlineConfig))
=======
>>>>>>> b0a04af1d (fix: Refactor Operator to deploy all feast services to the same Deployment/Pod (#4863))
			Expect(repoConfig.Registry).To(Equal(expectedRegistryConfig))

			By("Having kubernetes authorization")
			featureStore = minimalFeatureStore()
			featureStore.Spec.AuthzConfig = &feastdevv1alpha1.AuthzConfig{
				KubernetesAuthz: &feastdevv1alpha1.KubernetesAuthz{},
			}
			featureStore.Spec.Services = &feastdevv1alpha1.FeatureStoreServices{
				OfflineStore: &feastdevv1alpha1.OfflineStore{},
				OnlineStore:  &feastdevv1alpha1.OnlineStore{},
				Registry: &feastdevv1alpha1.Registry{
					Local: &feastdevv1alpha1.LocalRegistryConfig{},
				},
			}
			ApplyDefaultsToStatus(featureStore)

			expectedOfflineConfig = OfflineStoreConfig{
				Type: "dask",
			}

			repoConfig, err = getServiceRepoConfig(featureStore, mockExtractConfigFromSecret)
			Expect(err).NotTo(HaveOccurred())
			Expect(repoConfig.AuthzConfig.Type).To(Equal(KubernetesAuthType))
			Expect(repoConfig.OfflineStore).To(Equal(expectedOfflineConfig))
			Expect(repoConfig.OnlineStore).To(Equal(defaultOnlineStoreConfig(featureStore)))
			Expect(repoConfig.Registry).To(Equal(defaultRegistryConfig(featureStore)))

			By("Having oidc authorization")
			featureStore.Spec.AuthzConfig = &feastdevv1alpha1.AuthzConfig{
				OidcAuthz: &feastdevv1alpha1.OidcAuthz{
					SecretRef: corev1.LocalObjectReference{
						Name: "oidc-secret",
					},
				},
			}
			ApplyDefaultsToStatus(featureStore)

			secretExtractionFunc := mockOidcConfigFromSecret(map[string]interface{}{
				string(OidcAuthDiscoveryUrl): "discovery-url",
				string(OidcClientId):         "client-id",
				string(OidcClientSecret):     "client-secret",
				string(OidcUsername):         "username",
				string(OidcPassword):         "password"})
			repoConfig, err = getServiceRepoConfig(featureStore, secretExtractionFunc)
			Expect(err).NotTo(HaveOccurred())
			Expect(repoConfig.AuthzConfig.Type).To(Equal(OidcAuthType))
			Expect(repoConfig.AuthzConfig.OidcParameters).To(HaveLen(2))
			Expect(repoConfig.AuthzConfig.OidcParameters).To(HaveKey(string(OidcClientId)))
			Expect(repoConfig.AuthzConfig.OidcParameters).To(HaveKey(string(OidcAuthDiscoveryUrl)))
			Expect(repoConfig.OfflineStore).To(Equal(expectedOfflineConfig))
			Expect(repoConfig.OnlineStore).To(Equal(defaultOnlineStoreConfig(featureStore)))
			Expect(repoConfig.Registry).To(Equal(defaultRegistryConfig(featureStore)))

			repoConfig, err = getClientRepoConfig(featureStore, secretExtractionFunc)
			Expect(err).NotTo(HaveOccurred())
			Expect(repoConfig.AuthzConfig.Type).To(Equal(OidcAuthType))
			Expect(repoConfig.AuthzConfig.OidcParameters).To(HaveLen(3))
			Expect(repoConfig.AuthzConfig.OidcParameters).To(HaveKey(string(OidcClientSecret)))
			Expect(repoConfig.AuthzConfig.OidcParameters).To(HaveKey(string(OidcUsername)))
			Expect(repoConfig.AuthzConfig.OidcParameters).To(HaveKey(string(OidcPassword)))

			By("Having the all the db services")
			featureStore = minimalFeatureStore()
			featureStore.Spec.Services = &feastdevv1alpha1.FeatureStoreServices{
				OfflineStore: &feastdevv1alpha1.OfflineStore{
					Persistence: &feastdevv1alpha1.OfflineStorePersistence{
						DBPersistence: &feastdevv1alpha1.OfflineStoreDBStorePersistence{
							Type: string(OfflineDBPersistenceSnowflakeConfigType),
							SecretRef: corev1.LocalObjectReference{
								Name: "offline-test-secret",
							},
						},
					},
				},
				OnlineStore: &feastdevv1alpha1.OnlineStore{
					Persistence: &feastdevv1alpha1.OnlineStorePersistence{
						DBPersistence: &feastdevv1alpha1.OnlineStoreDBStorePersistence{
							Type: string(OnlineDBPersistenceSnowflakeConfigType),
							SecretRef: corev1.LocalObjectReference{
								Name: "online-test-secret",
							},
						},
					},
				},
				Registry: &feastdevv1alpha1.Registry{
					Local: &feastdevv1alpha1.LocalRegistryConfig{
						Persistence: &feastdevv1alpha1.RegistryPersistence{
							DBPersistence: &feastdevv1alpha1.RegistryDBStorePersistence{
								Type: string(RegistryDBPersistenceSnowflakeConfigType),
								SecretRef: corev1.LocalObjectReference{
									Name: "registry-test-secret",
								},
							},
						},
					},
				},
			}
			parameterMap := createParameterMap()
			ApplyDefaultsToStatus(featureStore)
			featureStore.Spec.Services.OfflineStore.Persistence.FilePersistence = nil
			featureStore.Spec.Services.OnlineStore.Persistence.FilePersistence = nil
			featureStore.Spec.Services.Registry.Local.Persistence.FilePersistence = nil
			repoConfig, err = getServiceRepoConfig(featureStore, mockExtractConfigFromSecret)
			Expect(err).NotTo(HaveOccurred())
			newMap := CopyMap(parameterMap)
			port := parameterMap["port"].(int)
			delete(newMap, "port")
			expectedOfflineConfig = OfflineStoreConfig{
				Type:         OfflineDBPersistenceSnowflakeConfigType,
				Port:         port,
				DBParameters: newMap,
			}
			expectedOnlineConfig = OnlineStoreConfig{
				Type:         OnlineDBPersistenceSnowflakeConfigType,
				DBParameters: CopyMap(parameterMap),
			}
			expectedRegistryConfig = RegistryConfig{
				RegistryType: RegistryDBPersistenceSnowflakeConfigType,
				DBParameters: parameterMap,
			}
			Expect(repoConfig.OfflineStore).To(Equal(expectedOfflineConfig))
			Expect(repoConfig.OnlineStore).To(Equal(expectedOnlineConfig))
			Expect(repoConfig.Registry).To(Equal(expectedRegistryConfig))
		})
	})
	It("should fail to create the repo configs", func() {
		featureStore := minimalFeatureStore()

		By("Having invalid server oidc authorization")
		featureStore.Spec.AuthzConfig = &feastdevv1alpha1.AuthzConfig{
			OidcAuthz: &feastdevv1alpha1.OidcAuthz{
				SecretRef: corev1.LocalObjectReference{
					Name: "oidc-secret",
				},
			},
		}
		ApplyDefaultsToStatus(featureStore)

		secretExtractionFunc := mockOidcConfigFromSecret(map[string]interface{}{
			string(OidcClientId):     "client-id",
			string(OidcClientSecret): "client-secret",
			string(OidcUsername):     "username",
			string(OidcPassword):     "password"})
<<<<<<< HEAD
<<<<<<< HEAD
		_, err := getServiceRepoConfig(featureStore, secretExtractionFunc)
		Expect(err).To(HaveOccurred())
		Expect(err.Error()).To(ContainSubstring("missing OIDC secret"))
		_, err = getServiceRepoConfig(featureStore, secretExtractionFunc)
		Expect(err).To(HaveOccurred())
		Expect(err.Error()).To(ContainSubstring("missing OIDC secret"))
		_, err = getServiceRepoConfig(featureStore, secretExtractionFunc)
=======
		_, err := getServiceRepoConfig(OfflineFeastType, featureStore, secretExtractionFunc)
=======
		_, err := getServiceRepoConfig(featureStore, secretExtractionFunc)
>>>>>>> b0a04af1d (fix: Refactor Operator to deploy all feast services to the same Deployment/Pod (#4863))
		Expect(err).To(HaveOccurred())
		Expect(err.Error()).To(ContainSubstring("missing OIDC secret"))
		_, err = getServiceRepoConfig(featureStore, secretExtractionFunc)
		Expect(err).To(HaveOccurred())
		Expect(err.Error()).To(ContainSubstring("missing OIDC secret"))
<<<<<<< HEAD
		_, err = getServiceRepoConfig(RegistryFeastType, featureStore, secretExtractionFunc)
>>>>>>> cd341f8f6 (feat: OIDC authorization in Feast Operator (#4801))
=======
		_, err = getServiceRepoConfig(featureStore, secretExtractionFunc)
>>>>>>> b0a04af1d (fix: Refactor Operator to deploy all feast services to the same Deployment/Pod (#4863))
		Expect(err).To(HaveOccurred())
		Expect(err.Error()).To(ContainSubstring("missing OIDC secret"))
		_, err = getClientRepoConfig(featureStore, secretExtractionFunc)
		Expect(err).ToNot(HaveOccurred())

		By("Having invalid client oidc authorization")
		featureStore.Spec.AuthzConfig = &feastdevv1alpha1.AuthzConfig{
			OidcAuthz: &feastdevv1alpha1.OidcAuthz{
				SecretRef: corev1.LocalObjectReference{
					Name: "oidc-secret",
				},
			},
		}
		ApplyDefaultsToStatus(featureStore)

		secretExtractionFunc = mockOidcConfigFromSecret(map[string]interface{}{
			string(OidcAuthDiscoveryUrl): "discovery-url",
			string(OidcClientId):         "client-id",
			string(OidcUsername):         "username",
			string(OidcPassword):         "password"})
<<<<<<< HEAD
<<<<<<< HEAD
		_, err = getServiceRepoConfig(featureStore, secretExtractionFunc)
		Expect(err).To(HaveOccurred())
		Expect(err.Error()).To(ContainSubstring("missing OIDC secret"))
		_, err = getServiceRepoConfig(featureStore, secretExtractionFunc)
		Expect(err).To(HaveOccurred())
		Expect(err.Error()).To(ContainSubstring("missing OIDC secret"))
		_, err = getServiceRepoConfig(featureStore, secretExtractionFunc)
=======
		_, err = getServiceRepoConfig(OfflineFeastType, featureStore, secretExtractionFunc)
=======
		_, err = getServiceRepoConfig(featureStore, secretExtractionFunc)
>>>>>>> b0a04af1d (fix: Refactor Operator to deploy all feast services to the same Deployment/Pod (#4863))
		Expect(err).To(HaveOccurred())
		Expect(err.Error()).To(ContainSubstring("missing OIDC secret"))
		_, err = getServiceRepoConfig(featureStore, secretExtractionFunc)
		Expect(err).To(HaveOccurred())
		Expect(err.Error()).To(ContainSubstring("missing OIDC secret"))
<<<<<<< HEAD
		_, err = getServiceRepoConfig(RegistryFeastType, featureStore, secretExtractionFunc)
>>>>>>> cd341f8f6 (feat: OIDC authorization in Feast Operator (#4801))
=======
		_, err = getServiceRepoConfig(featureStore, secretExtractionFunc)
>>>>>>> b0a04af1d (fix: Refactor Operator to deploy all feast services to the same Deployment/Pod (#4863))
		Expect(err).To(HaveOccurred())
		Expect(err.Error()).To(ContainSubstring("missing OIDC secret"))
		_, err = getClientRepoConfig(featureStore, secretExtractionFunc)
		Expect(err).To(HaveOccurred())
		Expect(err.Error()).To(ContainSubstring("missing OIDC secret"))
	})
})

var emptyOfflineStoreConfig = OfflineStoreConfig{}
var emptyRegistryConfig = RegistryConfig{}

func minimalFeatureStore() *feastdevv1alpha1.FeatureStore {
	return &feastdevv1alpha1.FeatureStore{
		ObjectMeta: metav1.ObjectMeta{Name: "test"},
		Spec: feastdevv1alpha1.FeatureStoreSpec{
			FeastProject: projectName,
		},
	}
}

<<<<<<< HEAD
<<<<<<< HEAD
=======
>>>>>>> 668d47b8e (feat: Add TLS support to the Operator (#4796))
func minimalFeatureStoreWithAllServices() *feastdevv1alpha1.FeatureStore {
	feast := minimalFeatureStore()
	feast.Spec.Services = &feastdevv1alpha1.FeatureStoreServices{
		OfflineStore: &feastdevv1alpha1.OfflineStore{},
		OnlineStore:  &feastdevv1alpha1.OnlineStore{},
		Registry:     &feastdevv1alpha1.Registry{},
	}
	return feast
}

<<<<<<< HEAD
<<<<<<< HEAD
func emptyMockExtractConfigFromSecret(storeType string, secretRef string, secretKeyName string) (map[string]interface{}, error) {
	return map[string]interface{}{}, nil
}

func mockExtractConfigFromSecret(storeType string, secretRef string, secretKeyName string) (map[string]interface{}, error) {
	return createParameterMap(), nil
}

func mockOidcConfigFromSecret(
	oidcProperties map[string]interface{}) func(storeType string, secretRef string, secretKeyName string) (map[string]interface{}, error) {
	return func(storeType string, secretRef string, secretKeyName string) (map[string]interface{}, error) {
		return oidcProperties, nil
	}
}

=======
=======
>>>>>>> 668d47b8e (feat: Add TLS support to the Operator (#4796))
func emptyMockExtractConfigFromSecret(secretRef string, secretKeyName string) (map[string]interface{}, error) {
=======
func emptyMockExtractConfigFromSecret(storeType string, secretRef string, secretKeyName string) (map[string]interface{}, error) {
>>>>>>> 4b8378c2a (fix: Made fixes to Go Operator DB persistence (#4830))
	return map[string]interface{}{}, nil
}

func mockExtractConfigFromSecret(storeType string, secretRef string, secretKeyName string) (map[string]interface{}, error) {
	return createParameterMap(), nil
}

<<<<<<< HEAD
>>>>>>> 863a82cb7 (feat: Added feast Go operator db stores support (#4771))
=======
func mockOidcConfigFromSecret(
	oidcProperties map[string]interface{}) func(storeType string, secretRef string, secretKeyName string) (map[string]interface{}, error) {
	return func(storeType string, secretRef string, secretKeyName string) (map[string]interface{}, error) {
		return oidcProperties, nil
	}
}

>>>>>>> cd341f8f6 (feat: OIDC authorization in Feast Operator (#4801))
func createParameterMap() map[string]interface{} {
	yamlString := `
hosts:
  - 192.168.1.1
  - 192.168.1.2
  - 192.168.1.3
keyspace: KeyspaceName
port: 9042                                                              
username: user                                                          
password: secret                                                        
protocol_version: 5                                                     
load_balancing:                                                         
  local_dc: datacenter1                                             
  load_balancing_policy: TokenAwarePolicy(DCAwareRoundRobinPolicy)
read_concurrency: 100                                                   
write_concurrency: 100
`
	var parameters map[string]interface{}

	err := yaml.Unmarshal([]byte(yamlString), &parameters)
	if err != nil {
		fmt.Println(err)
	}
	return parameters
}
