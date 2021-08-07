# SettingsSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Environment** | Pointer to [**SettingsEnvironment**](SettingsEnvironment.md) |  | [optional] [default to ENVIRONMENT_UNSPECIFIED]
**TokenIssuers** | Pointer to [**[]SettingsTokenIssuer**](SettingsTokenIssuer.md) |  | [optional] 

## Methods

### NewSettingsSettings

`func NewSettingsSettings() *SettingsSettings`

NewSettingsSettings instantiates a new SettingsSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSettingsSettingsWithDefaults

`func NewSettingsSettingsWithDefaults() *SettingsSettings`

NewSettingsSettingsWithDefaults instantiates a new SettingsSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnvironment

`func (o *SettingsSettings) GetEnvironment() SettingsEnvironment`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *SettingsSettings) GetEnvironmentOk() (*SettingsEnvironment, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *SettingsSettings) SetEnvironment(v SettingsEnvironment)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *SettingsSettings) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetTokenIssuers

`func (o *SettingsSettings) GetTokenIssuers() []SettingsTokenIssuer`

GetTokenIssuers returns the TokenIssuers field if non-nil, zero value otherwise.

### GetTokenIssuersOk

`func (o *SettingsSettings) GetTokenIssuersOk() (*[]SettingsTokenIssuer, bool)`

GetTokenIssuersOk returns a tuple with the TokenIssuers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenIssuers

`func (o *SettingsSettings) SetTokenIssuers(v []SettingsTokenIssuer)`

SetTokenIssuers sets TokenIssuers field to given value.

### HasTokenIssuers

`func (o *SettingsSettings) HasTokenIssuers() bool`

HasTokenIssuers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


