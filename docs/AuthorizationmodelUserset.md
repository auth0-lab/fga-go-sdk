# AuthorizationmodelUserset

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**This** | Pointer to **map[string]interface{}** | A DirectUserset is a sentinel message for referencing the direct members specified by an object/relation mapping. | [optional] 
**ComputedUserset** | Pointer to [**AuthorizationmodelObjectRelation**](AuthorizationmodelObjectRelation.md) |  | [optional] 
**TupleToUserset** | Pointer to [**AuthorizationmodelTupleToUserset**](AuthorizationmodelTupleToUserset.md) |  | [optional] 
**Union** | Pointer to [**AuthorizationmodelUsersets**](AuthorizationmodelUsersets.md) |  | [optional] 
**Intersection** | Pointer to [**AuthorizationmodelUsersets**](AuthorizationmodelUsersets.md) |  | [optional] 
**Difference** | Pointer to [**AuthorizationmodelDifference**](AuthorizationmodelDifference.md) |  | [optional] 

## Methods

### NewAuthorizationmodelUserset

`func NewAuthorizationmodelUserset() *AuthorizationmodelUserset`

NewAuthorizationmodelUserset instantiates a new AuthorizationmodelUserset object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthorizationmodelUsersetWithDefaults

`func NewAuthorizationmodelUsersetWithDefaults() *AuthorizationmodelUserset`

NewAuthorizationmodelUsersetWithDefaults instantiates a new AuthorizationmodelUserset object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetThis

`func (o *AuthorizationmodelUserset) GetThis() map[string]interface{}`

GetThis returns the This field if non-nil, zero value otherwise.

### GetThisOk

`func (o *AuthorizationmodelUserset) GetThisOk() (*map[string]interface{}, bool)`

GetThisOk returns a tuple with the This field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThis

`func (o *AuthorizationmodelUserset) SetThis(v map[string]interface{})`

SetThis sets This field to given value.

### HasThis

`func (o *AuthorizationmodelUserset) HasThis() bool`

HasThis returns a boolean if a field has been set.

### GetComputedUserset

`func (o *AuthorizationmodelUserset) GetComputedUserset() AuthorizationmodelObjectRelation`

GetComputedUserset returns the ComputedUserset field if non-nil, zero value otherwise.

### GetComputedUsersetOk

`func (o *AuthorizationmodelUserset) GetComputedUsersetOk() (*AuthorizationmodelObjectRelation, bool)`

GetComputedUsersetOk returns a tuple with the ComputedUserset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputedUserset

`func (o *AuthorizationmodelUserset) SetComputedUserset(v AuthorizationmodelObjectRelation)`

SetComputedUserset sets ComputedUserset field to given value.

### HasComputedUserset

`func (o *AuthorizationmodelUserset) HasComputedUserset() bool`

HasComputedUserset returns a boolean if a field has been set.

### GetTupleToUserset

`func (o *AuthorizationmodelUserset) GetTupleToUserset() AuthorizationmodelTupleToUserset`

GetTupleToUserset returns the TupleToUserset field if non-nil, zero value otherwise.

### GetTupleToUsersetOk

`func (o *AuthorizationmodelUserset) GetTupleToUsersetOk() (*AuthorizationmodelTupleToUserset, bool)`

GetTupleToUsersetOk returns a tuple with the TupleToUserset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTupleToUserset

`func (o *AuthorizationmodelUserset) SetTupleToUserset(v AuthorizationmodelTupleToUserset)`

SetTupleToUserset sets TupleToUserset field to given value.

### HasTupleToUserset

`func (o *AuthorizationmodelUserset) HasTupleToUserset() bool`

HasTupleToUserset returns a boolean if a field has been set.

### GetUnion

`func (o *AuthorizationmodelUserset) GetUnion() AuthorizationmodelUsersets`

GetUnion returns the Union field if non-nil, zero value otherwise.

### GetUnionOk

`func (o *AuthorizationmodelUserset) GetUnionOk() (*AuthorizationmodelUsersets, bool)`

GetUnionOk returns a tuple with the Union field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnion

`func (o *AuthorizationmodelUserset) SetUnion(v AuthorizationmodelUsersets)`

SetUnion sets Union field to given value.

### HasUnion

`func (o *AuthorizationmodelUserset) HasUnion() bool`

HasUnion returns a boolean if a field has been set.

### GetIntersection

`func (o *AuthorizationmodelUserset) GetIntersection() AuthorizationmodelUsersets`

GetIntersection returns the Intersection field if non-nil, zero value otherwise.

### GetIntersectionOk

`func (o *AuthorizationmodelUserset) GetIntersectionOk() (*AuthorizationmodelUsersets, bool)`

GetIntersectionOk returns a tuple with the Intersection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntersection

`func (o *AuthorizationmodelUserset) SetIntersection(v AuthorizationmodelUsersets)`

SetIntersection sets Intersection field to given value.

### HasIntersection

`func (o *AuthorizationmodelUserset) HasIntersection() bool`

HasIntersection returns a boolean if a field has been set.

### GetDifference

`func (o *AuthorizationmodelUserset) GetDifference() AuthorizationmodelDifference`

GetDifference returns the Difference field if non-nil, zero value otherwise.

### GetDifferenceOk

`func (o *AuthorizationmodelUserset) GetDifferenceOk() (*AuthorizationmodelDifference, bool)`

GetDifferenceOk returns a tuple with the Difference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDifference

`func (o *AuthorizationmodelUserset) SetDifference(v AuthorizationmodelDifference)`

SetDifference sets Difference field to given value.

### HasDifference

`func (o *AuthorizationmodelUserset) HasDifference() bool`

HasDifference returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


