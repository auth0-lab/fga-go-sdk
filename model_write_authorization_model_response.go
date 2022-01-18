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

// WriteAuthorizationModelResponse struct for WriteAuthorizationModelResponse
type WriteAuthorizationModelResponse struct {
	AuthorizationModelId *string `json:"authorization_model_id,omitempty"`
}

// NewWriteAuthorizationModelResponse instantiates a new WriteAuthorizationModelResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWriteAuthorizationModelResponse() *WriteAuthorizationModelResponse {
	this := WriteAuthorizationModelResponse{}
	return &this
}

// NewWriteAuthorizationModelResponseWithDefaults instantiates a new WriteAuthorizationModelResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWriteAuthorizationModelResponseWithDefaults() *WriteAuthorizationModelResponse {
	this := WriteAuthorizationModelResponse{}
	return &this
}

// GetAuthorizationModelId returns the AuthorizationModelId field value if set, zero value otherwise.
func (o *WriteAuthorizationModelResponse) GetAuthorizationModelId() string {
	if o == nil || o.AuthorizationModelId == nil {
		var ret string
		return ret
	}
	return *o.AuthorizationModelId
}

// GetAuthorizationModelIdOk returns a tuple with the AuthorizationModelId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WriteAuthorizationModelResponse) GetAuthorizationModelIdOk() (*string, bool) {
	if o == nil || o.AuthorizationModelId == nil {
		return nil, false
	}
	return o.AuthorizationModelId, true
}

// HasAuthorizationModelId returns a boolean if a field has been set.
func (o *WriteAuthorizationModelResponse) HasAuthorizationModelId() bool {
	if o != nil && o.AuthorizationModelId != nil {
		return true
	}

	return false
}

// SetAuthorizationModelId gets a reference to the given string and assigns it to the AuthorizationModelId field.
func (o *WriteAuthorizationModelResponse) SetAuthorizationModelId(v string) {
	o.AuthorizationModelId = &v
}

func (o WriteAuthorizationModelResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AuthorizationModelId != nil {
		toSerialize["authorization_model_id"] = o.AuthorizationModelId
	}
	return json.Marshal(toSerialize)
}

type NullableWriteAuthorizationModelResponse struct {
	value *WriteAuthorizationModelResponse
	isSet bool
}

func (v NullableWriteAuthorizationModelResponse) Get() *WriteAuthorizationModelResponse {
	return v.value
}

func (v *NullableWriteAuthorizationModelResponse) Set(val *WriteAuthorizationModelResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableWriteAuthorizationModelResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableWriteAuthorizationModelResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWriteAuthorizationModelResponse(val *WriteAuthorizationModelResponse) *NullableWriteAuthorizationModelResponse {
	return &NullableWriteAuthorizationModelResponse{value: val, isSet: true}
}

func (v NullableWriteAuthorizationModelResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWriteAuthorizationModelResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
