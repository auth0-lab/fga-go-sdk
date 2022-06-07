# Changelog

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
