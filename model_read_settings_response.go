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

// ReadSettingsResponse struct for ReadSettingsResponse
type ReadSettingsResponse struct {
	Environment  *Environment   `json:"environment,omitempty"`
	TokenIssuers *[]TokenIssuer `json:"token_issuers,omitempty"`
}

// NewReadSettingsResponse instantiates a new ReadSettingsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReadSettingsResponse() *ReadSettingsResponse {
	this := ReadSettingsResponse{}
	var environment Environment = ENVIRONMENT_UNSPECIFIED
	this.Environment = &environment
	return &this
}

// NewReadSettingsResponseWithDefaults instantiates a new ReadSettingsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReadSettingsResponseWithDefaults() *ReadSettingsResponse {
	this := ReadSettingsResponse{}
	var environment Environment = ENVIRONMENT_UNSPECIFIED
	this.Environment = &environment
	return &this
}

// GetEnvironment returns the Environment field value if set, zero value otherwise.
func (o *ReadSettingsResponse) GetEnvironment() Environment {
	if o == nil || o.Environment == nil {
		var ret Environment
		return ret
	}
	return *o.Environment
}

// GetEnvironmentOk returns a tuple with the Environment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReadSettingsResponse) GetEnvironmentOk() (*Environment, bool) {
	if o == nil || o.Environment == nil {
		return nil, false
	}
	return o.Environment, true
}

// HasEnvironment returns a boolean if a field has been set.
func (o *ReadSettingsResponse) HasEnvironment() bool {
	if o != nil && o.Environment != nil {
		return true
	}

	return false
}

// SetEnvironment gets a reference to the given Environment and assigns it to the Environment field.
func (o *ReadSettingsResponse) SetEnvironment(v Environment) {
	o.Environment = &v
}

// GetTokenIssuers returns the TokenIssuers field value if set, zero value otherwise.
func (o *ReadSettingsResponse) GetTokenIssuers() []TokenIssuer {
	if o == nil || o.TokenIssuers == nil {
		var ret []TokenIssuer
		return ret
	}
	return *o.TokenIssuers
}

// GetTokenIssuersOk returns a tuple with the TokenIssuers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReadSettingsResponse) GetTokenIssuersOk() (*[]TokenIssuer, bool) {
	if o == nil || o.TokenIssuers == nil {
		return nil, false
	}
	return o.TokenIssuers, true
}

// HasTokenIssuers returns a boolean if a field has been set.
func (o *ReadSettingsResponse) HasTokenIssuers() bool {
	if o != nil && o.TokenIssuers != nil {
		return true
	}

	return false
}

// SetTokenIssuers gets a reference to the given []TokenIssuer and assigns it to the TokenIssuers field.
func (o *ReadSettingsResponse) SetTokenIssuers(v []TokenIssuer) {
	o.TokenIssuers = &v
}

func (o ReadSettingsResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Environment != nil {
		toSerialize["environment"] = o.Environment
	}
	if o.TokenIssuers != nil {
		toSerialize["token_issuers"] = o.TokenIssuers
	}
	return json.Marshal(toSerialize)
}

type NullableReadSettingsResponse struct {
	value *ReadSettingsResponse
	isSet bool
}

func (v NullableReadSettingsResponse) Get() *ReadSettingsResponse {
	return v.value
}

func (v *NullableReadSettingsResponse) Set(val *ReadSettingsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableReadSettingsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableReadSettingsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReadSettingsResponse(val *ReadSettingsResponse) *NullableReadSettingsResponse {
	return &NullableReadSettingsResponse{value: val, isSet: true}
}

func (v NullableReadSettingsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReadSettingsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}