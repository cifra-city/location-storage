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

// checks if the DataByCityDataAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DataByCityDataAttributes{}

// DataByCityDataAttributes struct for DataByCityDataAttributes
type DataByCityDataAttributes struct {
	// City name
	Name string `json:"name"`
	// Country name
	Country string `json:"country"`
	Districts []DataByCityDataAttributesDistrictsInner `json:"districts"`
}

type _DataByCityDataAttributes DataByCityDataAttributes

// NewDataByCityDataAttributes instantiates a new DataByCityDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDataByCityDataAttributes(name string, country string, districts []DataByCityDataAttributesDistrictsInner) *DataByCityDataAttributes {
	this := DataByCityDataAttributes{}
	this.Name = name
	this.Country = country
	this.Districts = districts
	return &this
}

// NewDataByCityDataAttributesWithDefaults instantiates a new DataByCityDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDataByCityDataAttributesWithDefaults() *DataByCityDataAttributes {
	this := DataByCityDataAttributes{}
	return &this
}

// GetName returns the Name field value
func (o *DataByCityDataAttributes) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *DataByCityDataAttributes) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *DataByCityDataAttributes) SetName(v string) {
	o.Name = v
}

// GetCountry returns the Country field value
func (o *DataByCityDataAttributes) GetCountry() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Country
}

// GetCountryOk returns a tuple with the Country field value
// and a boolean to check if the value has been set.
func (o *DataByCityDataAttributes) GetCountryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Country, true
}

// SetCountry sets field value
func (o *DataByCityDataAttributes) SetCountry(v string) {
	o.Country = v
}

// GetDistricts returns the Districts field value
func (o *DataByCityDataAttributes) GetDistricts() []DataByCityDataAttributesDistrictsInner {
	if o == nil {
		var ret []DataByCityDataAttributesDistrictsInner
		return ret
	}

	return o.Districts
}

// GetDistrictsOk returns a tuple with the Districts field value
// and a boolean to check if the value has been set.
func (o *DataByCityDataAttributes) GetDistrictsOk() ([]DataByCityDataAttributesDistrictsInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Districts, true
}

// SetDistricts sets field value
func (o *DataByCityDataAttributes) SetDistricts(v []DataByCityDataAttributesDistrictsInner) {
	o.Districts = v
}

func (o DataByCityDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DataByCityDataAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["country"] = o.Country
	toSerialize["districts"] = o.Districts
	return toSerialize, nil
}

func (o *DataByCityDataAttributes) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"country",
		"districts",
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

	varDataByCityDataAttributes := _DataByCityDataAttributes{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDataByCityDataAttributes)

	if err != nil {
		return err
	}

	*o = DataByCityDataAttributes(varDataByCityDataAttributes)

	return err
}

type NullableDataByCityDataAttributes struct {
	value *DataByCityDataAttributes
	isSet bool
}

func (v NullableDataByCityDataAttributes) Get() *DataByCityDataAttributes {
	return v.value
}

func (v *NullableDataByCityDataAttributes) Set(val *DataByCityDataAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableDataByCityDataAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableDataByCityDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDataByCityDataAttributes(val *DataByCityDataAttributes) *NullableDataByCityDataAttributes {
	return &NullableDataByCityDataAttributes{value: val, isSet: true}
}

func (v NullableDataByCityDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDataByCityDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


