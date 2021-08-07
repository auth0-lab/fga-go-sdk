# Auth0FgaWriteRequestParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Writes** | Pointer to [**Auth0FgaTupleKeys**](Auth0FgaTupleKeys.md) |  | [optional] 
**Deletes** | Pointer to [**Auth0FgaTupleKeys**](Auth0FgaTupleKeys.md) |  | [optional] 
**AuthorizationModelId** | Pointer to **string** |  | [optional] 
**LockTuple** | Pointer to [**Auth0FgaTuple**](Auth0FgaTuple.md) |  | [optional] 

## Methods

### NewAuth0FgaWriteRequestParams

`func NewAuth0FgaWriteRequestParams() *Auth0FgaWriteRequestParams`

NewAuth0FgaWriteRequestParams instantiates a new Auth0FgaWriteRequestParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuth0FgaWriteRequestParamsWithDefaults

`func NewAuth0FgaWriteRequestParamsWithDefaults() *Auth0FgaWriteRequestParams`

NewAuth0FgaWriteRequestParamsWithDefaults instantiates a new Auth0FgaWriteRequestParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWrites

`func (o *Auth0FgaWriteRequestParams) GetWrites() Auth0FgaTupleKeys`

GetWrites returns the Writes field if non-nil, zero value otherwise.

### GetWritesOk

`func (o *Auth0FgaWriteRequestParams) GetWritesOk() (*Auth0FgaTupleKeys, bool)`

GetWritesOk returns a tuple with the Writes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWrites

`func (o *Auth0FgaWriteRequestParams) SetWrites(v Auth0FgaTupleKeys)`

SetWrites sets Writes field to given value.

### HasWrites

`func (o *Auth0FgaWriteRequestParams) HasWrites() bool`

HasWrites returns a boolean if a field has been set.

### GetDeletes

`func (o *Auth0FgaWriteRequestParams) GetDeletes() Auth0FgaTupleKeys`

GetDeletes returns the Deletes field if non-nil, zero value otherwise.

### GetDeletesOk

`func (o *Auth0FgaWriteRequestParams) GetDeletesOk() (*Auth0FgaTupleKeys, bool)`

GetDeletesOk returns a tuple with the Deletes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletes

`func (o *Auth0FgaWriteRequestParams) SetDeletes(v Auth0FgaTupleKeys)`

SetDeletes sets Deletes field to given value.

### HasDeletes

`func (o *Auth0FgaWriteRequestParams) HasDeletes() bool`

HasDeletes returns a boolean if a field has been set.

### GetAuthorizationModelId

`func (o *Auth0FgaWriteRequestParams) GetAuthorizationModelId() string`

GetAuthorizationModelId returns the AuthorizationModelId field if non-nil, zero value otherwise.

### GetAuthorizationModelIdOk

`func (o *Auth0FgaWriteRequestParams) GetAuthorizationModelIdOk() (*string, bool)`

GetAuthorizationModelIdOk returns a tuple with the AuthorizationModelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizationModelId

`func (o *Auth0FgaWriteRequestParams) SetAuthorizationModelId(v string)`

SetAuthorizationModelId sets AuthorizationModelId field to given value.

### HasAuthorizationModelId

`func (o *Auth0FgaWriteRequestParams) HasAuthorizationModelId() bool`

HasAuthorizationModelId returns a boolean if a field has been set.

### GetLockTuple

`func (o *Auth0FgaWriteRequestParams) GetLockTuple() Auth0FgaTuple`

GetLockTuple returns the LockTuple field if non-nil, zero value otherwise.

### GetLockTupleOk

`func (o *Auth0FgaWriteRequestParams) GetLockTupleOk() (*Auth0FgaTuple, bool)`

GetLockTupleOk returns a tuple with the LockTuple field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockTuple

`func (o *Auth0FgaWriteRequestParams) SetLockTuple(v Auth0FgaTuple)`

SetLockTuple sets LockTuple field to given value.

### HasLockTuple

`func (o *Auth0FgaWriteRequestParams) HasLockTuple() bool`

HasLockTuple returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


