# DataByCountryDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Country name | 
**Cities** | [**[]DataByCountryDataAttributesCitiesInner**](DataByCountryDataAttributesCitiesInner.md) |  | 

## Methods

### NewDataByCountryDataAttributes

`func NewDataByCountryDataAttributes(name string, cities []DataByCountryDataAttributesCitiesInner, ) *DataByCountryDataAttributes`

NewDataByCountryDataAttributes instantiates a new DataByCountryDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataByCountryDataAttributesWithDefaults

`func NewDataByCountryDataAttributesWithDefaults() *DataByCountryDataAttributes`

NewDataByCountryDataAttributesWithDefaults instantiates a new DataByCountryDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *DataByCountryDataAttributes) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DataByCountryDataAttributes) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DataByCountryDataAttributes) SetName(v string)`

SetName sets Name field to given value.


### GetCities

`func (o *DataByCountryDataAttributes) GetCities() []DataByCountryDataAttributesCitiesInner`

GetCities returns the Cities field if non-nil, zero value otherwise.

### GetCitiesOk

`func (o *DataByCountryDataAttributes) GetCitiesOk() (*[]DataByCountryDataAttributesCitiesInner, bool)`

GetCitiesOk returns a tuple with the Cities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCities

`func (o *DataByCountryDataAttributes) SetCities(v []DataByCountryDataAttributesCitiesInner)`

SetCities sets Cities field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


