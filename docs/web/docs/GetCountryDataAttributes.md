# GetCountryDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Country name | 
**Cities** | [**[]GetCountryDataAttributesCitiesInner**](GetCountryDataAttributesCitiesInner.md) |  | 

## Methods

### NewGetCountryDataAttributes

`func NewGetCountryDataAttributes(name string, cities []GetCountryDataAttributesCitiesInner, ) *GetCountryDataAttributes`

NewGetCountryDataAttributes instantiates a new GetCountryDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetCountryDataAttributesWithDefaults

`func NewGetCountryDataAttributesWithDefaults() *GetCountryDataAttributes`

NewGetCountryDataAttributesWithDefaults instantiates a new GetCountryDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *GetCountryDataAttributes) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetCountryDataAttributes) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetCountryDataAttributes) SetName(v string)`

SetName sets Name field to given value.


### GetCities

`func (o *GetCountryDataAttributes) GetCities() []GetCountryDataAttributesCitiesInner`

GetCities returns the Cities field if non-nil, zero value otherwise.

### GetCitiesOk

`func (o *GetCountryDataAttributes) GetCitiesOk() (*[]GetCountryDataAttributesCitiesInner, bool)`

GetCitiesOk returns a tuple with the Cities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCities

`func (o *GetCountryDataAttributes) SetCities(v []GetCountryDataAttributesCitiesInner)`

SetCities sets Cities field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


