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

// NotFoundErrorCode - no_not_found_error: no error  - undefined_endpoint: undefined endpoint.  - customer_id_not_found: customer ID is not found.  - store_id_not_found: store ID not found  - store_client_id_not_found: store client ID not found.  - resource_not_found: generic not found.  - unimplemented: method is unimplemented
type NotFoundErrorCode string

// List of NotFoundErrorCode
const (
	NO_NOT_FOUND_ERROR        NotFoundErrorCode = "no_not_found_error"
	UNDEFINED_ENDPOINT        NotFoundErrorCode = "undefined_endpoint"
	CUSTOMER_ID_NOT_FOUND     NotFoundErrorCode = "customer_id_not_found"
	STORE_ID_NOT_FOUND        NotFoundErrorCode = "store_id_not_found"
	STORE_CLIENT_ID_NOT_FOUND NotFoundErrorCode = "store_client_id_not_found"
	RESOURCE_NOT_FOUND        NotFoundErrorCode = "resource_not_found"
	UNIMPLEMENTED             NotFoundErrorCode = "unimplemented"
)

var allowedNotFoundErrorCodeEnumValues = []NotFoundErrorCode{
	"no_not_found_error",
	"undefined_endpoint",
	"customer_id_not_found",
	"store_id_not_found",
	"store_client_id_not_found",
	"resource_not_found",
	"unimplemented",
}

func (v *NotFoundErrorCode) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := NotFoundErrorCode(value)
	for _, existing := range allowedNotFoundErrorCodeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid NotFoundErrorCode", value)
}

// NewNotFoundErrorCodeFromValue returns a pointer to a valid NotFoundErrorCode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewNotFoundErrorCodeFromValue(v string) (*NotFoundErrorCode, error) {
	ev := NotFoundErrorCode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for NotFoundErrorCode: valid values are %v", v, allowedNotFoundErrorCodeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v NotFoundErrorCode) IsValid() bool {
	for _, existing := range allowedNotFoundErrorCodeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to NotFoundErrorCode value
func (v NotFoundErrorCode) Ptr() *NotFoundErrorCode {
	return &v
}

type NullableNotFoundErrorCode struct {
	value *NotFoundErrorCode
	isSet bool
}

func (v NullableNotFoundErrorCode) Get() *NotFoundErrorCode {
	return v.value
}

func (v *NullableNotFoundErrorCode) Set(val *NotFoundErrorCode) {
	v.value = val
	v.isSet = true
}

func (v NullableNotFoundErrorCode) IsSet() bool {
	return v.isSet
}

func (v *NullableNotFoundErrorCode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNotFoundErrorCode(val *NotFoundErrorCode) *NullableNotFoundErrorCode {
	return &NullableNotFoundErrorCode{value: val, isSet: true}
}

func (v NullableNotFoundErrorCode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNotFoundErrorCode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
