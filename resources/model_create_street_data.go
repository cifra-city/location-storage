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

// checks if the CreateStreetData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateStreetData{}

// CreateStreetData struct for CreateStreetData
type CreateStreetData struct {
	Type string `json:"type"`
	Attributes CreateStreetDataAttributes `json:"attributes"`
}

type _CreateStreetData CreateStreetData

// NewCreateStreetData instantiates a new CreateStreetData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateStreetData(type_ string, attributes CreateStreetDataAttributes) *CreateStreetData {
	this := CreateStreetData{}
	this.Type = type_
	this.Attributes = attributes
	return &this
}

// NewCreateStreetDataWithDefaults instantiates a new CreateStreetData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateStreetDataWithDefaults() *CreateStreetData {
	this := CreateStreetData{}
	return &this
}

// GetType returns the Type field value
func (o *CreateStreetData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *CreateStreetData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *CreateStreetData) SetType(v string) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *CreateStreetData) GetAttributes() CreateStreetDataAttributes {
	if o == nil {
		var ret CreateStreetDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *CreateStreetData) GetAttributesOk() (*CreateStreetDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *CreateStreetData) SetAttributes(v CreateStreetDataAttributes) {
	o.Attributes = v
}

func (o CreateStreetData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateStreetData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["attributes"] = o.Attributes
	return toSerialize, nil
}

func (o *CreateStreetData) UnmarshalJSON(data []byte) (err error) {
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

	varCreateStreetData := _CreateStreetData{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateStreetData)

	if err != nil {
		return err
	}

	*o = CreateStreetData(varCreateStreetData)

	return err
}

type NullableCreateStreetData struct {
	value *CreateStreetData
	isSet bool
}

func (v NullableCreateStreetData) Get() *CreateStreetData {
	return v.value
}

func (v *NullableCreateStreetData) Set(val *CreateStreetData) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateStreetData) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateStreetData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateStreetData(val *CreateStreetData) *NullableCreateStreetData {
	return &NullableCreateStreetData{value: val, isSet: true}
}

func (v NullableCreateStreetData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateStreetData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


