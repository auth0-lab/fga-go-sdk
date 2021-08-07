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

// ReadAuthorizationModelResponse struct for ReadAuthorizationModelResponse
type ReadAuthorizationModelResponse struct {
	AuthorizationModel *AuthorizationModel `json:"authorization_model,omitempty"`
}

// NewReadAuthorizationModelResponse instantiates a new ReadAuthorizationModelResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReadAuthorizationModelResponse() *ReadAuthorizationModelResponse {
	this := ReadAuthorizationModelResponse{}
	return &this
}

// NewReadAuthorizationModelResponseWithDefaults instantiates a new ReadAuthorizationModelResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReadAuthorizationModelResponseWithDefaults() *ReadAuthorizationModelResponse {
	this := ReadAuthorizationModelResponse{}
	return &this
}

// GetAuthorizationModel returns the AuthorizationModel field value if set, zero value otherwise.
func (o *ReadAuthorizationModelResponse) GetAuthorizationModel() AuthorizationModel {
	if o == nil || o.AuthorizationModel == nil {
		var ret AuthorizationModel
		return ret
	}
	return *o.AuthorizationModel
}

// GetAuthorizationModelOk returns a tuple with the AuthorizationModel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReadAuthorizationModelResponse) GetAuthorizationModelOk() (*AuthorizationModel, bool) {
	if o == nil || o.AuthorizationModel == nil {
		return nil, false
	}
	return o.AuthorizationModel, true
}

// HasAuthorizationModel returns a boolean if a field has been set.
func (o *ReadAuthorizationModelResponse) HasAuthorizationModel() bool {
	if o != nil && o.AuthorizationModel != nil {
		return true
	}

	return false
}

// SetAuthorizationModel gets a reference to the given AuthorizationModel and assigns it to the AuthorizationModel field.
func (o *ReadAuthorizationModelResponse) SetAuthorizationModel(v AuthorizationModel) {
	o.AuthorizationModel = &v
}

func (o ReadAuthorizationModelResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AuthorizationModel != nil {
		toSerialize["authorization_model"] = o.AuthorizationModel
	}
	return json.Marshal(toSerialize)
}

type NullableReadAuthorizationModelResponse struct {
	value *ReadAuthorizationModelResponse
	isSet bool
}

func (v NullableReadAuthorizationModelResponse) Get() *ReadAuthorizationModelResponse {
	return v.value
}

func (v *NullableReadAuthorizationModelResponse) Set(val *ReadAuthorizationModelResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableReadAuthorizationModelResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableReadAuthorizationModelResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReadAuthorizationModelResponse(val *ReadAuthorizationModelResponse) *NullableReadAuthorizationModelResponse {
	return &NullableReadAuthorizationModelResponse{value: val, isSet: true}
}

func (v NullableReadAuthorizationModelResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReadAuthorizationModelResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
