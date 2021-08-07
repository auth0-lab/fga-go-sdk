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

// WriteTokenIssuersRequestParams struct for WriteTokenIssuersRequestParams
type WriteTokenIssuersRequestParams struct {
	IssuerUrl *string `json:"issuer_url,omitempty"`
}

// NewWriteTokenIssuersRequestParams instantiates a new WriteTokenIssuersRequestParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWriteTokenIssuersRequestParams() *WriteTokenIssuersRequestParams {
	this := WriteTokenIssuersRequestParams{}
	return &this
}

// NewWriteTokenIssuersRequestParamsWithDefaults instantiates a new WriteTokenIssuersRequestParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWriteTokenIssuersRequestParamsWithDefaults() *WriteTokenIssuersRequestParams {
	this := WriteTokenIssuersRequestParams{}
	return &this
}

// GetIssuerUrl returns the IssuerUrl field value if set, zero value otherwise.
func (o *WriteTokenIssuersRequestParams) GetIssuerUrl() string {
	if o == nil || o.IssuerUrl == nil {
		var ret string
		return ret
	}
	return *o.IssuerUrl
}

// GetIssuerUrlOk returns a tuple with the IssuerUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WriteTokenIssuersRequestParams) GetIssuerUrlOk() (*string, bool) {
	if o == nil || o.IssuerUrl == nil {
		return nil, false
	}
	return o.IssuerUrl, true
}

// HasIssuerUrl returns a boolean if a field has been set.
func (o *WriteTokenIssuersRequestParams) HasIssuerUrl() bool {
	if o != nil && o.IssuerUrl != nil {
		return true
	}

	return false
}

// SetIssuerUrl gets a reference to the given string and assigns it to the IssuerUrl field.
func (o *WriteTokenIssuersRequestParams) SetIssuerUrl(v string) {
	o.IssuerUrl = &v
}

func (o WriteTokenIssuersRequestParams) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.IssuerUrl != nil {
		toSerialize["issuer_url"] = o.IssuerUrl
	}
	return json.Marshal(toSerialize)
}

type NullableWriteTokenIssuersRequestParams struct {
	value *WriteTokenIssuersRequestParams
	isSet bool
}

func (v NullableWriteTokenIssuersRequestParams) Get() *WriteTokenIssuersRequestParams {
	return v.value
}

func (v *NullableWriteTokenIssuersRequestParams) Set(val *WriteTokenIssuersRequestParams) {
	v.value = val
	v.isSet = true
}

func (v NullableWriteTokenIssuersRequestParams) IsSet() bool {
	return v.isSet
}

func (v *NullableWriteTokenIssuersRequestParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWriteTokenIssuersRequestParams(val *WriteTokenIssuersRequestParams) *NullableWriteTokenIssuersRequestParams {
	return &NullableWriteTokenIssuersRequestParams{value: val, isSet: true}
}

func (v NullableWriteTokenIssuersRequestParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWriteTokenIssuersRequestParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
