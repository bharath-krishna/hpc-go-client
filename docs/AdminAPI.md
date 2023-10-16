# \AdminAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddKcUserRole**](AdminAPI.md#AddKcUserRole) | **Put** /api/admin/users/{username}/rolemaps/{role_name} | Add Kc User Role
[**AddKcUserToGroup**](AdminAPI.md#AddKcUserToGroup) | **Put** /api/admin/users/{username}/groups/{group_name} | Add user to group
[**CheckUserRoles**](AdminAPI.md#CheckUserRoles) | **Get** /api/admin/users/{username}/roles/{role_name} | Check if users role exists
[**CreateKcGroup**](AdminAPI.md#CreateKcGroup) | **Post** /api/admin/groups | Create Group
[**CreateUsersRbac**](AdminAPI.md#CreateUsersRbac) | **Post** /api/admin/namespace/{namespace}/rbac | Create roles and rolebindings of namespace
[**ListUsers**](AdminAPI.md#ListUsers) | **Get** /api/admin/users | List Users
[**ListUsers_0**](AdminAPI.md#ListUsers_0) | **Get** /api/admin/users | List Users
[**OnboardUser**](AdminAPI.md#OnboardUser) | **Post** /api/admin/users | Onboard user



## AddKcUserRole

> AddKcUserRole(ctx, username, roleName).Execute()

Add Kc User Role

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/bharath-krishna/hpc-go-client"
)

func main() {
    username := "username_example" // string | 
    roleName := "roleName_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AdminAPI.AddKcUserRole(context.Background(), username, roleName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.AddKcUserRole``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**username** | **string** |  | 
**roleName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddKcUserRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddKcUserToGroup

> AddKcUserToGroup(ctx, username, groupName).Execute()

Add user to group

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/bharath-krishna/hpc-go-client"
)

func main() {
    username := "username_example" // string | 
    groupName := "groupName_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AdminAPI.AddKcUserToGroup(context.Background(), username, groupName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.AddKcUserToGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**username** | **string** |  | 
**groupName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddKcUserToGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CheckUserRoles

> CheckUserRoles(ctx, username, roleName).Execute()

Check if users role exists

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/bharath-krishna/hpc-go-client"
)

func main() {
    username := "username_example" // string | 
    roleName := "roleName_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AdminAPI.CheckUserRoles(context.Background(), username, roleName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.CheckUserRoles``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**username** | **string** |  | 
**roleName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCheckUserRolesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateKcGroup

> CreateKcGroup(ctx).CreateGroupBody(createGroupBody).Execute()

Create Group

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/bharath-krishna/hpc-go-client"
)

func main() {
    createGroupBody := *openapiclient.NewCreateGroupBody("GroupName_example") // CreateGroupBody | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AdminAPI.CreateKcGroup(context.Background()).CreateGroupBody(createGroupBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.CreateKcGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateKcGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createGroupBody** | [**CreateGroupBody**](CreateGroupBody.md) |  | 

### Return type

 (empty response body)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateUsersRbac

> CreateUsersRbac(ctx, namespace).Execute()

Create roles and rolebindings of namespace



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/bharath-krishna/hpc-go-client"
)

func main() {
    namespace := "namespace_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AdminAPI.CreateUsersRbac(context.Background(), namespace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.CreateUsersRbac``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**namespace** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateUsersRbacRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListUsers

> ListUsers(ctx).Execute()

List Users



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/bharath-krishna/hpc-go-client"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AdminAPI.ListUsers(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.ListUsers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListUsersRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListUsers_0

> ListUsers_0(ctx).Execute()

List Users



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/bharath-krishna/hpc-go-client"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AdminAPI.ListUsers_0(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.ListUsers_0``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListUsers_1Request struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OnboardUser

> OnboardUser(ctx).OnboardUser(onboardUser).Execute()

Onboard user



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/bharath-krishna/hpc-go-client"
)

func main() {
    onboardUser := *openapiclient.NewOnboardUser("Username_example", "Namespace_example", "Cluster_example", "RoleName_example") // OnboardUser | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AdminAPI.OnboardUser(context.Background()).OnboardUser(onboardUser).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.OnboardUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOnboardUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **onboardUser** | [**OnboardUser**](OnboardUser.md) |  | 

### Return type

 (empty response body)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

