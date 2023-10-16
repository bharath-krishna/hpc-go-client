/*
HPC Portal - API

Testing AdminAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package hpc

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/bharath-krishna/hpc-go-client"
)

func Test_hpc_AdminAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test AdminAPIService AddKcUserRole", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var username string
		var roleName string

		httpRes, err := apiClient.AdminAPI.AddKcUserRole(context.Background(), username, roleName).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AdminAPIService AddKcUserToGroup", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var username string
		var groupName string

		httpRes, err := apiClient.AdminAPI.AddKcUserToGroup(context.Background(), username, groupName).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AdminAPIService CheckUserRoles", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var username string
		var roleName string

		httpRes, err := apiClient.AdminAPI.CheckUserRoles(context.Background(), username, roleName).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AdminAPIService CreateKcGroup", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.AdminAPI.CreateKcGroup(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AdminAPIService CreateUsersRbac", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var namespace string

		httpRes, err := apiClient.AdminAPI.CreateUsersRbac(context.Background(), namespace).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AdminAPIService ListUsers", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.AdminAPI.ListUsers(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AdminAPIService ListUsers_1", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.AdminAPI.ListUsers_0(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AdminAPIService OnboardUser", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.AdminAPI.OnboardUser(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}