# Auth0FgaReadAssertionsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthorizationModelId** | Pointer to **string** | The authorization model ID | [optional] 
**Assertions** | Pointer to [**[]Auth0FgaAssertion**](Auth0FgaAssertion.md) |  | [optional] 

## Methods

### NewAuth0FgaReadAssertionsResponse

`func NewAuth0FgaReadAssertionsResponse() *Auth0FgaReadAssertionsResponse`

NewAuth0FgaReadAssertionsResponse instantiates a new Auth0FgaReadAssertionsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuth0FgaReadAssertionsResponseWithDefaults

`func NewAuth0FgaReadAssertionsResponseWithDefaults() *Auth0FgaReadAssertionsResponse`

NewAuth0FgaReadAssertionsResponseWithDefaults instantiates a new Auth0FgaReadAssertionsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthorizationModelId

`func (o *Auth0FgaReadAssertionsResponse) GetAuthorizationModelId() string`

GetAuthorizationModelId returns the AuthorizationModelId field if non-nil, zero value otherwise.

### GetAuthorizationModelIdOk

`func (o *Auth0FgaReadAssertionsResponse) GetAuthorizationModelIdOk() (*string, bool)`

GetAuthorizationModelIdOk returns a tuple with the AuthorizationModelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizationModelId

`func (o *Auth0FgaReadAssertionsResponse) SetAuthorizationModelId(v string)`

SetAuthorizationModelId sets AuthorizationModelId field to given value.

### HasAuthorizationModelId

`func (o *Auth0FgaReadAssertionsResponse) HasAuthorizationModelId() bool`

HasAuthorizationModelId returns a boolean if a field has been set.

### GetAssertions

`func (o *Auth0FgaReadAssertionsResponse) GetAssertions() []Auth0FgaAssertion`

GetAssertions returns the Assertions field if non-nil, zero value otherwise.

### GetAssertionsOk

`func (o *Auth0FgaReadAssertionsResponse) GetAssertionsOk() (*[]Auth0FgaAssertion, bool)`

GetAssertionsOk returns a tuple with the Assertions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssertions

`func (o *Auth0FgaReadAssertionsResponse) SetAssertions(v []Auth0FgaAssertion)`

SetAssertions sets Assertions field to given value.

### HasAssertions

`func (o *Auth0FgaReadAssertionsResponse) HasAssertions() bool`

HasAssertions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


