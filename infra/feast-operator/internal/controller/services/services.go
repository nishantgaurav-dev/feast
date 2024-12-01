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
	"errors"
	"strconv"
	"strings"

	feastdevv1alpha1 "github.com/feast-dev/feast/infra/feast-operator/api/v1alpha1"
	"github.com/feast-dev/feast/infra/feast-operator/internal/controller/handler"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	apimeta "k8s.io/apimachinery/pkg/api/meta"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/util/intstr"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
	"sigs.k8s.io/controller-runtime/pkg/log"
)

// Apply defaults and set service hostnames in FeatureStore status
func (feast *FeastServices) ApplyDefaults() error {
	ApplyDefaultsToStatus(feast.Handler.FeatureStore)
	if err := feast.setTlsDefaults(); err != nil {
		return err
	}
	if err := feast.setServiceHostnames(); err != nil {
		return err
	}
	return nil
}

<<<<<<< HEAD
<<<<<<< HEAD
// Deploy the feast services
func (feast *FeastServices) Deploy() error {
	openshiftTls, err := feast.checkOpenshiftTls()
	if err != nil {
		return err
	}
	if openshiftTls {
		if err := feast.createCaConfigMap(); err != nil {
			return err
		}
	} else {
		_ = feast.Handler.DeleteOwnedFeastObj(feast.initCaConfigMap())
	}

	services := feast.Handler.FeatureStore.Status.Applied.Services
	if feast.isOfflinStore() {
		err := feast.validateOfflineStorePersistence(services.OfflineStore.Persistence)
		if err != nil {
			return err
		}

		if err = feast.deployFeastServiceByType(OfflineFeastType); err != nil {
			return err
		}
	} else {
		if err := feast.removeFeastServiceByType(OfflineFeastType); err != nil {
			return err
=======
	services := feast.FeatureStore.Status.Applied.Services
=======
	services := feast.Handler.FeatureStore.Status.Applied.Services
>>>>>>> 39eb4d80c (feat: RBAC Authorization in Feast Operator (#4786))
	if services != nil {
		if services.OfflineStore != nil {
			offlinePersistence := services.OfflineStore.Persistence

			err := feast.validateOfflineStorePersistence(offlinePersistence)
			if err != nil {
				return err
			}

			if err = feast.deployFeastServiceByType(OfflineFeastType); err != nil {
				return err
			}
		} else {
			if err := feast.removeFeastServiceByType(OfflineFeastType); err != nil {
				return err
			}
		}

		if services.OnlineStore != nil {
			onlinePersistence := services.OnlineStore.Persistence

			err := feast.validateOnlineStorePersistence(onlinePersistence)
			if err != nil {
				return err
			}

			if err = feast.deployFeastServiceByType(OnlineFeastType); err != nil {
				return err
			}
		} else {
			if err := feast.removeFeastServiceByType(OnlineFeastType); err != nil {
				return err
			}
		}

		if feast.isLocalRegistry() {
			registryPersistence := services.Registry.Local.Persistence

			err := feast.validateRegistryPersistence(registryPersistence)
			if err != nil {
				return err
			}

			if err = feast.deployFeastServiceByType(RegistryFeastType); err != nil {
				return err
			}
		} else {
			if err := feast.removeFeastServiceByType(RegistryFeastType); err != nil {
				return err
			}
>>>>>>> 863a82cb7 (feat: Added feast Go operator db stores support (#4771))
		}
	}

	if feast.isOnlinStore() {
		err := feast.validateOnlineStorePersistence(services.OnlineStore.Persistence)
		if err != nil {
			return err
		}

		if err = feast.deployFeastServiceByType(OnlineFeastType); err != nil {
			return err
		}
	} else {
		if err := feast.removeFeastServiceByType(OnlineFeastType); err != nil {
			return err
		}
	}

	if feast.isLocalRegistry() {
		err := feast.validateRegistryPersistence(services.Registry.Local.Persistence)
		if err != nil {
			return err
		}

		if err = feast.deployFeastServiceByType(RegistryFeastType); err != nil {
			return err
		}
	} else {
		if err := feast.removeFeastServiceByType(RegistryFeastType); err != nil {
			return err
		}
	}
	if err := feast.createServiceAccount(); err != nil {
		return err
	}
	if err := feast.createDeployment(); err != nil {
		return err
	}
	if err := feast.deployClient(); err != nil {
		return err
	}

	return nil
}

<<<<<<< HEAD
<<<<<<< HEAD
=======
>>>>>>> 863a82cb7 (feat: Added feast Go operator db stores support (#4771))
func (feast *FeastServices) validateRegistryPersistence(registryPersistence *feastdevv1alpha1.RegistryPersistence) error {
	if registryPersistence != nil {
		dbPersistence := registryPersistence.DBPersistence

		if dbPersistence != nil && len(dbPersistence.Type) > 0 {
			if err := checkRegistryDBStorePersistenceType(dbPersistence.Type); err != nil {
				return err
			}

<<<<<<< HEAD
<<<<<<< HEAD
			if len(dbPersistence.SecretRef.Name) > 0 {
=======
			if dbPersistence.SecretRef != nil {
>>>>>>> 863a82cb7 (feat: Added feast Go operator db stores support (#4771))
=======
			if len(dbPersistence.SecretRef.Name) > 0 {
>>>>>>> cac619cfa (fix: Fix db store types in Operator CRD (#4798))
				secretRef := dbPersistence.SecretRef.Name
				if _, err := feast.getSecret(secretRef); err != nil {
					return err
				}
			}
		}
	}

	return nil
}

func (feast *FeastServices) validateOnlineStorePersistence(onlinePersistence *feastdevv1alpha1.OnlineStorePersistence) error {
	if onlinePersistence != nil {
		dbPersistence := onlinePersistence.DBPersistence

		if dbPersistence != nil && len(dbPersistence.Type) > 0 {
			if err := checkOnlineStoreDBStorePersistenceType(dbPersistence.Type); err != nil {
				return err
			}

<<<<<<< HEAD
<<<<<<< HEAD
			if len(dbPersistence.SecretRef.Name) > 0 {
=======
			if dbPersistence.SecretRef != nil {
>>>>>>> 863a82cb7 (feat: Added feast Go operator db stores support (#4771))
=======
			if len(dbPersistence.SecretRef.Name) > 0 {
>>>>>>> cac619cfa (fix: Fix db store types in Operator CRD (#4798))
				secretRef := dbPersistence.SecretRef.Name
				if _, err := feast.getSecret(secretRef); err != nil {
					return err
				}
			}
		}
	}

	return nil
}

func (feast *FeastServices) validateOfflineStorePersistence(offlinePersistence *feastdevv1alpha1.OfflineStorePersistence) error {
	if offlinePersistence != nil {
		filePersistence := offlinePersistence.FilePersistence
		dbPersistence := offlinePersistence.DBPersistence

		if filePersistence != nil && len(filePersistence.Type) > 0 {
			if err := checkOfflineStoreFilePersistenceType(filePersistence.Type); err != nil {
				return err
			}
		} else if dbPersistence != nil &&
			len(dbPersistence.Type) > 0 {
			if err := checkOfflineStoreDBStorePersistenceType(dbPersistence.Type); err != nil {
				return err
			}

<<<<<<< HEAD
<<<<<<< HEAD
			if len(dbPersistence.SecretRef.Name) > 0 {
=======
			if dbPersistence.SecretRef != nil {
>>>>>>> 863a82cb7 (feat: Added feast Go operator db stores support (#4771))
=======
			if len(dbPersistence.SecretRef.Name) > 0 {
>>>>>>> cac619cfa (fix: Fix db store types in Operator CRD (#4798))
				secretRef := dbPersistence.SecretRef.Name
				if _, err := feast.getSecret(secretRef); err != nil {
					return err
				}
			}
		}
	}

	return nil
}

func (feast *FeastServices) deployFeastServiceByType(feastType FeastServiceType) error {
	if pvcCreate, shouldCreate := shouldCreatePvc(feast.Handler.FeatureStore, feastType); shouldCreate {
<<<<<<< HEAD
		if err := feast.createPVC(pvcCreate, feastType); err != nil {
			return feast.setFeastServiceCondition(err, feastType)
		}
	} else {
		_ = feast.Handler.DeleteOwnedFeastObj(feast.initPVC(feastType))
	}
	if err := feast.createService(feastType); err != nil {
		return feast.setFeastServiceCondition(err, feastType)
	}
=======
func (feast *FeastServices) deployFeastServiceByType(feastType FeastServiceType) error {
>>>>>>> 48e3c47c2 (feat: Operator will create k8s serviceaccount for each feast service (#4767))
	if pvcCreate, shouldCreate := shouldCreatePvc(feast.FeatureStore, feastType); shouldCreate {
=======
>>>>>>> 39eb4d80c (feat: RBAC Authorization in Feast Operator (#4786))
		if err := feast.createPVC(pvcCreate, feastType); err != nil {
			return feast.setFeastServiceCondition(err, feastType)
		}
	} else {
		_ = feast.Handler.DeleteOwnedFeastObj(feast.initPVC(feastType))
	}
	if err := feast.createServiceAccount(feastType); err != nil {
		return feast.setFeastServiceCondition(err, feastType)
	}
	if err := feast.createDeployment(feastType); err != nil {
		return feast.setFeastServiceCondition(err, feastType)
	}
	if err := feast.createService(feastType); err != nil {
		return feast.setFeastServiceCondition(err, feastType)
	}
	return feast.setFeastServiceCondition(nil, feastType)
}

func (feast *FeastServices) removeFeastServiceByType(feastType FeastServiceType) error {
<<<<<<< HEAD
<<<<<<< HEAD
	if err := feast.Handler.DeleteOwnedFeastObj(feast.initFeastSvc(feastType)); err != nil {
		return err
	}
	if err := feast.Handler.DeleteOwnedFeastObj(feast.initPVC(feastType)); err != nil {
=======
	if err := feast.deleteOwnedFeastObj(feast.initFeastSvc(feastType)); err != nil {
=======
	if err := feast.Handler.DeleteOwnedFeastObj(feast.initFeastSvc(feastType)); err != nil {
>>>>>>> 39eb4d80c (feat: RBAC Authorization in Feast Operator (#4786))
		return err
	}
	if err := feast.Handler.DeleteOwnedFeastObj(feast.initFeastDeploy(feastType)); err != nil {
		return err
	}
<<<<<<< HEAD
	if err := feast.deleteOwnedFeastObj(feast.initFeastSA(feastType)); err != nil {
>>>>>>> 48e3c47c2 (feat: Operator will create k8s serviceaccount for each feast service (#4767))
		return err
	}
<<<<<<< HEAD
	apimeta.RemoveStatusCondition(&feast.Handler.FeatureStore.Status.Conditions, FeastServiceConditions[feastType][metav1.ConditionTrue].Type)
=======
	if err := feast.deleteOwnedFeastObj(feast.initPVC(feastType)); err != nil {
		return err
	}
	apimeta.RemoveStatusCondition(&feast.FeatureStore.Status.Conditions, FeastServiceConditions[feastType][metav1.ConditionTrue].Type)
>>>>>>> 6c1a66ea8 (feat: PVC configuration and impl (#4750))
=======
	if err := feast.Handler.DeleteOwnedFeastObj(feast.initFeastSA(feastType)); err != nil {
		return err
	}
	if err := feast.Handler.DeleteOwnedFeastObj(feast.initPVC(feastType)); err != nil {
		return err
	}
	apimeta.RemoveStatusCondition(&feast.Handler.FeatureStore.Status.Conditions, FeastServiceConditions[feastType][metav1.ConditionTrue].Type)
>>>>>>> 39eb4d80c (feat: RBAC Authorization in Feast Operator (#4786))
	return nil
}

func (feast *FeastServices) createService(feastType FeastServiceType) error {
	logger := log.FromContext(feast.Handler.Context)
	svc := feast.initFeastSvc(feastType)
	if op, err := controllerutil.CreateOrUpdate(feast.Handler.Context, feast.Handler.Client, svc, controllerutil.MutateFn(func() error {
		return feast.setService(svc, feastType)
	})); err != nil {
		return err
	} else if op == controllerutil.OperationResultCreated || op == controllerutil.OperationResultUpdated {
		logger.Info("Successfully reconciled", "Service", svc.Name, "operation", op)
	}
	return nil
}

<<<<<<< HEAD
func (feast *FeastServices) createServiceAccount() error {
	logger := log.FromContext(feast.Handler.Context)
	sa := feast.initFeastSA()
	if op, err := controllerutil.CreateOrUpdate(feast.Handler.Context, feast.Handler.Client, sa, controllerutil.MutateFn(func() error {
		return feast.setServiceAccount(sa)
=======
func (feast *FeastServices) createServiceAccount(feastType FeastServiceType) error {
	logger := log.FromContext(feast.Handler.Context)
	sa := feast.initFeastSA(feastType)
	if op, err := controllerutil.CreateOrUpdate(feast.Handler.Context, feast.Handler.Client, sa, controllerutil.MutateFn(func() error {
		return feast.setServiceAccount(sa, feastType)
>>>>>>> 48e3c47c2 (feat: Operator will create k8s serviceaccount for each feast service (#4767))
	})); err != nil {
		return err
	} else if op == controllerutil.OperationResultCreated || op == controllerutil.OperationResultUpdated {
		logger.Info("Successfully reconciled", "ServiceAccount", sa.Name, "operation", op)
	}
	return nil
}

<<<<<<< HEAD
func (feast *FeastServices) createDeployment() error {
	logger := log.FromContext(feast.Handler.Context)
	deploy := feast.initFeastDeploy()
	if op, err := controllerutil.CreateOrUpdate(feast.Handler.Context, feast.Handler.Client, deploy, controllerutil.MutateFn(func() error {
		return feast.setDeployment(deploy)
=======
func (feast *FeastServices) createDeployment(feastType FeastServiceType) error {
	logger := log.FromContext(feast.Handler.Context)
	deploy := feast.initFeastDeploy(feastType)
	if op, err := controllerutil.CreateOrUpdate(feast.Handler.Context, feast.Handler.Client, deploy, controllerutil.MutateFn(func() error {
		return feast.setDeployment(deploy, feastType)
>>>>>>> 48e3c47c2 (feat: Operator will create k8s serviceaccount for each feast service (#4767))
	})); err != nil {
		return err
	} else if op == controllerutil.OperationResultCreated || op == controllerutil.OperationResultUpdated {
		logger.Info("Successfully reconciled", "Deployment", deploy.Name, "operation", op)
	}

	return nil
}

func (feast *FeastServices) createPVC(pvcCreate *feastdevv1alpha1.PvcCreate, feastType FeastServiceType) error {
<<<<<<< HEAD
<<<<<<< HEAD
	logger := log.FromContext(feast.Handler.Context)
	pvc, err := feast.createNewPVC(pvcCreate, feastType)
=======
	logger := log.FromContext(feast.Context)
=======
	logger := log.FromContext(feast.Handler.Context)
>>>>>>> 39eb4d80c (feat: RBAC Authorization in Feast Operator (#4786))
	pvc, err := feast.createNewPVC(pvcCreate, feastType)
	if err != nil {
		return err
	}

	err = feast.Handler.Client.Get(feast.Handler.Context, client.ObjectKeyFromObject(pvc), pvc)
	if err != nil && apierrors.IsNotFound(err) {
		err = feast.Handler.Client.Create(feast.Handler.Context, pvc)
		if err != nil {
			return err
		}
		logger.Info("Successfully created", "PersistentVolumeClaim", pvc.Name)
	}

	return nil
}

func (feast *FeastServices) setDeployment(deploy *appsv1.Deployment, feastType FeastServiceType) error {
	fsYamlB64, err := feast.GetServiceFeatureStoreYamlBase64(feastType)
>>>>>>> 6c1a66ea8 (feat: PVC configuration and impl (#4750))
	if err != nil {
		return err
	}
<<<<<<< HEAD
=======
	deploy.Labels = feast.getLabels(feastType)
	deploySettings := FeastServiceConstants[feastType]
	serviceConfigs := feast.getServiceConfigs(feastType)
	defaultServiceConfigs := serviceConfigs.DefaultConfigs
	sa := feast.initFeastSA(feastType)
>>>>>>> 48e3c47c2 (feat: Operator will create k8s serviceaccount for each feast service (#4767))

	// PVCs are immutable, so we only create... we don't update an existing one.
	err = feast.Handler.Client.Get(feast.Handler.Context, client.ObjectKeyFromObject(pvc), pvc)
	if err != nil && apierrors.IsNotFound(err) {
		err = feast.Handler.Client.Create(feast.Handler.Context, pvc)
		if err != nil {
			return err
		}
		logger.Info("Successfully created", "PersistentVolumeClaim", pvc.Name)
	}

	return nil
}

func (feast *FeastServices) setDeployment(deploy *appsv1.Deployment) error {
	deploy.Labels = feast.getLabels()
	deploy.Spec = appsv1.DeploymentSpec{
		Replicas: &DefaultReplicas,
		Selector: metav1.SetAsLabelSelector(deploy.GetLabels()),
		Strategy: appsv1.DeploymentStrategy{
			Type: appsv1.RecreateDeploymentStrategyType,
		},
		Template: corev1.PodTemplateSpec{
			ObjectMeta: metav1.ObjectMeta{
				Labels: deploy.GetLabels(),
			},
			Spec: corev1.PodSpec{
<<<<<<< HEAD
				ServiceAccountName: feast.initFeastSA().Name,
=======
				ServiceAccountName: sa.Name,
				Containers: []corev1.Container{
					{
						Name:    string(feastType),
						Image:   *defaultServiceConfigs.Image,
						Command: deploySettings.Command,
						Ports: []corev1.ContainerPort{
							{
								Name:          string(feastType),
								ContainerPort: deploySettings.TargetPort,
								Protocol:      corev1.ProtocolTCP,
							},
						},
						Env: []corev1.EnvVar{
							{
								Name:  FeatureStoreYamlEnvVar,
								Value: fsYamlB64,
							},
						},
						LivenessProbe: &corev1.Probe{
							ProbeHandler:        probeHandler,
							InitialDelaySeconds: 30,
							PeriodSeconds:       30,
						},
						ReadinessProbe: &corev1.Probe{
							ProbeHandler:        probeHandler,
							InitialDelaySeconds: 20,
							PeriodSeconds:       10,
						},
					},
				},
>>>>>>> 48e3c47c2 (feat: Operator will create k8s serviceaccount for each feast service (#4767))
			},
		},
	}
<<<<<<< HEAD
	if err := feast.setPod(&deploy.Spec.Template.Spec); err != nil {
		return err
	}
	return controllerutil.SetControllerReference(feast.Handler.FeatureStore, deploy, feast.Handler.Scheme)
=======

	// configs are applied here
	container := &deploy.Spec.Template.Spec.Containers[0]
	applyOptionalContainerConfigs(container, serviceConfigs.OptionalConfigs)
	if pvcConfig, hasPvcConfig := hasPvcConfig(feast.Handler.FeatureStore, feastType); hasPvcConfig {
		mountPvcConfig(&deploy.Spec.Template.Spec, pvcConfig, deploy.Name)
	}

<<<<<<< HEAD
	return controllerutil.SetControllerReference(feast.FeatureStore, deploy, feast.Scheme)
>>>>>>> 6c1a66ea8 (feat: PVC configuration and impl (#4750))
=======
	return controllerutil.SetControllerReference(feast.Handler.FeatureStore, deploy, feast.Handler.Scheme)
>>>>>>> 39eb4d80c (feat: RBAC Authorization in Feast Operator (#4786))
}

func (feast *FeastServices) setPod(podSpec *corev1.PodSpec) error {
	if err := feast.setContainers(podSpec); err != nil {
		return err
	}
	feast.setRegistryClientInitContainer(podSpec)
	feast.mountTlsConfigs(podSpec)
	feast.mountPvcConfigs(podSpec)
	feast.mountEmptyDirVolumes(podSpec)

	return nil
}

func (feast *FeastServices) setContainers(podSpec *corev1.PodSpec) error {
	fsYamlB64, err := feast.GetServiceFeatureStoreYamlBase64()
	if err != nil {
		return err
	}
	feastProject := feast.Handler.FeatureStore.Status.Applied.FeastProject
	workingDir := getOfflineMountPath(feast.Handler.FeatureStore)
	podSpec.InitContainers = append(podSpec.InitContainers, corev1.Container{
		Name:  "feast-init",
		Image: DefaultImage,
		Env: []corev1.EnvVar{
			{
				Name:  TmpFeatureStoreYamlEnvVar,
				Value: fsYamlB64,
			},
		},
		Command: []string{"/bin/sh", "-c"},
		Args: []string{"echo \"Starting feast initialization job...\";\n[ -d " +
			feastProject + " ] || feast init " + feastProject + ";\necho $" +
			TmpFeatureStoreYamlEnvVar + " | base64 -d \u003e " + workingDir + "/" + feastProject +
			"/feature_repo/feature_store.yaml;\necho \"Feast initialization complete\";\n"},
		WorkingDir: workingDir,
	})
	if feast.isLocalRegistry() {
		feast.setContainer(&podSpec.Containers, RegistryFeastType, fsYamlB64)
	}
	if feast.isOfflinStore() {
		feast.setContainer(&podSpec.Containers, OfflineFeastType, fsYamlB64)
	}
	if feast.isOnlinStore() {
		feast.setContainer(&podSpec.Containers, OnlineFeastType, fsYamlB64)
	}
	return nil
}

func (feast *FeastServices) setContainer(containers *[]corev1.Container, feastType FeastServiceType, fsYamlB64 string) {
	tls := feast.getTlsConfigs(feastType)
	serviceConfigs := feast.getServiceConfigs(feastType)
	defaultServiceConfigs := serviceConfigs.DefaultConfigs
	probeHandler := getProbeHandler(feastType, tls)
	container := &corev1.Container{
		Name:       string(feastType),
		Image:      *defaultServiceConfigs.Image,
		WorkingDir: getOfflineMountPath(feast.Handler.FeatureStore) + "/" + feast.Handler.FeatureStore.Status.Applied.FeastProject + "/feature_repo",
		Command:    feast.getContainerCommand(feastType),
		Ports: []corev1.ContainerPort{
			{
				Name:          string(feastType),
				ContainerPort: getTargetPort(feastType, tls),
				Protocol:      corev1.ProtocolTCP,
			},
		},
		Env: []corev1.EnvVar{
			{
				Name:  TmpFeatureStoreYamlEnvVar,
				Value: fsYamlB64,
			},
			/*
				{
					Name:  mlpConfigVar,
					Value: DefaultMlpConfigPath,
				},
			*/
		},
		StartupProbe: &corev1.Probe{
			ProbeHandler:     probeHandler,
			PeriodSeconds:    3,
			FailureThreshold: 40,
		},
		LivenessProbe: &corev1.Probe{
			ProbeHandler:     probeHandler,
			PeriodSeconds:    20,
			FailureThreshold: 6,
		},
		ReadinessProbe: &corev1.Probe{
			ProbeHandler:  probeHandler,
			PeriodSeconds: 10,
		},
	}
	applyOptionalContainerConfigs(container, serviceConfigs.OptionalConfigs)
	*containers = append(*containers, *container)
}

func (feast *FeastServices) getContainerCommand(feastType FeastServiceType) []string {
	baseCommand := "feast"
	options := []string{}
	logLevel := feast.getLogLevelForType(feastType)
	if logLevel != nil {
		options = append(options, "--log-level", strings.ToUpper(*logLevel))
	}

	deploySettings := FeastServiceConstants[feastType]
	targetPort := deploySettings.TargetHttpPort
	tls := feast.getTlsConfigs(feastType)
	if tls.IsTLS() {
		targetPort = deploySettings.TargetHttpsPort
		feastTlsPath := GetTlsPath(feastType)
		deploySettings.Args = append(deploySettings.Args, []string{"--key", feastTlsPath + tls.SecretKeyNames.TlsKey,
			"--cert", feastTlsPath + tls.SecretKeyNames.TlsCrt}...)
	}
	deploySettings.Args = append(deploySettings.Args, []string{"-p", strconv.Itoa(int(targetPort))}...)

	// Combine base command, options, and arguments
	feastCommand := append([]string{baseCommand}, options...)
	feastCommand = append(feastCommand, deploySettings.Args...)

	return feastCommand
}

func (feast *FeastServices) setRegistryClientInitContainer(podSpec *corev1.PodSpec) {
	hostname := feast.Handler.FeatureStore.Status.ServiceHostnames.Registry
	// add grpc init container if remote registry reference (feastRef) is configured
	if len(hostname) > 0 && feast.IsRemoteRefRegistry() {
		grpcurlFlag := "-plaintext"
		hostSplit := strings.Split(hostname, ":")
		if len(hostSplit) > 1 && hostSplit[1] == "443" {
			grpcurlFlag = "-insecure"
		}
		podSpec.InitContainers = append(podSpec.InitContainers, corev1.Container{
			Name:  "init-registry",
			Image: "fullstorydev/grpcurl:v1.9.1-alpine",
			Command: []string{
				"sh", "-c",
				"until grpcurl -H \"authorization: Bearer $(cat /var/run/secrets/kubernetes.io/serviceaccount/token)\" " +
					grpcurlFlag + " -d '' -format text " + hostname + " grpc.health.v1.Health/Check; do echo waiting for registry; sleep 2; done",
			},
		})
	}
}
func (feast *FeastServices) setService(svc *corev1.Service, feastType FeastServiceType) error {
	svc.Labels = feast.getFeastTypeLabels(feastType)
	if feast.isOpenShiftTls(feastType) {
		svc.Annotations = map[string]string{
			"service.beta.openshift.io/serving-cert-secret-name": svc.Name + tlsNameSuffix,
		}
	}

	var port int32 = HttpPort
	scheme := HttpScheme
	tls := feast.getTlsConfigs(feastType)
	if tls.IsTLS() {
		port = HttpsPort
		scheme = HttpsScheme
	}
	svc.Spec = corev1.ServiceSpec{
		Selector: feast.getLabels(),
		Type:     corev1.ServiceTypeClusterIP,
		Ports: []corev1.ServicePort{
			{
				Name:       scheme,
				Port:       port,
				Protocol:   corev1.ProtocolTCP,
				TargetPort: intstr.FromInt(int(getTargetPort(feastType, tls))),
			},
		},
	}

	return controllerutil.SetControllerReference(feast.Handler.FeatureStore, svc, feast.Handler.Scheme)
<<<<<<< HEAD
}

func (feast *FeastServices) setServiceAccount(sa *corev1.ServiceAccount) error {
	sa.Labels = feast.getLabels()
	return controllerutil.SetControllerReference(feast.Handler.FeatureStore, sa, feast.Handler.Scheme)
}

func (feast *FeastServices) createNewPVC(pvcCreate *feastdevv1alpha1.PvcCreate, feastType FeastServiceType) (*corev1.PersistentVolumeClaim, error) {
	pvc := feast.initPVC(feastType)

	pvc.Spec = corev1.PersistentVolumeClaimSpec{
		AccessModes: pvcCreate.AccessModes,
		Resources:   pvcCreate.Resources,
	}
	if pvcCreate.StorageClassName != nil {
		pvc.Spec.StorageClassName = pvcCreate.StorageClassName
	}
	return pvc, controllerutil.SetControllerReference(feast.Handler.FeatureStore, pvc, feast.Handler.Scheme)
=======
>>>>>>> 39eb4d80c (feat: RBAC Authorization in Feast Operator (#4786))
}

func (feast *FeastServices) setServiceAccount(sa *corev1.ServiceAccount, feastType FeastServiceType) error {
	sa.Labels = feast.getLabels(feastType)
	return controllerutil.SetControllerReference(feast.Handler.FeatureStore, sa, feast.Handler.Scheme)
}

func (feast *FeastServices) createNewPVC(pvcCreate *feastdevv1alpha1.PvcCreate, feastType FeastServiceType) (*corev1.PersistentVolumeClaim, error) {
	pvc := feast.initPVC(feastType)

	pvc.Spec = corev1.PersistentVolumeClaimSpec{
		AccessModes: []corev1.PersistentVolumeAccessMode{corev1.ReadWriteMany},
		Resources:   pvcCreate.Resources,
	}
	if pvcCreate.StorageClassName != nil {
		pvc.Spec.StorageClassName = pvcCreate.StorageClassName
	}
	return pvc, controllerutil.SetControllerReference(feast.Handler.FeatureStore, pvc, feast.Handler.Scheme)
}

func (feast *FeastServices) getServiceConfigs(feastType FeastServiceType) feastdevv1alpha1.ServiceConfigs {
<<<<<<< HEAD
<<<<<<< HEAD
	appliedServices := feast.Handler.FeatureStore.Status.Applied.Services
	switch feastType {
	case OfflineFeastType:
		if feast.isOfflinStore() {
			return appliedServices.OfflineStore.ServiceConfigs
		}
	case OnlineFeastType:
		if feast.isOnlinStore() {
			return appliedServices.OnlineStore.ServiceConfigs
		}
	case RegistryFeastType:
		if feast.isLocalRegistry() {
			return appliedServices.Registry.Local.ServiceConfigs
=======
	appliedSpec := feast.FeatureStore.Status.Applied
=======
	appliedSpec := feast.Handler.FeatureStore.Status.Applied
>>>>>>> 39eb4d80c (feat: RBAC Authorization in Feast Operator (#4786))
	switch feastType {
	case OfflineFeastType:
		if appliedSpec.Services.OfflineStore != nil {
			return appliedSpec.Services.OfflineStore.ServiceConfigs
		}
	case OnlineFeastType:
		if appliedSpec.Services.OnlineStore != nil {
			return appliedSpec.Services.OnlineStore.ServiceConfigs
		}
	case RegistryFeastType:
		if appliedSpec.Services.Registry != nil {
			if appliedSpec.Services.Registry.Local != nil {
				return appliedSpec.Services.Registry.Local.ServiceConfigs
			}
>>>>>>> 6c1a66ea8 (feat: PVC configuration and impl (#4750))
		}
	}
	return feastdevv1alpha1.ServiceConfigs{}
}

<<<<<<< HEAD
func (feast *FeastServices) getLogLevelForType(feastType FeastServiceType) *string {
	services := feast.Handler.FeatureStore.Status.Applied.Services
	switch feastType {
	case OfflineFeastType:
		if services.OfflineStore != nil && services.OfflineStore.LogLevel != "" {
			return &services.OfflineStore.LogLevel
		}
	case OnlineFeastType:
		if services.OnlineStore != nil && services.OnlineStore.LogLevel != "" {
			return &services.OnlineStore.LogLevel
		}
	case RegistryFeastType:
		if services.Registry != nil && services.Registry.Local.LogLevel != "" {
			return &services.Registry.Local.LogLevel
		}
	}
	return nil
}

// GetObjectMeta returns the feast k8s object metadata with type
func (feast *FeastServices) GetObjectMeta() metav1.ObjectMeta {
	return metav1.ObjectMeta{Name: GetFeastName(feast.Handler.FeatureStore), Namespace: feast.Handler.FeatureStore.Namespace}
}

// GetObjectMeta returns the feast k8s object metadata with type
func (feast *FeastServices) GetObjectMetaType(feastType FeastServiceType) metav1.ObjectMeta {
=======
// GetObjectMeta returns the feast k8s object metadata
func (feast *FeastServices) GetObjectMeta(feastType FeastServiceType) metav1.ObjectMeta {
>>>>>>> 39eb4d80c (feat: RBAC Authorization in Feast Operator (#4786))
	return metav1.ObjectMeta{Name: feast.GetFeastServiceName(feastType), Namespace: feast.Handler.FeatureStore.Namespace}
}

func (feast *FeastServices) GetFeastServiceName(feastType FeastServiceType) string {
	return GetFeastServiceName(feast.Handler.FeatureStore, feastType)
}

// GetFeastServiceName returns the feast service object name based on service type
func GetFeastServiceName(featureStore *feastdevv1alpha1.FeatureStore, feastType FeastServiceType) string {
	return GetFeastName(featureStore) + "-" + string(feastType)
}

func GetFeastName(featureStore *feastdevv1alpha1.FeatureStore) string {
	return handler.FeastPrefix + featureStore.Name
}

func (feast *FeastServices) getFeastTypeLabels(feastType FeastServiceType) map[string]string {
	labels := feast.getLabels()
	labels[ServiceTypeLabelKey] = string(feastType)
	return labels
}

func (feast *FeastServices) getLabels() map[string]string {
	return map[string]string{
<<<<<<< HEAD
		NameLabelKey: feast.Handler.FeatureStore.Name,
=======
		NameLabelKey:        feast.Handler.FeatureStore.Name,
		ServiceTypeLabelKey: string(feastType),
>>>>>>> 39eb4d80c (feat: RBAC Authorization in Feast Operator (#4786))
	}
}

func (feast *FeastServices) setServiceHostnames() error {
	feast.Handler.FeatureStore.Status.ServiceHostnames = feastdevv1alpha1.ServiceHostnames{}
<<<<<<< HEAD
	domain := svcDomain + ":"
	if feast.isOfflinStore() {
		objMeta := feast.initFeastSvc(OfflineFeastType)
		feast.Handler.FeatureStore.Status.ServiceHostnames.OfflineStore = objMeta.Name + "." + objMeta.Namespace + domain +
			getPortStr(feast.Handler.FeatureStore.Status.Applied.Services.OfflineStore.TLS)
	}
	if feast.isOnlinStore() {
		objMeta := feast.initFeastSvc(OnlineFeastType)
		feast.Handler.FeatureStore.Status.ServiceHostnames.OnlineStore = objMeta.Name + "." + objMeta.Namespace + domain +
			getPortStr(feast.Handler.FeatureStore.Status.Applied.Services.OnlineStore.TLS)
	}
	if feast.isLocalRegistry() {
		objMeta := feast.initFeastSvc(RegistryFeastType)
		feast.Handler.FeatureStore.Status.ServiceHostnames.Registry = objMeta.Name + "." + objMeta.Namespace + domain +
			getPortStr(feast.Handler.FeatureStore.Status.Applied.Services.Registry.Local.TLS)
	} else if feast.isRemoteRegistry() {
		return feast.setRemoteRegistryURL()
=======
	services := feast.Handler.FeatureStore.Status.Applied.Services
	if services != nil {
		domain := svcDomain + ":" + strconv.Itoa(HttpPort)
		if services.OfflineStore != nil {
			objMeta := feast.GetObjectMeta(OfflineFeastType)
			feast.Handler.FeatureStore.Status.ServiceHostnames.OfflineStore = objMeta.Name + "." + objMeta.Namespace + domain
		}
		if services.OnlineStore != nil {
			objMeta := feast.GetObjectMeta(OnlineFeastType)
			feast.Handler.FeatureStore.Status.ServiceHostnames.OnlineStore = objMeta.Name + "." + objMeta.Namespace + domain
		}
		if feast.isLocalRegistry() {
			objMeta := feast.GetObjectMeta(RegistryFeastType)
			feast.Handler.FeatureStore.Status.ServiceHostnames.Registry = objMeta.Name + "." + objMeta.Namespace + domain
		} else if feast.isRemoteRegistry() {
			return feast.setRemoteRegistryURL()
		}
>>>>>>> 39eb4d80c (feat: RBAC Authorization in Feast Operator (#4786))
	}
	return nil
}

func (feast *FeastServices) setFeastServiceCondition(err error, feastType FeastServiceType) error {
	conditionMap := FeastServiceConditions[feastType]
	if err != nil {
		logger := log.FromContext(feast.Handler.Context)
		cond := conditionMap[metav1.ConditionFalse]
		cond.Message = "Error: " + err.Error()
		apimeta.SetStatusCondition(&feast.Handler.FeatureStore.Status.Conditions, cond)
		logger.Error(err, "Error deploying the FeatureStore "+string(ClientFeastType)+" service")
		return err
	} else {
		apimeta.SetStatusCondition(&feast.Handler.FeatureStore.Status.Conditions, conditionMap[metav1.ConditionTrue])
	}
	return nil
}

func (feast *FeastServices) setRemoteRegistryURL() error {
	if feast.isRemoteHostnameRegistry() {
		feast.Handler.FeatureStore.Status.ServiceHostnames.Registry = *feast.Handler.FeatureStore.Status.Applied.Services.Registry.Remote.Hostname
	} else if feast.IsRemoteRefRegistry() {
<<<<<<< HEAD
		remoteFeast, err := feast.getRemoteRegistryFeastHandler()
		if err != nil {
			return err
		}
		// referenced/remote registry must use the local install option and be in a 'Ready' state.
		if remoteFeast != nil &&
			remoteFeast.isLocalRegistry() &&
			apimeta.IsStatusConditionTrue(remoteFeast.Handler.FeatureStore.Status.Conditions, feastdevv1alpha1.RegistryReadyType) {
			feast.Handler.FeatureStore.Status.ServiceHostnames.Registry = remoteFeast.Handler.FeatureStore.Status.ServiceHostnames.Registry
=======
		feastRemoteRef := feast.Handler.FeatureStore.Status.Applied.Services.Registry.Remote.FeastRef
		// default to FeatureStore namespace if not set
		if len(feastRemoteRef.Namespace) == 0 {
			feastRemoteRef.Namespace = feast.Handler.FeatureStore.Namespace
		}

		nsName := types.NamespacedName{Name: feastRemoteRef.Name, Namespace: feastRemoteRef.Namespace}
		crNsName := client.ObjectKeyFromObject(feast.Handler.FeatureStore)
		if nsName == crNsName {
			return errors.New("FeatureStore '" + crNsName.Name + "' can't reference itself in `spec.services.registry.remote.feastRef`")
		}

		remoteFeastObj := &feastdevv1alpha1.FeatureStore{}
		if err := feast.Handler.Client.Get(feast.Handler.Context, nsName, remoteFeastObj); err != nil {
			if apierrors.IsNotFound(err) {
				return errors.New("Referenced FeatureStore '" + feastRemoteRef.Name + "' was not found")
			}
			return err
		}

		remoteFeast := FeastServices{
			Handler: handler.FeastHandler{
				Client:       feast.Handler.Client,
				Context:      feast.Handler.Context,
				FeatureStore: remoteFeastObj,
				Scheme:       feast.Handler.Scheme,
			},
		}
		// referenced/remote registry must use the local install option and be in a 'Ready' state.
		if remoteFeast.isLocalRegistry() && apimeta.IsStatusConditionTrue(remoteFeastObj.Status.Conditions, feastdevv1alpha1.RegistryReadyType) {
			feast.Handler.FeatureStore.Status.ServiceHostnames.Registry = remoteFeastObj.Status.ServiceHostnames.Registry
>>>>>>> 39eb4d80c (feat: RBAC Authorization in Feast Operator (#4786))
		} else {
			return errors.New("Remote feast registry of referenced FeatureStore '" + remoteFeast.Handler.FeatureStore.Name + "' is not ready")
		}
	}
	return nil
}

func (feast *FeastServices) getRemoteRegistryFeastHandler() (*FeastServices, error) {
	if feast.IsRemoteRefRegistry() {
		feastRemoteRef := feast.Handler.FeatureStore.Status.Applied.Services.Registry.Remote.FeastRef
		nsName := types.NamespacedName{Name: feastRemoteRef.Name, Namespace: feastRemoteRef.Namespace}
		crNsName := client.ObjectKeyFromObject(feast.Handler.FeatureStore)
		if nsName == crNsName {
			return nil, errors.New("FeatureStore '" + crNsName.Name + "' can't reference itself in `spec.services.registry.remote.feastRef`")
		}
		remoteFeastObj := &feastdevv1alpha1.FeatureStore{}
		if err := feast.Handler.Client.Get(feast.Handler.Context, nsName, remoteFeastObj); err != nil {
			if apierrors.IsNotFound(err) {
				return nil, errors.New("Referenced FeatureStore '" + feastRemoteRef.Name + "' was not found")
			}
			return nil, err
		}
		return &FeastServices{
			Handler: handler.FeastHandler{
				Client:       feast.Handler.Client,
				Context:      feast.Handler.Context,
				FeatureStore: remoteFeastObj,
				Scheme:       feast.Handler.Scheme,
			},
		}, nil
	}
	return nil, nil
}

func (feast *FeastServices) isLocalRegistry() bool {
	return IsLocalRegistry(feast.Handler.FeatureStore)
}

func (feast *FeastServices) isRemoteRegistry() bool {
<<<<<<< HEAD
	return isRemoteRegistry(feast.Handler.FeatureStore)
}

func (feast *FeastServices) IsRemoteRefRegistry() bool {
	return feast.isRemoteRegistry() &&
		feast.Handler.FeatureStore.Status.Applied.Services.Registry.Remote.FeastRef != nil
}

func (feast *FeastServices) isRemoteHostnameRegistry() bool {
	return feast.isRemoteRegistry() &&
		feast.Handler.FeatureStore.Status.Applied.Services.Registry.Remote.Hostname != nil
=======
	appliedServices := feast.Handler.FeatureStore.Status.Applied.Services
	return appliedServices != nil && appliedServices.Registry != nil && appliedServices.Registry.Remote != nil
}

func (feast *FeastServices) IsRemoteRefRegistry() bool {
	if feast.isRemoteRegistry() {
		remote := feast.Handler.FeatureStore.Status.Applied.Services.Registry.Remote
		return remote != nil && remote.FeastRef != nil
	}
	return false
}

func (feast *FeastServices) isRemoteHostnameRegistry() bool {
	if feast.isRemoteRegistry() {
		remote := feast.Handler.FeatureStore.Status.Applied.Services.Registry.Remote
		return remote != nil && remote.Hostname != nil
	}
	return false
>>>>>>> 39eb4d80c (feat: RBAC Authorization in Feast Operator (#4786))
}

func (feast *FeastServices) isOfflinStore() bool {
	appliedServices := feast.Handler.FeatureStore.Status.Applied.Services
	return appliedServices != nil && appliedServices.OfflineStore != nil
}

func (feast *FeastServices) isOnlinStore() bool {
	appliedServices := feast.Handler.FeatureStore.Status.Applied.Services
	return appliedServices != nil && appliedServices.OnlineStore != nil
}

func (feast *FeastServices) initFeastDeploy() *appsv1.Deployment {
	deploy := &appsv1.Deployment{
		ObjectMeta: feast.GetObjectMeta(),
	}
	deploy.SetGroupVersionKind(appsv1.SchemeGroupVersion.WithKind("Deployment"))
	return deploy
}

func (feast *FeastServices) initFeastSvc(feastType FeastServiceType) *corev1.Service {
	svc := &corev1.Service{
		ObjectMeta: feast.GetObjectMetaType(feastType),
	}
	svc.SetGroupVersionKind(corev1.SchemeGroupVersion.WithKind("Service"))
	return svc
}

<<<<<<< HEAD
<<<<<<< HEAD
func (feast *FeastServices) initFeastSA() *corev1.ServiceAccount {
	sa := &corev1.ServiceAccount{
		ObjectMeta: feast.GetObjectMeta(),
=======
func (feast *FeastServices) initFeastSA(feastType FeastServiceType) *corev1.ServiceAccount {
	sa := &corev1.ServiceAccount{
		ObjectMeta: feast.GetObjectMeta(feastType),
>>>>>>> 48e3c47c2 (feat: Operator will create k8s serviceaccount for each feast service (#4767))
	}
	sa.SetGroupVersionKind(corev1.SchemeGroupVersion.WithKind("ServiceAccount"))
	return sa
}

<<<<<<< HEAD
func (feast *FeastServices) initPVC(feastType FeastServiceType) *corev1.PersistentVolumeClaim {
	pvc := &corev1.PersistentVolumeClaim{
		ObjectMeta: feast.GetObjectMetaType(feastType),
=======
=======
>>>>>>> 48e3c47c2 (feat: Operator will create k8s serviceaccount for each feast service (#4767))
func (feast *FeastServices) initPVC(feastType FeastServiceType) *corev1.PersistentVolumeClaim {
	pvc := &corev1.PersistentVolumeClaim{
		ObjectMeta: feast.GetObjectMeta(feastType),
	}
	pvc.SetGroupVersionKind(corev1.SchemeGroupVersion.WithKind("PersistentVolumeClaim"))
	return pvc
}

<<<<<<< HEAD
// delete an object if the FeatureStore is set as the object's controller/owner
func (feast *FeastServices) deleteOwnedFeastObj(obj client.Object) error {
	name := obj.GetName()
	kind := obj.GetObjectKind().GroupVersionKind().Kind
	if err := feast.Client.Get(feast.Context, client.ObjectKeyFromObject(obj), obj); err != nil {
		if apierrors.IsNotFound(err) {
			return nil
		}
		return err
	}
	for _, ref := range obj.GetOwnerReferences() {
		if *ref.Controller && ref.UID == feast.FeatureStore.UID {
			if err := feast.Client.Delete(feast.Context, obj); err != nil {
				return err
			}
			log.FromContext(feast.Context).Info("Successfully deleted", kind, name)
		}
>>>>>>> 6c1a66ea8 (feat: PVC configuration and impl (#4750))
	}
	pvc.SetGroupVersionKind(corev1.SchemeGroupVersion.WithKind("PersistentVolumeClaim"))
	return pvc
}

=======
>>>>>>> 39eb4d80c (feat: RBAC Authorization in Feast Operator (#4786))
func applyOptionalContainerConfigs(container *corev1.Container, optionalConfigs feastdevv1alpha1.OptionalConfigs) {
	if optionalConfigs.Env != nil {
		container.Env = envOverride(container.Env, *optionalConfigs.Env)
	}
	if optionalConfigs.EnvFrom != nil {
		container.EnvFrom = *optionalConfigs.EnvFrom
	}
	if optionalConfigs.ImagePullPolicy != nil {
		container.ImagePullPolicy = *optionalConfigs.ImagePullPolicy
	}
	if optionalConfigs.Resources != nil {
		container.Resources = *optionalConfigs.Resources
	}
}

func (feast *FeastServices) mountPvcConfigs(podSpec *corev1.PodSpec) {
	for _, feastType := range feastServerTypes {
		if pvcConfig, hasPvcConfig := hasPvcConfig(feast.Handler.FeatureStore, feastType); hasPvcConfig {
			feast.mountPvcConfig(podSpec, pvcConfig, feastType)
		}
	}
}

func (feast *FeastServices) mountPvcConfig(podSpec *corev1.PodSpec, pvcConfig *feastdevv1alpha1.PvcConfig, feastType FeastServiceType) {
	if podSpec != nil && pvcConfig != nil {
		volName := feast.initPVC(feastType).Name
		pvcName := volName
		if pvcConfig.Ref != nil {
			pvcName = pvcConfig.Ref.Name
		}
		podSpec.Volumes = append(podSpec.Volumes, corev1.Volume{
			Name: volName,
			VolumeSource: corev1.VolumeSource{
				PersistentVolumeClaim: &corev1.PersistentVolumeClaimVolumeSource{
					ClaimName: pvcName,
				},
			},
		})
		if feastType == OfflineFeastType {
			for i := range podSpec.InitContainers {
				podSpec.InitContainers[i].VolumeMounts = append(podSpec.InitContainers[i].VolumeMounts, corev1.VolumeMount{
					Name:      volName,
					MountPath: pvcConfig.MountPath,
				})
			}
		}
		for i := range podSpec.Containers {
			podSpec.Containers[i].VolumeMounts = append(podSpec.Containers[i].VolumeMounts, corev1.VolumeMount{
				Name:      volName,
				MountPath: pvcConfig.MountPath,
			})
		}
	}
}

func (feast *FeastServices) mountEmptyDirVolumes(podSpec *corev1.PodSpec) {
	if shouldMountEmptyDir(feast.Handler.FeatureStore) {
		mountEmptyDirVolume(podSpec)
	}
}

func mountEmptyDirVolume(podSpec *corev1.PodSpec) {
	if podSpec != nil {
		volName := strings.TrimPrefix(EphemeralPath, "/")
		podSpec.Volumes = append(podSpec.Volumes, corev1.Volume{
			Name: volName,
			VolumeSource: corev1.VolumeSource{
				EmptyDir: &corev1.EmptyDirVolumeSource{},
			},
		})
		for i := range podSpec.InitContainers {
			podSpec.InitContainers[i].VolumeMounts = append(podSpec.InitContainers[i].VolumeMounts, corev1.VolumeMount{
				Name:      volName,
				MountPath: EphemeralPath,
			})
		}
		for i := range podSpec.Containers {
			podSpec.Containers[i].VolumeMounts = append(podSpec.Containers[i].VolumeMounts, corev1.VolumeMount{
				Name:      volName,
				MountPath: EphemeralPath,
			})
		}
	}
}

func getTargetPort(feastType FeastServiceType, tls *feastdevv1alpha1.TlsConfigs) int32 {
	if tls.IsTLS() {
		return FeastServiceConstants[feastType].TargetHttpsPort
	}
	return FeastServiceConstants[feastType].TargetHttpPort
}

func getProbeHandler(feastType FeastServiceType, tls *feastdevv1alpha1.TlsConfigs) corev1.ProbeHandler {
	targetPort := getTargetPort(feastType, tls)
	if feastType == OnlineFeastType {
		probeHandler := corev1.ProbeHandler{
			HTTPGet: &corev1.HTTPGetAction{
				Path: "/health",
				Port: intstr.FromInt(int(targetPort)),
			},
		}
		if tls.IsTLS() {
			probeHandler.HTTPGet.Scheme = corev1.URISchemeHTTPS
		}
		return probeHandler
	}
	return corev1.ProbeHandler{
		TCPSocket: &corev1.TCPSocketAction{
			Port: intstr.FromInt(int(targetPort)),
		},
	}
}

func mountPvcConfig(podSpec *corev1.PodSpec, pvcConfig *feastdevv1alpha1.PvcConfig, deployName string) {
	container := &podSpec.Containers[0]
	var pvcName string
	if pvcConfig.Create != nil {
		pvcName = deployName
	} else {
		pvcName = pvcConfig.Ref.Name
	}

	podSpec.Volumes = []corev1.Volume{
		{
			Name: pvcName,
			VolumeSource: corev1.VolumeSource{
				PersistentVolumeClaim: &corev1.PersistentVolumeClaimVolumeSource{
					ClaimName: pvcName,
				},
			},
		},
	}
	container.VolumeMounts = []corev1.VolumeMount{
		{
			Name:      pvcName,
			MountPath: pvcConfig.MountPath,
		},
	}
}
