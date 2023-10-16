# CreateJobBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Description** | **string** |  | 
**Image** | **string** |  | 
**Namespace** | **string** |  | 
**Gpus** | **string** |  | 
**Pvc** | **string** |  | 
**Command** | **string** |  | 
**Args** | **string** |  | 
**SharedMem** | **string** |  | 
**SharedMemUnit** | **string** |  | 
**Distributed** | **bool** |  | 
**Interactive** | **bool** |  | 

## Methods

### NewCreateJobBody

`func NewCreateJobBody(name string, description string, image string, namespace string, gpus string, pvc string, command string, args string, sharedMem string, sharedMemUnit string, distributed bool, interactive bool, ) *CreateJobBody`

NewCreateJobBody instantiates a new CreateJobBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateJobBodyWithDefaults

`func NewCreateJobBodyWithDefaults() *CreateJobBody`

NewCreateJobBodyWithDefaults instantiates a new CreateJobBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateJobBody) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateJobBody) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateJobBody) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *CreateJobBody) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateJobBody) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateJobBody) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetImage

`func (o *CreateJobBody) GetImage() string`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *CreateJobBody) GetImageOk() (*string, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *CreateJobBody) SetImage(v string)`

SetImage sets Image field to given value.


### GetNamespace

`func (o *CreateJobBody) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *CreateJobBody) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *CreateJobBody) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.


### GetGpus

`func (o *CreateJobBody) GetGpus() string`

GetGpus returns the Gpus field if non-nil, zero value otherwise.

### GetGpusOk

`func (o *CreateJobBody) GetGpusOk() (*string, bool)`

GetGpusOk returns a tuple with the Gpus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpus

`func (o *CreateJobBody) SetGpus(v string)`

SetGpus sets Gpus field to given value.


### GetPvc

`func (o *CreateJobBody) GetPvc() string`

GetPvc returns the Pvc field if non-nil, zero value otherwise.

### GetPvcOk

`func (o *CreateJobBody) GetPvcOk() (*string, bool)`

GetPvcOk returns a tuple with the Pvc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPvc

`func (o *CreateJobBody) SetPvc(v string)`

SetPvc sets Pvc field to given value.


### GetCommand

`func (o *CreateJobBody) GetCommand() string`

GetCommand returns the Command field if non-nil, zero value otherwise.

### GetCommandOk

`func (o *CreateJobBody) GetCommandOk() (*string, bool)`

GetCommandOk returns a tuple with the Command field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommand

`func (o *CreateJobBody) SetCommand(v string)`

SetCommand sets Command field to given value.


### GetArgs

`func (o *CreateJobBody) GetArgs() string`

GetArgs returns the Args field if non-nil, zero value otherwise.

### GetArgsOk

`func (o *CreateJobBody) GetArgsOk() (*string, bool)`

GetArgsOk returns a tuple with the Args field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArgs

`func (o *CreateJobBody) SetArgs(v string)`

SetArgs sets Args field to given value.


### GetSharedMem

`func (o *CreateJobBody) GetSharedMem() string`

GetSharedMem returns the SharedMem field if non-nil, zero value otherwise.

### GetSharedMemOk

`func (o *CreateJobBody) GetSharedMemOk() (*string, bool)`

GetSharedMemOk returns a tuple with the SharedMem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedMem

`func (o *CreateJobBody) SetSharedMem(v string)`

SetSharedMem sets SharedMem field to given value.


### GetSharedMemUnit

`func (o *CreateJobBody) GetSharedMemUnit() string`

GetSharedMemUnit returns the SharedMemUnit field if non-nil, zero value otherwise.

### GetSharedMemUnitOk

`func (o *CreateJobBody) GetSharedMemUnitOk() (*string, bool)`

GetSharedMemUnitOk returns a tuple with the SharedMemUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedMemUnit

`func (o *CreateJobBody) SetSharedMemUnit(v string)`

SetSharedMemUnit sets SharedMemUnit field to given value.


### GetDistributed

`func (o *CreateJobBody) GetDistributed() bool`

GetDistributed returns the Distributed field if non-nil, zero value otherwise.

### GetDistributedOk

`func (o *CreateJobBody) GetDistributedOk() (*bool, bool)`

GetDistributedOk returns a tuple with the Distributed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistributed

`func (o *CreateJobBody) SetDistributed(v bool)`

SetDistributed sets Distributed field to given value.


### GetInteractive

`func (o *CreateJobBody) GetInteractive() bool`

GetInteractive returns the Interactive field if non-nil, zero value otherwise.

### GetInteractiveOk

`func (o *CreateJobBody) GetInteractiveOk() (*bool, bool)`

GetInteractiveOk returns a tuple with the Interactive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInteractive

`func (o *CreateJobBody) SetInteractive(v bool)`

SetInteractive sets Interactive field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


