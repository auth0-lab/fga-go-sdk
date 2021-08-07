# Auth0FgaCheckRequestParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TupleKey** | Pointer to [**Auth0FgaTupleKey**](Auth0FgaTupleKey.md) |  | [optional] 
**AuthorizationModelId** | Pointer to **string** |  | [optional] 
**Trace** | Pointer to **bool** | defaults to false. making it true has performance implications. only use for debugging purposes, etc. | [optional] [readonly] 

## Methods

### NewAuth0FgaCheckRequestParams

`func NewAuth0FgaCheckRequestParams() *Auth0FgaCheckRequestParams`

NewAuth0FgaCheckRequestParams instantiates a new Auth0FgaCheckRequestParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuth0FgaCheckRequestParamsWithDefaults

`func NewAuth0FgaCheckRequestParamsWithDefaults() *Auth0FgaCheckRequestParams`

NewAuth0FgaCheckRequestParamsWithDefaults instantiates a new Auth0FgaCheckRequestParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTupleKey

`func (o *Auth0FgaCheckRequestParams) GetTupleKey() Auth0FgaTupleKey`

GetTupleKey returns the TupleKey field if non-nil, zero value otherwise.

### GetTupleKeyOk

`func (o *Auth0FgaCheckRequestParams) GetTupleKeyOk() (*Auth0FgaTupleKey, bool)`

GetTupleKeyOk returns a tuple with the TupleKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTupleKey

`func (o *Auth0FgaCheckRequestParams) SetTupleKey(v Auth0FgaTupleKey)`

SetTupleKey sets TupleKey field to given value.

### HasTupleKey

`func (o *Auth0FgaCheckRequestParams) HasTupleKey() bool`

HasTupleKey returns a boolean if a field has been set.

### GetAuthorizationModelId

`func (o *Auth0FgaCheckRequestParams) GetAuthorizationModelId() string`

GetAuthorizationModelId returns the AuthorizationModelId field if non-nil, zero value otherwise.

### GetAuthorizationModelIdOk

`func (o *Auth0FgaCheckRequestParams) GetAuthorizationModelIdOk() (*string, bool)`

GetAuthorizationModelIdOk returns a tuple with the AuthorizationModelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizationModelId

`func (o *Auth0FgaCheckRequestParams) SetAuthorizationModelId(v string)`

SetAuthorizationModelId sets AuthorizationModelId field to given value.

### HasAuthorizationModelId

`func (o *Auth0FgaCheckRequestParams) HasAuthorizationModelId() bool`

HasAuthorizationModelId returns a boolean if a field has been set.

### GetTrace

`func (o *Auth0FgaCheckRequestParams) GetTrace() bool`

GetTrace returns the Trace field if non-nil, zero value otherwise.

### GetTraceOk

`func (o *Auth0FgaCheckRequestParams) GetTraceOk() (*bool, bool)`

GetTraceOk returns a tuple with the Trace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrace

`func (o *Auth0FgaCheckRequestParams) SetTrace(v bool)`

SetTrace sets Trace field to given value.

### HasTrace

`func (o *Auth0FgaCheckRequestParams) HasTrace() bool`

HasTrace returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


