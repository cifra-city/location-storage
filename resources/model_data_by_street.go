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

// checks if the DataByStreet type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DataByStreet{}

// DataByStreet struct for DataByStreet
type DataByStreet struct {
	Data DataByStreetData `json:"data"`
}

type _DataByStreet DataByStreet

// NewDataByStreet instantiates a new DataByStreet object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDataByStreet(data DataByStreetData) *DataByStreet {
	this := DataByStreet{}
	this.Data = data
	return &this
}

// NewDataByStreetWithDefaults instantiates a new DataByStreet object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDataByStreetWithDefaults() *DataByStreet {
	this := DataByStreet{}
	return &this
}

// GetData returns the Data field value
func (o *DataByStreet) GetData() DataByStreetData {
	if o == nil {
		var ret DataByStreetData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *DataByStreet) GetDataOk() (*DataByStreetData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *DataByStreet) SetData(v DataByStreetData) {
	o.Data = v
}

func (o DataByStreet) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DataByStreet) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *DataByStreet) UnmarshalJSON(data []byte) (err error) {
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

	varDataByStreet := _DataByStreet{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDataByStreet)

	if err != nil {
		return err
	}

	*o = DataByStreet(varDataByStreet)

	return err
}

type NullableDataByStreet struct {
	value *DataByStreet
	isSet bool
}

func (v NullableDataByStreet) Get() *DataByStreet {
	return v.value
}

func (v *NullableDataByStreet) Set(val *DataByStreet) {
	v.value = val
	v.isSet = true
}

func (v NullableDataByStreet) IsSet() bool {
	return v.isSet
}

func (v *NullableDataByStreet) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDataByStreet(val *DataByStreet) *NullableDataByStreet {
	return &NullableDataByStreet{value: val, isSet: true}
}

func (v NullableDataByStreet) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDataByStreet) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


