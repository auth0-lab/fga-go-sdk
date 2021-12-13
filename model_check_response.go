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

// CheckResponse struct for CheckResponse
type CheckResponse struct {
	Allowed    *bool   `json:"allowed,omitempty"`
	Resolution *string `json:"resolution,omitempty"`
}

// NewCheckResponse instantiates a new CheckResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCheckResponse() *CheckResponse {
	this := CheckResponse{}
	return &this
}

// NewCheckResponseWithDefaults instantiates a new CheckResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCheckResponseWithDefaults() *CheckResponse {
	this := CheckResponse{}
	return &this
}

// GetAllowed returns the Allowed field value if set, zero value otherwise.
func (o *CheckResponse) GetAllowed() bool {
	if o == nil || o.Allowed == nil {
		var ret bool
		return ret
	}
	return *o.Allowed
}

// GetAllowedOk returns a tuple with the Allowed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckResponse) GetAllowedOk() (*bool, bool) {
	if o == nil || o.Allowed == nil {
		return nil, false
	}
	return o.Allowed, true
}

// HasAllowed returns a boolean if a field has been set.
func (o *CheckResponse) HasAllowed() bool {
	if o != nil && o.Allowed != nil {
		return true
	}

	return false
}

// SetAllowed gets a reference to the given bool and assigns it to the Allowed field.
func (o *CheckResponse) SetAllowed(v bool) {
	o.Allowed = &v
}

// GetResolution returns the Resolution field value if set, zero value otherwise.
func (o *CheckResponse) GetResolution() string {
	if o == nil || o.Resolution == nil {
		var ret string
		return ret
	}
	return *o.Resolution
}

// GetResolutionOk returns a tuple with the Resolution field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckResponse) GetResolutionOk() (*string, bool) {
	if o == nil || o.Resolution == nil {
		return nil, false
	}
	return o.Resolution, true
}

// HasResolution returns a boolean if a field has been set.
func (o *CheckResponse) HasResolution() bool {
	if o != nil && o.Resolution != nil {
		return true
	}

	return false
}

// SetResolution gets a reference to the given string and assigns it to the Resolution field.
func (o *CheckResponse) SetResolution(v string) {
	o.Resolution = &v
}

func (o CheckResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Allowed != nil {
		toSerialize["allowed"] = o.Allowed
	}
	if o.Resolution != nil {
		toSerialize["resolution"] = o.Resolution
	}
	return json.Marshal(toSerialize)
}

type NullableCheckResponse struct {
	value *CheckResponse
	isSet bool
}

func (v NullableCheckResponse) Get() *CheckResponse {
	return v.value
}

func (v *NullableCheckResponse) Set(val *CheckResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCheckResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCheckResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCheckResponse(val *CheckResponse) *NullableCheckResponse {
	return &NullableCheckResponse{value: val, isSet: true}
}

func (v NullableCheckResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCheckResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}