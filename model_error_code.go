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

// ErrorCode - no_error: no error  - validation_error: generic validation error.  - authorization_model_not_found: authorization model not found.  - authorization_model_resolution_too_complex: too many rewrite rules to be resolved.  - invalid_write_input: invalid write input.  - cannot_allow_duplicate_tuples_in_one_request: duplicate tuples in one request.  - cannot_allow_duplicate_types_in_one_request: duplicate types in one request.  - cannot_allow_multiple_references_to_one_relation: cannot use a relation to define itself.  - invalid_continuation_token: invalid continuation token.  - invalid_tuple_set: invalid tuple set.  - invalid_check_input: invalid check input.  - invalid_expand_input: invalid expand input.  - unsupported_user_set: unsupported user set.  - invalid_object_format: invalid object format.  - immutable_store: operation on immutable store.  - max_number_token_issuers: reaching maximum number of token issuers.  - token_issuer_already_registered: token issuers already registered.  - tos_agreement_already_signed: agreement already signed.  - write_failed_due_to_invalid_input: write request failed due to invalid input.  - authorization_model_assertions_not_found: no assertions found for authorization model.  - settings_not_found: settings not found.  - latest_authorization_model_not_found: latest authorization model not found.  - type_not_found: type not found.  - relation_not_found: relation not found.  - empty_relation_definition: empty relation definition.  - too_many_types: too many types.  - invalid_user: invalid user.  - invalid_token_issuer: invalid token issuer.  - invalid_tuple: invalid tuple.  - unknown_relation: unknown relation.  - max_clients_exceeded: maximum clients exceeded.  - store_id_invalid_length: store id has invalid length.  - issuer_url_invalid_uri: issuer url has invalid URI.  - issuer_url_required_absolute_path: issuer url is not absolute path.  - assertions_too_many_items: assertions have too many items.  - id_too_long: ID is too long.  - invalid_environment: invalid environment is specified.  - authorization_model_id_too_long: authorization model id is too long.  - tuple_key_value_not_specified: tuple key value is not specified.  - tuple_keys_too_many_or_too_few_items: tuple keys have too few or too many items.  - page_size_invalid: page size is outside of acceptable range.  - param_missing_value: params value is missing.  - difference_base_missing_value: difference's base value is missing.  - subtract_base_missing_value: subtract base value is missing.  - object_too_long: object length is too long.  - relation_too_long: relation length is too long.  - type_definitions_too_few_items: type definitions do not have enough item.  - type_invalid_length: type length is invalid.  - type_invalid_pattern: type does not match expected pattern.  - relations_too_few_items: relations have too few items.  - relations_too_long: relations' length is too long.  - relations_invalid_pattern: relations do not match expected pattern.  - object_invalid_pattern: object does not match expected pattern.  - query_string_type_continuation_token_mismatch: type in the query string and the continuation token don't match.
type ErrorCode string

// List of ErrorCode
const (
	NO_ERROR                                         ErrorCode = "no_error"
	VALIDATION_ERROR                                 ErrorCode = "validation_error"
	AUTHORIZATION_MODEL_NOT_FOUND                    ErrorCode = "authorization_model_not_found"
	AUTHORIZATION_MODEL_RESOLUTION_TOO_COMPLEX       ErrorCode = "authorization_model_resolution_too_complex"
	INVALID_WRITE_INPUT                              ErrorCode = "invalid_write_input"
	CANNOT_ALLOW_DUPLICATE_TUPLES_IN_ONE_REQUEST     ErrorCode = "cannot_allow_duplicate_tuples_in_one_request"
	CANNOT_ALLOW_DUPLICATE_TYPES_IN_ONE_REQUEST      ErrorCode = "cannot_allow_duplicate_types_in_one_request"
	CANNOT_ALLOW_MULTIPLE_REFERENCES_TO_ONE_RELATION ErrorCode = "cannot_allow_multiple_references_to_one_relation"
	INVALID_CONTINUATION_TOKEN                       ErrorCode = "invalid_continuation_token"
	INVALID_TUPLE_SET                                ErrorCode = "invalid_tuple_set"
	INVALID_CHECK_INPUT                              ErrorCode = "invalid_check_input"
	INVALID_EXPAND_INPUT                             ErrorCode = "invalid_expand_input"
	UNSUPPORTED_USER_SET                             ErrorCode = "unsupported_user_set"
	INVALID_OBJECT_FORMAT                            ErrorCode = "invalid_object_format"
	IMMUTABLE_STORE                                  ErrorCode = "immutable_store"
	MAX_NUMBER_TOKEN_ISSUERS                         ErrorCode = "max_number_token_issuers"
	TOKEN_ISSUER_ALREADY_REGISTERED                  ErrorCode = "token_issuer_already_registered"
	TOS_AGREEMENT_ALREADY_SIGNED                     ErrorCode = "tos_agreement_already_signed"
	WRITE_FAILED_DUE_TO_INVALID_INPUT                ErrorCode = "write_failed_due_to_invalid_input"
	AUTHORIZATION_MODEL_ASSERTIONS_NOT_FOUND         ErrorCode = "authorization_model_assertions_not_found"
	SETTINGS_NOT_FOUND                               ErrorCode = "settings_not_found"
	LATEST_AUTHORIZATION_MODEL_NOT_FOUND             ErrorCode = "latest_authorization_model_not_found"
	TYPE_NOT_FOUND                                   ErrorCode = "type_not_found"
	RELATION_NOT_FOUND                               ErrorCode = "relation_not_found"
	EMPTY_RELATION_DEFINITION                        ErrorCode = "empty_relation_definition"
	TOO_MANY_TYPES                                   ErrorCode = "too_many_types"
	INVALID_USER                                     ErrorCode = "invalid_user"
	INVALID_TOKEN_ISSUER                             ErrorCode = "invalid_token_issuer"
	INVALID_TUPLE                                    ErrorCode = "invalid_tuple"
	UNKNOWN_RELATION                                 ErrorCode = "unknown_relation"
	MAX_CLIENTS_EXCEEDED                             ErrorCode = "max_clients_exceeded"
	STORE_ID_INVALID_LENGTH                          ErrorCode = "store_id_invalid_length"
	ISSUER_URL_INVALID_URI                           ErrorCode = "issuer_url_invalid_uri"
	ISSUER_URL_REQUIRED_ABSOLUTE_PATH                ErrorCode = "issuer_url_required_absolute_path"
	ASSERTIONS_TOO_MANY_ITEMS                        ErrorCode = "assertions_too_many_items"
	ID_TOO_LONG                                      ErrorCode = "id_too_long"
	INVALID_ENVIRONMENT                              ErrorCode = "invalid_environment"
	AUTHORIZATION_MODEL_ID_TOO_LONG                  ErrorCode = "authorization_model_id_too_long"
	TUPLE_KEY_VALUE_NOT_SPECIFIED                    ErrorCode = "tuple_key_value_not_specified"
	TUPLE_KEYS_TOO_MANY_OR_TOO_FEW_ITEMS             ErrorCode = "tuple_keys_too_many_or_too_few_items"
	PAGE_SIZE_INVALID                                ErrorCode = "page_size_invalid"
	PARAM_MISSING_VALUE                              ErrorCode = "param_missing_value"
	DIFFERENCE_BASE_MISSING_VALUE                    ErrorCode = "difference_base_missing_value"
	SUBTRACT_BASE_MISSING_VALUE                      ErrorCode = "subtract_base_missing_value"
	OBJECT_TOO_LONG                                  ErrorCode = "object_too_long"
	RELATION_TOO_LONG                                ErrorCode = "relation_too_long"
	TYPE_DEFINITIONS_TOO_FEW_ITEMS                   ErrorCode = "type_definitions_too_few_items"
	TYPE_INVALID_LENGTH                              ErrorCode = "type_invalid_length"
	TYPE_INVALID_PATTERN                             ErrorCode = "type_invalid_pattern"
	RELATIONS_TOO_FEW_ITEMS                          ErrorCode = "relations_too_few_items"
	RELATIONS_TOO_LONG                               ErrorCode = "relations_too_long"
	RELATIONS_INVALID_PATTERN                        ErrorCode = "relations_invalid_pattern"
	OBJECT_INVALID_PATTERN                           ErrorCode = "object_invalid_pattern"
	QUERY_STRING_TYPE_CONTINUATION_TOKEN_MISMATCH    ErrorCode = "query_string_type_continuation_token_mismatch"
)

