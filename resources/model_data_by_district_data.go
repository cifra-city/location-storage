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

// checks if the DataByDistrictData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DataByDistrictData{}

// DataByDistrictData struct for DataByDistrictData
type DataByDistrictData struct {
	// district id
	Id string `json:"id"`
	Type string `json:"type"`
	Attributes DataByDistrictDataAttributes `json:"attributes"`
}

type _DataByDistrictData DataByDistrictData

// NewDataByDistrictData instantiates a new DataByDistrictData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDataByDistrictData(id string, type_ string, attributes DataByDistrictDataAttributes) *DataByDistrictData {
	this := DataByDistrictData{}
	this.Id = id
	this.Type = type_
	this.Attributes = attributes
	return &this
}

// NewDataByDistrictDataWithDefaults instantiates a new DataByDistrictData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDataByDistrictDataWithDefaults() *DataByDistrictData {
	this := DataByDistrictData{}
	return &this
}

// GetId returns the Id field value
func (o *DataByDistrictData) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *DataByDistrictData) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *DataByDistrictData) SetId(v string) {
	o.Id = v
}

// GetType returns the Type field value
func (o *DataByDistrictData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *DataByDistrictData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *DataByDistrictData) SetType(v string) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *DataByDistrictData) GetAttributes() DataByDistrictDataAttributes {
	if o == nil {
		var ret DataByDistrictDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *DataByDistrictData) GetAttributesOk() (*DataByDistrictDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *DataByDistrictData) SetAttributes(v DataByDistrictDataAttributes) {
	o.Attributes = v
}

func (o DataByDistrictData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DataByDistrictData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["type"] = o.Type
	toSerialize["attributes"] = o.Attributes
	return toSerialize, nil
}

func (o *DataByDistrictData) UnmarshalJSON(data []byte) (err error) {
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

	varDataByDistrictData := _DataByDistrictData{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDataByDistrictData)

	if err != nil {
		return err
	}

	*o = DataByDistrictData(varDataByDistrictData)

	return err
}

type NullableDataByDistrictData struct {
	value *DataByDistrictData
	isSet bool
}

func (v NullableDataByDistrictData) Get() *DataByDistrictData {
	return v.value
}

func (v *NullableDataByDistrictData) Set(val *DataByDistrictData) {
	v.value = val
	v.isSet = true
}

func (v NullableDataByDistrictData) IsSet() bool {
	return v.isSet
}

func (v *NullableDataByDistrictData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDataByDistrictData(val *DataByDistrictData) *NullableDataByDistrictData {
	return &NullableDataByDistrictData{value: val, isSet: true}
}

func (v NullableDataByDistrictData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDataByDistrictData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

