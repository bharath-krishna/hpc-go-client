# JobDetailsResp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cluster** | **string** |  | 
**Name** | **string** |  | 
**Namespace** | **string** |  | 
**Command** | **string** |  | 
**Image** | **string** |  | 
**CreationTimestamp** | **string** |  | 
**JobId** | Pointer to **string** |  | [optional] 
**LogShellUrl** | **string** |  | 

## Methods

### NewJobDetailsResp

`func NewJobDetailsResp(cluster string, name string, namespace string, command string, image string, creationTimestamp string, logShellUrl string, ) *JobDetailsResp`

NewJobDetailsResp instantiates a new JobDetailsResp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJobDetailsRespWithDefaults

`func NewJobDetailsRespWithDefaults() *JobDetailsResp`

NewJobDetailsRespWithDefaults instantiates a new JobDetailsResp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCluster

`func (o *JobDetailsResp) GetCluster() string`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *JobDetailsResp) GetClusterOk() (*string, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *JobDetailsResp) SetCluster(v string)`

SetCluster sets Cluster field to given value.


### GetName

`func (o *JobDetailsResp) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *JobDetailsResp) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *JobDetailsResp) SetName(v string)`

SetName sets Name field to given value.


### GetNamespace

`func (o *JobDetailsResp) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *JobDetailsResp) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *JobDetailsResp) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.


### GetCommand

`func (o *JobDetailsResp) GetCommand() string`

GetCommand returns the Command field if non-nil, zero value otherwise.

### GetCommandOk

`func (o *JobDetailsResp) GetCommandOk() (*string, bool)`

GetCommandOk returns a tuple with the Command field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommand

`func (o *JobDetailsResp) SetCommand(v string)`

SetCommand sets Command field to given value.


### GetImage

`func (o *JobDetailsResp) GetImage() string`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *JobDetailsResp) GetImageOk() (*string, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *JobDetailsResp) SetImage(v string)`

SetImage sets Image field to given value.


### GetCreationTimestamp

`func (o *JobDetailsResp) GetCreationTimestamp() string`

GetCreationTimestamp returns the CreationTimestamp field if non-nil, zero value otherwise.

### GetCreationTimestampOk

`func (o *JobDetailsResp) GetCreationTimestampOk() (*string, bool)`

GetCreationTimestampOk returns a tuple with the CreationTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTimestamp

`func (o *JobDetailsResp) SetCreationTimestamp(v string)`

SetCreationTimestamp sets CreationTimestamp field to given value.


### GetJobId

`func (o *JobDetailsResp) GetJobId() string`

GetJobId returns the JobId field if non-nil, zero value otherwise.

### GetJobIdOk

`func (o *JobDetailsResp) GetJobIdOk() (*string, bool)`

GetJobIdOk returns a tuple with the JobId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobId

`func (o *JobDetailsResp) SetJobId(v string)`

SetJobId sets JobId field to given value.

### HasJobId

`func (o *JobDetailsResp) HasJobId() bool`

HasJobId returns a boolean if a field has been set.

### GetLogShellUrl

`func (o *JobDetailsResp) GetLogShellUrl() string`

GetLogShellUrl returns the LogShellUrl field if non-nil, zero value otherwise.

### GetLogShellUrlOk

`func (o *JobDetailsResp) GetLogShellUrlOk() (*string, bool)`

GetLogShellUrlOk returns a tuple with the LogShellUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogShellUrl

`func (o *JobDetailsResp) SetLogShellUrl(v string)`

SetLogShellUrl sets LogShellUrl field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


