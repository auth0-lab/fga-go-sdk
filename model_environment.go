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
	"fmt"
)

// Environment the model 'Environment'
type Environment string

// List of Environment
const (
	ENVIRONMENT_UNSPECIFIED Environment = "ENVIRONMENT_UNSPECIFIED"
	DEVELOPMENT             Environment = "DEVELOPMENT"
	STAGING                 Environment = "STAGING"
	PRODUCTION              Environment = "PRODUCTION"
)

var allowedEnvironmentEnumValues = []Environment{
	"ENVIRONMENT_UNSPECIFIED",
	"DEVELOPMENT",
	"STAGING",
	"PRODUCTION",
}

func (v *Environment) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := Environment(value)
	for _, existing := range allowedEnvironmentEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid Environment", value)
}

// NewEnvironmentFromValue returns a pointer to a valid Environment
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnvironmentFromValue(v string) (*Environment, error) {
	ev := Environment(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for Environment: valid values are %v", v, allowedEnvironmentEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v Environment) IsValid() bool {
	for _, existing := range allowedEnvironmentEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Environment value
func (v Environment) Ptr() *Environment {
	return &v
}

type NullableEnvironment struct {
	value *Environment
	isSet bool
}

func (v NullableEnvironment) Get() *Environment {
	return v.value
}

func (v *NullableEnvironment) Set(val *Environment) {
	v.value = val
	v.isSet = true
}

func (v NullableEnvironment) IsSet() bool {
	return v.isSet
}

func (v *NullableEnvironment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnvironment(val *Environment) *NullableEnvironment {
	return &NullableEnvironment{value: val, isSet: true}
}

func (v NullableEnvironment) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnvironment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
