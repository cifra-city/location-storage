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

// checks if the UpdateStreetData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateStreetData{}

// UpdateStreetData struct for UpdateStreetData
type UpdateStreetData struct {
	Type string `json:"type"`
	Attributes UpdateStreetDataAttributes `json:"attributes"`
}

type _UpdateStreetData UpdateStreetData

// NewUpdateStreetData instantiates a new UpdateStreetData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateStreetData(type_ string, attributes UpdateStreetDataAttributes) *UpdateStreetData {
	this := UpdateStreetData{}
	this.Type = type_
	this.Attributes = attributes
	return &this
}

// NewUpdateStreetDataWithDefaults instantiates a new UpdateStreetData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateStreetDataWithDefaults() *UpdateStreetData {
	this := UpdateStreetData{}
	return &this
}

// GetType returns the Type field value
func (o *UpdateStreetData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *UpdateStreetData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *UpdateStreetData) SetType(v string) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *UpdateStreetData) GetAttributes() UpdateStreetDataAttributes {
	if o == nil {
		var ret UpdateStreetDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *UpdateStreetData) GetAttributesOk() (*UpdateStreetDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *UpdateStreetData) SetAttributes(v UpdateStreetDataAttributes) {
	o.Attributes = v
}

func (o UpdateStreetData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateStreetData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["attributes"] = o.Attributes
	return toSerialize, nil
}

func (o *UpdateStreetData) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
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

	varUpdateStreetData := _UpdateStreetData{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUpdateStreetData)

	if err != nil {
		return err
	}

	*o = UpdateStreetData(varUpdateStreetData)

	return err
}

type NullableUpdateStreetData struct {
	value *UpdateStreetData
	isSet bool
}

func (v NullableUpdateStreetData) Get() *UpdateStreetData {
	return v.value
}

func (v *NullableUpdateStreetData) Set(val *UpdateStreetData) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateStreetData) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateStreetData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateStreetData(val *UpdateStreetData) *NullableUpdateStreetData {
	return &NullableUpdateStreetData{value: val, isSet: true}
}

func (v NullableUpdateStreetData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateStreetData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


