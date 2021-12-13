/*
 * Auth0 Fine Grained Authorization (FGA)
 *
 * Auth0 Fine Grained Authorization (FGA) is an early-stage product we are building at Auth0 as part of Auth0Lab to solve fine-grained authorization at scale. If you are interested in learning more about our plans, please reach out via our Discord chat.  The limits and information described in this document is subject to change.
 *
 * API version: 0.1
 * Contact: https://discord.gg/8naAwJfWN6
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package auth0fga

import (
	"encoding/json"
)

// TypeDefinition struct for TypeDefinition
type TypeDefinition struct {
	Type      string             `json:"type"`
	Relations map[string]Userset `json:"relations"`
}

// NewTypeDefinition instantiates a new TypeDefinition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTypeDefinition(type_ string, relations map[string]Userset) *TypeDefinition {
	this := TypeDefinition{}
	this.Type = type_
	this.Relations = relations
	return &this
}

// NewTypeDefinitionWithDefaults instantiates a new TypeDefinition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTypeDefinitionWithDefaults() *TypeDefinition {
	this := TypeDefinition{}
	return &this
}

// GetType returns the Type field value
func (o *TypeDefinition) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *TypeDefinition) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *TypeDefinition) SetType(v string) {
	o.Type = v
}

// GetRelations returns the Relations field value
func (o *TypeDefinition) GetRelations() map[string]Userset {
	if o == nil {
		var ret map[string]Userset
		return ret
	}

	return o.Relations
}

// GetRelationsOk returns a tuple with the Relations field value
// and a boolean to check if the value has been set.
func (o *TypeDefinition) GetRelationsOk() (*map[string]Userset, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Relations, true
}

// SetRelations sets field value
func (o *TypeDefinition) SetRelations(v map[string]Userset) {
	o.Relations = v
}

func (o TypeDefinition) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["relations"] = o.Relations
	}
	return json.Marshal(toSerialize)
}

type NullableTypeDefinition struct {
	value *TypeDefinition
	isSet bool
}

func (v NullableTypeDefinition) Get() *TypeDefinition {
	return v.value
}

func (v *NullableTypeDefinition) Set(val *TypeDefinition) {
	v.value = val
	v.isSet = true
}

func (v NullableTypeDefinition) IsSet() bool {
	return v.isSet
}

func (v *NullableTypeDefinition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTypeDefinition(val *TypeDefinition) *NullableTypeDefinition {
	return &NullableTypeDefinition{value: val, isSet: true}
}

func (v NullableTypeDefinition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTypeDefinition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}