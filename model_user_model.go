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

// checks if the UserModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserModel{}

// UserModel struct for UserModel
type UserModel struct {
	Sub string `json:"sub"`
	EmailVerified bool `json:"email_verified"`
	Name string `json:"name"`
	PreferredUsername string `json:"preferred_username"`
	GivenName string `json:"given_name"`
	FamilyName string `json:"family_name"`
	Email string `json:"email"`
}

// NewUserModel instantiates a new UserModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserModel(sub string, emailVerified bool, name string, preferredUsername string, givenName string, familyName string, email string) *UserModel {
	this := UserModel{}
	this.Sub = sub
	this.EmailVerified = emailVerified
	this.Name = name
	this.PreferredUsername = preferredUsername
	this.GivenName = givenName
	this.FamilyName = familyName
	this.Email = email
	return &this
}

// NewUserModelWithDefaults instantiates a new UserModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserModelWithDefaults() *UserModel {
	this := UserModel{}
	return &this
}

// GetSub returns the Sub field value
func (o *UserModel) GetSub() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Sub
}

// GetSubOk returns a tuple with the Sub field value
// and a boolean to check if the value has been set.
func (o *UserModel) GetSubOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Sub, true
}

// SetSub sets field value
func (o *UserModel) SetSub(v string) {
	o.Sub = v
}

// GetEmailVerified returns the EmailVerified field value
func (o *UserModel) GetEmailVerified() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.EmailVerified
}

// GetEmailVerifiedOk returns a tuple with the EmailVerified field value
// and a boolean to check if the value has been set.
func (o *UserModel) GetEmailVerifiedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EmailVerified, true
}

// SetEmailVerified sets field value
func (o *UserModel) SetEmailVerified(v bool) {
	o.EmailVerified = v
}

// GetName returns the Name field value
func (o *UserModel) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *UserModel) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *UserModel) SetName(v string) {
	o.Name = v
}

// GetPreferredUsername returns the PreferredUsername field value
func (o *UserModel) GetPreferredUsername() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PreferredUsername
}

// GetPreferredUsernameOk returns a tuple with the PreferredUsername field value
// and a boolean to check if the value has been set.
func (o *UserModel) GetPreferredUsernameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PreferredUsername, true
}

// SetPreferredUsername sets field value
func (o *UserModel) SetPreferredUsername(v string) {
	o.PreferredUsername = v
}

// GetGivenName returns the GivenName field value
func (o *UserModel) GetGivenName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.GivenName
}

// GetGivenNameOk returns a tuple with the GivenName field value
// and a boolean to check if the value has been set.
func (o *UserModel) GetGivenNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GivenName, true
}

// SetGivenName sets field value
func (o *UserModel) SetGivenName(v string) {
	o.GivenName = v
}

// GetFamilyName returns the FamilyName field value
func (o *UserModel) GetFamilyName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FamilyName
}

// GetFamilyNameOk returns a tuple with the FamilyName field value
// and a boolean to check if the value has been set.
func (o *UserModel) GetFamilyNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FamilyName, true
}

// SetFamilyName sets field value
func (o *UserModel) SetFamilyName(v string) {
	o.FamilyName = v
}

// GetEmail returns the Email field value
func (o *UserModel) GetEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *UserModel) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value
func (o *UserModel) SetEmail(v string) {
	o.Email = v
}

func (o UserModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["sub"] = o.Sub
	toSerialize["email_verified"] = o.EmailVerified
	toSerialize["name"] = o.Name
	toSerialize["preferred_username"] = o.PreferredUsername
	toSerialize["given_name"] = o.GivenName
	toSerialize["family_name"] = o.FamilyName
	toSerialize["email"] = o.Email
	return toSerialize, nil
}

type NullableUserModel struct {
	value *UserModel
	isSet bool
}

func (v NullableUserModel) Get() *UserModel {
	return v.value
}

func (v *NullableUserModel) Set(val *UserModel) {
	v.value = val
	v.isSet = true
}

func (v NullableUserModel) IsSet() bool {
	return v.isSet
}

func (v *NullableUserModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserModel(val *UserModel) *NullableUserModel {
	return &NullableUserModel{value: val, isSet: true}
}

func (v NullableUserModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

