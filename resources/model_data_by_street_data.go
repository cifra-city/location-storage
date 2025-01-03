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

// checks if the DataByStreetData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DataByStreetData{}

// DataByStreetData struct for DataByStreetData
type DataByStreetData struct {
	// street id
	Id string `json:"id"`
	Type string `json:"type"`
	Attributes DataByStreetDataAttributes `json:"attributes"`
}

type _DataByStreetData DataByStreetData

// NewDataByStreetData instantiates a new DataByStreetData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDataByStreetData(id string, type_ string, attributes DataByStreetDataAttributes) *DataByStreetData {
	this := DataByStreetData{}
	this.Id = id
	this.Type = type_
	this.Attributes = attributes
	return &this
}

// NewDataByStreetDataWithDefaults instantiates a new DataByStreetData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDataByStreetDataWithDefaults() *DataByStreetData {
	this := DataByStreetData{}
	return &this
}

// GetId returns the Id field value
func (o *DataByStreetData) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *DataByStreetData) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *DataByStreetData) SetId(v string) {
	o.Id = v
}

// GetType returns the Type field value
func (o *DataByStreetData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *DataByStreetData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *DataByStreetData) SetType(v string) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *DataByStreetData) GetAttributes() DataByStreetDataAttributes {
	if o == nil {
		var ret DataByStreetDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *DataByStreetData) GetAttributesOk() (*DataByStreetDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *DataByStreetData) SetAttributes(v DataByStreetDataAttributes) {
	o.Attributes = v
}

func (o DataByStreetData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DataByStreetData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["type"] = o.Type
	toSerialize["attributes"] = o.Attributes
	return toSerialize, nil
}

func (o *DataByStreetData) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"type",
		"attributes",
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

	varDataByStreetData := _DataByStreetData{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDataByStreetData)

	if err != nil {
		return err
	}

	*o = DataByStreetData(varDataByStreetData)

	return err
}

type NullableDataByStreetData struct {
	value *DataByStreetData
	isSet bool
}

func (v NullableDataByStreetData) Get() *DataByStreetData {
	return v.value
}

func (v *NullableDataByStreetData) Set(val *DataByStreetData) {
	v.value = val
	v.isSet = true
}

func (v NullableDataByStreetData) IsSet() bool {
	return v.isSet
}

func (v *NullableDataByStreetData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDataByStreetData(val *DataByStreetData) *NullableDataByStreetData {
	return &NullableDataByStreetData{value: val, isSet: true}
}

func (v NullableDataByStreetData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDataByStreetData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


