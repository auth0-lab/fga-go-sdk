# \Auth0FgaApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Check**](Auth0FgaApi.md#Check) | **Post** /{store_id}/check | Check whether a user is authorized to access an object
[**DeleteTokenIssuer**](Auth0FgaApi.md#DeleteTokenIssuer) | **Delete** /{store_id}/settings/token-issuers/{id} | Remove 3rd party token issuer for Auth0 FGA read and write operation
[**Expand**](Auth0FgaApi.md#Expand) | **Post** /{store_id}/expand | Expand all relationships in userset tree format, and following userset rewrite rules.  Useful to reason about and debug a certain relationship
[**Read**](Auth0FgaApi.md#Read) | **Post** /{store_id}/read | Get tuples from the store that matches a query, without following userset rewrite rules
[**ReadAssertions**](Auth0FgaApi.md#ReadAssertions) | **Get** /{store_id}/assertions/{authorization_model_id} | Read assertions for an authorization model ID
[**ReadAuthorizationModel**](Auth0FgaApi.md#ReadAuthorizationModel) | **Get** /{store_id}/authorization-models/{id} | Return a particular version of an authorization model
[**ReadAuthorizationModels**](Auth0FgaApi.md#ReadAuthorizationModels) | **Get** /{store_id}/authorization-models | Return all the authorization model IDs for a particular store
[**ReadSettings**](Auth0FgaApi.md#ReadSettings) | **Get** /{store_id}/settings | Return store settings, including the environment tag
[**Write**](Auth0FgaApi.md#Write) | **Post** /{store_id}/write | Add or delete tuples from the store
[**WriteAssertions**](Auth0FgaApi.md#WriteAssertions) | **Post** /{store_id}/assertions/{authorization_model_id} | Upsert assertions for an authorization model ID
[**WriteAuthorizationModel**](Auth0FgaApi.md#WriteAuthorizationModel) | **Post** /{store_id}/authorization-models | Create a new authorization model
[**WriteSettings**](Auth0FgaApi.md#WriteSettings) | **Patch** /{store_id}/settings | Update the environment tag for a store
[**WriteTokenIssuer**](Auth0FgaApi.md#WriteTokenIssuer) | **Post** /{store_id}/settings/token-issuers | Add 3rd party token issuer for Auth0 FGA read and write operations



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
    
    body := *openapiclient.NewCheckRequestParams() // CheckRequestParams | 

    configuration := auth0fga.NewConfiguration(UserConfiguration{
        StoreId:      os.Getenv("AUTH0_FGA_STORE_ID"),
        ClientId:     os.Getenv("AUTH0_FGA_CLIENT_ID"),
        ClientSecret: os.Getenv("AUTH0_FGA_CLIENT_SECRET"),
        Environment:  os.Getenv("AUTH0_FGA_ENVIRONMENT"),
    })

    apiClient := auth0fga.NewAPIClient(configuration)

    resp, r, err := apiClient.Auth0FgaApi.Check(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Auth0FgaApi.Check``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
        switch v := err.(type) {
        case Auth0FgaApiAuthenticationError:
            // Handle authentication error
        case Auth0FgaApiValidationError:
            // Handle parameter validation error
        case Auth0FgaApiInternalError:
            // Handle API internal error
        case Auth0FgaApiRateLimitError:
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
**body** | [**CheckRequestParams**](CheckRequestParams.md) |  | 

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


## DeleteTokenIssuer

> map[string]interface{} DeleteTokenIssuer(ctx, id).Execute()

Remove 3rd party token issuer for Auth0 FGA read and write operation



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
    
    id := "id_example" // string | Id of token issuer to be removed

    configuration := auth0fga.NewConfiguration(UserConfiguration{
        StoreId:      os.Getenv("AUTH0_FGA_STORE_ID"),
        ClientId:     os.Getenv("AUTH0_FGA_CLIENT_ID"),
        ClientSecret: os.Getenv("AUTH0_FGA_CLIENT_SECRET"),
        Environment:  os.Getenv("AUTH0_FGA_ENVIRONMENT"),
    })

    apiClient := auth0fga.NewAPIClient(configuration)

    resp, r, err := apiClient.Auth0FgaApi.DeleteTokenIssuer(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Auth0FgaApi.DeleteTokenIssuer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
        switch v := err.(type) {
        case Auth0FgaApiAuthenticationError:
            // Handle authentication error
        case Auth0FgaApiValidationError:
            // Handle parameter validation error
        case Auth0FgaApiInternalError:
            // Handle API internal error
        case Auth0FgaApiRateLimitError:
            // Exponential backoff in handling rate limit error
        default:
            // Handle unknown/undefined error
        }
    }
    // response from `DeleteTokenIssuer`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `Auth0FgaApi.DeleteTokenIssuer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Id of token issuer to be removed | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTokenIssuerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

### Return type

**map[string]interface{}**

### Authorization

[ClientCredentials](../README.md#ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
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
    
    body := *openapiclient.NewExpandRequestParams() // ExpandRequestParams | 

    configuration := auth0fga.NewConfiguration(UserConfiguration{
        StoreId:      os.Getenv("AUTH0_FGA_STORE_ID"),
        ClientId:     os.Getenv("AUTH0_FGA_CLIENT_ID"),
        ClientSecret: os.Getenv("AUTH0_FGA_CLIENT_SECRET"),
        Environment:  os.Getenv("AUTH0_FGA_ENVIRONMENT"),
    })

    apiClient := auth0fga.NewAPIClient(configuration)

    resp, r, err := apiClient.Auth0FgaApi.Expand(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Auth0FgaApi.Expand``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
        switch v := err.(type) {
        case Auth0FgaApiAuthenticationError:
            // Handle authentication error
        case Auth0FgaApiValidationError:
            // Handle parameter validation error
        case Auth0FgaApiInternalError:
            // Handle API internal error
        case Auth0FgaApiRateLimitError:
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
**body** | [**ExpandRequestParams**](ExpandRequestParams.md) |  | 

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
    
    body := *openapiclient.NewReadRequestParams() // ReadRequestParams | 

    configuration := auth0fga.NewConfiguration(UserConfiguration{
        StoreId:      os.Getenv("AUTH0_FGA_STORE_ID"),
        ClientId:     os.Getenv("AUTH0_FGA_CLIENT_ID"),
        ClientSecret: os.Getenv("AUTH0_FGA_CLIENT_SECRET"),
        Environment:  os.Getenv("AUTH0_FGA_ENVIRONMENT"),
    })

    apiClient := auth0fga.NewAPIClient(configuration)

    resp, r, err := apiClient.Auth0FgaApi.Read(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Auth0FgaApi.Read``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
        switch v := err.(type) {
        case Auth0FgaApiAuthenticationError:
            // Handle authentication error
        case Auth0FgaApiValidationError:
            // Handle parameter validation error
        case Auth0FgaApiInternalError:
            // Handle API internal error
        case Auth0FgaApiRateLimitError:
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
**body** | [**ReadRequestParams**](ReadRequestParams.md) |  | 

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

    configuration := auth0fga.NewConfiguration(UserConfiguration{
        StoreId:      os.Getenv("AUTH0_FGA_STORE_ID"),
        ClientId:     os.Getenv("AUTH0_FGA_CLIENT_ID"),
        ClientSecret: os.Getenv("AUTH0_FGA_CLIENT_SECRET"),
        Environment:  os.Getenv("AUTH0_FGA_ENVIRONMENT"),
    })

    apiClient := auth0fga.NewAPIClient(configuration)

    resp, r, err := apiClient.Auth0FgaApi.ReadAssertions(context.Background(), authorizationModelId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Auth0FgaApi.ReadAssertions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
        switch v := err.(type) {
        case Auth0FgaApiAuthenticationError:
            // Handle authentication error
        case Auth0FgaApiValidationError:
            // Handle parameter validation error
        case Auth0FgaApiInternalError:
            // Handle API internal error
        case Auth0FgaApiRateLimitError:
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

    configuration := auth0fga.NewConfiguration(UserConfiguration{
        StoreId:      os.Getenv("AUTH0_FGA_STORE_ID"),
        ClientId:     os.Getenv("AUTH0_FGA_CLIENT_ID"),
        ClientSecret: os.Getenv("AUTH0_FGA_CLIENT_SECRET"),
        Environment:  os.Getenv("AUTH0_FGA_ENVIRONMENT"),
    })

    apiClient := auth0fga.NewAPIClient(configuration)

    resp, r, err := apiClient.Auth0FgaApi.ReadAuthorizationModel(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Auth0FgaApi.ReadAuthorizationModel``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
        switch v := err.(type) {
        case Auth0FgaApiAuthenticationError:
            // Handle authentication error
        case Auth0FgaApiValidationError:
            // Handle parameter validation error
        case Auth0FgaApiInternalError:
            // Handle API internal error
        case Auth0FgaApiRateLimitError:
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

    configuration := auth0fga.NewConfiguration(UserConfiguration{
        StoreId:      os.Getenv("AUTH0_FGA_STORE_ID"),
        ClientId:     os.Getenv("AUTH0_FGA_CLIENT_ID"),
        ClientSecret: os.Getenv("AUTH0_FGA_CLIENT_SECRET"),
        Environment:  os.Getenv("AUTH0_FGA_ENVIRONMENT"),
    })

    apiClient := auth0fga.NewAPIClient(configuration)

    resp, r, err := apiClient.Auth0FgaApi.ReadAuthorizationModels(context.Background()).PageSize(pageSize).ContinuationToken(continuationToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Auth0FgaApi.ReadAuthorizationModels``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
        switch v := err.(type) {
        case Auth0FgaApiAuthenticationError:
            // Handle authentication error
        case Auth0FgaApiValidationError:
            // Handle parameter validation error
        case Auth0FgaApiInternalError:
            // Handle API internal error
        case Auth0FgaApiRateLimitError:
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


## ReadSettings

> Settings ReadSettings(ctx).Execute()

Return store settings, including the environment tag



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
    

    configuration := auth0fga.NewConfiguration(UserConfiguration{
        StoreId:      os.Getenv("AUTH0_FGA_STORE_ID"),
        ClientId:     os.Getenv("AUTH0_FGA_CLIENT_ID"),
        ClientSecret: os.Getenv("AUTH0_FGA_CLIENT_SECRET"),
        Environment:  os.Getenv("AUTH0_FGA_ENVIRONMENT"),
    })

    apiClient := auth0fga.NewAPIClient(configuration)

    resp, r, err := apiClient.Auth0FgaApi.ReadSettings(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Auth0FgaApi.ReadSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
        switch v := err.(type) {
        case Auth0FgaApiAuthenticationError:
            // Handle authentication error
        case Auth0FgaApiValidationError:
            // Handle parameter validation error
        case Auth0FgaApiInternalError:
            // Handle API internal error
        case Auth0FgaApiRateLimitError:
            // Exponential backoff in handling rate limit error
        default:
            // Handle unknown/undefined error
        }
    }
    // response from `ReadSettings`: Settings
    fmt.Fprintf(os.Stdout, "Response from `Auth0FgaApi.ReadSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.

### Other Parameters

Other parameters are passed through a pointer to a apiReadSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

### Return type

[**Settings**](Settings.md)

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
    
    body := *openapiclient.NewWriteRequestParams() // WriteRequestParams | 

    configuration := auth0fga.NewConfiguration(UserConfiguration{
        StoreId:      os.Getenv("AUTH0_FGA_STORE_ID"),
        ClientId:     os.Getenv("AUTH0_FGA_CLIENT_ID"),
        ClientSecret: os.Getenv("AUTH0_FGA_CLIENT_SECRET"),
        Environment:  os.Getenv("AUTH0_FGA_ENVIRONMENT"),
    })

    apiClient := auth0fga.NewAPIClient(configuration)

    resp, r, err := apiClient.Auth0FgaApi.Write(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Auth0FgaApi.Write``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
        switch v := err.(type) {
        case Auth0FgaApiAuthenticationError:
            // Handle authentication error
        case Auth0FgaApiValidationError:
            // Handle parameter validation error
        case Auth0FgaApiInternalError:
            // Handle API internal error
        case Auth0FgaApiRateLimitError:
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
**body** | [**WriteRequestParams**](WriteRequestParams.md) |  | 

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

> map[string]interface{} WriteAssertions(ctx, authorizationModelId).Body(body).Execute()

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
    body := *openapiclient.NewWriteAssertionsRequestParams([]openapiclient.Assertion{*openapiclient.NewAssertion(*openapiclient.NewTupleKey(), false)}) // WriteAssertionsRequestParams | 

    configuration := auth0fga.NewConfiguration(UserConfiguration{
        StoreId:      os.Getenv("AUTH0_FGA_STORE_ID"),
        ClientId:     os.Getenv("AUTH0_FGA_CLIENT_ID"),
        ClientSecret: os.Getenv("AUTH0_FGA_CLIENT_SECRET"),
        Environment:  os.Getenv("AUTH0_FGA_ENVIRONMENT"),
    })

    apiClient := auth0fga.NewAPIClient(configuration)

    resp, r, err := apiClient.Auth0FgaApi.WriteAssertions(context.Background(), authorizationModelId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Auth0FgaApi.WriteAssertions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
        switch v := err.(type) {
        case Auth0FgaApiAuthenticationError:
            // Handle authentication error
        case Auth0FgaApiValidationError:
            // Handle parameter validation error
        case Auth0FgaApiInternalError:
            // Handle API internal error
        case Auth0FgaApiRateLimitError:
            // Exponential backoff in handling rate limit error
        default:
            // Handle unknown/undefined error
        }
    }
    // response from `WriteAssertions`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `Auth0FgaApi.WriteAssertions`: %v\n", resp)
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
**body** | [**WriteAssertionsRequestParams**](WriteAssertionsRequestParams.md) |  | 

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


## WriteAuthorizationModel

> WriteAuthorizationModelResponse WriteAuthorizationModel(ctx).Body(body).Execute()

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
    
    body := *openapiclient.NewTypeDefinitions() // TypeDefinitions | 

    configuration := auth0fga.NewConfiguration(UserConfiguration{
        StoreId:      os.Getenv("AUTH0_FGA_STORE_ID"),
        ClientId:     os.Getenv("AUTH0_FGA_CLIENT_ID"),
        ClientSecret: os.Getenv("AUTH0_FGA_CLIENT_SECRET"),
        Environment:  os.Getenv("AUTH0_FGA_ENVIRONMENT"),
    })

    apiClient := auth0fga.NewAPIClient(configuration)

    resp, r, err := apiClient.Auth0FgaApi.WriteAuthorizationModel(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Auth0FgaApi.WriteAuthorizationModel``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
        switch v := err.(type) {
        case Auth0FgaApiAuthenticationError:
            // Handle authentication error
        case Auth0FgaApiValidationError:
            // Handle parameter validation error
        case Auth0FgaApiInternalError:
            // Handle API internal error
        case Auth0FgaApiRateLimitError:
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
**body** | [**TypeDefinitions**](TypeDefinitions.md) |  | 

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


## WriteSettings

> Settings WriteSettings(ctx).Body(body).Execute()

Update the environment tag for a store



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
    
    body := *openapiclient.NewWriteSettingsRequestParams() // WriteSettingsRequestParams | 

    configuration := auth0fga.NewConfiguration(UserConfiguration{
        StoreId:      os.Getenv("AUTH0_FGA_STORE_ID"),
        ClientId:     os.Getenv("AUTH0_FGA_CLIENT_ID"),
        ClientSecret: os.Getenv("AUTH0_FGA_CLIENT_SECRET"),
        Environment:  os.Getenv("AUTH0_FGA_ENVIRONMENT"),
    })

    apiClient := auth0fga.NewAPIClient(configuration)

    resp, r, err := apiClient.Auth0FgaApi.WriteSettings(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Auth0FgaApi.WriteSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
        switch v := err.(type) {
        case Auth0FgaApiAuthenticationError:
            // Handle authentication error
        case Auth0FgaApiValidationError:
            // Handle parameter validation error
        case Auth0FgaApiInternalError:
            // Handle API internal error
        case Auth0FgaApiRateLimitError:
            // Exponential backoff in handling rate limit error
        default:
            // Handle unknown/undefined error
        }
    }
    // response from `WriteSettings`: Settings
    fmt.Fprintf(os.Stdout, "Response from `Auth0FgaApi.WriteSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.

### Other Parameters

Other parameters are passed through a pointer to a apiWriteSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**body** | [**WriteSettingsRequestParams**](WriteSettingsRequestParams.md) |  | 

### Return type

[**Settings**](Settings.md)

### Authorization

[ClientCredentials](../README.md#ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WriteTokenIssuer

> TokenIssuer WriteTokenIssuer(ctx).Body(body).Execute()

Add 3rd party token issuer for Auth0 FGA read and write operations



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
    
    body := *openapiclient.NewWriteTokenIssuersRequestParams() // WriteTokenIssuersRequestParams | 

    configuration := auth0fga.NewConfiguration(UserConfiguration{
        StoreId:      os.Getenv("AUTH0_FGA_STORE_ID"),
        ClientId:     os.Getenv("AUTH0_FGA_CLIENT_ID"),
        ClientSecret: os.Getenv("AUTH0_FGA_CLIENT_SECRET"),
        Environment:  os.Getenv("AUTH0_FGA_ENVIRONMENT"),
    })

    apiClient := auth0fga.NewAPIClient(configuration)

    resp, r, err := apiClient.Auth0FgaApi.WriteTokenIssuer(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Auth0FgaApi.WriteTokenIssuer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
        switch v := err.(type) {
        case Auth0FgaApiAuthenticationError:
            // Handle authentication error
        case Auth0FgaApiValidationError:
            // Handle parameter validation error
        case Auth0FgaApiInternalError:
            // Handle API internal error
        case Auth0FgaApiRateLimitError:
            // Exponential backoff in handling rate limit error
        default:
            // Handle unknown/undefined error
        }
    }
    // response from `WriteTokenIssuer`: TokenIssuer
    fmt.Fprintf(os.Stdout, "Response from `Auth0FgaApi.WriteTokenIssuer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.

### Other Parameters

Other parameters are passed through a pointer to a apiWriteTokenIssuerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**body** | [**WriteTokenIssuersRequestParams**](WriteTokenIssuersRequestParams.md) |  | 

### Return type

[**TokenIssuer**](TokenIssuer.md)

### Authorization

[ClientCredentials](../README.md#ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

