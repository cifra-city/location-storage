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

// checks if the CreateDistrict type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateDistrict{}

// CreateDistrict struct for CreateDistrict
type CreateDistrict struct {
	Data CreateDistrictData `json:"data"`
}

type _CreateDistrict CreateDistrict

// NewCreateDistrict instantiates a new CreateDistrict object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateDistrict(data CreateDistrictData) *CreateDistrict {
	this := CreateDistrict{}
	this.Data = data
	return &this
}

// NewCreateDistrictWithDefaults instantiates a new CreateDistrict object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateDistrictWithDefaults() *CreateDistrict {
	this := CreateDistrict{}
	return &this
}

// GetData returns the Data field value
func (o *CreateDistrict) GetData() CreateDistrictData {
	if o == nil {
		var ret CreateDistrictData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *CreateDistrict) GetDataOk() (*CreateDistrictData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *CreateDistrict) SetData(v CreateDistrictData) {
	o.Data = v
}

func (o CreateDistrict) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateDistrict) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *CreateDistrict) UnmarshalJSON(data []byte) (err error) {
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

	varCreateDistrict := _CreateDistrict{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateDistrict)

	if err != nil {
		return err
	}

	*o = CreateDistrict(varCreateDistrict)

	return err
}

type NullableCreateDistrict struct {
	value *CreateDistrict
	isSet bool
}

func (v NullableCreateDistrict) Get() *CreateDistrict {
	return v.value
}

func (v *NullableCreateDistrict) Set(val *CreateDistrict) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateDistrict) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateDistrict) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateDistrict(val *CreateDistrict) *NullableCreateDistrict {
	return &NullableCreateDistrict{value: val, isSet: true}
}

func (v NullableCreateDistrict) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateDistrict) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


