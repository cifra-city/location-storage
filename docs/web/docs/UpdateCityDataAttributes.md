# UpdateCityDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CityId** | **string** | City uuid | 
**NewName** | Pointer to **string** | City new name | [optional] 
**CountryId** | Pointer to **string** | Country uuid if you want to change country | [optional] 

## Methods

### NewUpdateCityDataAttributes

`func NewUpdateCityDataAttributes(cityId string, ) *UpdateCityDataAttributes`

NewUpdateCityDataAttributes instantiates a new UpdateCityDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateCityDataAttributesWithDefaults

`func NewUpdateCityDataAttributesWithDefaults() *UpdateCityDataAttributes`

NewUpdateCityDataAttributesWithDefaults instantiates a new UpdateCityDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCityId

`func (o *UpdateCityDataAttributes) GetCityId() string`

GetCityId returns the CityId field if non-nil, zero value otherwise.

### GetCityIdOk

`func (o *UpdateCityDataAttributes) GetCityIdOk() (*string, bool)`

GetCityIdOk returns a tuple with the CityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCityId

`func (o *UpdateCityDataAttributes) SetCityId(v string)`

SetCityId sets CityId field to given value.


### GetNewName

`func (o *UpdateCityDataAttributes) GetNewName() string`

GetNewName returns the NewName field if non-nil, zero value otherwise.

### GetNewNameOk

`func (o *UpdateCityDataAttributes) GetNewNameOk() (*string, bool)`

GetNewNameOk returns a tuple with the NewName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewName

`func (o *UpdateCityDataAttributes) SetNewName(v string)`

SetNewName sets NewName field to given value.

### HasNewName

`func (o *UpdateCityDataAttributes) HasNewName() bool`

HasNewName returns a boolean if a field has been set.

### GetCountryId

`func (o *UpdateCityDataAttributes) GetCountryId() string`

GetCountryId returns the CountryId field if non-nil, zero value otherwise.

### GetCountryIdOk

`func (o *UpdateCityDataAttributes) GetCountryIdOk() (*string, bool)`

GetCountryIdOk returns a tuple with the CountryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryId

`func (o *UpdateCityDataAttributes) SetCountryId(v string)`

SetCountryId sets CountryId field to given value.

### HasCountryId

`func (o *UpdateCityDataAttributes) HasCountryId() bool`

HasCountryId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


