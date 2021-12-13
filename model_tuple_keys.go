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

// TupleKeys struct for TupleKeys
type TupleKeys struct {
	TupleKeys []TupleKey `json:"tuple_keys"`
}

// NewTupleKeys instantiates a new TupleKeys object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTupleKeys(tupleKeys []TupleKey) *TupleKeys {
	this := TupleKeys{}
	this.TupleKeys = tupleKeys
	return &this
}

// NewTupleKeysWithDefaults instantiates a new TupleKeys object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTupleKeysWithDefaults() *TupleKeys {
	this := TupleKeys{}
	return &this
}

// GetTupleKeys returns the TupleKeys field value
func (o *TupleKeys) GetTupleKeys() []TupleKey {
	if o == nil {
		var ret []TupleKey
		return ret
	}

	return o.TupleKeys
}

// GetTupleKeysOk returns a tuple with the TupleKeys field value
// and a boolean to check if the value has been set.
func (o *TupleKeys) GetTupleKeysOk() (*[]TupleKey, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TupleKeys, true
}

// SetTupleKeys sets field value
func (o *TupleKeys) SetTupleKeys(v []TupleKey) {
	o.TupleKeys = v
}

func (o TupleKeys) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["tuple_keys"] = o.TupleKeys
	}
	return json.Marshal(toSerialize)
}

type NullableTupleKeys struct {
	value *TupleKeys
	isSet bool
}

func (v NullableTupleKeys) Get() *TupleKeys {
	return v.value
}

func (v *NullableTupleKeys) Set(val *TupleKeys) {
	v.value = val
	v.isSet = true
}

func (v NullableTupleKeys) IsSet() bool {
	return v.isSet
}

func (v *NullableTupleKeys) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTupleKeys(val *TupleKeys) *NullableTupleKeys {
	return &NullableTupleKeys{value: val, isSet: true}
}

func (v NullableTupleKeys) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTupleKeys) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}