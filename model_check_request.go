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

// CheckRequest struct for CheckRequest
type CheckRequest struct {
	TupleKey             TupleKey             `json:"tuple_key"`
	ContextualTuples     *ContextualTupleKeys `json:"contextual_tuples,omitempty"`
	AuthorizationModelId *string              `json:"authorization_model_id,omitempty"`
	// Defaults to false. Making it true has performance implications.
	Trace *bool `json:"trace,omitempty"`
}

// NewCheckRequest instantiates a new CheckRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCheckRequest(tupleKey TupleKey) *CheckRequest {
	this := CheckRequest{}
	this.TupleKey = tupleKey
	return &this
}

// NewCheckRequestWithDefaults instantiates a new CheckRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCheckRequestWithDefaults() *CheckRequest {
	this := CheckRequest{}
	return &this
}

// GetTupleKey returns the TupleKey field value
func (o *CheckRequest) GetTupleKey() TupleKey {
	if o == nil {
		var ret TupleKey
		return ret
	}

	return o.TupleKey
}

// GetTupleKeyOk returns a tuple with the TupleKey field value
// and a boolean to check if the value has been set.
func (o *CheckRequest) GetTupleKeyOk() (*TupleKey, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TupleKey, true
}

// SetTupleKey sets field value
func (o *CheckRequest) SetTupleKey(v TupleKey) {
	o.TupleKey = v
}

// GetContextualTuples returns the ContextualTuples field value if set, zero value otherwise.
func (o *CheckRequest) GetContextualTuples() ContextualTupleKeys {
	if o == nil || o.ContextualTuples == nil {
		var ret ContextualTupleKeys
		return ret
	}
	return *o.ContextualTuples
}

// GetContextualTuplesOk returns a tuple with the ContextualTuples field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckRequest) GetContextualTuplesOk() (*ContextualTupleKeys, bool) {
	if o == nil || o.ContextualTuples == nil {
		return nil, false
	}
	return o.ContextualTuples, true
}

// HasContextualTuples returns a boolean if a field has been set.
func (o *CheckRequest) HasContextualTuples() bool {
	if o != nil && o.ContextualTuples != nil {
		return true
	}

	return false
}

// SetContextualTuples gets a reference to the given ContextualTupleKeys and assigns it to the ContextualTuples field.
func (o *CheckRequest) SetContextualTuples(v ContextualTupleKeys) {
	o.ContextualTuples = &v
}

// GetAuthorizationModelId returns the AuthorizationModelId field value if set, zero value otherwise.
func (o *CheckRequest) GetAuthorizationModelId() string {
	if o == nil || o.AuthorizationModelId == nil {
		var ret string
		return ret
	}
	return *o.AuthorizationModelId
}

// GetAuthorizationModelIdOk returns a tuple with the AuthorizationModelId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckRequest) GetAuthorizationModelIdOk() (*string, bool) {
	if o == nil || o.AuthorizationModelId == nil {
		return nil, false
	}
	return o.AuthorizationModelId, true
}

// HasAuthorizationModelId returns a boolean if a field has been set.
func (o *CheckRequest) HasAuthorizationModelId() bool {
	if o != nil && o.AuthorizationModelId != nil {
		return true
	}

	return false
}

// SetAuthorizationModelId gets a reference to the given string and assigns it to the AuthorizationModelId field.
func (o *CheckRequest) SetAuthorizationModelId(v string) {
	o.AuthorizationModelId = &v
}

// GetTrace returns the Trace field value if set, zero value otherwise.
func (o *CheckRequest) GetTrace() bool {
	if o == nil || o.Trace == nil {
		var ret bool
		return ret
	}
	return *o.Trace
}

// GetTraceOk returns a tuple with the Trace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckRequest) GetTraceOk() (*bool, bool) {
	if o == nil || o.Trace == nil {
		return nil, false
	}
	return o.Trace, true
}

// HasTrace returns a boolean if a field has been set.
func (o *CheckRequest) HasTrace() bool {
	if o != nil && o.Trace != nil {
		return true
	}

	return false
}

// SetTrace gets a reference to the given bool and assigns it to the Trace field.
func (o *CheckRequest) SetTrace(v bool) {
	o.Trace = &v
}

func (o CheckRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["tuple_key"] = o.TupleKey
	}
	if o.ContextualTuples != nil {
		toSerialize["contextual_tuples"] = o.ContextualTuples
	}
	if o.AuthorizationModelId != nil {
		toSerialize["authorization_model_id"] = o.AuthorizationModelId
	}
	if o.Trace != nil {
		toSerialize["trace"] = o.Trace
	}
	return json.Marshal(toSerialize)
}

type NullableCheckRequest struct {
	value *CheckRequest
	isSet bool
}

func (v NullableCheckRequest) Get() *CheckRequest {
	return v.value
}

func (v *NullableCheckRequest) Set(val *CheckRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCheckRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCheckRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCheckRequest(val *CheckRequest) *NullableCheckRequest {
	return &NullableCheckRequest{value: val, isSet: true}
}

func (v NullableCheckRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCheckRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}