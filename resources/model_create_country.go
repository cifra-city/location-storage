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

// checks if the CreateCountry type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateCountry{}

// CreateCountry struct for CreateCountry
type CreateCountry struct {
	Data CreateCountryData `json:"data"`
}

type _CreateCountry CreateCountry

// NewCreateCountry instantiates a new CreateCountry object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateCountry(data CreateCountryData) *CreateCountry {
	this := CreateCountry{}
	this.Data = data
	return &this
}

// NewCreateCountryWithDefaults instantiates a new CreateCountry object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateCountryWithDefaults() *CreateCountry {
	this := CreateCountry{}
	return &this
}

// GetData returns the Data field value
func (o *CreateCountry) GetData() CreateCountryData {
	if o == nil {
		var ret CreateCountryData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *CreateCountry) GetDataOk() (*CreateCountryData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *CreateCountry) SetData(v CreateCountryData) {
	o.Data = v
}

func (o CreateCountry) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateCountry) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *CreateCountry) UnmarshalJSON(data []byte) (err error) {
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

	varCreateCountry := _CreateCountry{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateCountry)

	if err != nil {
		return err
	}

	*o = CreateCountry(varCreateCountry)

	return err
}

type NullableCreateCountry struct {
	value *CreateCountry
	isSet bool
}

func (v NullableCreateCountry) Get() *CreateCountry {
	return v.value
}

func (v *NullableCreateCountry) Set(val *CreateCountry) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateCountry) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateCountry) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateCountry(val *CreateCountry) *NullableCreateCountry {
	return &NullableCreateCountry{value: val, isSet: true}
}

func (v NullableCreateCountry) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateCountry) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

