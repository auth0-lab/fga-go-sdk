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
	"time"
)

// Tuple struct for Tuple
type Tuple struct {
	Key       *TupleKey  `json:"key,omitempty"`
	Timestamp *time.Time `json:"timestamp,omitempty"`
}

// NewTuple instantiates a new Tuple object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTuple() *Tuple {
	this := Tuple{}
	return &this
}

// NewTupleWithDefaults instantiates a new Tuple object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTupleWithDefaults() *Tuple {
	this := Tuple{}
	return &this
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *Tuple) GetKey() TupleKey {
	if o == nil || o.Key == nil {
		var ret TupleKey
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Tuple) GetKeyOk() (*TupleKey, bool) {
	if o == nil || o.Key == nil {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *Tuple) HasKey() bool {
	if o != nil && o.Key != nil {
		return true
	}

	return false
}

// SetKey gets a reference to the given TupleKey and assigns it to the Key field.
func (o *Tuple) SetKey(v TupleKey) {
	o.Key = &v
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise.
func (o *Tuple) GetTimestamp() time.Time {
	if o == nil || o.Timestamp == nil {
		var ret time.Time
		return ret
	}
	return *o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Tuple) GetTimestampOk() (*time.Time, bool) {
	if o == nil || o.Timestamp == nil {
		return nil, false
	}
	return o.Timestamp, true
}

// HasTimestamp returns a boolean if a field has been set.
func (o *Tuple) HasTimestamp() bool {
	if o != nil && o.Timestamp != nil {
		return true
	}

	return false
}

// SetTimestamp gets a reference to the given time.Time and assigns it to the Timestamp field.
func (o *Tuple) SetTimestamp(v time.Time) {
	o.Timestamp = &v
}

func (o Tuple) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Key != nil {
		toSerialize["key"] = o.Key
	}
	if o.Timestamp != nil {
		toSerialize["timestamp"] = o.Timestamp
	}
	return json.Marshal(toSerialize)
}

type NullableTuple struct {
	value *Tuple
	isSet bool
}

func (v NullableTuple) Get() *Tuple {
	return v.value
}

func (v *NullableTuple) Set(val *Tuple) {
	v.value = val
	v.isSet = true
}

func (v NullableTuple) IsSet() bool {
	return v.isSet
}

func (v *NullableTuple) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTuple(val *Tuple) *NullableTuple {
	return &NullableTuple{value: val, isSet: true}
}

func (v NullableTuple) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTuple) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
