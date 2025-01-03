# DataByCityDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | City name | 
**Locations** | **string** | City location | 
**Streets** | [**[]DataByCityDataAttributesStreetsInner**](DataByCityDataAttributesStreetsInner.md) |  | 

## Methods

### NewDataByCityDataAttributes

`func NewDataByCityDataAttributes(name string, locations string, streets []DataByCityDataAttributesStreetsInner, ) *DataByCityDataAttributes`

NewDataByCityDataAttributes instantiates a new DataByCityDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataByCityDataAttributesWithDefaults

`func NewDataByCityDataAttributesWithDefaults() *DataByCityDataAttributes`

NewDataByCityDataAttributesWithDefaults instantiates a new DataByCityDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *DataByCityDataAttributes) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DataByCityDataAttributes) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DataByCityDataAttributes) SetName(v string)`

SetName sets Name field to given value.


### GetLocations

`func (o *DataByCityDataAttributes) GetLocations() string`

GetLocations returns the Locations field if non-nil, zero value otherwise.

### GetLocationsOk

`func (o *DataByCityDataAttributes) GetLocationsOk() (*string, bool)`

GetLocationsOk returns a tuple with the Locations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocations

`func (o *DataByCityDataAttributes) SetLocations(v string)`

SetLocations sets Locations field to given value.


### GetStreets

`func (o *DataByCityDataAttributes) GetStreets() []DataByCityDataAttributesStreetsInner`

GetStreets returns the Streets field if non-nil, zero value otherwise.

### GetStreetsOk

`func (o *DataByCityDataAttributes) GetStreetsOk() (*[]DataByCityDataAttributesStreetsInner, bool)`

GetStreetsOk returns a tuple with the Streets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreets

`func (o *DataByCityDataAttributes) SetStreets(v []DataByCityDataAttributesStreetsInner)`

SetStreets sets Streets field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


