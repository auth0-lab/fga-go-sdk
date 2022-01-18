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

// ExpandRequestParams struct for ExpandRequestParams
type ExpandRequestParams struct {
	TupleKey             *TupleKey `json:"tuple_key,omitempty"`
	AuthorizationModelId *string   `json:"authorization_model_id,omitempty"`
}

// NewExpandRequestParams instantiates a new ExpandRequestParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExpandRequestParams() *ExpandRequestParams {
	this := ExpandRequestParams{}
	return &this
}

// NewExpandRequestParamsWithDefaults instantiates a new ExpandRequestParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExpandRequestParamsWithDefaults() *ExpandRequestParams {
	this := ExpandRequestParams{}
	return &this
}

// GetTupleKey returns the TupleKey field value if set, zero value otherwise.
func (o *ExpandRequestParams) GetTupleKey() TupleKey {
	if o == nil || o.TupleKey == nil {
		var ret TupleKey
		return ret
	}
	return *o.TupleKey
}

// GetTupleKeyOk returns a tuple with the TupleKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExpandRequestParams) GetTupleKeyOk() (*TupleKey, bool) {
	if o == nil || o.TupleKey == nil {
		return nil, false
	}
	return o.TupleKey, true
}

// HasTupleKey returns a boolean if a field has been set.
func (o *ExpandRequestParams) HasTupleKey() bool {
	if o != nil && o.TupleKey != nil {
		return true
	}

	return false
}

// SetTupleKey gets a reference to the given TupleKey and assigns it to the TupleKey field.
func (o *ExpandRequestParams) SetTupleKey(v TupleKey) {
	o.TupleKey = &v
}

// GetAuthorizationModelId returns the AuthorizationModelId field value if set, zero value otherwise.
func (o *ExpandRequestParams) GetAuthorizationModelId() string {
	if o == nil || o.AuthorizationModelId == nil {
		var ret string
		return ret
	}
	return *o.AuthorizationModelId
}

// GetAuthorizationModelIdOk returns a tuple with the AuthorizationModelId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExpandRequestParams) GetAuthorizationModelIdOk() (*string, bool) {
	if o == nil || o.AuthorizationModelId == nil {
		return nil, false
	}
	return o.AuthorizationModelId, true
}

// HasAuthorizationModelId returns a boolean if a field has been set.
func (o *ExpandRequestParams) HasAuthorizationModelId() bool {
	if o != nil && o.AuthorizationModelId != nil {
		return true
	}

	return false
}

// SetAuthorizationModelId gets a reference to the given string and assigns it to the AuthorizationModelId field.
func (o *ExpandRequestParams) SetAuthorizationModelId(v string) {
	o.AuthorizationModelId = &v
}

func (o ExpandRequestParams) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.TupleKey != nil {
		toSerialize["tuple_key"] = o.TupleKey
	}
	if o.AuthorizationModelId != nil {
		toSerialize["authorization_model_id"] = o.AuthorizationModelId
	}
	return json.Marshal(toSerialize)
}

type NullableExpandRequestParams struct {
	value *ExpandRequestParams
	isSet bool
}

func (v NullableExpandRequestParams) Get() *ExpandRequestParams {
	return v.value
}

func (v *NullableExpandRequestParams) Set(val *ExpandRequestParams) {
	v.value = val
	v.isSet = true
}

func (v NullableExpandRequestParams) IsSet() bool {
	return v.isSet
}

func (v *NullableExpandRequestParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExpandRequestParams(val *ExpandRequestParams) *NullableExpandRequestParams {
	return &NullableExpandRequestParams{value: val, isSet: true}
}

func (v NullableExpandRequestParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExpandRequestParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
