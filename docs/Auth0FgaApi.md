# \Auth0FgaApi

All URIs are relative to *https://api.us1.fga.dev*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Check**](Auth0FgaApi.md#Check) | **Post** /stores/{store_id}/check | Check whether a user is authorized to access an object
[**Expand**](Auth0FgaApi.md#Expand) | **Post** /stores/{store_id}/expand | Expand all relationships in userset tree format, and following userset rewrite rules.  Useful to reason about and debug a certain relationship
[**Read**](Auth0FgaApi.md#Read) | **Post** /stores/{store_id}/read | Get tuples from the store that matches a query, without following userset rewrite rules
[**ReadAssertions**](Auth0FgaApi.md#ReadAssertions) | **Get** /stores/{store_id}/assertions/{authorization_model_id} | Read assertions for an authorization model ID
[**ReadAuthorizationModel**](Auth0FgaApi.md#ReadAuthorizationModel) | **Get** /stores/{store_id}/authorization-models/{id} | Return a particular version of an authorization model
[**ReadAuthorizationModels**](Auth0FgaApi.md#ReadAuthorizationModels) | **Get** /stores/{store_id}/authorization-models | Return all the authorization model IDs for a particular store
[**ReadChanges**](Auth0FgaApi.md#ReadChanges) | **Get** /stores/{store_id}/changes | Return a list of all the tuple changes
[**Write**](Auth0FgaApi.md#Write) | **Post** /stores/{store_id}/write | Add or delete tuples from the store
[**WriteAssertions**](Auth0FgaApi.md#WriteAssertions) | **Put** /stores/{store_id}/assertions/{authorization_model_id} | Upsert assertions for an authorization model ID
[**WriteAuthorizationModel**](Auth0FgaApi.md#WriteAuthorizationModel) | **Post** /stores/{store_id}/authorization-models | Create a new authorization model



## Check

> CheckResponse Check(ctx).Body(body).Execute()

Check whether a user is authorized to access an object



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    auth0fga "github.com/auth0-lab/fga-go-sdk"
)

func main() {
    
    body := *openapiclient.NewCheckRequest() // CheckRequest | 

    // See https://github.com/auth0-lab/fga-go-sdk#getting-your-store-id-client-id-and-client-secret
    configuration, err := auth0fga.NewConfiguration(auth0fga.Configuration{
        Environment:  os.Getenv("AUTH0_FGA_ENVIRONMENT"), // can be: "us"/"staging"/"playground"
        StoreId:      os.Getenv("AUTH0_FGA_STORE_ID"),
        ClientId:     os.Getenv("AUTH0_FGA_CLIENT_ID"), // Required for all environments except playground
        ClientSecret: os.Getenv("AUTH0_FGA_CLIENT_SECRET"), // Required for all environments except playground
    })

    apiClient := auth0fga.NewAPIClient(configuration)

    resp, r, err := apiClient.Auth0FgaApi.Check(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Auth0FgaApi.Check``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
        switch v := err.(type) {
        case FgaApiAuthenticationError:
            // Handle authentication error
        case FgaApiValidationError:
            // Handle parameter validation error
        case FgaApiNotFoundError:
            // Handle not found error
        case FgaApiInternalError:
            // Handle API internal error
        case FgaApiRateLimitError:
            // Exponential backoff in handling rate limit error
        default:
            // Handle unknown/undefined error
        }
    }
    // response from `Check`: CheckResponse
    fmt.Fprintf(os.Stdout, "Response from `Auth0FgaApi.Check`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.

### Other Parameters

Other parameters are passed through a pointer to a apiCheckRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**body** | [**CheckRequest**](CheckRequest.md) |  | 

### Return type

[**CheckResponse**](CheckResponse.md)

### Authorization

[ClientCredentials](../README.md#ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Expand

> ExpandResponse Expand(ctx).Body(body).Execute()

Expand all relationships in userset tree format, and following userset rewrite rules.  Useful to reason about and debug a certain relationship



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    auth0fga "github.com/auth0-lab/fga-go-sdk"
)

func main() {
    
    body := *openapiclient.NewExpandRequest() // ExpandRequest | 

    // See https://github.com/auth0-lab/fga-go-sdk#getting-your-store-id-client-id-and-client-secret
    configuration, err := auth0fga.NewConfiguration(auth0fga.Configuration{
        Environment:  os.Getenv("AUTH0_FGA_ENVIRONMENT"), // can be: "us"/"staging"/"playground"
        StoreId:      os.Getenv("AUTH0_FGA_STORE_ID"),
        ClientId:     os.Getenv("AUTH0_FGA_CLIENT_ID"), // Required for all environments except playground
        ClientSecret: os.Getenv("AUTH0_FGA_CLIENT_SECRET"), // Required for all environments except playground
    })

    apiClient := auth0fga.NewAPIClient(configuration)

    resp, r, err := apiClient.Auth0FgaApi.Expand(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Auth0FgaApi.Expand``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
        switch v := err.(type) {
        case FgaApiAuthenticationError:
            // Handle authentication error
        case FgaApiValidationError:
            // Handle parameter validation error
        case FgaApiNotFoundError:
            // Handle not found error
        case FgaApiInternalError:
            // Handle API internal error
        case FgaApiRateLimitError:
            // Exponential backoff in handling rate limit error
        default:
            // Handle unknown/undefined error
        }
    }
    // response from `Expand`: ExpandResponse
    fmt.Fprintf(os.Stdout, "Response from `Auth0FgaApi.Expand`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.

### Other Parameters

Other parameters are passed through a pointer to a apiExpandRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**body** | [**ExpandRequest**](ExpandRequest.md) |  | 

### Return type

[**ExpandResponse**](ExpandResponse.md)

### Authorization

[ClientCredentials](../README.md#ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Read

> ReadResponse Read(ctx).Body(body).Execute()

Get tuples from the store that matches a query, without following userset rewrite rules



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    auth0fga "github.com/auth0-lab/fga-go-sdk"
)

func main() {
    
    body := *openapiclient.NewReadRequest() // ReadRequest | 

    // See https://github.com/auth0-lab/fga-go-sdk#getting-your-store-id-client-id-and-client-secret
    configuration, err := auth0fga.NewConfiguration(auth0fga.Configuration{
        Environment:  os.Getenv("AUTH0_FGA_ENVIRONMENT"), // can be: "us"/"staging"/"playground"
        StoreId:      os.Getenv("AUTH0_FGA_STORE_ID"),
        ClientId:     os.Getenv("AUTH0_FGA_CLIENT_ID"), // Required for all environments except playground
        ClientSecret: os.Getenv("AUTH0_FGA_CLIENT_SECRET"), // Required for all environments except playground
    })

    apiClient := auth0fga.NewAPIClient(configuration)

    resp, r, err := apiClient.Auth0FgaApi.Read(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Auth0FgaApi.Read``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
        switch v := err.(type) {
        case FgaApiAuthenticationError:
            // Handle authentication error
        case FgaApiValidationError:
            // Handle parameter validation error
        case FgaApiNotFoundError:
            // Handle not found error
        case FgaApiInternalError:
            // Handle API internal error
        case FgaApiRateLimitError:
            // Exponential backoff in handling rate limit error
        default:
            // Handle unknown/undefined error
        }
    }
    // response from `Read`: ReadResponse
    fmt.Fprintf(os.Stdout, "Response from `Auth0FgaApi.Read`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.

### Other Parameters

Other parameters are passed through a pointer to a apiReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**body** | [**ReadRequest**](ReadRequest.md) |  | 

### Return type

[**ReadResponse**](ReadResponse.md)

### Authorization

[ClientCredentials](../README.md#ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadAssertions

> ReadAssertionsResponse ReadAssertions(ctx, authorizationModelId).Execute()

Read assertions for an authorization model ID



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    auth0fga "github.com/auth0-lab/fga-go-sdk"
)

func main() {
    
    authorizationModelId := "authorizationModelId_example" // string | 

    // See https://github.com/auth0-lab/fga-go-sdk#getting-your-store-id-client-id-and-client-secret
    configuration, err := auth0fga.NewConfiguration(auth0fga.Configuration{
        Environment:  os.Getenv("AUTH0_FGA_ENVIRONMENT"), // can be: "us"/"staging"/"playground"
        StoreId:      os.Getenv("AUTH0_FGA_STORE_ID"),
        ClientId:     os.Getenv("AUTH0_FGA_CLIENT_ID"), // Required for all environments except playground
        ClientSecret: os.Getenv("AUTH0_FGA_CLIENT_SECRET"), // Required for all environments except playground
    })

    apiClient := auth0fga.NewAPIClient(configuration)

    resp, r, err := apiClient.Auth0FgaApi.ReadAssertions(context.Background(), authorizationModelId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Auth0FgaApi.ReadAssertions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
        switch v := err.(type) {
        case FgaApiAuthenticationError:
            // Handle authentication error
        case FgaApiValidationError:
            // Handle parameter validation error
        case FgaApiNotFoundError:
            // Handle not found error
        case FgaApiInternalError:
            // Handle API internal error
        case FgaApiRateLimitError:
            // Exponential backoff in handling rate limit error
        default:
            // Handle unknown/undefined error
        }
    }
    // response from `ReadAssertions`: ReadAssertionsResponse
    fmt.Fprintf(os.Stdout, "Response from `Auth0FgaApi.ReadAssertions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**authorizationModelId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadAssertionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

### Return type

[**ReadAssertionsResponse**](ReadAssertionsResponse.md)

### Authorization

[ClientCredentials](../README.md#ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadAuthorizationModel

> ReadAuthorizationModelResponse ReadAuthorizationModel(ctx, id).Execute()

Return a particular version of an authorization model



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    auth0fga "github.com/auth0-lab/fga-go-sdk"
)

func main() {
    
    id := "id_example" // string | 

    // See https://github.com/auth0-lab/fga-go-sdk#getting-your-store-id-client-id-and-client-secret
    configuration, err := auth0fga.NewConfiguration(auth0fga.Configuration{
        Environment:  os.Getenv("AUTH0_FGA_ENVIRONMENT"), // can be: "us"/"staging"/"playground"
        StoreId:      os.Getenv("AUTH0_FGA_STORE_ID"),
        ClientId:     os.Getenv("AUTH0_FGA_CLIENT_ID"), // Required for all environments except playground
        ClientSecret: os.Getenv("AUTH0_FGA_CLIENT_SECRET"), // Required for all environments except playground
    })

    apiClient := auth0fga.NewAPIClient(configuration)

    resp, r, err := apiClient.Auth0FgaApi.ReadAuthorizationModel(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Auth0FgaApi.ReadAuthorizationModel``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
        switch v := err.(type) {
        case FgaApiAuthenticationError:
            // Handle authentication error
        case FgaApiValidationError:
            // Handle parameter validation error
        case FgaApiNotFoundError:
            // Handle not found error
        case FgaApiInternalError:
            // Handle API internal error
        case FgaApiRateLimitError:
            // Exponential backoff in handling rate limit error
        default:
            // Handle unknown/undefined error
        }
    }
    // response from `ReadAuthorizationModel`: ReadAuthorizationModelResponse
    fmt.Fprintf(os.Stdout, "Response from `Auth0FgaApi.ReadAuthorizationModel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadAuthorizationModelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

### Return type

[**ReadAuthorizationModelResponse**](ReadAuthorizationModelResponse.md)

### Authorization

[ClientCredentials](../README.md#ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadAuthorizationModels

> ReadAuthorizationModelsResponse ReadAuthorizationModels(ctx).PageSize(pageSize).ContinuationToken(continuationToken).Execute()

Return all the authorization model IDs for a particular store



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    auth0fga "github.com/auth0-lab/fga-go-sdk"
)

func main() {
    
    pageSize := int32(56) // int32 |  (optional)
    continuationToken := "continuationToken_example" // string |  (optional)

    // See https://github.com/auth0-lab/fga-go-sdk#getting-your-store-id-client-id-and-client-secret
    configuration, err := auth0fga.NewConfiguration(auth0fga.Configuration{
        Environment:  os.Getenv("AUTH0_FGA_ENVIRONMENT"), // can be: "us"/"staging"/"playground"
        StoreId:      os.Getenv("AUTH0_FGA_STORE_ID"),
        ClientId:     os.Getenv("AUTH0_FGA_CLIENT_ID"), // Required for all environments except playground
        ClientSecret: os.Getenv("AUTH0_FGA_CLIENT_SECRET"), // Required for all environments except playground
    })

    apiClient := auth0fga.NewAPIClient(configuration)

    resp, r, err := apiClient.Auth0FgaApi.ReadAuthorizationModels(context.Background()).PageSize(pageSize).ContinuationToken(continuationToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Auth0FgaApi.ReadAuthorizationModels``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
        switch v := err.(type) {
        case FgaApiAuthenticationError:
            // Handle authentication error
        case FgaApiValidationError:
            // Handle parameter validation error
        case FgaApiNotFoundError:
            // Handle not found error
        case FgaApiInternalError:
            // Handle API internal error
        case FgaApiRateLimitError:
            // Exponential backoff in handling rate limit error
        default:
            // Handle unknown/undefined error
        }
    }
    // response from `ReadAuthorizationModels`: ReadAuthorizationModelsResponse
    fmt.Fprintf(os.Stdout, "Response from `Auth0FgaApi.ReadAuthorizationModels`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.

### Other Parameters

Other parameters are passed through a pointer to a apiReadAuthorizationModelsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**pageSize** | **int32** |  | 
**continuationToken** | **string** |  | 

### Return type

[**ReadAuthorizationModelsResponse**](ReadAuthorizationModelsResponse.md)

### Authorization

[ClientCredentials](../README.md#ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadChanges

> ReadChangesResponse ReadChanges(ctx).Type_(type_).PageSize(pageSize).ContinuationToken(continuationToken).Execute()

Return a list of all the tuple changes



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    auth0fga "github.com/auth0-lab/fga-go-sdk"
)

func main() {
    
    type_ := "type__example" // string |  (optional)
    pageSize := int32(56) // int32 |  (optional)
    continuationToken := "continuationToken_example" // string |  (optional)

    // See https://github.com/auth0-lab/fga-go-sdk#getting-your-store-id-client-id-and-client-secret
    configuration, err := auth0fga.NewConfiguration(auth0fga.Configuration{
        Environment:  os.Getenv("AUTH0_FGA_ENVIRONMENT"), // can be: "us"/"staging"/"playground"
        StoreId:      os.Getenv("AUTH0_FGA_STORE_ID"),
        ClientId:     os.Getenv("AUTH0_FGA_CLIENT_ID"), // Required for all environments except playground
        ClientSecret: os.Getenv("AUTH0_FGA_CLIENT_SECRET"), // Required for all environments except playground
    })

    apiClient := auth0fga.NewAPIClient(configuration)

    resp, r, err := apiClient.Auth0FgaApi.ReadChanges(context.Background()).Type_(type_).PageSize(pageSize).ContinuationToken(continuationToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Auth0FgaApi.ReadChanges``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
        switch v := err.(type) {
        case FgaApiAuthenticationError:
            // Handle authentication error
        case FgaApiValidationError:
            // Handle parameter validation error
        case FgaApiNotFoundError:
            // Handle not found error
        case FgaApiInternalError:
            // Handle API internal error
        case FgaApiRateLimitError:
            // Exponential backoff in handling rate limit error
        default:
            // Handle unknown/undefined error
        }
    }
    // response from `ReadChanges`: ReadChangesResponse
    fmt.Fprintf(os.Stdout, "Response from `Auth0FgaApi.ReadChanges`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.

### Other Parameters

Other parameters are passed through a pointer to a apiReadChangesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**type_** | **string** |  | 
**pageSize** | **int32** |  | 
**continuationToken** | **string** |  | 

### Return type

[**ReadChangesResponse**](ReadChangesResponse.md)

### Authorization

[ClientCredentials](../README.md#ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Write

> map[string]interface{} Write(ctx).Body(body).Execute()

Add or delete tuples from the store



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    auth0fga "github.com/auth0-lab/fga-go-sdk"
)

func main() {
    
    body := *openapiclient.NewWriteRequest() // WriteRequest | 

    // See https://github.com/auth0-lab/fga-go-sdk#getting-your-store-id-client-id-and-client-secret
    configuration, err := auth0fga.NewConfiguration(auth0fga.Configuration{
        Environment:  os.Getenv("AUTH0_FGA_ENVIRONMENT"), // can be: "us"/"staging"/"playground"
        StoreId:      os.Getenv("AUTH0_FGA_STORE_ID"),
        ClientId:     os.Getenv("AUTH0_FGA_CLIENT_ID"), // Required for all environments except playground
        ClientSecret: os.Getenv("AUTH0_FGA_CLIENT_SECRET"), // Required for all environments except playground
    })

    apiClient := auth0fga.NewAPIClient(configuration)

    resp, r, err := apiClient.Auth0FgaApi.Write(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Auth0FgaApi.Write``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
        switch v := err.(type) {
        case FgaApiAuthenticationError:
            // Handle authentication error
        case FgaApiValidationError:
            // Handle parameter validation error
        case FgaApiNotFoundError:
            // Handle not found error
        case FgaApiInternalError:
            // Handle API internal error
        case FgaApiRateLimitError:
            // Exponential backoff in handling rate limit error
        default:
            // Handle unknown/undefined error
        }
    }
    // response from `Write`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `Auth0FgaApi.Write`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.

### Other Parameters

Other parameters are passed through a pointer to a apiWriteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**body** | [**WriteRequest**](WriteRequest.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

[ClientCredentials](../README.md#ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WriteAssertions

> WriteAssertions(ctx, authorizationModelId).Body(body).Execute()

Upsert assertions for an authorization model ID



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    auth0fga "github.com/auth0-lab/fga-go-sdk"
)

func main() {
    
    authorizationModelId := "authorizationModelId_example" // string | 
    body := *openapiclient.NewWriteAssertionsRequest([]openapiclient.Assertion{*openapiclient.NewAssertion(false)}) // WriteAssertionsRequest | 

    // See https://github.com/auth0-lab/fga-go-sdk#getting-your-store-id-client-id-and-client-secret
    configuration, err := auth0fga.NewConfiguration(auth0fga.Configuration{
        Environment:  os.Getenv("AUTH0_FGA_ENVIRONMENT"), // can be: "us"/"staging"/"playground"
        StoreId:      os.Getenv("AUTH0_FGA_STORE_ID"),
        ClientId:     os.Getenv("AUTH0_FGA_CLIENT_ID"), // Required for all environments except playground
        ClientSecret: os.Getenv("AUTH0_FGA_CLIENT_SECRET"), // Required for all environments except playground
    })

    apiClient := auth0fga.NewAPIClient(configuration)

    resp, r, err := apiClient.Auth0FgaApi.WriteAssertions(context.Background(), authorizationModelId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Auth0FgaApi.WriteAssertions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
        switch v := err.(type) {
        case FgaApiAuthenticationError:
            // Handle authentication error
        case FgaApiValidationError:
            // Handle parameter validation error
        case FgaApiNotFoundError:
            // Handle not found error
        case FgaApiInternalError:
            // Handle API internal error
        case FgaApiRateLimitError:
            // Exponential backoff in handling rate limit error
        default:
            // Handle unknown/undefined error
        }
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**authorizationModelId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiWriteAssertionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**body** | [**WriteAssertionsRequest**](WriteAssertionsRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[ClientCredentials](../README.md#ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WriteAuthorizationModel

> WriteAuthorizationModelResponse WriteAuthorizationModel(ctx).TypeDefinitions(typeDefinitions).Execute()

Create a new authorization model



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    auth0fga "github.com/auth0-lab/fga-go-sdk"
)

func main() {
    
    typeDefinitions := *openapiclient.NewTypeDefinitions() // TypeDefinitions | 

    // See https://github.com/auth0-lab/fga-go-sdk#getting-your-store-id-client-id-and-client-secret
    configuration, err := auth0fga.NewConfiguration(auth0fga.Configuration{
        Environment:  os.Getenv("AUTH0_FGA_ENVIRONMENT"), // can be: "us"/"staging"/"playground"
        StoreId:      os.Getenv("AUTH0_FGA_STORE_ID"),
        ClientId:     os.Getenv("AUTH0_FGA_CLIENT_ID"), // Required for all environments except playground
        ClientSecret: os.Getenv("AUTH0_FGA_CLIENT_SECRET"), // Required for all environments except playground
    })

    apiClient := auth0fga.NewAPIClient(configuration)

    resp, r, err := apiClient.Auth0FgaApi.WriteAuthorizationModel(context.Background()).TypeDefinitions(typeDefinitions).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Auth0FgaApi.WriteAuthorizationModel``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
        switch v := err.(type) {
        case FgaApiAuthenticationError:
            // Handle authentication error
        case FgaApiValidationError:
            // Handle parameter validation error
        case FgaApiNotFoundError:
            // Handle not found error
        case FgaApiInternalError:
            // Handle API internal error
        case FgaApiRateLimitError:
            // Exponential backoff in handling rate limit error
        default:
            // Handle unknown/undefined error
        }
    }
    // response from `WriteAuthorizationModel`: WriteAuthorizationModelResponse
    fmt.Fprintf(os.Stdout, "Response from `Auth0FgaApi.WriteAuthorizationModel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.

### Other Parameters

Other parameters are passed through a pointer to a apiWriteAuthorizationModelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**typeDefinitions** | [**TypeDefinitions**](TypeDefinitions.md) |  | 

### Return type

[**WriteAuthorizationModelResponse**](WriteAuthorizationModelResponse.md)

### Authorization

[ClientCredentials](../README.md#ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

