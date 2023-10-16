# UserModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sub** | **string** |  | 
**EmailVerified** | **bool** |  | 
**Name** | **string** |  | 
**PreferredUsername** | **string** |  | 
**GivenName** | **string** |  | 
**FamilyName** | **string** |  | 
**Email** | **string** |  | 

## Methods

### NewUserModel

`func NewUserModel(sub string, emailVerified bool, name string, preferredUsername string, givenName string, familyName string, email string, ) *UserModel`

NewUserModel instantiates a new UserModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserModelWithDefaults

`func NewUserModelWithDefaults() *UserModel`

NewUserModelWithDefaults instantiates a new UserModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSub

`func (o *UserModel) GetSub() string`

GetSub returns the Sub field if non-nil, zero value otherwise.

### GetSubOk

`func (o *UserModel) GetSubOk() (*string, bool)`

GetSubOk returns a tuple with the Sub field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSub

`func (o *UserModel) SetSub(v string)`

SetSub sets Sub field to given value.


### GetEmailVerified

`func (o *UserModel) GetEmailVerified() bool`

GetEmailVerified returns the EmailVerified field if non-nil, zero value otherwise.

### GetEmailVerifiedOk

`func (o *UserModel) GetEmailVerifiedOk() (*bool, bool)`

GetEmailVerifiedOk returns a tuple with the EmailVerified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailVerified

`func (o *UserModel) SetEmailVerified(v bool)`

SetEmailVerified sets EmailVerified field to given value.


### GetName

`func (o *UserModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UserModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UserModel) SetName(v string)`

SetName sets Name field to given value.


### GetPreferredUsername

`func (o *UserModel) GetPreferredUsername() string`

GetPreferredUsername returns the PreferredUsername field if non-nil, zero value otherwise.

### GetPreferredUsernameOk

`func (o *UserModel) GetPreferredUsernameOk() (*string, bool)`

GetPreferredUsernameOk returns a tuple with the PreferredUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreferredUsername

`func (o *UserModel) SetPreferredUsername(v string)`

SetPreferredUsername sets PreferredUsername field to given value.


### GetGivenName

`func (o *UserModel) GetGivenName() string`

GetGivenName returns the GivenName field if non-nil, zero value otherwise.

### GetGivenNameOk

`func (o *UserModel) GetGivenNameOk() (*string, bool)`

GetGivenNameOk returns a tuple with the GivenName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGivenName

`func (o *UserModel) SetGivenName(v string)`

SetGivenName sets GivenName field to given value.


### GetFamilyName

`func (o *UserModel) GetFamilyName() string`

GetFamilyName returns the FamilyName field if non-nil, zero value otherwise.

### GetFamilyNameOk

`func (o *UserModel) GetFamilyNameOk() (*string, bool)`

GetFamilyNameOk returns a tuple with the FamilyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFamilyName

`func (o *UserModel) SetFamilyName(v string)`

SetFamilyName sets FamilyName field to given value.


### GetEmail

`func (o *UserModel) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UserModel) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UserModel) SetEmail(v string)`

SetEmail sets Email field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


