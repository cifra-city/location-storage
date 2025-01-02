# GetCityDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | City name | 
**Country** | **string** | Country name | 
**Districts** | [**[]GetCityDataAttributesDistrictsInner**](GetCityDataAttributesDistrictsInner.md) |  | 

## Methods

### NewGetCityDataAttributes

`func NewGetCityDataAttributes(name string, country string, districts []GetCityDataAttributesDistrictsInner, ) *GetCityDataAttributes`

NewGetCityDataAttributes instantiates a new GetCityDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetCityDataAttributesWithDefaults

`func NewGetCityDataAttributesWithDefaults() *GetCityDataAttributes`

NewGetCityDataAttributesWithDefaults instantiates a new GetCityDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *GetCityDataAttributes) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetCityDataAttributes) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetCityDataAttributes) SetName(v string)`

SetName sets Name field to given value.


### GetCountry

`func (o *GetCityDataAttributes) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *GetCityDataAttributes) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *GetCityDataAttributes) SetCountry(v string)`

SetCountry sets Country field to given value.


### GetDistricts

`func (o *GetCityDataAttributes) GetDistricts() []GetCityDataAttributesDistrictsInner`

GetDistricts returns the Districts field if non-nil, zero value otherwise.

### GetDistrictsOk

`func (o *GetCityDataAttributes) GetDistrictsOk() (*[]GetCityDataAttributesDistrictsInner, bool)`

GetDistrictsOk returns a tuple with the Districts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistricts

`func (o *GetCityDataAttributes) SetDistricts(v []GetCityDataAttributesDistrictsInner)`

SetDistricts sets Districts field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


