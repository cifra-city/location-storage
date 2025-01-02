# CreateCityDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | City name | 
**CityId** | **string** | City uuid | 
**CountryId** | Pointer to **string** | Country uuid if u need to change country for the city | [optional] 

## Methods

### NewCreateCityDataAttributes

`func NewCreateCityDataAttributes(name string, cityId string, ) *CreateCityDataAttributes`

NewCreateCityDataAttributes instantiates a new CreateCityDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateCityDataAttributesWithDefaults

`func NewCreateCityDataAttributesWithDefaults() *CreateCityDataAttributes`

NewCreateCityDataAttributesWithDefaults instantiates a new CreateCityDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateCityDataAttributes) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateCityDataAttributes) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateCityDataAttributes) SetName(v string)`

SetName sets Name field to given value.


### GetCityId

`func (o *CreateCityDataAttributes) GetCityId() string`

GetCityId returns the CityId field if non-nil, zero value otherwise.

### GetCityIdOk

`func (o *CreateCityDataAttributes) GetCityIdOk() (*string, bool)`

GetCityIdOk returns a tuple with the CityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCityId

`func (o *CreateCityDataAttributes) SetCityId(v string)`

SetCityId sets CityId field to given value.


### GetCountryId

`func (o *CreateCityDataAttributes) GetCountryId() string`

GetCountryId returns the CountryId field if non-nil, zero value otherwise.

### GetCountryIdOk

`func (o *CreateCityDataAttributes) GetCountryIdOk() (*string, bool)`

GetCountryIdOk returns a tuple with the CountryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryId

`func (o *CreateCityDataAttributes) SetCountryId(v string)`

SetCountryId sets CountryId field to given value.

### HasCountryId

`func (o *CreateCityDataAttributes) HasCountryId() bool`

HasCountryId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


