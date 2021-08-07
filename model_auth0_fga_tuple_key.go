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

// Auth0FgaTupleKey struct for Auth0FgaTupleKey
type Auth0FgaTupleKey struct {
	Object   *string `json:"object,omitempty"`
	Relation *string `json:"relation,omitempty"`
	User     *string `json:"user,omitempty"`
}

// NewAuth0FgaTupleKey instantiates a new Auth0FgaTupleKey object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuth0FgaTupleKey() *Auth0FgaTupleKey {
	this := Auth0FgaTupleKey{}
	return &this
}

// NewAuth0FgaTupleKeyWithDefaults instantiates a new Auth0FgaTupleKey object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuth0FgaTupleKeyWithDefaults() *Auth0FgaTupleKey {
	this := Auth0FgaTupleKey{}
	return &this
}

// GetObject returns the Object field value if set, zero value otherwise.
func (o *Auth0FgaTupleKey) GetObject() string {
	if o == nil || o.Object == nil {
		var ret string
		return ret
	}
	return *o.Object
}

// GetObjectOk returns a tuple with the Object field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Auth0FgaTupleKey) GetObjectOk() (*string, bool) {
	if o == nil || o.Object == nil {
		return nil, false
	}
	return o.Object, true
}

// HasObject returns a boolean if a field has been set.
func (o *Auth0FgaTupleKey) HasObject() bool {
	if o != nil && o.Object != nil {
		return true
	}

	return false
}

// SetObject gets a reference to the given string and assigns it to the Object field.
func (o *Auth0FgaTupleKey) SetObject(v string) {
	o.Object = &v
}

// GetRelation returns the Relation field value if set, zero value otherwise.
func (o *Auth0FgaTupleKey) GetRelation() string {
	if o == nil || o.Relation == nil {
		var ret string
		return ret
	}
	return *o.Relation
}

// GetRelationOk returns a tuple with the Relation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Auth0FgaTupleKey) GetRelationOk() (*string, bool) {
	if o == nil || o.Relation == nil {
		return nil, false
	}
	return o.Relation, true
}

// HasRelation returns a boolean if a field has been set.
func (o *Auth0FgaTupleKey) HasRelation() bool {
	if o != nil && o.Relation != nil {
		return true
	}

	return false
}

// SetRelation gets a reference to the given string and assigns it to the Relation field.
func (o *Auth0FgaTupleKey) SetRelation(v string) {
	o.Relation = &v
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *Auth0FgaTupleKey) GetUser() string {
	if o == nil || o.User == nil {
		var ret string
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Auth0FgaTupleKey) GetUserOk() (*string, bool) {
	if o == nil || o.User == nil {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *Auth0FgaTupleKey) HasUser() bool {
	if o != nil && o.User != nil {
		return true
	}

	return false
}

// SetUser gets a reference to the given string and assigns it to the User field.
func (o *Auth0FgaTupleKey) SetUser(v string) {
	o.User = &v
}

func (o Auth0FgaTupleKey) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Object != nil {
		toSerialize["object"] = o.Object
	}
	if o.Relation != nil {
		toSerialize["relation"] = o.Relation
	}
	if o.User != nil {
		toSerialize["user"] = o.User
	}
	return json.Marshal(toSerialize)
}

type NullableAuth0FgaTupleKey struct {
	value *Auth0FgaTupleKey
	isSet bool
}

func (v NullableAuth0FgaTupleKey) Get() *Auth0FgaTupleKey {
	return v.value
}

func (v *NullableAuth0FgaTupleKey) Set(val *Auth0FgaTupleKey) {
	v.value = val
	v.isSet = true
}

func (v NullableAuth0FgaTupleKey) IsSet() bool {
	return v.isSet
}

func (v *NullableAuth0FgaTupleKey) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuth0FgaTupleKey(val *Auth0FgaTupleKey) *NullableAuth0FgaTupleKey {
	return &NullableAuth0FgaTupleKey{value: val, isSet: true}
}

func (v NullableAuth0FgaTupleKey) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuth0FgaTupleKey) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
