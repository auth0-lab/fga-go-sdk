# UsersetTreeLeaf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Users** | Pointer to [**UsersetTreeUsers**](UsersetTreeUsers.md) |  | [optional] 
**Computed** | Pointer to [**UsersetTreeComputed**](UsersetTreeComputed.md) |  | [optional] 
**TupleToUserset** | Pointer to [**Auth0FgaUsersetTreeTupleToUserset**](Auth0FgaUsersetTreeTupleToUserset.md) |  | [optional] 

## Methods

### NewUsersetTreeLeaf

`func NewUsersetTreeLeaf() *UsersetTreeLeaf`

NewUsersetTreeLeaf instantiates a new UsersetTreeLeaf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUsersetTreeLeafWithDefaults

`func NewUsersetTreeLeafWithDefaults() *UsersetTreeLeaf`

NewUsersetTreeLeafWithDefaults instantiates a new UsersetTreeLeaf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsers

`func (o *UsersetTreeLeaf) GetUsers() UsersetTreeUsers`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *UsersetTreeLeaf) GetUsersOk() (*UsersetTreeUsers, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *UsersetTreeLeaf) SetUsers(v UsersetTreeUsers)`

SetUsers sets Users field to given value.

### HasUsers

`func (o *UsersetTreeLeaf) HasUsers() bool`

HasUsers returns a boolean if a field has been set.

### GetComputed

`func (o *UsersetTreeLeaf) GetComputed() UsersetTreeComputed`

GetComputed returns the Computed field if non-nil, zero value otherwise.

### GetComputedOk

`func (o *UsersetTreeLeaf) GetComputedOk() (*UsersetTreeComputed, bool)`

GetComputedOk returns a tuple with the Computed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputed

`func (o *UsersetTreeLeaf) SetComputed(v UsersetTreeComputed)`

SetComputed sets Computed field to given value.

### HasComputed

`func (o *UsersetTreeLeaf) HasComputed() bool`

HasComputed returns a boolean if a field has been set.

### GetTupleToUserset

`func (o *UsersetTreeLeaf) GetTupleToUserset() Auth0FgaUsersetTreeTupleToUserset`

GetTupleToUserset returns the TupleToUserset field if non-nil, zero value otherwise.

### GetTupleToUsersetOk

`func (o *UsersetTreeLeaf) GetTupleToUsersetOk() (*Auth0FgaUsersetTreeTupleToUserset, bool)`

GetTupleToUsersetOk returns a tuple with the TupleToUserset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTupleToUserset

`func (o *UsersetTreeLeaf) SetTupleToUserset(v Auth0FgaUsersetTreeTupleToUserset)`

SetTupleToUserset sets TupleToUserset field to given value.

### HasTupleToUserset

`func (o *UsersetTreeLeaf) HasTupleToUserset() bool`

HasTupleToUserset returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


