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

// AuthorizationmodelAuthorizationModel struct for AuthorizationmodelAuthorizationModel
type AuthorizationmodelAuthorizationModel struct {
	Id              *string                             `json:"id,omitempty"`
	TypeDefinitions *[]AuthorizationmodelTypeDefinition `json:"type_definitions,omitempty"`
}

// NewAuthorizationmodelAuthorizationModel instantiates a new AuthorizationmodelAuthorizationModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthorizationmodelAuthorizationModel() *AuthorizationmodelAuthorizationModel {
	this := AuthorizationmodelAuthorizationModel{}
	return &this
}

// NewAuthorizationmodelAuthorizationModelWithDefaults instantiates a new AuthorizationmodelAuthorizationModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthorizationmodelAuthorizationModelWithDefaults() *AuthorizationmodelAuthorizationModel {
	this := AuthorizationmodelAuthorizationModel{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AuthorizationmodelAuthorizationModel) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizationmodelAuthorizationModel) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AuthorizationmodelAuthorizationModel) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *AuthorizationmodelAuthorizationModel) SetId(v string) {
	o.Id = &v
}

// GetTypeDefinitions returns the TypeDefinitions field value if set, zero value otherwise.
func (o *AuthorizationmodelAuthorizationModel) GetTypeDefinitions() []AuthorizationmodelTypeDefinition {
	if o == nil || o.TypeDefinitions == nil {
		var ret []AuthorizationmodelTypeDefinition
		return ret
	}
	return *o.TypeDefinitions
}

// GetTypeDefinitionsOk returns a tuple with the TypeDefinitions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizationmodelAuthorizationModel) GetTypeDefinitionsOk() (*[]AuthorizationmodelTypeDefinition, bool) {
	if o == nil || o.TypeDefinitions == nil {
		return nil, false
	}
	return o.TypeDefinitions, true
}

// HasTypeDefinitions returns a boolean if a field has been set.
func (o *AuthorizationmodelAuthorizationModel) HasTypeDefinitions() bool {
	if o != nil && o.TypeDefinitions != nil {
		return true
	}

	return false
}

// SetTypeDefinitions gets a reference to the given []AuthorizationmodelTypeDefinition and assigns it to the TypeDefinitions field.
func (o *AuthorizationmodelAuthorizationModel) SetTypeDefinitions(v []AuthorizationmodelTypeDefinition) {
	o.TypeDefinitions = &v
}

func (o AuthorizationmodelAuthorizationModel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.TypeDefinitions != nil {
		toSerialize["type_definitions"] = o.TypeDefinitions
	}
	return json.Marshal(toSerialize)
}

type NullableAuthorizationmodelAuthorizationModel struct {
	value *AuthorizationmodelAuthorizationModel
	isSet bool
}

func (v NullableAuthorizationmodelAuthorizationModel) Get() *AuthorizationmodelAuthorizationModel {
	return v.value
}

func (v *NullableAuthorizationmodelAuthorizationModel) Set(val *AuthorizationmodelAuthorizationModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthorizationmodelAuthorizationModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthorizationmodelAuthorizationModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthorizationmodelAuthorizationModel(val *AuthorizationmodelAuthorizationModel) *NullableAuthorizationmodelAuthorizationModel {
	return &NullableAuthorizationmodelAuthorizationModel{value: val, isSet: true}
}

func (v NullableAuthorizationmodelAuthorizationModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthorizationmodelAuthorizationModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
