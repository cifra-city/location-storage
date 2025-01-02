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

// checks if the DataByDistrict type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DataByDistrict{}

// DataByDistrict struct for DataByDistrict
type DataByDistrict struct {
	Data DataByDistrictData `json:"data"`
}

type _DataByDistrict DataByDistrict

// NewDataByDistrict instantiates a new DataByDistrict object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDataByDistrict(data DataByDistrictData) *DataByDistrict {
	this := DataByDistrict{}
	this.Data = data
	return &this
}

// NewDataByDistrictWithDefaults instantiates a new DataByDistrict object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDataByDistrictWithDefaults() *DataByDistrict {
	this := DataByDistrict{}
	return &this
}

// GetData returns the Data field value
func (o *DataByDistrict) GetData() DataByDistrictData {
	if o == nil {
		var ret DataByDistrictData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *DataByDistrict) GetDataOk() (*DataByDistrictData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *DataByDistrict) SetData(v DataByDistrictData) {
	o.Data = v
}

func (o DataByDistrict) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DataByDistrict) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *DataByDistrict) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"data",
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

	varDataByDistrict := _DataByDistrict{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDataByDistrict)

	if err != nil {
		return err
	}

	*o = DataByDistrict(varDataByDistrict)

	return err
}

type NullableDataByDistrict struct {
	value *DataByDistrict
	isSet bool
}

func (v NullableDataByDistrict) Get() *DataByDistrict {
	return v.value
}

func (v *NullableDataByDistrict) Set(val *DataByDistrict) {
	v.value = val
	v.isSet = true
}

func (v NullableDataByDistrict) IsSet() bool {
	return v.isSet
}

func (v *NullableDataByDistrict) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDataByDistrict(val *DataByDistrict) *NullableDataByDistrict {
	return &NullableDataByDistrict{value: val, isSet: true}
}

func (v NullableDataByDistrict) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDataByDistrict) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

