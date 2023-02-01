// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

import (
	apiconsolev1 "github.com/openshift/api/console/v1"
	internal "github.com/openshift/client-go/console/applyconfigurations/internal"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	managedfields "k8s.io/apimachinery/pkg/util/managedfields"
	v1 "k8s.io/client-go/applyconfigurations/meta/v1"
)

// ConsoleLinkApplyConfiguration represents an declarative configuration of the ConsoleLink type for use
// with apply.
type ConsoleLinkApplyConfiguration struct {
	v1.TypeMetaApplyConfiguration    `json:",inline"`
	*v1.ObjectMetaApplyConfiguration `json:"metadata,omitempty"`
	Spec                             *ConsoleLinkSpecApplyConfiguration `json:"spec,omitempty"`
}

// ConsoleLink constructs an declarative configuration of the ConsoleLink type for use with
// apply.
func ConsoleLink(name string) *ConsoleLinkApplyConfiguration {
	b := &ConsoleLinkApplyConfiguration{}
	b.WithName(name)
	b.WithKind("ConsoleLink")
	b.WithAPIVersion("console.openshift.io/v1")
	return b
}

// ExtractConsoleLink extracts the applied configuration owned by fieldManager from
// consoleLink. If no managedFields are found in consoleLink for fieldManager, a
// ConsoleLinkApplyConfiguration is returned with only the Name, Namespace (if applicable),
// APIVersion and Kind populated. It is possible that no managed fields were found for because other
// field managers have taken ownership of all the fields previously owned by fieldManager, or because
// the fieldManager never owned fields any fields.
// consoleLink must be a unmodified ConsoleLink API object that was retrieved from the Kubernetes API.
// ExtractConsoleLink provides a way to perform a extract/modify-in-place/apply workflow.
// Note that an extracted apply configuration will contain fewer fields than what the fieldManager previously
// applied if another fieldManager has updated or force applied any of the previously applied fields.
// Experimental!
func ExtractConsoleLink(consoleLink *apiconsolev1.ConsoleLink, fieldManager string) (*ConsoleLinkApplyConfiguration, error) {
	return extractConsoleLink(consoleLink, fieldManager, "")
}

// ExtractConsoleLinkStatus is the same as ExtractConsoleLink except
// that it extracts the status subresource applied configuration.
// Experimental!
func ExtractConsoleLinkStatus(consoleLink *apiconsolev1.ConsoleLink, fieldManager string) (*ConsoleLinkApplyConfiguration, error) {
	return extractConsoleLink(consoleLink, fieldManager, "status")
}

func extractConsoleLink(consoleLink *apiconsolev1.ConsoleLink, fieldManager string, subresource string) (*ConsoleLinkApplyConfiguration, error) {
	b := &ConsoleLinkApplyConfiguration{}
	err := managedfields.ExtractInto(consoleLink, internal.Parser().Type("com.github.openshift.api.console.v1.ConsoleLink"), fieldManager, b, subresource)
	if err != nil {
		return nil, err
	}
	b.WithName(consoleLink.Name)

	b.WithKind("ConsoleLink")
	b.WithAPIVersion("console.openshift.io/v1")
	return b, nil
}

// WithKind sets the Kind field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Kind field is set to the value of the last call.
func (b *ConsoleLinkApplyConfiguration) WithKind(value string) *ConsoleLinkApplyConfiguration {
	b.Kind = &value
	return b
}

// WithAPIVersion sets the APIVersion field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the APIVersion field is set to the value of the last call.
func (b *ConsoleLinkApplyConfiguration) WithAPIVersion(value string) *ConsoleLinkApplyConfiguration {
	b.APIVersion = &value
	return b
}

// WithName sets the Name field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Name field is set to the value of the last call.
func (b *ConsoleLinkApplyConfiguration) WithName(value string) *ConsoleLinkApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.Name = &value
	return b
}

// WithGenerateName sets the GenerateName field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the GenerateName field is set to the value of the last call.
func (b *ConsoleLinkApplyConfiguration) WithGenerateName(value string) *ConsoleLinkApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.GenerateName = &value
	return b
}

