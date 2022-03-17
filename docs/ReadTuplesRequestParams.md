# ReadTuplesRequestParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PageSize** | Pointer to **int32** |  | [optional] 
**ContinuationToken** | Pointer to **string** |  | [optional] 

## Methods

### NewReadTuplesRequestParams

`func NewReadTuplesRequestParams() *ReadTuplesRequestParams`

NewReadTuplesRequestParams instantiates a new ReadTuplesRequestParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReadTuplesRequestParamsWithDefaults

`func NewReadTuplesRequestParamsWithDefaults() *ReadTuplesRequestParams`

NewReadTuplesRequestParamsWithDefaults instantiates a new ReadTuplesRequestParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPageSize

`func (o *ReadTuplesRequestParams) GetPageSize() int32`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *ReadTuplesRequestParams) GetPageSizeOk() (*int32, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *ReadTuplesRequestParams) SetPageSize(v int32)`

SetPageSize sets PageSize field to given value.

### HasPageSize

`func (o *ReadTuplesRequestParams) HasPageSize() bool`

HasPageSize returns a boolean if a field has been set.

### GetContinuationToken

`func (o *ReadTuplesRequestParams) GetContinuationToken() string`

GetContinuationToken returns the ContinuationToken field if non-nil, zero value otherwise.

### GetContinuationTokenOk

`func (o *ReadTuplesRequestParams) GetContinuationTokenOk() (*string, bool)`

GetContinuationTokenOk returns a tuple with the ContinuationToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContinuationToken

`func (o *ReadTuplesRequestParams) SetContinuationToken(v string)`

SetContinuationToken sets ContinuationToken field to given value.

### HasContinuationToken

`func (o *ReadTuplesRequestParams) HasContinuationToken() bool`

HasContinuationToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


