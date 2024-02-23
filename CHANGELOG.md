# Changelog

## v0.7.0

### [0.7.0](https://github.com/auth0-lab/fga-go-sdk/compare/v0.6.1...v0.7.0) (2024-02-23)

[Breaking]

As of this point this SDK is deprecated and should no longer be used. Please use https://github.com/openfga/go-sdk instead.

We strongly recommend you use the [OpenFGA Go SDK](https://github.com/openfga/go-sdk) directly instead with the following configuration:

For US1 (Production US) environment, use the following values:
- API URL: `https://api.us1.fga.dev`
- Credential Method: ClientCredentials
- API Token Issuer: `fga.us.auth0.com`
- API Audience: `https://api.us1.fga.dev/`

You can get the rest of the necessary variables from the FGA Dashboard. See [here](https://docs.fga.dev/intro/dashboard#create-api-credentials).

```go
package main

import (
    "os"

    . "github.com/openfga/go-sdk/client"
    "github.com/openfga/go-sdk/credentials"
)

func main() {
	fgaClient, err := NewSdkClient(&ClientConfiguration{
        ApiUrl:         "https://api.us1.fga.dev",
        StoreId:        os.Getenv("FGA_STORE_ID"),
        AuthorizationModelId: os.Getenv("FGA_MODEL_ID"),
        Credentials: &credentials.Credentials{
            Method: credentials.CredentialsMethodClientCredentials,
            Config: &credentials.Config{
                ClientCredentialsClientId:       os.Getenv("FGA_CLIENT_ID"),
                ClientCredentialsClientSecret:   os.Getenv("FGA_CLIENT_SECRET"),
                ClientCredentialsApiAudience:    "https://api.us1.fga.dev/",
                ClientCredentialsApiTokenIssuer: "fga.us.auth0.com",
            },
        },
    })

    if err != nil {
        // .. Handle error
    }
}
```

## v0.6.1

### [0.6.1](https://github.com/auth0-lab/fga-go-sdk/compare/v0.6.0...v0.6.1) (2023-03-10)

- fix: retry on 5xx errors
- chore(deps): drop `golang.org/x/net` dependency

## v0.6.0

### [0.6.0](https://github.com/auth0-lab/fga-go-sdk/compare/v0.5.0...v0.6.0) (2022-12-16)

Changes:
- [BREAKING] feat(list-objects)!: response has been changed to include the object type
    e.g. response that was `{"object_ids":["roadmap"]}`, will now be `{"objects":["document:roadmap"]}`

Fixes:
- [BREAKING] fix(models)!: update interfaces that had incorrectly optional fields to make them required

Chore:
- chore(deps): update dependencies

## v0.5.0

### [0.5.0](https://github.com/auth0-lab/fga-go-sdk/compare/v0.4.0...v0.5.0) (2022-10-13)
- BREAKING: exported type `TypeDefinitions` is now `WriteAuthorizationModelRequest`
    Note: This is only a breaking change on the SDK, not the API.
- Support ListObjects
    Support for [ListObjects API](https://docs.fga.dev/api/service#/Relationship%20Queries/ListObjects)

    You call the API and receive the list of object ids from a particular type that the user has a certain relation with.

    For example, to find the list of documents that Anne can read:

    ```golang
    body := auth0fga.ListObjectsRequest{
        AuthorizationModelId: PtrString(""),
        User:                 PtrString("user:anne"),
        Relation:             PtrString("can_view"),
        Type:                 PtrString("document"),
    }
    data, response, err := apiClient.Auth0FgaApi.ListObjects(context.Background()).Body(body).Execute()

    // response.object_ids = ["roadmap"]
    ```
- Use [govulncheck](https://go.dev/blog/vuln) in CI to check for issues
- chore(deps): upgrade dependencies
- Target go 1.19

## v0.4.0

### [0.4.0](https://github.com/auth0-lab/fga-go-sdk/compare/v0.3.0...v0.4.0) (2022-06-07)

#### Changes
- feat!: [ReadAuthorizationModels](https://docs.fga.dev/api/service#/Store%20Models/ReadAuthorizationModels) now returns an array of Authorization Models instead of IDs [BREAKING CHANGE]

    The response will become similar to:
    ```json
    {
        "authorization_models": [
            {
              "id": (string),
              "type_definitions": [...]
            }
        ]
    }
    ```
- feat!: drop support for all settings endpoints [BREAKING CHANGE]
- feat!: Simplify error prefix to `Fga` [BREAKING CHANGE]

    Possible Errors:
    - `FgaApiError`: All errors returned by the API extend this error
    - `FgaApiValidationError`: 400 and 422 Validation Errors returned by the API
    - `FgaApiNotFoundError`: 404 errors returned by the API
    - `FgaApiRateLimitExceededError`: 429 errors returned by the API
    - `FgaApiInternalError`: 5xx errors returned by the API
    - `FgaApiAuthenticationError`: Error during authentication
- feat!: drop `Params` postfix from the name of the request interface [BREAKING CHANGE]

  e.g. `ReadRequestParams` will become `ReadRequest`
- feat: add support for [contextual tuples](https://docs.fga.dev/intro/auth0-fga-concepts#what-are-contextual-tuples) in the [Check](https://docs.fga.dev/api/service#/Tuples/Check) request

    You can call them like so:
    ```go
    body := fgaSdk.CheckRequest{
        TupleKey: &fgaSdk.TupleKey{
            User: fgaSdk.PtrString("anne"),
            Relation: fgaSdk.PtrString("can_view"),
            Object: fgaSdk.PtrString("transaction:A"),
        },
        ContextualTuples: &fgaSdk.ContextualTuples{
            TupleKeys: []fgaSdk.TupleKey{
                {
                    User: fgaSdk.PtrString("anne"),
                    Relation: fgaSdk.PtrString("user"),
                    Object: fgaSdk.PtrString("ip-address-range:10.0.0.0/16"),
                },
                {
                    User: fgaSdk.PtrString("anne"),
                    Relation: fgaSdk.PtrString("user"),
                    Object: fgaSdk.PtrString("timeslot:18_19"),
                }
            }
        }
    }
    data, response, err := fgaClient.Auth0FgaApi.Check(context.Background()).Body(body).Execute()
    ```
- chore: upgrade dependencies
- chore: internal refactor

## v0.3.0

### [0.3.0](https://github.com/auth0-lab/fga-go-sdk/compare/v0.2.2...v0.3.0) (2022-03-17)

#### Changes
- chore!: change method names used to set request body (BREAKING CHANGE)
- chore: fix user agent header format
- feat: add support for the Watch API

## v0.2.2

### [0.2.2](https://github.com/auth0-lab/fga-go-sdk/compare/v0.2.1...v0.2.2) (2022-03-09)

#### Changes
- fix: fix for return types on 204 no content

## v0.2.1

### [0.2.1](https://github.com/auth0-lab/fga-go-sdk/compare/v0.2.0...v0.2.1) (2022-03-07)

#### Changes
- feat(error-handling): expose new api error codes

## v0.2.0

### 0.2.0 (2022-02-09)

#### Changes
- feat: update interfaces for latest api breaking changes
- chore(deps): update dependencies
