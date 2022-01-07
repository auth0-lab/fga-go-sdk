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

// TupleKey struct for TupleKey
type TupleKey struct {
	Object   *string `json:"object,omitempty"`
	Relation *string `json:"relation,omitempty"`
	User     *string `json:"user,omitempty"`
}

// NewTupleKey instantiates a new TupleKey object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTupleKey() *TupleKey {
	this := TupleKey{}
	return &this
}

// NewTupleKeyWithDefaults instantiates a new TupleKey object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTupleKeyWithDefaults() *TupleKey {
	this := TupleKey{}
	return &this
}

// GetObject returns the Object field value if set, zero value otherwise.
func (o *TupleKey) GetObject() string {
	if o == nil || o.Object == nil {
		var ret string
		return ret
	}
	return *o.Object
}

// GetObjectOk returns a tuple with the Object field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TupleKey) GetObjectOk() (*string, bool) {
	if o == nil || o.Object == nil {
		return nil, false
	}
	return o.Object, true
}

// HasObject returns a boolean if a field has been set.
func (o *TupleKey) HasObject() bool {
	if o != nil && o.Object != nil {
		return true
	}

	return false
}

// SetObject gets a reference to the given string and assigns it to the Object field.
func (o *TupleKey) SetObject(v string) {
	o.Object = &v
}

// GetRelation returns the Relation field value if set, zero value otherwise.
func (o *TupleKey) GetRelation() string {
	if o == nil || o.Relation == nil {
		var ret string
		return ret
	}
	return *o.Relation
}

// GetRelationOk returns a tuple with the Relation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TupleKey) GetRelationOk() (*string, bool) {
	if o == nil || o.Relation == nil {
		return nil, false
	}
	return o.Relation, true
}

// HasRelation returns a boolean if a field has been set.
func (o *TupleKey) HasRelation() bool {
	if o != nil && o.Relation != nil {
		return true
	}

	return false
}

// SetRelation gets a reference to the given string and assigns it to the Relation field.
func (o *TupleKey) SetRelation(v string) {
	o.Relation = &v
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *TupleKey) GetUser() string {
	if o == nil || o.User == nil {
		var ret string
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TupleKey) GetUserOk() (*string, bool) {
	if o == nil || o.User == nil {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *TupleKey) HasUser() bool {
	if o != nil && o.User != nil {
		return true
	}

	return false
}

// SetUser gets a reference to the given string and assigns it to the User field.
func (o *TupleKey) SetUser(v string) {
	o.User = &v
}

func (o TupleKey) MarshalJSON() ([]byte, error) {
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

type NullableTupleKey struct {
	value *TupleKey
	isSet bool
}

func (v NullableTupleKey) Get() *TupleKey {
	return v.value
}

func (v *NullableTupleKey) Set(val *TupleKey) {
	v.value = val
	v.isSet = true
}

func (v NullableTupleKey) IsSet() bool {
	return v.isSet
}

func (v *NullableTupleKey) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTupleKey(val *TupleKey) *NullableTupleKey {
	return &NullableTupleKey{value: val, isSet: true}
}

func (v NullableTupleKey) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTupleKey) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
