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
	"fmt"
)

// SettingsEnvironment the model 'SettingsEnvironment'
type SettingsEnvironment string

// List of settingsEnvironment
const (
	ENVIRONMENT_UNSPECIFIED SettingsEnvironment = "ENVIRONMENT_UNSPECIFIED"
	DEVELOPMENT             SettingsEnvironment = "DEVELOPMENT"
	STAGING                 SettingsEnvironment = "STAGING"
	PRODUCTION              SettingsEnvironment = "PRODUCTION"
)

var allowedSettingsEnvironmentEnumValues = []SettingsEnvironment{
	"ENVIRONMENT_UNSPECIFIED",
	"DEVELOPMENT",
	"STAGING",
	"PRODUCTION",
}

func (v *SettingsEnvironment) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SettingsEnvironment(value)
	for _, existing := range allowedSettingsEnvironmentEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SettingsEnvironment", value)
}

// NewSettingsEnvironmentFromValue returns a pointer to a valid SettingsEnvironment
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSettingsEnvironmentFromValue(v string) (*SettingsEnvironment, error) {
	ev := SettingsEnvironment(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SettingsEnvironment: valid values are %v", v, allowedSettingsEnvironmentEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SettingsEnvironment) IsValid() bool {
	for _, existing := range allowedSettingsEnvironmentEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to settingsEnvironment value
func (v SettingsEnvironment) Ptr() *SettingsEnvironment {
	return &v
}

type NullableSettingsEnvironment struct {
	value *SettingsEnvironment
	isSet bool
}

func (v NullableSettingsEnvironment) Get() *SettingsEnvironment {
	return v.value
}

func (v *NullableSettingsEnvironment) Set(val *SettingsEnvironment) {
	v.value = val
	v.isSet = true
}

func (v NullableSettingsEnvironment) IsSet() bool {
	return v.isSet
}

func (v *NullableSettingsEnvironment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSettingsEnvironment(val *SettingsEnvironment) *NullableSettingsEnvironment {
	return &NullableSettingsEnvironment{value: val, isSet: true}
}

func (v NullableSettingsEnvironment) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSettingsEnvironment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
