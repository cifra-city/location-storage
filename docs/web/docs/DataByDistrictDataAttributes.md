# DataByDistrictDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | district name | 
**City** | **string** | city uuid | 
**Streets** | [**[]DataByDistrictDataAttributesStreetsInner**](DataByDistrictDataAttributesStreetsInner.md) |  | 

## Methods

### NewDataByDistrictDataAttributes

`func NewDataByDistrictDataAttributes(name string, city string, streets []DataByDistrictDataAttributesStreetsInner, ) *DataByDistrictDataAttributes`

NewDataByDistrictDataAttributes instantiates a new DataByDistrictDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataByDistrictDataAttributesWithDefaults

`func NewDataByDistrictDataAttributesWithDefaults() *DataByDistrictDataAttributes`

NewDataByDistrictDataAttributesWithDefaults instantiates a new DataByDistrictDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *DataByDistrictDataAttributes) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DataByDistrictDataAttributes) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DataByDistrictDataAttributes) SetName(v string)`

SetName sets Name field to given value.


### GetCity

`func (o *DataByDistrictDataAttributes) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *DataByDistrictDataAttributes) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *DataByDistrictDataAttributes) SetCity(v string)`

SetCity sets City field to given value.


### GetStreets

`func (o *DataByDistrictDataAttributes) GetStreets() []DataByDistrictDataAttributesStreetsInner`

GetStreets returns the Streets field if non-nil, zero value otherwise.

### GetStreetsOk

`func (o *DataByDistrictDataAttributes) GetStreetsOk() (*[]DataByDistrictDataAttributesStreetsInner, bool)`

GetStreetsOk returns a tuple with the Streets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreets

`func (o *DataByDistrictDataAttributes) SetStreets(v []DataByDistrictDataAttributesStreetsInner)`

SetStreets sets Streets field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


