// Copyright 2021 Chaos Mesh Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// See the License for the specific language governing permissions and
// limitations under the License.

package v1alpha1

import (
	"fmt"

	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/validation/field"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
)

// log is for logging in this package.
var gcpchaoslog = logf.Log.WithName("gcpchaos-resource")

// +kubebuilder:webhook:path=/mutate-chaos-mesh-org-v1alpha1-gcpchaos,mutating=true,failurePolicy=fail,groups=chaos-mesh.org,resources=gcpchaos,verbs=create;update,versions=v1alpha1,name=mgcpchaos.kb.io

var _ webhook.Defaulter = &GcpChaos{}

// Default implements webhook.Defaulter so a webhook will be registered for the type
func (in *GcpChaos) Default() {
	gcpchaoslog.Info("default", "name", in.Name)
}

// +kubebuilder:webhook:verbs=create;update,path=/validate-chaos-mesh-org-v1alpha1-gcpchaos,mutating=false,failurePolicy=fail,groups=chaos-mesh.org,resources=gcpchaos,versions=v1alpha1,name=vgcpchaos.kb.io

var _ ChaosValidator = &GcpChaos{}

// ValidateCreate implements webhook.Validator so a webhook will be registered for the type
func (in *GcpChaos) ValidateCreate() error {
	gcpchaoslog.Info("validate create", "name", in.Name)
	return in.Validate()
}

// ValidateUpdate implements webhook.Validator so a webhook will be registered for the type
func (in *GcpChaos) ValidateUpdate(old runtime.Object) error {
	gcpchaoslog.Info("validate update", "name", in.Name)
	return in.Validate()
}

// ValidateDelete implements webhook.Validator so a webhook will be registered for the type
func (in *GcpChaos) ValidateDelete() error {
	gcpchaoslog.Info("validate delete", "name", in.Name)

	// Nothing to do?
	return nil
}

// SelectSpec returns the selector config for authority validate
func (in *GcpChaos) GetSelectSpec() []SelectSpec {
	return nil
}

// Validate validates chaos object
func (in *GcpChaos) Validate() error {
	specField := field.NewPath("spec")
	allErrs := in.ValidateScheduler(specField)
	allErrs = append(allErrs, in.ValidatePodMode(specField)...)
	allErrs = append(allErrs, in.Spec.validateDeviceName(specField.Child("deviceName"))...)

	if len(allErrs) > 0 {
		return fmt.Errorf(allErrs.ToAggregate().Error())
	}
	return nil
}

// ValidateScheduler validates the scheduler and duration
func (in *GcpChaos) ValidateScheduler(spec *field.Path) field.ErrorList {
	allErrs := field.ErrorList{}
	schedulerField := spec.Child("scheduler")

	switch in.Spec.Action {
	case NodeStop, DiskLoss:
		allErrs = append(allErrs, ValidateScheduler(in, spec)...)
	case NodeReset:
		// We choose to ignore the Duration property even user define it
		if in.Spec.Scheduler != nil {
			_, err := ParseCron(in.Spec.Scheduler.Cron, schedulerField.Child("cron"))
			allErrs = append(allErrs, err...)
		}
	default:
		err := fmt.Errorf("awschaos[%s/%s] have unknown action type", in.Namespace, in.Name)
		log.Error(err, "Wrong AwsChaos Action type")

		actionField := spec.Child("action")
		allErrs = append(allErrs, field.Invalid(actionField, in.Spec.Action, err.Error()))
	}
	return allErrs
}

// ValidatePodMode validates the value with podmode
func (in *GcpChaos) ValidatePodMode(spec *field.Path) field.ErrorList {
	// Because gcp chaos does not need a pod mode, so return nil here.
	return nil
}

// validateDeviceName validates the DeviceName
func (in *GcpChaosSpec) validateDeviceName(containerField *field.Path) field.ErrorList {
	allErrs := field.ErrorList{}
	if in.Action == DiskLoss {
		if in.DeviceNames == nil {
			err := fmt.Errorf("at least one device name is required on %s action", in.Action)
			allErrs = append(allErrs, field.Invalid(containerField, in.DeviceNames, err.Error()))
		}
	}
	return allErrs
}
