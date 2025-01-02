# GetStreetData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | street id | 
**Type** | **string** |  | 
**Attributes** | [**GetStreetDataAttributes**](GetStreetDataAttributes.md) |  | 

## Methods

### NewGetStreetData

`func NewGetStreetData(id string, type_ string, attributes GetStreetDataAttributes, ) *GetStreetData`

NewGetStreetData instantiates a new GetStreetData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetStreetDataWithDefaults

`func NewGetStreetDataWithDefaults() *GetStreetData`

NewGetStreetDataWithDefaults instantiates a new GetStreetData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetStreetData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetStreetData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetStreetData) SetId(v string)`

SetId sets Id field to given value.


### GetType

`func (o *GetStreetData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetStreetData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetStreetData) SetType(v string)`

SetType sets Type field to given value.


### GetAttributes

`func (o *GetStreetData) GetAttributes() GetStreetDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *GetStreetData) GetAttributesOk() (*GetStreetDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *GetStreetData) SetAttributes(v GetStreetDataAttributes)`

SetAttributes sets Attributes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


