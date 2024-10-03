/*
Copyright The Kubernetes Authors.

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

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

import (
	apiextensionsv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
)

// GeneratePatternApplyConfiguration represents an declarative configuration of the GeneratePattern type for use
// with apply.
type GeneratePatternApplyConfiguration struct {
	*ResourceSpecApplyConfiguration `json:"ResourceSpec,omitempty"`
	RawData                         *apiextensionsv1.JSON        `json:"data,omitempty"`
	Clone                           *CloneFromApplyConfiguration `json:"clone,omitempty"`
	CloneList                       *CloneListApplyConfiguration `json:"cloneList,omitempty"`
}

// GeneratePatternApplyConfiguration constructs an declarative configuration of the GeneratePattern type for use with
// apply.
func GeneratePattern() *GeneratePatternApplyConfiguration {
	return &GeneratePatternApplyConfiguration{}
}

// WithAPIVersion sets the APIVersion field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the APIVersion field is set to the value of the last call.
func (b *GeneratePatternApplyConfiguration) WithAPIVersion(value string) *GeneratePatternApplyConfiguration {
	b.ensureResourceSpecApplyConfigurationExists()
	b.APIVersion = &value
	return b
}

// WithKind sets the Kind field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Kind field is set to the value of the last call.
func (b *GeneratePatternApplyConfiguration) WithKind(value string) *GeneratePatternApplyConfiguration {
	b.ensureResourceSpecApplyConfigurationExists()
	b.Kind = &value
	return b
}

// WithNamespace sets the Namespace field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Namespace field is set to the value of the last call.
func (b *GeneratePatternApplyConfiguration) WithNamespace(value string) *GeneratePatternApplyConfiguration {
	b.ensureResourceSpecApplyConfigurationExists()
	b.Namespace = &value
	return b
}

// WithName sets the Name field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Name field is set to the value of the last call.
func (b *GeneratePatternApplyConfiguration) WithName(value string) *GeneratePatternApplyConfiguration {
	b.ensureResourceSpecApplyConfigurationExists()
	b.Name = &value
	return b
}

// WithSelector sets the Selector field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Selector field is set to the value of the last call.
func (b *GeneratePatternApplyConfiguration) WithSelector(value metav1.LabelSelector) *GeneratePatternApplyConfiguration {
	b.ensureResourceSpecApplyConfigurationExists()
	b.Selector = &value
	return b
}

// WithUID sets the UID field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the UID field is set to the value of the last call.
func (b *GeneratePatternApplyConfiguration) WithUID(value types.UID) *GeneratePatternApplyConfiguration {
	b.ensureResourceSpecApplyConfigurationExists()
	b.UID = &value
	return b
}

func (b *GeneratePatternApplyConfiguration) ensureResourceSpecApplyConfigurationExists() {
	if b.ResourceSpecApplyConfiguration == nil {
		b.ResourceSpecApplyConfiguration = &ResourceSpecApplyConfiguration{}
	}
}

// WithRawData sets the RawData field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the RawData field is set to the value of the last call.
func (b *GeneratePatternApplyConfiguration) WithRawData(value apiextensionsv1.JSON) *GeneratePatternApplyConfiguration {
	b.RawData = &value
	return b
}

// WithClone sets the Clone field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Clone field is set to the value of the last call.
func (b *GeneratePatternApplyConfiguration) WithClone(value *CloneFromApplyConfiguration) *GeneratePatternApplyConfiguration {
	b.Clone = value
	return b
}

// WithCloneList sets the CloneList field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the CloneList field is set to the value of the last call.
func (b *GeneratePatternApplyConfiguration) WithCloneList(value *CloneListApplyConfiguration) *GeneratePatternApplyConfiguration {
	b.CloneList = value
	return b
}
