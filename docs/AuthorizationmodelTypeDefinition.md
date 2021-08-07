# AuthorizationmodelTypeDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Relations** | [**map[string]AuthorizationmodelUserset**](AuthorizationmodelUserset.md) |  | 

## Methods

### NewAuthorizationmodelTypeDefinition

`func NewAuthorizationmodelTypeDefinition(type_ string, relations map[string]AuthorizationmodelUserset, ) *AuthorizationmodelTypeDefinition`

NewAuthorizationmodelTypeDefinition instantiates a new AuthorizationmodelTypeDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthorizationmodelTypeDefinitionWithDefaults

`func NewAuthorizationmodelTypeDefinitionWithDefaults() *AuthorizationmodelTypeDefinition`

NewAuthorizationmodelTypeDefinitionWithDefaults instantiates a new AuthorizationmodelTypeDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *AuthorizationmodelTypeDefinition) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AuthorizationmodelTypeDefinition) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AuthorizationmodelTypeDefinition) SetType(v string)`

SetType sets Type field to given value.


### GetRelations

`func (o *AuthorizationmodelTypeDefinition) GetRelations() map[string]AuthorizationmodelUserset`

GetRelations returns the Relations field if non-nil, zero value otherwise.

### GetRelationsOk

`func (o *AuthorizationmodelTypeDefinition) GetRelationsOk() (*map[string]AuthorizationmodelUserset, bool)`

GetRelationsOk returns a tuple with the Relations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelations

`func (o *AuthorizationmodelTypeDefinition) SetRelations(v map[string]AuthorizationmodelUserset)`

SetRelations sets Relations field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


