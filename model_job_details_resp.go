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

// checks if the JobDetailsResp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &JobDetailsResp{}

// JobDetailsResp struct for JobDetailsResp
type JobDetailsResp struct {
	Cluster string `json:"cluster"`
	Name string `json:"name"`
	Namespace string `json:"namespace"`
	Command string `json:"command"`
	Image string `json:"image"`
	CreationTimestamp string `json:"creation_timestamp"`
	JobId *string `json:"job_id,omitempty"`
	LogShellUrl string `json:"logShellUrl"`
}

// NewJobDetailsResp instantiates a new JobDetailsResp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewJobDetailsResp(cluster string, name string, namespace string, command string, image string, creationTimestamp string, logShellUrl string) *JobDetailsResp {
	this := JobDetailsResp{}
	this.Cluster = cluster
	this.Name = name
	this.Namespace = namespace
	this.Command = command
	this.Image = image
	this.CreationTimestamp = creationTimestamp
	this.LogShellUrl = logShellUrl
	return &this
}

// NewJobDetailsRespWithDefaults instantiates a new JobDetailsResp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewJobDetailsRespWithDefaults() *JobDetailsResp {
	this := JobDetailsResp{}
	return &this
}

// GetCluster returns the Cluster field value
func (o *JobDetailsResp) GetCluster() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Cluster
}

// GetClusterOk returns a tuple with the Cluster field value
// and a boolean to check if the value has been set.
func (o *JobDetailsResp) GetClusterOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Cluster, true
}

// SetCluster sets field value
func (o *JobDetailsResp) SetCluster(v string) {
	o.Cluster = v
}

// GetName returns the Name field value
func (o *JobDetailsResp) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *JobDetailsResp) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *JobDetailsResp) SetName(v string) {
	o.Name = v
}

// GetNamespace returns the Namespace field value
func (o *JobDetailsResp) GetNamespace() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Namespace
}

// GetNamespaceOk returns a tuple with the Namespace field value
// and a boolean to check if the value has been set.
func (o *JobDetailsResp) GetNamespaceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Namespace, true
}

// SetNamespace sets field value
func (o *JobDetailsResp) SetNamespace(v string) {
	o.Namespace = v
}

// GetCommand returns the Command field value
func (o *JobDetailsResp) GetCommand() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Command
}

// GetCommandOk returns a tuple with the Command field value
// and a boolean to check if the value has been set.
func (o *JobDetailsResp) GetCommandOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Command, true
}

// SetCommand sets field value
func (o *JobDetailsResp) SetCommand(v string) {
	o.Command = v
}

// GetImage returns the Image field value
func (o *JobDetailsResp) GetImage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Image
}

// GetImageOk returns a tuple with the Image field value
// and a boolean to check if the value has been set.
func (o *JobDetailsResp) GetImageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Image, true
}

// SetImage sets field value
func (o *JobDetailsResp) SetImage(v string) {
	o.Image = v
}

// GetCreationTimestamp returns the CreationTimestamp field value
func (o *JobDetailsResp) GetCreationTimestamp() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreationTimestamp
}

// GetCreationTimestampOk returns a tuple with the CreationTimestamp field value
// and a boolean to check if the value has been set.
func (o *JobDetailsResp) GetCreationTimestampOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreationTimestamp, true
}

// SetCreationTimestamp sets field value
func (o *JobDetailsResp) SetCreationTimestamp(v string) {
	o.CreationTimestamp = v
}

// GetJobId returns the JobId field value if set, zero value otherwise.
func (o *JobDetailsResp) GetJobId() string {
	if o == nil || IsNil(o.JobId) {
		var ret string
		return ret
	}
	return *o.JobId
}

// GetJobIdOk returns a tuple with the JobId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JobDetailsResp) GetJobIdOk() (*string, bool) {
	if o == nil || IsNil(o.JobId) {
		return nil, false
	}
	return o.JobId, true
}

// HasJobId returns a boolean if a field has been set.
func (o *JobDetailsResp) HasJobId() bool {
	if o != nil && !IsNil(o.JobId) {
		return true
	}

	return false
}

// SetJobId gets a reference to the given string and assigns it to the JobId field.
func (o *JobDetailsResp) SetJobId(v string) {
	o.JobId = &v
}

// GetLogShellUrl returns the LogShellUrl field value
func (o *JobDetailsResp) GetLogShellUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LogShellUrl
}

// GetLogShellUrlOk returns a tuple with the LogShellUrl field value
// and a boolean to check if the value has been set.
func (o *JobDetailsResp) GetLogShellUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LogShellUrl, true
}

// SetLogShellUrl sets field value
func (o *JobDetailsResp) SetLogShellUrl(v string) {
	o.LogShellUrl = v
}

func (o JobDetailsResp) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o JobDetailsResp) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["cluster"] = o.Cluster
	toSerialize["name"] = o.Name
	toSerialize["namespace"] = o.Namespace
	toSerialize["command"] = o.Command
	toSerialize["image"] = o.Image
	toSerialize["creation_timestamp"] = o.CreationTimestamp
	if !IsNil(o.JobId) {
		toSerialize["job_id"] = o.JobId
	}
	toSerialize["logShellUrl"] = o.LogShellUrl
	return toSerialize, nil
}

type NullableJobDetailsResp struct {
	value *JobDetailsResp
	isSet bool
}

func (v NullableJobDetailsResp) Get() *JobDetailsResp {
	return v.value
}

func (v *NullableJobDetailsResp) Set(val *JobDetailsResp) {
	v.value = val
	v.isSet = true
}

func (v NullableJobDetailsResp) IsSet() bool {
	return v.isSet
}

func (v *NullableJobDetailsResp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJobDetailsResp(val *JobDetailsResp) *NullableJobDetailsResp {
	return &NullableJobDetailsResp{value: val, isSet: true}
}

func (v NullableJobDetailsResp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJobDetailsResp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


