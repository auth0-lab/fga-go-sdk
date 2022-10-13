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

// ResourceExhaustedErrorMessageResponse struct for ResourceExhaustedErrorMessageResponse
type ResourceExhaustedErrorMessageResponse struct {
	Code    *ResourceExhaustedErrorCode `json:"code,omitempty"`
	Message *string                     `json:"message,omitempty"`
}

// NewResourceExhaustedErrorMessageResponse instantiates a new ResourceExhaustedErrorMessageResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResourceExhaustedErrorMessageResponse() *ResourceExhaustedErrorMessageResponse {
	this := ResourceExhaustedErrorMessageResponse{}
	var code ResourceExhaustedErrorCode = NO_RESOURCE_EXHAUSTED_ERROR
	this.Code = &code
	return &this
}

// NewResourceExhaustedErrorMessageResponseWithDefaults instantiates a new ResourceExhaustedErrorMessageResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResourceExhaustedErrorMessageResponseWithDefaults() *ResourceExhaustedErrorMessageResponse {
	this := ResourceExhaustedErrorMessageResponse{}
	var code ResourceExhaustedErrorCode = NO_RESOURCE_EXHAUSTED_ERROR
	this.Code = &code
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *ResourceExhaustedErrorMessageResponse) GetCode() ResourceExhaustedErrorCode {
	if o == nil || o.Code == nil {
		var ret ResourceExhaustedErrorCode
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceExhaustedErrorMessageResponse) GetCodeOk() (*ResourceExhaustedErrorCode, bool) {
	if o == nil || o.Code == nil {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *ResourceExhaustedErrorMessageResponse) HasCode() bool {
	if o != nil && o.Code != nil {
		return true
	}

	return false
}

// SetCode gets a reference to the given ResourceExhaustedErrorCode and assigns it to the Code field.
func (o *ResourceExhaustedErrorMessageResponse) SetCode(v ResourceExhaustedErrorCode) {
	o.Code = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *ResourceExhaustedErrorMessageResponse) GetMessage() string {
	if o == nil || o.Message == nil {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceExhaustedErrorMessageResponse) GetMessageOk() (*string, bool) {
	if o == nil || o.Message == nil {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *ResourceExhaustedErrorMessageResponse) HasMessage() bool {
	if o != nil && o.Message != nil {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *ResourceExhaustedErrorMessageResponse) SetMessage(v string) {
	o.Message = &v
}

func (o ResourceExhaustedErrorMessageResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Code != nil {
		toSerialize["code"] = o.Code
	}
	if o.Message != nil {
		toSerialize["message"] = o.Message
	}
	return json.Marshal(toSerialize)
}

type NullableResourceExhaustedErrorMessageResponse struct {
	value *ResourceExhaustedErrorMessageResponse
	isSet bool
}

func (v NullableResourceExhaustedErrorMessageResponse) Get() *ResourceExhaustedErrorMessageResponse {
	return v.value
}

func (v *NullableResourceExhaustedErrorMessageResponse) Set(val *ResourceExhaustedErrorMessageResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableResourceExhaustedErrorMessageResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableResourceExhaustedErrorMessageResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResourceExhaustedErrorMessageResponse(val *ResourceExhaustedErrorMessageResponse) *NullableResourceExhaustedErrorMessageResponse {
	return &NullableResourceExhaustedErrorMessageResponse{value: val, isSet: true}
}

func (v NullableResourceExhaustedErrorMessageResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResourceExhaustedErrorMessageResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
