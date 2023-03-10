/**
 * Go SDK for Auth0 Fine Grained Authorization (FGA)
 *
 * API version: 0.1
 * Website: <https://fga.dev>
 * Documentation: <https://docs.fga.dev>
 * Support: <https://discord.gg/8naAwJfWN6>
 * License: [MIT](https://github.com/auth0-lab/fga-go-sdk/blob/main/LICENSE)
 *
 * NOTE: This file was auto generated by [OpenAPI Generator](https://openapi-generator.tech). DO NOT EDIT.
 */

package auth0fga

import (
	"encoding/json"
)

// ExpandRequest struct for ExpandRequest
type ExpandRequest struct {
	TupleKey             TupleKey `json:"tuple_key"`
	AuthorizationModelId *string  `json:"authorization_model_id,omitempty"`
}

// NewExpandRequest instantiates a new ExpandRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExpandRequest(tupleKey TupleKey) *ExpandRequest {
	this := ExpandRequest{}
	this.TupleKey = tupleKey
	return &this
}

// NewExpandRequestWithDefaults instantiates a new ExpandRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExpandRequestWithDefaults() *ExpandRequest {
	this := ExpandRequest{}
	return &this
}

// GetTupleKey returns the TupleKey field value
func (o *ExpandRequest) GetTupleKey() TupleKey {
	if o == nil {
		var ret TupleKey
		return ret
	}

	return o.TupleKey
}

// GetTupleKeyOk returns a tuple with the TupleKey field value
// and a boolean to check if the value has been set.
func (o *ExpandRequest) GetTupleKeyOk() (*TupleKey, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TupleKey, true
}

// SetTupleKey sets field value
func (o *ExpandRequest) SetTupleKey(v TupleKey) {
	o.TupleKey = v
}

// GetAuthorizationModelId returns the AuthorizationModelId field value if set, zero value otherwise.
func (o *ExpandRequest) GetAuthorizationModelId() string {
	if o == nil || o.AuthorizationModelId == nil {
		var ret string
		return ret
	}
	return *o.AuthorizationModelId
}

// GetAuthorizationModelIdOk returns a tuple with the AuthorizationModelId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExpandRequest) GetAuthorizationModelIdOk() (*string, bool) {
	if o == nil || o.AuthorizationModelId == nil {
		return nil, false
	}
	return o.AuthorizationModelId, true
}

// HasAuthorizationModelId returns a boolean if a field has been set.
func (o *ExpandRequest) HasAuthorizationModelId() bool {
	if o != nil && o.AuthorizationModelId != nil {
		return true
	}

	return false
}

// SetAuthorizationModelId gets a reference to the given string and assigns it to the AuthorizationModelId field.
func (o *ExpandRequest) SetAuthorizationModelId(v string) {
	o.AuthorizationModelId = &v
}

func (o ExpandRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["tuple_key"] = o.TupleKey
	}
	if o.AuthorizationModelId != nil {
		toSerialize["authorization_model_id"] = o.AuthorizationModelId
	}
	return json.Marshal(toSerialize)
}

type NullableExpandRequest struct {
	value *ExpandRequest
	isSet bool
}

func (v NullableExpandRequest) Get() *ExpandRequest {
	return v.value
}

func (v *NullableExpandRequest) Set(val *ExpandRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableExpandRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableExpandRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExpandRequest(val *ExpandRequest) *NullableExpandRequest {
	return &NullableExpandRequest{value: val, isSet: true}
}

func (v NullableExpandRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExpandRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
