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

// Any struct for Any
type Any struct {
	Type *string `json:"@type,omitempty"`
}

// NewAny instantiates a new Any object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAny() *Any {
	this := Any{}
	return &this
}

// NewAnyWithDefaults instantiates a new Any object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAnyWithDefaults() *Any {
	this := Any{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *Any) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Any) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *Any) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *Any) SetType(v string) {
	o.Type = &v
}

func (o Any) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["@type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableAny struct {
	value *Any
	isSet bool
}

func (v NullableAny) Get() *Any {
	return v.value
}

func (v *NullableAny) Set(val *Any) {
	v.value = val
	v.isSet = true
}

func (v NullableAny) IsSet() bool {
	return v.isSet
}

func (v *NullableAny) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAny(val *Any) *NullableAny {
	return &NullableAny{value: val, isSet: true}
}

func (v NullableAny) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAny) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}