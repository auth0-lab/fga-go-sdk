# UsersetTreeNode

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Leaf** | Pointer to [**UsersetTreeLeaf**](UsersetTreeLeaf.md) |  | [optional] 
**Difference** | Pointer to [**Auth0FgaUsersetTreeDifference**](Auth0FgaUsersetTreeDifference.md) |  | [optional] 
**Union** | Pointer to [**UsersetTreeNodes**](UsersetTreeNodes.md) |  | [optional] 
**Intersection** | Pointer to [**UsersetTreeNodes**](UsersetTreeNodes.md) |  | [optional] 

## Methods

### NewUsersetTreeNode

`func NewUsersetTreeNode() *UsersetTreeNode`

NewUsersetTreeNode instantiates a new UsersetTreeNode object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUsersetTreeNodeWithDefaults

`func NewUsersetTreeNodeWithDefaults() *UsersetTreeNode`

NewUsersetTreeNodeWithDefaults instantiates a new UsersetTreeNode object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UsersetTreeNode) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UsersetTreeNode) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UsersetTreeNode) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UsersetTreeNode) HasName() bool`

HasName returns a boolean if a field has been set.

### GetLeaf

`func (o *UsersetTreeNode) GetLeaf() UsersetTreeLeaf`

GetLeaf returns the Leaf field if non-nil, zero value otherwise.

### GetLeafOk

`func (o *UsersetTreeNode) GetLeafOk() (*UsersetTreeLeaf, bool)`

GetLeafOk returns a tuple with the Leaf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeaf

`func (o *UsersetTreeNode) SetLeaf(v UsersetTreeLeaf)`

SetLeaf sets Leaf field to given value.

### HasLeaf

`func (o *UsersetTreeNode) HasLeaf() bool`

HasLeaf returns a boolean if a field has been set.

### GetDifference

`func (o *UsersetTreeNode) GetDifference() Auth0FgaUsersetTreeDifference`

GetDifference returns the Difference field if non-nil, zero value otherwise.

### GetDifferenceOk

`func (o *UsersetTreeNode) GetDifferenceOk() (*Auth0FgaUsersetTreeDifference, bool)`

GetDifferenceOk returns a tuple with the Difference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDifference

`func (o *UsersetTreeNode) SetDifference(v Auth0FgaUsersetTreeDifference)`

SetDifference sets Difference field to given value.

### HasDifference

`func (o *UsersetTreeNode) HasDifference() bool`

HasDifference returns a boolean if a field has been set.

### GetUnion

`func (o *UsersetTreeNode) GetUnion() UsersetTreeNodes`

GetUnion returns the Union field if non-nil, zero value otherwise.

### GetUnionOk

`func (o *UsersetTreeNode) GetUnionOk() (*UsersetTreeNodes, bool)`

GetUnionOk returns a tuple with the Union field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnion

`func (o *UsersetTreeNode) SetUnion(v UsersetTreeNodes)`

SetUnion sets Union field to given value.

### HasUnion

`func (o *UsersetTreeNode) HasUnion() bool`

HasUnion returns a boolean if a field has been set.

### GetIntersection

`func (o *UsersetTreeNode) GetIntersection() UsersetTreeNodes`

GetIntersection returns the Intersection field if non-nil, zero value otherwise.

### GetIntersectionOk

`func (o *UsersetTreeNode) GetIntersectionOk() (*UsersetTreeNodes, bool)`

GetIntersectionOk returns a tuple with the Intersection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntersection

`func (o *UsersetTreeNode) SetIntersection(v UsersetTreeNodes)`

SetIntersection sets Intersection field to given value.

### HasIntersection

`func (o *UsersetTreeNode) HasIntersection() bool`

HasIntersection returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


