/**
 * Go SDK for Auth0 Fine Grained Authorization (FGA)
 *
 * Auth0 Fine Grained Authorization (FGA) is an early-stage product we are building at Auth0 as part of Auth0Lab to solve fine-grained authorization at scale. If you are interested in learning more about our plans, please reach out via our Discord chat.  The limits and information described in this document is subject to change.
 *
 * API version: 0.1
 * Website: https://fga.dev
 * Documentation: https://docs.fga.dev
 * Support: https://discord.gg/8naAwJfWN6
 *
 * NOTE: This file was auto generated by OpenAPI Generator (https://openapi-generator.tech). DO NOT EDIT.
 */

package auth0fga

import (
	"encoding/json"
)

// AuthorizationmodelTupleToUserset struct for AuthorizationmodelTupleToUserset
type AuthorizationmodelTupleToUserset struct {
	Tupleset        *ObjectRelation `json:"tupleset,omitempty"`
	ComputedUserset *ObjectRelation `json:"computedUserset,omitempty"`
}

// NewAuthorizationmodelTupleToUserset instantiates a new AuthorizationmodelTupleToUserset object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthorizationmodelTupleToUserset() *AuthorizationmodelTupleToUserset {
	this := AuthorizationmodelTupleToUserset{}
	return &this
}

// NewAuthorizationmodelTupleToUsersetWithDefaults instantiates a new AuthorizationmodelTupleToUserset object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthorizationmodelTupleToUsersetWithDefaults() *AuthorizationmodelTupleToUserset {
	this := AuthorizationmodelTupleToUserset{}
	return &this
}

// GetTupleset returns the Tupleset field value if set, zero value otherwise.
func (o *AuthorizationmodelTupleToUserset) GetTupleset() ObjectRelation {
	if o == nil || o.Tupleset == nil {
		var ret ObjectRelation
		return ret
	}
	return *o.Tupleset
}

// GetTuplesetOk returns a tuple with the Tupleset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizationmodelTupleToUserset) GetTuplesetOk() (*ObjectRelation, bool) {
	if o == nil || o.Tupleset == nil {
		return nil, false
	}
	return o.Tupleset, true
}

// HasTupleset returns a boolean if a field has been set.
func (o *AuthorizationmodelTupleToUserset) HasTupleset() bool {
	if o != nil && o.Tupleset != nil {
		return true
	}

	return false
}

// SetTupleset gets a reference to the given ObjectRelation and assigns it to the Tupleset field.
func (o *AuthorizationmodelTupleToUserset) SetTupleset(v ObjectRelation) {
	o.Tupleset = &v
}

// GetComputedUserset returns the ComputedUserset field value if set, zero value otherwise.
func (o *AuthorizationmodelTupleToUserset) GetComputedUserset() ObjectRelation {
	if o == nil || o.ComputedUserset == nil {
		var ret ObjectRelation
		return ret
	}
	return *o.ComputedUserset
}

// GetComputedUsersetOk returns a tuple with the ComputedUserset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizationmodelTupleToUserset) GetComputedUsersetOk() (*ObjectRelation, bool) {
	if o == nil || o.ComputedUserset == nil {
		return nil, false
	}
	return o.ComputedUserset, true
}

// HasComputedUserset returns a boolean if a field has been set.
func (o *AuthorizationmodelTupleToUserset) HasComputedUserset() bool {
	if o != nil && o.ComputedUserset != nil {
		return true
	}

	return false
}

// SetComputedUserset gets a reference to the given ObjectRelation and assigns it to the ComputedUserset field.
func (o *AuthorizationmodelTupleToUserset) SetComputedUserset(v ObjectRelation) {
	o.ComputedUserset = &v
}

func (o AuthorizationmodelTupleToUserset) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Tupleset != nil {
		toSerialize["tupleset"] = o.Tupleset
	}
	if o.ComputedUserset != nil {
		toSerialize["computedUserset"] = o.ComputedUserset
	}
	return json.Marshal(toSerialize)
}

type NullableAuthorizationmodelTupleToUserset struct {
	value *AuthorizationmodelTupleToUserset
	isSet bool
}

func (v NullableAuthorizationmodelTupleToUserset) Get() *AuthorizationmodelTupleToUserset {
	return v.value
}

func (v *NullableAuthorizationmodelTupleToUserset) Set(val *AuthorizationmodelTupleToUserset) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthorizationmodelTupleToUserset) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthorizationmodelTupleToUserset) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthorizationmodelTupleToUserset(val *AuthorizationmodelTupleToUserset) *NullableAuthorizationmodelTupleToUserset {
	return &NullableAuthorizationmodelTupleToUserset{value: val, isSet: true}
}

func (v NullableAuthorizationmodelTupleToUserset) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthorizationmodelTupleToUserset) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
