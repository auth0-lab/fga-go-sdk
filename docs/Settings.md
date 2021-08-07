# Settings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Environment** | Pointer to [**Environment**](Environment.md) |  | [optional] [default to ENVIRONMENT_UNSPECIFIED]
**TokenIssuers** | Pointer to [**[]TokenIssuer**](TokenIssuer.md) |  | [optional] 

## Methods

### NewSettings

`func NewSettings() *Settings`

NewSettings instantiates a new Settings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSettingsWithDefaults

`func NewSettingsWithDefaults() *Settings`

NewSettingsWithDefaults instantiates a new Settings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnvironment

`func (o *Settings) GetEnvironment() Environment`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *Settings) GetEnvironmentOk() (*Environment, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *Settings) SetEnvironment(v Environment)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *Settings) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetTokenIssuers

`func (o *Settings) GetTokenIssuers() []TokenIssuer`

GetTokenIssuers returns the TokenIssuers field if non-nil, zero value otherwise.

### GetTokenIssuersOk

`func (o *Settings) GetTokenIssuersOk() (*[]TokenIssuer, bool)`

GetTokenIssuersOk returns a tuple with the TokenIssuers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenIssuers

`func (o *Settings) SetTokenIssuers(v []TokenIssuer)`

SetTokenIssuers sets TokenIssuers field to given value.

### HasTokenIssuers

`func (o *Settings) HasTokenIssuers() bool`

HasTokenIssuers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


