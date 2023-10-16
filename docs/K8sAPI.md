# \K8sAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateNamespace**](K8sAPI.md#CreateNamespace) | **Post** /api/k8s/namespaces | Create Namespace
[**CreateNamespacedJob**](K8sAPI.md#CreateNamespacedJob) | **Post** /api/k8s/clusters/{cluster}/namespaces/{namespace}/jobs | Create K8s Job
[**CreateNamespacedJobTemplate**](K8sAPI.md#CreateNamespacedJobTemplate) | **Post** /api/k8s/cluster/{cluster}/namespace/{namespace}/jobs/template | Create K8s Job Template
[**GetJobDetails**](K8sAPI.md#GetJobDetails) | **Get** /api/k8s/clusters/{cluster}/namespaces/{namespace}/jobsdetails | Get job details
[**GetLogs**](K8sAPI.md#GetLogs) | **Get** /api/k8s/clusters/{cluster}/namespaces/{namespace}/jobs/{job_name}/logs | Get static logs
[**GetLogshell**](K8sAPI.md#GetLogshell) | **Get** /api/k8s/clusters/{cluster}/namespaces/{namespace}/jobs/{job_name}/logshell | Launch a Cloud shell
[**GetNamespacedPvcs**](K8sAPI.md#GetNamespacedPvcs) | **Get** /api/k8s/cluster/{cluster}/namespace/{namespace}/pvcs | Get PVCs
[**GetUsersNamespaces**](K8sAPI.md#GetUsersNamespaces) | **Get** /api/k8s/clusters/{cluster}/users/{username}/namespaces | Get users namespaces
[**LaunchLogshell**](K8sAPI.md#LaunchLogshell) | **Post** /api/k8s/clusters/{cluster}/namespaces/{namespace}/jobs/{job_name}/logshell | Launch a Cloud shell
[**ListAllJobDetails**](K8sAPI.md#ListAllJobDetails) | **Get** /api/k8s/jobsdetails | Get all job&#39;s details in a cluster
[**ListAllJobs**](K8sAPI.md#ListAllJobs) | **Get** /api/k8s/cluster/{cluster}/jobs | Get K8s Job
[**ListClusters**](K8sAPI.md#ListClusters) | **Get** /api/k8s/clusters | List Clusters
[**ListClusters_0**](K8sAPI.md#ListClusters_0) | **Get** /api/k8s/clusters | List Clusters
[**ListNamespacedJobs**](K8sAPI.md#ListNamespacedJobs) | **Get** /api/k8s/clusters/{cluster}/namespaces/{namespace}/jobs | List all jobs in a namespace



## CreateNamespace

> CreateNamespace(ctx).CreateNamespace(createNamespace).Execute()

Create Namespace



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
    createNamespace := *openapiclient.NewCreateNamespace("Namespace_example", "Cluster_example") // CreateNamespace | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.K8sAPI.CreateNamespace(context.Background()).CreateNamespace(createNamespace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `K8sAPI.CreateNamespace``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateNamespaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createNamespace** | [**CreateNamespace**](CreateNamespace.md) |  | 

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


## CreateNamespacedJob

> CreateNamespacedJob(ctx, cluster, namespace).CreateJobBody(createJobBody).Execute()

Create K8s Job



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
    cluster := "cluster_example" // string | 
    namespace := "namespace_example" // string | 
    createJobBody := *openapiclient.NewCreateJobBody("Name_example", "Description_example", "Image_example", "Namespace_example", "Gpus_example", "Pvc_example", "Command_example", "Args_example", "SharedMem_example", "SharedMemUnit_example", false, false) // CreateJobBody | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.K8sAPI.CreateNamespacedJob(context.Background(), cluster, namespace).CreateJobBody(createJobBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `K8sAPI.CreateNamespacedJob``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cluster** | **string** |  | 
**namespace** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateNamespacedJobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **createJobBody** | [**CreateJobBody**](CreateJobBody.md) |  | 

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


## CreateNamespacedJobTemplate

> CreateNamespacedJobTemplate(ctx, cluster, namespace).CreateJobBody(createJobBody).Execute()

Create K8s Job Template



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
    cluster := "cluster_example" // string | 
    namespace := "namespace_example" // string | 
    createJobBody := *openapiclient.NewCreateJobBody("Name_example", "Description_example", "Image_example", "Namespace_example", "Gpus_example", "Pvc_example", "Command_example", "Args_example", "SharedMem_example", "SharedMemUnit_example", false, false) // CreateJobBody | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.K8sAPI.CreateNamespacedJobTemplate(context.Background(), cluster, namespace).CreateJobBody(createJobBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `K8sAPI.CreateNamespacedJobTemplate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cluster** | **string** |  | 
**namespace** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateNamespacedJobTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **createJobBody** | [**CreateJobBody**](CreateJobBody.md) |  | 

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


## GetJobDetails

> []JobDetailsResp GetJobDetails(ctx, cluster, namespace).Execute()

Get job details



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
    cluster := "cluster_example" // string | 
    namespace := "namespace_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.K8sAPI.GetJobDetails(context.Background(), cluster, namespace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `K8sAPI.GetJobDetails``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetJobDetails`: []JobDetailsResp
    fmt.Fprintf(os.Stdout, "Response from `K8sAPI.GetJobDetails`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cluster** | **string** |  | 
**namespace** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetJobDetailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]JobDetailsResp**](JobDetailsResp.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLogs

> string GetLogs(ctx, cluster, namespace, jobName).Execute()

Get static logs



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
    cluster := "cluster_example" // string | 
    namespace := "namespace_example" // string | 
    jobName := "jobName_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.K8sAPI.GetLogs(context.Background(), cluster, namespace, jobName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `K8sAPI.GetLogs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLogs`: string
    fmt.Fprintf(os.Stdout, "Response from `K8sAPI.GetLogs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cluster** | **string** |  | 
**namespace** | **string** |  | 
**jobName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLogsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

**string**

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLogshell

> CloudShellResp GetLogshell(ctx, cluster, namespace, jobName).Execute()

Launch a Cloud shell



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
    cluster := "cluster_example" // string | 
    namespace := "namespace_example" // string | 
    jobName := "jobName_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.K8sAPI.GetLogshell(context.Background(), cluster, namespace, jobName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `K8sAPI.GetLogshell``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLogshell`: CloudShellResp
    fmt.Fprintf(os.Stdout, "Response from `K8sAPI.GetLogshell`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cluster** | **string** |  | 
**namespace** | **string** |  | 
**jobName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLogshellRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**CloudShellResp**](CloudShellResp.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNamespacedPvcs

> []string GetNamespacedPvcs(ctx, cluster, namespace).Execute()

Get PVCs



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
    cluster := "cluster_example" // string | 
    namespace := "namespace_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.K8sAPI.GetNamespacedPvcs(context.Background(), cluster, namespace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `K8sAPI.GetNamespacedPvcs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNamespacedPvcs`: []string
    fmt.Fprintf(os.Stdout, "Response from `K8sAPI.GetNamespacedPvcs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cluster** | **string** |  | 
**namespace** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNamespacedPvcsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**[]string**

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUsersNamespaces

> GetUsersNamespaces(ctx, cluster, username).Execute()

Get users namespaces



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
    cluster := "cluster_example" // string | 
    username := "username_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.K8sAPI.GetUsersNamespaces(context.Background(), cluster, username).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `K8sAPI.GetUsersNamespaces``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cluster** | **string** |  | 
**username** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUsersNamespacesRequest struct via the builder pattern


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


## LaunchLogshell

> LaunchLogshell(ctx, cluster, namespace, jobName).Execute()

Launch a Cloud shell



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
    cluster := "cluster_example" // string | 
    namespace := "namespace_example" // string | 
    jobName := "jobName_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.K8sAPI.LaunchLogshell(context.Background(), cluster, namespace, jobName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `K8sAPI.LaunchLogshell``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cluster** | **string** |  | 
**namespace** | **string** |  | 
**jobName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiLaunchLogshellRequest struct via the builder pattern


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


## ListAllJobDetails

> ListAllJobDetails(ctx).Execute()

Get all job's details in a cluster



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
    r, err := apiClient.K8sAPI.ListAllJobDetails(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `K8sAPI.ListAllJobDetails``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListAllJobDetailsRequest struct via the builder pattern


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


## ListAllJobs

> ListAllJobs(ctx, cluster).Execute()

Get K8s Job



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
    cluster := "cluster_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.K8sAPI.ListAllJobs(context.Background(), cluster).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `K8sAPI.ListAllJobs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cluster** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListAllJobsRequest struct via the builder pattern


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


## ListClusters

> ListClusters(ctx).Execute()

List Clusters



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
    r, err := apiClient.K8sAPI.ListClusters(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `K8sAPI.ListClusters``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListClustersRequest struct via the builder pattern


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


## ListClusters_0

> ListClusters_0(ctx).Execute()

List Clusters



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
    r, err := apiClient.K8sAPI.ListClusters_0(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `K8sAPI.ListClusters_0``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListClusters_1Request struct via the builder pattern


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


## ListNamespacedJobs

> ListNamespacedJobs(ctx, cluster, namespace).Execute()

List all jobs in a namespace



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
    cluster := "cluster_example" // string | 
    namespace := "namespace_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.K8sAPI.ListNamespacedJobs(context.Background(), cluster, namespace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `K8sAPI.ListNamespacedJobs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cluster** | **string** |  | 
**namespace** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListNamespacedJobsRequest struct via the builder pattern


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

