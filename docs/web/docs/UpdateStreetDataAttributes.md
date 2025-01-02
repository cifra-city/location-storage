# UpdateStreetDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NewName** | Pointer to **string** | City name | [optional] 
**StreetId** | **string** | Street uuid | 
**DistrictId** | Pointer to **string** | District uuid if u need to change district for the street | [optional] 

## Methods

### NewUpdateStreetDataAttributes

`func NewUpdateStreetDataAttributes(streetId string, ) *UpdateStreetDataAttributes`

NewUpdateStreetDataAttributes instantiates a new UpdateStreetDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateStreetDataAttributesWithDefaults

`func NewUpdateStreetDataAttributesWithDefaults() *UpdateStreetDataAttributes`

NewUpdateStreetDataAttributesWithDefaults instantiates a new UpdateStreetDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNewName

`func (o *UpdateStreetDataAttributes) GetNewName() string`

GetNewName returns the NewName field if non-nil, zero value otherwise.

### GetNewNameOk

`func (o *UpdateStreetDataAttributes) GetNewNameOk() (*string, bool)`

GetNewNameOk returns a tuple with the NewName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewName

`func (o *UpdateStreetDataAttributes) SetNewName(v string)`

SetNewName sets NewName field to given value.

### HasNewName

`func (o *UpdateStreetDataAttributes) HasNewName() bool`

HasNewName returns a boolean if a field has been set.

### GetStreetId

`func (o *UpdateStreetDataAttributes) GetStreetId() string`

GetStreetId returns the StreetId field if non-nil, zero value otherwise.

### GetStreetIdOk

`func (o *UpdateStreetDataAttributes) GetStreetIdOk() (*string, bool)`

GetStreetIdOk returns a tuple with the StreetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreetId

`func (o *UpdateStreetDataAttributes) SetStreetId(v string)`

SetStreetId sets StreetId field to given value.


### GetDistrictId

`func (o *UpdateStreetDataAttributes) GetDistrictId() string`

GetDistrictId returns the DistrictId field if non-nil, zero value otherwise.

### GetDistrictIdOk

`func (o *UpdateStreetDataAttributes) GetDistrictIdOk() (*string, bool)`

GetDistrictIdOk returns a tuple with the DistrictId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistrictId

`func (o *UpdateStreetDataAttributes) SetDistrictId(v string)`

SetDistrictId sets DistrictId field to given value.

### HasDistrictId

`func (o *UpdateStreetDataAttributes) HasDistrictId() bool`

HasDistrictId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


