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

> Auth0FgaCheckResponse Check(ctx).Body(body).Execute()

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
    
    body := *auth0fga.NewAuth0FgaCheckRequestParams() // Auth0FgaCheckRequestParams | 

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
    }
    // response from `Check`: Auth0FgaCheckResponse
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
**body** | [**Auth0FgaCheckRequestParams**](Auth0FgaCheckRequestParams.md) |  | 

### Return type

[**Auth0FgaCheckResponse**](Auth0FgaCheckResponse.md)

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

> Auth0FgaExpandResponse Expand(ctx).Body(body).Execute()

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
    
    body := *auth0fga.NewAuth0FgaExpandRequestParams() // Auth0FgaExpandRequestParams | 

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
    }
    // response from `Expand`: Auth0FgaExpandResponse
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
**body** | [**Auth0FgaExpandRequestParams**](Auth0FgaExpandRequestParams.md) |  | 

### Return type

[**Auth0FgaExpandResponse**](Auth0FgaExpandResponse.md)

### Authorization

[ClientCredentials](../README.md#ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Read

> Auth0FgaReadResponse Read(ctx).Body(body).Execute()

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
    
    body := *auth0fga.NewAuth0FgaReadRequestParams() // Auth0FgaReadRequestParams | 

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
    }
    // response from `Read`: Auth0FgaReadResponse
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
**body** | [**Auth0FgaReadRequestParams**](Auth0FgaReadRequestParams.md) |  | 

### Return type

[**Auth0FgaReadResponse**](Auth0FgaReadResponse.md)

### Authorization

[ClientCredentials](../README.md#ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadAssertions

> Auth0FgaReadAssertionsResponse ReadAssertions(ctx, authorizationModelId).Execute()

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
    
    authorizationModelId := "authorizationModelId_example" // string | The authorization model ID

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
    }
    // response from `ReadAssertions`: Auth0FgaReadAssertionsResponse
    fmt.Fprintf(os.Stdout, "Response from `Auth0FgaApi.ReadAssertions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**authorizationModelId** | **string** | The authorization model ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadAssertionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

### Return type

[**Auth0FgaReadAssertionsResponse**](Auth0FgaReadAssertionsResponse.md)

### Authorization

[ClientCredentials](../README.md#ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadAuthorizationModel

> Auth0FgaReadAuthorizationModelResponse ReadAuthorizationModel(ctx, id).Execute()

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
    
    id := "id_example" // string | The authorization model ID

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
    }
    // response from `ReadAuthorizationModel`: Auth0FgaReadAuthorizationModelResponse
    fmt.Fprintf(os.Stdout, "Response from `Auth0FgaApi.ReadAuthorizationModel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The authorization model ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadAuthorizationModelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

### Return type

[**Auth0FgaReadAuthorizationModelResponse**](Auth0FgaReadAuthorizationModelResponse.md)

### Authorization

[ClientCredentials](../README.md#ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadAuthorizationModels

> Auth0FgaReadAuthorizationModelsResponse ReadAuthorizationModels(ctx).PageSize(pageSize).ContinuationToken(continuationToken).Execute()

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
    }
    // response from `ReadAuthorizationModels`: Auth0FgaReadAuthorizationModelsResponse
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

[**Auth0FgaReadAuthorizationModelsResponse**](Auth0FgaReadAuthorizationModelsResponse.md)

### Authorization

[ClientCredentials](../README.md#ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadSettings

> SettingsSettings ReadSettings(ctx).Execute()

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
    }
    // response from `ReadSettings`: SettingsSettings
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

[**SettingsSettings**](SettingsSettings.md)

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
    
    body := *auth0fga.NewAuth0FgaWriteRequestParams() // Auth0FgaWriteRequestParams | 

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
**body** | [**Auth0FgaWriteRequestParams**](Auth0FgaWriteRequestParams.md) |  | 

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
    
    authorizationModelId := "authorizationModelId_example" // string | The authorization model ID
    body := *auth0fga.NewAuth0FgaWriteAssertionsRequestParams([]auth0fga.Auth0FgaAssertion{*auth0fga.NewAuth0FgaAssertion(*auth0fga.NewAuth0FgaTupleKey(), false)}) // Auth0FgaWriteAssertionsRequestParams | 

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
    }
    // response from `WriteAssertions`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `Auth0FgaApi.WriteAssertions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**authorizationModelId** | **string** | The authorization model ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiWriteAssertionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**body** | [**Auth0FgaWriteAssertionsRequestParams**](Auth0FgaWriteAssertionsRequestParams.md) |  | 

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

> Auth0FgaWriteAuthorizationModelResponse WriteAuthorizationModel(ctx).Body(body).Execute()

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
    
    body := *auth0fga.NewAuthorizationmodelTypeDefinitions() // AuthorizationmodelTypeDefinitions | 

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
    }
    // response from `WriteAuthorizationModel`: Auth0FgaWriteAuthorizationModelResponse
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
**body** | [**AuthorizationmodelTypeDefinitions**](AuthorizationmodelTypeDefinitions.md) |  | 

### Return type

[**Auth0FgaWriteAuthorizationModelResponse**](Auth0FgaWriteAuthorizationModelResponse.md)

### Authorization

[ClientCredentials](../README.md#ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WriteSettings

> SettingsSettings WriteSettings(ctx).Body(body).Execute()

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
    
    body := *auth0fga.NewAuth0FgaWriteSettingsRequestParams() // Auth0FgaWriteSettingsRequestParams | 

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
    }
    // response from `WriteSettings`: SettingsSettings
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
**body** | [**Auth0FgaWriteSettingsRequestParams**](Auth0FgaWriteSettingsRequestParams.md) |  | 

### Return type

[**SettingsSettings**](SettingsSettings.md)

### Authorization

[ClientCredentials](../README.md#ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WriteTokenIssuer

> SettingsTokenIssuer WriteTokenIssuer(ctx).Body(body).Execute()

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
    
    body := *auth0fga.NewAuth0FgaWriteTokenIssuersRequestParams() // Auth0FgaWriteTokenIssuersRequestParams | 

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
    }
    // response from `WriteTokenIssuer`: SettingsTokenIssuer
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
**body** | [**Auth0FgaWriteTokenIssuersRequestParams**](Auth0FgaWriteTokenIssuersRequestParams.md) |  | 

### Return type

[**SettingsTokenIssuer**](SettingsTokenIssuer.md)

### Authorization

[ClientCredentials](../README.md#ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

