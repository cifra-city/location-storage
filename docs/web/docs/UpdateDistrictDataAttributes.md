# UpdateDistrictDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NewName** | Pointer to **string** | City new name | [optional] 
**DistrictId** | **string** | district uuid | 
**CityId** | Pointer to **string** | City uuid if u need to change city for the district | [optional] 

## Methods

### NewUpdateDistrictDataAttributes

`func NewUpdateDistrictDataAttributes(districtId string, ) *UpdateDistrictDataAttributes`

NewUpdateDistrictDataAttributes instantiates a new UpdateDistrictDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateDistrictDataAttributesWithDefaults

`func NewUpdateDistrictDataAttributesWithDefaults() *UpdateDistrictDataAttributes`

NewUpdateDistrictDataAttributesWithDefaults instantiates a new UpdateDistrictDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNewName

`func (o *UpdateDistrictDataAttributes) GetNewName() string`

GetNewName returns the NewName field if non-nil, zero value otherwise.

### GetNewNameOk

`func (o *UpdateDistrictDataAttributes) GetNewNameOk() (*string, bool)`

GetNewNameOk returns a tuple with the NewName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewName

`func (o *UpdateDistrictDataAttributes) SetNewName(v string)`

SetNewName sets NewName field to given value.

### HasNewName

`func (o *UpdateDistrictDataAttributes) HasNewName() bool`

HasNewName returns a boolean if a field has been set.

### GetDistrictId

`func (o *UpdateDistrictDataAttributes) GetDistrictId() string`

GetDistrictId returns the DistrictId field if non-nil, zero value otherwise.

### GetDistrictIdOk

`func (o *UpdateDistrictDataAttributes) GetDistrictIdOk() (*string, bool)`

GetDistrictIdOk returns a tuple with the DistrictId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistrictId

`func (o *UpdateDistrictDataAttributes) SetDistrictId(v string)`

SetDistrictId sets DistrictId field to given value.


### GetCityId

`func (o *UpdateDistrictDataAttributes) GetCityId() string`

GetCityId returns the CityId field if non-nil, zero value otherwise.

### GetCityIdOk

`func (o *UpdateDistrictDataAttributes) GetCityIdOk() (*string, bool)`

GetCityIdOk returns a tuple with the CityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCityId

`func (o *UpdateDistrictDataAttributes) SetCityId(v string)`

SetCityId sets CityId field to given value.

### HasCityId

`func (o *UpdateDistrictDataAttributes) HasCityId() bool`

HasCityId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