var allowedErrorCodeEnumValues = []ErrorCode{
	"no_error",
	"validation_error",
	"authorization_model_not_found",
	"authorization_model_resolution_too_complex",
	"invalid_write_input",
	"cannot_allow_duplicate_tuples_in_one_request",
	"cannot_allow_duplicate_types_in_one_request",
	"cannot_allow_multiple_references_to_one_relation",
	"invalid_continuation_token",
	"invalid_tuple_set",
	"invalid_check_input",
	"invalid_expand_input",
	"unsupported_user_set",
	"invalid_object_format",
	"immutable_store",
	"max_number_token_issuers",
	"token_issuer_already_registered",
	"tos_agreement_already_signed",
	"write_failed_due_to_invalid_input",
	"authorization_model_assertions_not_found",
	"settings_not_found",
	"latest_authorization_model_not_found",
	"type_not_found",
	"relation_not_found",
	"empty_relation_definition",
	"too_many_types",
	"invalid_user",
	"invalid_token_issuer",
	"invalid_tuple",
	"unknown_relation",
	"max_clients_exceeded",
	"store_id_invalid_length",
	"issuer_url_invalid_uri",
	"issuer_url_required_absolute_path",
	"assertions_too_many_items",
	"id_too_long",
	"invalid_environment",
	"authorization_model_id_too_long",
	"tuple_key_value_not_specified",
	"tuple_keys_too_many_or_too_few_items",
	"page_size_invalid",
	"param_missing_value",
	"difference_base_missing_value",
	"subtract_base_missing_value",
	"object_too_long",
	"relation_too_long",
	"type_definitions_too_few_items",
	"type_invalid_length",
	"type_invalid_pattern",
	"relations_too_few_items",
	"relations_too_long",
	"relations_invalid_pattern",
	"object_invalid_pattern",
	"query_string_type_continuation_token_mismatch",
}

func (v *ErrorCode) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ErrorCode(value)
	for _, existing := range allowedErrorCodeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ErrorCode", value)
}

// NewErrorCodeFromValue returns a pointer to a valid ErrorCode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewErrorCodeFromValue(v string) (*ErrorCode, error) {
	ev := ErrorCode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ErrorCode: valid values are %v", v, allowedErrorCodeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ErrorCode) IsValid() bool {
	for _, existing := range allowedErrorCodeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ErrorCode value
func (v ErrorCode) Ptr() *ErrorCode {
	return &v
}

type NullableErrorCode struct {
	value *ErrorCode
	isSet bool
}

func (v NullableErrorCode) Get() *ErrorCode {
	return v.value
}

func (v *NullableErrorCode) Set(val *ErrorCode) {
	v.value = val
	v.isSet = true
}

func (v NullableErrorCode) IsSet() bool {
	return v.isSet
}

func (v *NullableErrorCode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableErrorCode(val *ErrorCode) *NullableErrorCode {
	return &NullableErrorCode{value: val, isSet: true}
}

func (v NullableErrorCode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableErrorCode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