// WithNamespace sets the Namespace field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Namespace field is set to the value of the last call.
func (b *ConsoleLinkApplyConfiguration) WithNamespace(value string) *ConsoleLinkApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.Namespace = &value
	return b
}

// WithUID sets the UID field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the UID field is set to the value of the last call.
func (b *ConsoleLinkApplyConfiguration) WithUID(value types.UID) *ConsoleLinkApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.UID = &value
	return b
}

// WithResourceVersion sets the ResourceVersion field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ResourceVersion field is set to the value of the last call.
func (b *ConsoleLinkApplyConfiguration) WithResourceVersion(value string) *ConsoleLinkApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.ResourceVersion = &value
	return b
}

// WithGeneration sets the Generation field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Generation field is set to the value of the last call.
func (b *ConsoleLinkApplyConfiguration) WithGeneration(value int64) *ConsoleLinkApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.Generation = &value
	return b
}

// WithCreationTimestamp sets the CreationTimestamp field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the CreationTimestamp field is set to the value of the last call.
func (b *ConsoleLinkApplyConfiguration) WithCreationTimestamp(value metav1.Time) *ConsoleLinkApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.CreationTimestamp = &value
	return b
}

// WithDeletionTimestamp sets the DeletionTimestamp field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the DeletionTimestamp field is set to the value of the last call.
func (b *ConsoleLinkApplyConfiguration) WithDeletionTimestamp(value metav1.Time) *ConsoleLinkApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.DeletionTimestamp = &value
	return b
}

// WithDeletionGracePeriodSeconds sets the DeletionGracePeriodSeconds field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the DeletionGracePeriodSeconds field is set to the value of the last call.
func (b *ConsoleLinkApplyConfiguration) WithDeletionGracePeriodSeconds(value int64) *ConsoleLinkApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.DeletionGracePeriodSeconds = &value
	return b
}

// WithLabels puts the entries into the Labels field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, the entries provided by each call will be put on the Labels field,
// overwriting an existing map entries in Labels field with the same key.
func (b *ConsoleLinkApplyConfiguration) WithLabels(entries map[string]string) *ConsoleLinkApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	if b.Labels == nil && len(entries) > 0 {
		b.Labels = make(map[string]string, len(entries))
	}
	for k, v := range entries {
		b.Labels[k] = v
	}
	return b
}

// WithAnnotations puts the entries into the Annotations field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, the entries provided by each call will be put on the Annotations field,
// overwriting an existing map entries in Annotations field with the same key.
func (b *ConsoleLinkApplyConfiguration) WithAnnotations(entries map[string]string) *ConsoleLinkApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	if b.Annotations == nil && len(entries) > 0 {
		b.Annotations = make(map[string]string, len(entries))
	}
	for k, v := range entries {
		b.Annotations[k] = v
	}
	return b
}

// WithOwnerReferences adds the given value to the OwnerReferences field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the OwnerReferences field.
func (b *ConsoleLinkApplyConfiguration) WithOwnerReferences(values ...*v1.OwnerReferenceApplyConfiguration) *ConsoleLinkApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithOwnerReferences")
		}
		b.OwnerReferences = append(b.OwnerReferences, *values[i])
	}
	return b
}

// WithFinalizers adds the given value to the Finalizers field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Finalizers field.
func (b *ConsoleLinkApplyConfiguration) WithFinalizers(values ...string) *ConsoleLinkApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	for i := range values {
		b.Finalizers = append(b.Finalizers, values[i])
	}
	return b
}

func (b *ConsoleLinkApplyConfiguration) ensureObjectMetaApplyConfigurationExists() {
	if b.ObjectMetaApplyConfiguration == nil {
		b.ObjectMetaApplyConfiguration = &v1.ObjectMetaApplyConfiguration{}
	}
}

// WithSpec sets the Spec field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Spec field is set to the value of the last call.
func (b *ConsoleLinkApplyConfiguration) WithSpec(value *ConsoleLinkSpecApplyConfiguration) *ConsoleLinkApplyConfiguration {
	b.Spec = value
	return b
}