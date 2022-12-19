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
	"fmt"
)

// InternalErrorCode - no_internal_error: no error  - internal_error: generic internal error.  - auth_internal_error: internal error due to authentication error.  - auth_failed_error_fetching_well_known_jwks: authentication failure due to internal error in fetching well-known/jwks.json.  - cancelled: internal error - request being cancelled.  - deadline_exceeded: internal error - deadline exceeded.  - already_exists: internal error - already exists.  - resource_exhausted: internal error - resource exhausted.  - failed_precondition: internal error - failed precondition.  - aborted: internal error - aborted.  - out_of_range: internal error - out of range.  - unavailable: internal error - unavailable.  - data_loss: internal error - data loss.
type InternalErrorCode string

// List of InternalErrorCode
const (
	NO_INTERNAL_ERROR                          InternalErrorCode = "no_internal_error"
	INTERNAL_ERROR                             InternalErrorCode = "internal_error"
	AUTH_INTERNAL_ERROR                        InternalErrorCode = "auth_internal_error"
	AUTH_FAILED_ERROR_FETCHING_WELL_KNOWN_JWKS InternalErrorCode = "auth_failed_error_fetching_well_known_jwks"
	CANCELLED                                  InternalErrorCode = "cancelled"
	DEADLINE_EXCEEDED                          InternalErrorCode = "deadline_exceeded"
	ALREADY_EXISTS                             InternalErrorCode = "already_exists"
	RESOURCE_EXHAUSTED                         InternalErrorCode = "resource_exhausted"
	FAILED_PRECONDITION                        InternalErrorCode = "failed_precondition"
	ABORTED                                    InternalErrorCode = "aborted"
	OUT_OF_RANGE                               InternalErrorCode = "out_of_range"
	UNAVAILABLE                                InternalErrorCode = "unavailable"
	DATA_LOSS                                  InternalErrorCode = "data_loss"
)

var allowedInternalErrorCodeEnumValues = []InternalErrorCode{
	"no_internal_error",
	"internal_error",
	"auth_internal_error",
	"auth_failed_error_fetching_well_known_jwks",
	"cancelled",
	"deadline_exceeded",
	"already_exists",
	"resource_exhausted",
	"failed_precondition",
	"aborted",
	"out_of_range",
	"unavailable",
	"data_loss",
}

func (v *InternalErrorCode) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := InternalErrorCode(value)
	for _, existing := range allowedInternalErrorCodeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid InternalErrorCode", value)
}

// NewInternalErrorCodeFromValue returns a pointer to a valid InternalErrorCode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewInternalErrorCodeFromValue(v string) (*InternalErrorCode, error) {
	ev := InternalErrorCode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for InternalErrorCode: valid values are %v", v, allowedInternalErrorCodeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v InternalErrorCode) IsValid() bool {
	for _, existing := range allowedInternalErrorCodeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to InternalErrorCode value
func (v InternalErrorCode) Ptr() *InternalErrorCode {
	return &v
}

type NullableInternalErrorCode struct {
	value *InternalErrorCode
	isSet bool
}

func (v NullableInternalErrorCode) Get() *InternalErrorCode {
	return v.value
}

func (v *NullableInternalErrorCode) Set(val *InternalErrorCode) {
	v.value = val
	v.isSet = true
}

func (v NullableInternalErrorCode) IsSet() bool {
	return v.isSet
}

func (v *NullableInternalErrorCode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInternalErrorCode(val *InternalErrorCode) *NullableInternalErrorCode {
	return &NullableInternalErrorCode{value: val, isSet: true}
}

func (v NullableInternalErrorCode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInternalErrorCode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
