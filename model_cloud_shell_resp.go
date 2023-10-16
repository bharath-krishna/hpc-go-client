/*
HPC Portal - API

Code for logging, authentication with JWT token, e2e tests are added

API version: 0.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package hpc

import (
	"encoding/json"
)

// checks if the CloudShellResp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CloudShellResp{}

// CloudShellResp struct for CloudShellResp
type CloudShellResp struct {
	Name string `json:"name"`
	Namespace string `json:"namespace"`
	Command string `json:"command"`
	CreationTimestamp string `json:"creation_timestamp"`
	Url string `json:"url"`
}

// NewCloudShellResp instantiates a new CloudShellResp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCloudShellResp(name string, namespace string, command string, creationTimestamp string, url string) *CloudShellResp {
	this := CloudShellResp{}
	this.Name = name
	this.Namespace = namespace
	this.Command = command
	this.CreationTimestamp = creationTimestamp
	this.Url = url
	return &this
}

// NewCloudShellRespWithDefaults instantiates a new CloudShellResp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCloudShellRespWithDefaults() *CloudShellResp {
	this := CloudShellResp{}
	return &this
}

// GetName returns the Name field value
func (o *CloudShellResp) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CloudShellResp) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CloudShellResp) SetName(v string) {
	o.Name = v
}

// GetNamespace returns the Namespace field value
func (o *CloudShellResp) GetNamespace() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Namespace
}

// GetNamespaceOk returns a tuple with the Namespace field value
// and a boolean to check if the value has been set.
func (o *CloudShellResp) GetNamespaceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Namespace, true
}

// SetNamespace sets field value
func (o *CloudShellResp) SetNamespace(v string) {
	o.Namespace = v
}

// GetCommand returns the Command field value
func (o *CloudShellResp) GetCommand() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Command
}

// GetCommandOk returns a tuple with the Command field value
// and a boolean to check if the value has been set.
func (o *CloudShellResp) GetCommandOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Command, true
}

// SetCommand sets field value
func (o *CloudShellResp) SetCommand(v string) {
	o.Command = v
}

// GetCreationTimestamp returns the CreationTimestamp field value
func (o *CloudShellResp) GetCreationTimestamp() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreationTimestamp
}

// GetCreationTimestampOk returns a tuple with the CreationTimestamp field value
// and a boolean to check if the value has been set.
func (o *CloudShellResp) GetCreationTimestampOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreationTimestamp, true
}

// SetCreationTimestamp sets field value
func (o *CloudShellResp) SetCreationTimestamp(v string) {
	o.CreationTimestamp = v
}

// GetUrl returns the Url field value
func (o *CloudShellResp) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *CloudShellResp) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *CloudShellResp) SetUrl(v string) {
	o.Url = v
}

func (o CloudShellResp) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CloudShellResp) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["namespace"] = o.Namespace
	toSerialize["command"] = o.Command
	toSerialize["creation_timestamp"] = o.CreationTimestamp
	toSerialize["url"] = o.Url
	return toSerialize, nil
}

type NullableCloudShellResp struct {
	value *CloudShellResp
	isSet bool
}

func (v NullableCloudShellResp) Get() *CloudShellResp {
	return v.value
}

func (v *NullableCloudShellResp) Set(val *CloudShellResp) {
	v.value = val
	v.isSet = true
}

func (v NullableCloudShellResp) IsSet() bool {
	return v.isSet
}

func (v *NullableCloudShellResp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCloudShellResp(val *CloudShellResp) *NullableCloudShellResp {
	return &NullableCloudShellResp{value: val, isSet: true}
}

func (v NullableCloudShellResp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCloudShellResp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


