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
	"fmt"
)

// ResourceExhaustedErrorCode - no_resource_exhausted_error: no error  - rate_limit_exceeded: operation failed due to exceeding rate limit.  - auth_rate_limit_exceeded: rate limit error during authentication.
type ResourceExhaustedErrorCode string

// List of ResourceExhaustedErrorCode
const (
	NO_RESOURCE_EXHAUSTED_ERROR ResourceExhaustedErrorCode = "no_resource_exhausted_error"
	RATE_LIMIT_EXCEEDED         ResourceExhaustedErrorCode = "rate_limit_exceeded"
	AUTH_RATE_LIMIT_EXCEEDED    ResourceExhaustedErrorCode = "auth_rate_limit_exceeded"
)

var allowedResourceExhaustedErrorCodeEnumValues = []ResourceExhaustedErrorCode{
	"no_resource_exhausted_error",
	"rate_limit_exceeded",
	"auth_rate_limit_exceeded",
}

func (v *ResourceExhaustedErrorCode) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ResourceExhaustedErrorCode(value)
	for _, existing := range allowedResourceExhaustedErrorCodeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ResourceExhaustedErrorCode", value)
}

// NewResourceExhaustedErrorCodeFromValue returns a pointer to a valid ResourceExhaustedErrorCode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewResourceExhaustedErrorCodeFromValue(v string) (*ResourceExhaustedErrorCode, error) {
	ev := ResourceExhaustedErrorCode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ResourceExhaustedErrorCode: valid values are %v", v, allowedResourceExhaustedErrorCodeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ResourceExhaustedErrorCode) IsValid() bool {
	for _, existing := range allowedResourceExhaustedErrorCodeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ResourceExhaustedErrorCode value
func (v ResourceExhaustedErrorCode) Ptr() *ResourceExhaustedErrorCode {
	return &v
}

type NullableResourceExhaustedErrorCode struct {
	value *ResourceExhaustedErrorCode
	isSet bool
}

func (v NullableResourceExhaustedErrorCode) Get() *ResourceExhaustedErrorCode {
	return v.value
}

func (v *NullableResourceExhaustedErrorCode) Set(val *ResourceExhaustedErrorCode) {
	v.value = val
	v.isSet = true
}

func (v NullableResourceExhaustedErrorCode) IsSet() bool {
	return v.isSet
}

func (v *NullableResourceExhaustedErrorCode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResourceExhaustedErrorCode(val *ResourceExhaustedErrorCode) *NullableResourceExhaustedErrorCode {
	return &NullableResourceExhaustedErrorCode{value: val, isSet: true}
}

func (v NullableResourceExhaustedErrorCode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResourceExhaustedErrorCode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
