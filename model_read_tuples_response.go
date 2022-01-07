/**
 * Auth0 Fine Grained Authorization (FGA)/Auth0 FGA SDK for Go
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

// ReadTuplesResponse struct for ReadTuplesResponse
type ReadTuplesResponse struct {
	Tuples            *[]Tuple `json:"tuples,omitempty"`
	ContinuationToken *string  `json:"continuation_token,omitempty"`
}

// NewReadTuplesResponse instantiates a new ReadTuplesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReadTuplesResponse() *ReadTuplesResponse {
	this := ReadTuplesResponse{}
	return &this
}

// NewReadTuplesResponseWithDefaults instantiates a new ReadTuplesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReadTuplesResponseWithDefaults() *ReadTuplesResponse {
	this := ReadTuplesResponse{}
	return &this
}

// GetTuples returns the Tuples field value if set, zero value otherwise.
func (o *ReadTuplesResponse) GetTuples() []Tuple {
	if o == nil || o.Tuples == nil {
		var ret []Tuple
		return ret
	}
	return *o.Tuples
}

// GetTuplesOk returns a tuple with the Tuples field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReadTuplesResponse) GetTuplesOk() (*[]Tuple, bool) {
	if o == nil || o.Tuples == nil {
		return nil, false
	}
	return o.Tuples, true
}

// HasTuples returns a boolean if a field has been set.
func (o *ReadTuplesResponse) HasTuples() bool {
	if o != nil && o.Tuples != nil {
		return true
	}

	return false
}

// SetTuples gets a reference to the given []Tuple and assigns it to the Tuples field.
func (o *ReadTuplesResponse) SetTuples(v []Tuple) {
	o.Tuples = &v
}

// GetContinuationToken returns the ContinuationToken field value if set, zero value otherwise.
func (o *ReadTuplesResponse) GetContinuationToken() string {
	if o == nil || o.ContinuationToken == nil {
		var ret string
		return ret
	}
	return *o.ContinuationToken
}

// GetContinuationTokenOk returns a tuple with the ContinuationToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReadTuplesResponse) GetContinuationTokenOk() (*string, bool) {
	if o == nil || o.ContinuationToken == nil {
		return nil, false
	}
	return o.ContinuationToken, true
}

// HasContinuationToken returns a boolean if a field has been set.
func (o *ReadTuplesResponse) HasContinuationToken() bool {
	if o != nil && o.ContinuationToken != nil {
		return true
	}

	return false
}

// SetContinuationToken gets a reference to the given string and assigns it to the ContinuationToken field.
func (o *ReadTuplesResponse) SetContinuationToken(v string) {
	o.ContinuationToken = &v
}

func (o ReadTuplesResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Tuples != nil {
		toSerialize["tuples"] = o.Tuples
	}
	if o.ContinuationToken != nil {
		toSerialize["continuation_token"] = o.ContinuationToken
	}
	return json.Marshal(toSerialize)
}

type NullableReadTuplesResponse struct {
	value *ReadTuplesResponse
	isSet bool
}

func (v NullableReadTuplesResponse) Get() *ReadTuplesResponse {
	return v.value
}

func (v *NullableReadTuplesResponse) Set(val *ReadTuplesResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableReadTuplesResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableReadTuplesResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReadTuplesResponse(val *ReadTuplesResponse) *NullableReadTuplesResponse {
	return &NullableReadTuplesResponse{value: val, isSet: true}
}

func (v NullableReadTuplesResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReadTuplesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
