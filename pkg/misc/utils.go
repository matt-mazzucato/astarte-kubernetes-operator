/*
  This file is part of Astarte.

  Copyright 2020 Ispirata Srl

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

package misc

import (
	"context"
	"fmt"

	apiv1alpha1 "github.com/astarte-platform/astarte-kubernetes-operator/pkg/apis/api/v1alpha1"
	"github.com/go-logr/logr"
	"github.com/openlyinc/pointy"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
)

const (
	// RabbitMQDefaultUserCredentialsUsernameKey is the default Username key for RabbitMQ Secret
	RabbitMQDefaultUserCredentialsUsernameKey = "admin-username"
	// RabbitMQDefaultUserCredentialsPasswordKey is the default Password key for RabbitMQ Secret
	RabbitMQDefaultUserCredentialsPasswordKey = "admin-password"
)

type allocationCoefficients struct {
	CPUCoefficient    float64
	MemoryCoefficient float64
}

var defaultComponentAllocations = map[apiv1alpha1.AstarteComponent]allocationCoefficients{
	apiv1alpha1.AppEngineAPI:       {CPUCoefficient: 0.19, MemoryCoefficient: 0.19},
	apiv1alpha1.DataUpdaterPlant:   {CPUCoefficient: 0.22, MemoryCoefficient: 0.22},
	apiv1alpha1.Housekeeping:       {CPUCoefficient: 0.05, MemoryCoefficient: 0.05},
	apiv1alpha1.HousekeepingAPI:    {CPUCoefficient: 0.05, MemoryCoefficient: 0.05},
	apiv1alpha1.Pairing:            {CPUCoefficient: 0.07, MemoryCoefficient: 0.07},
	apiv1alpha1.PairingAPI:         {CPUCoefficient: 0.14, MemoryCoefficient: 0.14},
	apiv1alpha1.RealmManagement:    {CPUCoefficient: 0.07, MemoryCoefficient: 0.07},
	apiv1alpha1.RealmManagementAPI: {CPUCoefficient: 0.07, MemoryCoefficient: 0.07},
	apiv1alpha1.TriggerEngine:      {CPUCoefficient: 0.08, MemoryCoefficient: 0.08},
	apiv1alpha1.Dashboard:          {CPUCoefficient: 0.06, MemoryCoefficient: 0.06},
}

// ReconcileConfigMap creates or updates a ConfigMap through controllerutil through its data map
func ReconcileConfigMap(objName string, data map[string]string, cr metav1.Object, c client.Client, scheme *runtime.Scheme, log logr.Logger) (controllerutil.OperationResult, error) {
	configMap := &v1.ConfigMap{ObjectMeta: metav1.ObjectMeta{Name: objName, Namespace: cr.GetNamespace()}}
	result, err := controllerutil.CreateOrUpdate(context.TODO(), c, configMap, func() error {
		if err := controllerutil.SetControllerReference(cr, configMap, scheme); err != nil {
			return err
		}
		// Set the ConfigMap data to the requested map
		configMap.Data = data
		return nil
	})

	if err != nil {
		return controllerutil.OperationResultNone, err
	}

	LogCreateOrUpdateOperationResult(log, result, cr, configMap)
	return result, err
}

// ReconcileSecret creates or updates a Secret through controllerutil through its data
func ReconcileSecret(objName string, data map[string][]byte, cr metav1.Object, c client.Client, scheme *runtime.Scheme, log logr.Logger) (controllerutil.OperationResult, error) {
	secret := &v1.Secret{ObjectMeta: metav1.ObjectMeta{Name: objName, Namespace: cr.GetNamespace()}}
	result, err := controllerutil.CreateOrUpdate(context.TODO(), c, secret, func() error {
		if err := controllerutil.SetControllerReference(cr, secret, scheme); err != nil {
			return err
		}

		secret.Type = v1.SecretTypeOpaque
		secret.Data = data
		return nil
	})

	if err != nil {
		return controllerutil.OperationResultNone, err
	}

	LogCreateOrUpdateOperationResult(log, result, cr, secret)
	return result, err
}

// ReconcileSecretString creates or updates a Secret through controllerutil by using StringData
func ReconcileSecretString(objName string, data map[string]string, cr metav1.Object, c client.Client, scheme *runtime.Scheme, log logr.Logger) (controllerutil.OperationResult, error) {
	return ReconcileSecretStringWithLabels(objName, data, map[string]string{}, cr, c, scheme, log)
}

// ReconcileSecretStringWithLabels creates or updates a Secret through controllerutil by using StringData, and adding a set of Labels
func ReconcileSecretStringWithLabels(objName string, data, labels map[string]string, cr metav1.Object, c client.Client, scheme *runtime.Scheme, log logr.Logger) (controllerutil.OperationResult, error) {
	secret := &v1.Secret{ObjectMeta: metav1.ObjectMeta{Name: objName, Namespace: cr.GetNamespace(), Labels: labels}}
	result, err := controllerutil.CreateOrUpdate(context.TODO(), c, secret, func() error {
		if err := controllerutil.SetControllerReference(cr, secret, scheme); err != nil {
			return err
		}

		secret.Type = v1.SecretTypeOpaque
		secret.StringData = data
		return nil
	})

	if err != nil {
		return controllerutil.OperationResultNone, err
	}

	LogCreateOrUpdateOperationResult(log, result, cr, secret)
	return result, err
}

// LogCreateOrUpdateOperationResult logs conveniently a controllerutil Operation
func LogCreateOrUpdateOperationResult(log logr.Logger, result controllerutil.OperationResult, cr metav1.Object, obj metav1.Object) {
	reqLogger := log.WithValues("Request.Namespace", cr.GetNamespace(), "Request.Name", cr.GetName())
	switch result {
	case controllerutil.OperationResultCreated:
		reqLogger.Info("Resource created", "Resource", obj.GetName())
	case controllerutil.OperationResultUpdated:
		reqLogger.Info("Resource updated", "Resource", obj.GetName())
	case controllerutil.OperationResultNone:
		// Debug level logging, we don't want to clutter
		reqLogger.V(1).Info("Resource unchanged", "Resource", obj.GetName())
	}
}

// GetVerneMQBrokerURL returns the complete URL for VerneMQ (MQTT) for an Astarte resource
func GetVerneMQBrokerURL(cr *apiv1alpha1.Astarte) string {
	return fmt.Sprintf("mqtts://%s:%d", cr.Spec.VerneMQ.Host, pointy.Int16Value(cr.Spec.VerneMQ.Port, 8883))
}

// GetResourcesForAstarteComponent returns the allocated resources for a given Astarte component, taking into account both the
// directive from Components, and the directive from the individual component (if any).
// It will compute a ResourceRequirements for the component based on said values and internal logic.
func GetResourcesForAstarteComponent(cr *apiv1alpha1.Astarte, requestedResources *v1.ResourceRequirements, component apiv1alpha1.AstarteComponent) v1.ResourceRequirements {
	if requestedResources != nil {
		// There has been an explicit allocation, so return that
		return *requestedResources
	}

	// Do we have any resources set?
	if cr.Spec.Components.Resources == nil {
		// All burst. If you say so...
		return v1.ResourceRequirements{}
	}

	// TODO: Kill this and integrate it into the whole shenanigan
	if component == apiv1alpha1.FlowComponent {
		return v1.ResourceRequirements{
			Limits: v1.ResourceList{
				v1.ResourceCPU:    resource.NewScaledQuantity(1000, resource.Milli).DeepCopy(),
				v1.ResourceMemory: resource.NewScaledQuantity(512, resource.Mega).DeepCopy(),
			},
			Requests: v1.ResourceList{
				v1.ResourceCPU:    resource.NewScaledQuantity(500, resource.Milli).DeepCopy(),
				v1.ResourceMemory: resource.NewScaledQuantity(256, resource.Mega).DeepCopy(),
			},
		}
	}

	// Ok, let's do the distribution dance.
	a := getWeightedDefaultAllocationFor(cr, component)

	if a.MemoryCoefficient == 0 || a.CPUCoefficient == 0 {
		return v1.ResourceRequirements{}
	}

	cpuLimits := getAllocationScaledQuantity(cr.Spec.Components.Resources.Limits.Cpu(), resource.Milli, a.CPUCoefficient)
	cpuRequests := getAllocationScaledQuantity(cr.Spec.Components.Resources.Requests.Cpu(), resource.Milli, a.CPUCoefficient)
	memoryLimits := getAllocationScaledQuantity(cr.Spec.Components.Resources.Limits.Memory(), resource.Mega, a.MemoryCoefficient)
	memoryRequests := getAllocationScaledQuantity(cr.Spec.Components.Resources.Requests.Memory(), resource.Mega, a.MemoryCoefficient)

	realRequests := v1.ResourceList{}

	// Adjust the requests to ensure we don't starve our services.
	// If the CPU is <150m, we're better off bursting the component. Otherwise, add it to requests.
	if cpuRequests.MilliValue() >= 150 {
		cpuRequests = resource.NewScaledQuantity(0, resource.Milli)
	}
	realRequests[v1.ResourceCPU] = *cpuRequests
	// If the RAM is less than 128M, increase it even though we're getting out of boundaries.
	// We can't really afford to trigger the OOM killer in the cluster, better making some components unschedulable.
	if memoryRequests.ScaledValue(resource.Mega) < 128 {
		memoryRequests = resource.NewScaledQuantity(128, resource.Mega)
	}
	realRequests[v1.ResourceMemory] = *memoryRequests

	// Same goes for limits (for CPU). Instead, though, set a higher value. That would be 300m.
	// For memory, we have to trust the user here.
	if cpuLimits.MilliValue() < 300 {
		cpuLimits = resource.NewScaledQuantity(300, resource.Milli)
	}

	return v1.ResourceRequirements{
		Limits: v1.ResourceList{
			v1.ResourceCPU:    cpuLimits.DeepCopy(),
			v1.ResourceMemory: memoryLimits.DeepCopy(),
		},
		Requests: realRequests,
	}
}

func getAllocationScaledQuantity(qty *resource.Quantity, scale resource.Scale, coefficient float64) *resource.Quantity {
	return resource.NewScaledQuantity(int64(float64(qty.ScaledValue(scale))*coefficient), scale)
}

func getNumberOfDeployedAstarteComponentsAsFloat(cr *apiv1alpha1.Astarte) float64 {
	var deployedComponents int

	if pointy.BoolValue(cr.Spec.Components.AppengineAPI.Deploy, true) {
		deployedComponents++
	}
	if pointy.BoolValue(cr.Spec.Components.Dashboard.Deploy, true) {
		deployedComponents++
	}
	if pointy.BoolValue(cr.Spec.Components.DataUpdaterPlant.Deploy, true) {
		deployedComponents++
	}
	if pointy.BoolValue(cr.Spec.Components.Housekeeping.Backend.Deploy, true) {
		deployedComponents++
	}
	if pointy.BoolValue(cr.Spec.Components.Housekeeping.API.Deploy, true) {
		deployedComponents++
	}
	if pointy.BoolValue(cr.Spec.Components.Pairing.Backend.Deploy, true) {
		deployedComponents++
	}
	if pointy.BoolValue(cr.Spec.Components.Pairing.API.Deploy, true) {
		deployedComponents++
	}
	if pointy.BoolValue(cr.Spec.Components.RealmManagement.Backend.Deploy, true) {
		deployedComponents++
	}
	if pointy.BoolValue(cr.Spec.Components.RealmManagement.API.Deploy, true) {
		deployedComponents++
	}
	if pointy.BoolValue(cr.Spec.Components.TriggerEngine.Deploy, true) {
		deployedComponents++
	}

	return float64(deployedComponents)
}

func getLeftoverCoefficients(cr *apiv1alpha1.Astarte) allocationCoefficients {
	aC := allocationCoefficients{}

	aC = checkComponentForLeftoverAllocations(cr.Spec.Components.AppengineAPI.AstarteGenericClusteredResource, apiv1alpha1.AppEngineAPI, aC)
	aC = checkComponentForLeftoverAllocations(cr.Spec.Components.Dashboard.AstarteGenericClusteredResource, apiv1alpha1.Dashboard, aC)
	aC = checkComponentForLeftoverAllocations(cr.Spec.Components.DataUpdaterPlant.AstarteGenericClusteredResource, apiv1alpha1.DataUpdaterPlant, aC)
	aC = checkComponentForLeftoverAllocations(cr.Spec.Components.Housekeeping.Backend, apiv1alpha1.Housekeeping, aC)
	aC = checkComponentForLeftoverAllocations(cr.Spec.Components.Housekeeping.API.AstarteGenericClusteredResource, apiv1alpha1.HousekeepingAPI, aC)
	aC = checkComponentForLeftoverAllocations(cr.Spec.Components.Pairing.Backend, apiv1alpha1.Pairing, aC)
	aC = checkComponentForLeftoverAllocations(cr.Spec.Components.Pairing.API.AstarteGenericClusteredResource, apiv1alpha1.PairingAPI, aC)
	aC = checkComponentForLeftoverAllocations(cr.Spec.Components.RealmManagement.Backend, apiv1alpha1.RealmManagement, aC)
	aC = checkComponentForLeftoverAllocations(cr.Spec.Components.RealmManagement.API.AstarteGenericClusteredResource, apiv1alpha1.RealmManagementAPI, aC)
	aC = checkComponentForLeftoverAllocations(cr.Spec.Components.TriggerEngine.AstarteGenericClusteredResource, apiv1alpha1.TriggerEngine, aC)

	return aC
}

func checkComponentForLeftoverAllocations(clusteredResource apiv1alpha1.AstarteGenericClusteredResource,
	component apiv1alpha1.AstarteComponent, aC allocationCoefficients) allocationCoefficients {
	if !pointy.BoolValue(clusteredResource.Deploy, true) {
		aC.CPUCoefficient += defaultComponentAllocations[component].CPUCoefficient
		aC.MemoryCoefficient += defaultComponentAllocations[component].MemoryCoefficient
	}

	return aC
}

func getWeightedDefaultAllocationFor(cr *apiv1alpha1.Astarte, component apiv1alpha1.AstarteComponent) allocationCoefficients {
	// Add other percentages proportionally
	leftovers := getLeftoverCoefficients(cr)
	defaultAllocation := defaultComponentAllocations[component]

	if leftovers.CPUCoefficient > 0 {
		defaultAllocation.CPUCoefficient += (leftovers.CPUCoefficient / getNumberOfDeployedAstarteComponentsAsFloat(cr))
	}
	if leftovers.MemoryCoefficient > 0 {
		defaultAllocation.MemoryCoefficient += (leftovers.MemoryCoefficient / getNumberOfDeployedAstarteComponentsAsFloat(cr))
	}

	return defaultAllocation
}

// IsAstarteComponentDeployed returns whether an Astarte component is deployed by cr
func IsAstarteComponentDeployed(cr *apiv1alpha1.Astarte, component apiv1alpha1.AstarteComponent) bool {
	switch component {
	case apiv1alpha1.AppEngineAPI:
		return pointy.BoolValue(cr.Spec.Components.AppengineAPI.Deploy, true)
	case apiv1alpha1.Dashboard:
		return pointy.BoolValue(cr.Spec.Components.Dashboard.Deploy, true)
	case apiv1alpha1.DataUpdaterPlant:
		return pointy.BoolValue(cr.Spec.Components.DataUpdaterPlant.Deploy, true)
	case apiv1alpha1.FlowComponent:
		return pointy.BoolValue(cr.Spec.Components.Flow.Deploy, true)
	case apiv1alpha1.Housekeeping:
		return pointy.BoolValue(cr.Spec.Components.Housekeeping.Backend.Deploy, true)
	case apiv1alpha1.HousekeepingAPI:
		return pointy.BoolValue(cr.Spec.Components.Housekeeping.API.Deploy, true)
	case apiv1alpha1.Pairing:
		return pointy.BoolValue(cr.Spec.Components.Pairing.Backend.Deploy, true)
	case apiv1alpha1.PairingAPI:
		return pointy.BoolValue(cr.Spec.Components.Pairing.API.Deploy, true)
	case apiv1alpha1.RealmManagement:
		return pointy.BoolValue(cr.Spec.Components.RealmManagement.Backend.Deploy, true)
	case apiv1alpha1.RealmManagementAPI:
		return pointy.BoolValue(cr.Spec.Components.RealmManagement.API.Deploy, true)
	case apiv1alpha1.TriggerEngine:
		return pointy.BoolValue(cr.Spec.Components.TriggerEngine.Deploy, true)
	}

	// We should not have gotten here
	return false
}

// GetRabbitMQHostnameAndPort returns the Cluster-accessible Hostname and AMQP port for RabbitMQ
func GetRabbitMQHostnameAndPort(cr *apiv1alpha1.Astarte) (string, int16) {
	if cr.Spec.RabbitMQ.Connection != nil {
		if cr.Spec.RabbitMQ.Connection.Host != "" {
			return cr.Spec.RabbitMQ.Connection.Host, pointy.Int16Value(cr.Spec.RabbitMQ.Connection.Port, 5672)
		}
	}

	// We're on defaults then. Give the standard hostname + port for our service
	return fmt.Sprintf("%s-rabbitmq.%s.svc.cluster.local", cr.Name, cr.Namespace), 5672
}

// GetRabbitMQUserCredentialsSecret gets the secret holding RabbitMQ credentials in the form <secret name>, <username key>, <password key>
func GetRabbitMQUserCredentialsSecret(cr *apiv1alpha1.Astarte) (string, string, string) {
	if cr.Spec.RabbitMQ.Connection != nil {
		if cr.Spec.RabbitMQ.Connection.Secret != nil {
			return cr.Spec.RabbitMQ.Connection.Secret.Name, cr.Spec.RabbitMQ.Connection.Secret.UsernameKey, cr.Spec.RabbitMQ.Connection.Secret.PasswordKey
		}
	}

	// Standard setup
	return cr.Name + "-rabbitmq-user-credentials", RabbitMQDefaultUserCredentialsUsernameKey, RabbitMQDefaultUserCredentialsPasswordKey
}

// GetRabbitMQCredentialsFor returns the RabbitMQ host, username and password for a given CR. This information
// can be used for connecting to RabbitMQ from the Operator or an external agent, and it should not be used for
// any other purpose.
func GetRabbitMQCredentialsFor(cr *apiv1alpha1.Astarte, c client.Client) (string, int16, string, string, error) {
	host, port := GetRabbitMQHostnameAndPort(cr)
	secretName, usernameKey, passwordKey := GetRabbitMQUserCredentialsSecret(cr)

	// Fetch the Secret
	secret := &v1.Secret{}
	if err := c.Get(context.TODO(), types.NamespacedName{Name: secretName, Namespace: cr.Namespace}, secret); err != nil {
		return "", 0, "", "", err
	}

	return host, port, string(secret.Data[usernameKey]), string(secret.Data[passwordKey]), nil
}
