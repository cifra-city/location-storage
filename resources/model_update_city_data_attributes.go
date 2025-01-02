/*
Cifra SSO REST API

SSO REST API for Cifra app

API version: 0.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package resources

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the UpdateCityDataAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateCityDataAttributes{}

// UpdateCityDataAttributes struct for UpdateCityDataAttributes
type UpdateCityDataAttributes struct {
	// City uuid
	CityId string `json:"city_id"`
	// City new name
	NewName *string `json:"new_name,omitempty"`
	// Country uuid if you want to change country
	CountryId *string `json:"country_id,omitempty"`
}

type _UpdateCityDataAttributes UpdateCityDataAttributes

// NewUpdateCityDataAttributes instantiates a new UpdateCityDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateCityDataAttributes(cityId string) *UpdateCityDataAttributes {
	this := UpdateCityDataAttributes{}
	this.CityId = cityId
	return &this
}

// NewUpdateCityDataAttributesWithDefaults instantiates a new UpdateCityDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateCityDataAttributesWithDefaults() *UpdateCityDataAttributes {
	this := UpdateCityDataAttributes{}
	return &this
}

// GetCityId returns the CityId field value
func (o *UpdateCityDataAttributes) GetCityId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CityId
}

// GetCityIdOk returns a tuple with the CityId field value
// and a boolean to check if the value has been set.
func (o *UpdateCityDataAttributes) GetCityIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CityId, true
}

// SetCityId sets field value
func (o *UpdateCityDataAttributes) SetCityId(v string) {
	o.CityId = v
}

// GetNewName returns the NewName field value if set, zero value otherwise.
func (o *UpdateCityDataAttributes) GetNewName() string {
	if o == nil || IsNil(o.NewName) {
		var ret string
		return ret
	}
	return *o.NewName
}

// GetNewNameOk returns a tuple with the NewName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateCityDataAttributes) GetNewNameOk() (*string, bool) {
	if o == nil || IsNil(o.NewName) {
		return nil, false
	}
	return o.NewName, true
}

// HasNewName returns a boolean if a field has been set.
func (o *UpdateCityDataAttributes) HasNewName() bool {
	if o != nil && !IsNil(o.NewName) {
		return true
	}

	return false
}

// SetNewName gets a reference to the given string and assigns it to the NewName field.
func (o *UpdateCityDataAttributes) SetNewName(v string) {
	o.NewName = &v
}

// GetCountryId returns the CountryId field value if set, zero value otherwise.
func (o *UpdateCityDataAttributes) GetCountryId() string {
	if o == nil || IsNil(o.CountryId) {
		var ret string
		return ret
	}
	return *o.CountryId
}

// GetCountryIdOk returns a tuple with the CountryId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateCityDataAttributes) GetCountryIdOk() (*string, bool) {
	if o == nil || IsNil(o.CountryId) {
		return nil, false
	}
	return o.CountryId, true
}

// HasCountryId returns a boolean if a field has been set.
func (o *UpdateCityDataAttributes) HasCountryId() bool {
	if o != nil && !IsNil(o.CountryId) {
		return true
	}

	return false
}

// SetCountryId gets a reference to the given string and assigns it to the CountryId field.
func (o *UpdateCityDataAttributes) SetCountryId(v string) {
	o.CountryId = &v
}

func (o UpdateCityDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateCityDataAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["city_id"] = o.CityId
	if !IsNil(o.NewName) {
		toSerialize["new_name"] = o.NewName
	}
	if !IsNil(o.CountryId) {
		toSerialize["country_id"] = o.CountryId
	}
	return toSerialize, nil
}

func (o *UpdateCityDataAttributes) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"city_id",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varUpdateCityDataAttributes := _UpdateCityDataAttributes{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUpdateCityDataAttributes)

	if err != nil {
		return err
	}

	*o = UpdateCityDataAttributes(varUpdateCityDataAttributes)

	return err
}

type NullableUpdateCityDataAttributes struct {
	value *UpdateCityDataAttributes
	isSet bool
}

func (v NullableUpdateCityDataAttributes) Get() *UpdateCityDataAttributes {
	return v.value
}

func (v *NullableUpdateCityDataAttributes) Set(val *UpdateCityDataAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateCityDataAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateCityDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateCityDataAttributes(val *UpdateCityDataAttributes) *NullableUpdateCityDataAttributes {
	return &NullableUpdateCityDataAttributes{value: val, isSet: true}
}

func (v NullableUpdateCityDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateCityDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

