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

// TokenIssuer struct for TokenIssuer
type TokenIssuer struct {
	Id        *string `json:"id,omitempty"`
	IssuerUrl *string `json:"issuer_url,omitempty"`
}

// NewTokenIssuer instantiates a new TokenIssuer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTokenIssuer() *TokenIssuer {
	this := TokenIssuer{}
	return &this
}

// NewTokenIssuerWithDefaults instantiates a new TokenIssuer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTokenIssuerWithDefaults() *TokenIssuer {
	this := TokenIssuer{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *TokenIssuer) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenIssuer) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *TokenIssuer) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *TokenIssuer) SetId(v string) {
	o.Id = &v
}

// GetIssuerUrl returns the IssuerUrl field value if set, zero value otherwise.
func (o *TokenIssuer) GetIssuerUrl() string {
	if o == nil || o.IssuerUrl == nil {
		var ret string
		return ret
	}
	return *o.IssuerUrl
}

// GetIssuerUrlOk returns a tuple with the IssuerUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenIssuer) GetIssuerUrlOk() (*string, bool) {
	if o == nil || o.IssuerUrl == nil {
		return nil, false
	}
	return o.IssuerUrl, true
}

// HasIssuerUrl returns a boolean if a field has been set.
func (o *TokenIssuer) HasIssuerUrl() bool {
	if o != nil && o.IssuerUrl != nil {
		return true
	}

	return false
}

// SetIssuerUrl gets a reference to the given string and assigns it to the IssuerUrl field.
func (o *TokenIssuer) SetIssuerUrl(v string) {
	o.IssuerUrl = &v
}

func (o TokenIssuer) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.IssuerUrl != nil {
		toSerialize["issuer_url"] = o.IssuerUrl
	}
	return json.Marshal(toSerialize)
}

type NullableTokenIssuer struct {
	value *TokenIssuer
	isSet bool
}

func (v NullableTokenIssuer) Get() *TokenIssuer {
	return v.value
}

func (v *NullableTokenIssuer) Set(val *TokenIssuer) {
	v.value = val
	v.isSet = true
}

func (v NullableTokenIssuer) IsSet() bool {
	return v.isSet
}

func (v *NullableTokenIssuer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTokenIssuer(val *TokenIssuer) *NullableTokenIssuer {
	return &NullableTokenIssuer{value: val, isSet: true}
}

func (v NullableTokenIssuer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTokenIssuer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
