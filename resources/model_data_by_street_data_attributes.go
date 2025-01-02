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

// checks if the DataByStreetDataAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DataByStreetDataAttributes{}

// DataByStreetDataAttributes struct for DataByStreetDataAttributes
type DataByStreetDataAttributes struct {
	// city uuid
	City string `json:"city"`
	// district uuid
	District string `json:"district"`
	// Street name
	Name string `json:"name"`
}

type _DataByStreetDataAttributes DataByStreetDataAttributes

// NewDataByStreetDataAttributes instantiates a new DataByStreetDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDataByStreetDataAttributes(city string, district string, name string) *DataByStreetDataAttributes {
	this := DataByStreetDataAttributes{}
	this.City = city
	this.District = district
	this.Name = name
	return &this
}

// NewDataByStreetDataAttributesWithDefaults instantiates a new DataByStreetDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDataByStreetDataAttributesWithDefaults() *DataByStreetDataAttributes {
	this := DataByStreetDataAttributes{}
	return &this
}

// GetCity returns the City field value
func (o *DataByStreetDataAttributes) GetCity() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.City
}

// GetCityOk returns a tuple with the City field value
// and a boolean to check if the value has been set.
func (o *DataByStreetDataAttributes) GetCityOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.City, true
}

// SetCity sets field value
func (o *DataByStreetDataAttributes) SetCity(v string) {
	o.City = v
}

// GetDistrict returns the District field value
func (o *DataByStreetDataAttributes) GetDistrict() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.District
}

// GetDistrictOk returns a tuple with the District field value
// and a boolean to check if the value has been set.
func (o *DataByStreetDataAttributes) GetDistrictOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.District, true
}

// SetDistrict sets field value
func (o *DataByStreetDataAttributes) SetDistrict(v string) {
	o.District = v
}

// GetName returns the Name field value
func (o *DataByStreetDataAttributes) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *DataByStreetDataAttributes) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *DataByStreetDataAttributes) SetName(v string) {
	o.Name = v
}

func (o DataByStreetDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DataByStreetDataAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["city"] = o.City
	toSerialize["district"] = o.District
	toSerialize["name"] = o.Name
	return toSerialize, nil
}

func (o *DataByStreetDataAttributes) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"city",
		"district",
		"name",
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

	varDataByStreetDataAttributes := _DataByStreetDataAttributes{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDataByStreetDataAttributes)

	if err != nil {
		return err
	}

	*o = DataByStreetDataAttributes(varDataByStreetDataAttributes)

	return err
}

type NullableDataByStreetDataAttributes struct {
	value *DataByStreetDataAttributes
	isSet bool
}

func (v NullableDataByStreetDataAttributes) Get() *DataByStreetDataAttributes {
	return v.value
}

func (v *NullableDataByStreetDataAttributes) Set(val *DataByStreetDataAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableDataByStreetDataAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableDataByStreetDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDataByStreetDataAttributes(val *DataByStreetDataAttributes) *NullableDataByStreetDataAttributes {
	return &NullableDataByStreetDataAttributes{value: val, isSet: true}
}

func (v NullableDataByStreetDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDataByStreetDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

