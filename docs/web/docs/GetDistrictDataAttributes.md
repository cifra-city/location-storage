# GetDistrictDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | district name | 
**City** | **string** | city uuid | 
**Streets** | [**[]GetDistrictDataAttributesStreetsInner**](GetDistrictDataAttributesStreetsInner.md) |  | 

## Methods

### NewGetDistrictDataAttributes

`func NewGetDistrictDataAttributes(name string, city string, streets []GetDistrictDataAttributesStreetsInner, ) *GetDistrictDataAttributes`

NewGetDistrictDataAttributes instantiates a new GetDistrictDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetDistrictDataAttributesWithDefaults

`func NewGetDistrictDataAttributesWithDefaults() *GetDistrictDataAttributes`

NewGetDistrictDataAttributesWithDefaults instantiates a new GetDistrictDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *GetDistrictDataAttributes) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetDistrictDataAttributes) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetDistrictDataAttributes) SetName(v string)`

SetName sets Name field to given value.


### GetCity

`func (o *GetDistrictDataAttributes) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *GetDistrictDataAttributes) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *GetDistrictDataAttributes) SetCity(v string)`

SetCity sets City field to given value.


### GetStreets

`func (o *GetDistrictDataAttributes) GetStreets() []GetDistrictDataAttributesStreetsInner`

GetStreets returns the Streets field if non-nil, zero value otherwise.

### GetStreetsOk

`func (o *GetDistrictDataAttributes) GetStreetsOk() (*[]GetDistrictDataAttributesStreetsInner, bool)`

GetStreetsOk returns a tuple with the Streets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreets

`func (o *GetDistrictDataAttributes) SetStreets(v []GetDistrictDataAttributesStreetsInner)`

SetStreets sets Streets field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


