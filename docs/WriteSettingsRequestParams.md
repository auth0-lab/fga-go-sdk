# WriteSettingsRequestParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Environment** | Pointer to [**Environment**](Environment.md) |  | [optional] [default to ENVIRONMENT_UNSPECIFIED]

## Methods

### NewWriteSettingsRequestParams

`func NewWriteSettingsRequestParams() *WriteSettingsRequestParams`

NewWriteSettingsRequestParams instantiates a new WriteSettingsRequestParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWriteSettingsRequestParamsWithDefaults

`func NewWriteSettingsRequestParamsWithDefaults() *WriteSettingsRequestParams`

NewWriteSettingsRequestParamsWithDefaults instantiates a new WriteSettingsRequestParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnvironment

`func (o *WriteSettingsRequestParams) GetEnvironment() Environment`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *WriteSettingsRequestParams) GetEnvironmentOk() (*Environment, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *WriteSettingsRequestParams) SetEnvironment(v Environment)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *WriteSettingsRequestParams) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


