# Go SDK for Auth0 Fine Grained Authorization (FGA)

[![FOSSA Status](https://app.fossa.com/api/projects/custom%2B4989%2Fgithub.com%2Fauth0-lab%2Ffga-go-sdk.svg?type=shield)](https://app.fossa.com/projects/custom%2B4989%2Fgithub.com%2Fauth0-lab%2Ffga-go-sdk?ref=badge_shield)

This is an autogenerated Go SDK for Auth0 Fine Grained Authorization (FGA). It provides a wrapper around the [Auth0 Fine Grained Authorization API](https://docs.fga.dev/api/service).

Warning: This SDK comes with no SLAs and is not production-ready!

## Table of Contents

- [About Auth0 Fine Grained Authorization](#about-auth0-fine-grained-authorization)
- [Resources](#resources)
- [Installation](#installation)
- [Getting Started](#getting-started)
  - [Initializing the API Client](#initializing-the-api-client)
  - [Getting your Store ID, Client ID and Client Secret](#getting-your-store-id-client-id-and-client-secret)
  - [Calling the API](#calling-the-api)
    - [Write Authorization Model](#write-authorization-model)
    - [Read a Single Authorization Model](#read-a-single-authorization-model)
    - [Read Authorization Model IDs](#read-authorization-model-ids)
    - [Check](#check)
    - [Write Tuples](#write-tuples)
    - [Delete Tuples](#delete-tuples)
    - [Expand](#expand)
    - [Read](#read)
  - [API Endpoints](#api-endpoints)
  - [Models](#models)
- [Contributing](#contributing)
  - [Issues](#issues)
  - [Pull Requests](#pull-requests) [Note: We are not accepting Pull Requests at this time!]
- [License](#license)

## About Auth0 Fine Grained Authorization

[Auth0 Fine Grained Authorization (FGA)](https://dashboard.fga.dev) is the  early-stage product we are building at Auth0 as part of [Auth0Lab](https://twitter.com/Auth0Lab/) to solve fine-grained authorization at scale.
If you are interested in learning more about our plans, please reach out via our <a target="_blank" href="https://discord.gg/8naAwJfWN6" rel="noreferrer">Discord chat</a>.

Please note:
* At this point in time, Auth0 Fine Grained Authorization does not come with any SLAs; availability and uptime are not guaranteed.
* While this project is in its early stages, the SDK methods are in flux and might change without a major bump

## Resources

- [The Auth0 FGA Playground](https://play.fga.dev)
- [The Auth0 FGA Documentation](https://docs.fga.dev)
- [Zanzibar Academy](https://zanzibar.academy)
- [Auth0Lab on Twitter](https://twitter.com/Auth0Lab/)
- [Discord Community](https://discord.gg/pvbNmqC)

## Installation

To install:

```
go get -u github.com/auth0-lab/fga-go-sdk
```

In your code, import the module and use it:

```go
import "github.com/auth0-lab/fga-go-sdk"

func Main() {
	configuration, err := auth0fga.NewConfiguration(auth0fga.UserConfiguration{})
}
```

You can then run

```shell
go mod tidy
```

to update `go.mod` and `go.sum` if you are using them.

## Getting Started

### Initializing the API Client

```golang

import (
    auth0fga "github.com/auth0-lab/fga-go-sdk"
    "os"
)

func Main() {
    configuration, err := auth0fga.NewConfiguration(auth0fga.UserConfiguration{
        Environment:  os.Getenv("AUTH0_FGA_ENVIRONMENT"),
        StoreId:      os.Getenv("AUTH0_FGA_STORE_ID"),
        ClientId:     os.Getenv("AUTH0_FGA_CLIENT_ID"),
        ClientSecret: os.Getenv("AUTH0_FGA_CLIENT_SECRET"),
    })

    if err != nil {
    // .. Handle error
    }

    apiClient := auth0fga.NewAPIClient(configuration)
}
```

### Getting your Store ID, Client ID and Client Secret

#### Production

Make sure you have created your credentials on the Auth0 FGA Dashboard. [Learn how ➡](https://docs.fga.dev/intro/dashboard#create-api-credentials)

You will need to set the `AUTH0_FGA_ENVIRONMENT` variable to `"us"`. Provide the store id, client id and client secret you have created on the Dashboard.

#### Playground

If you are testing this on the public playground, you need to set your `AUTH0_FGA_ENVIRONMENT` to `"playground"`.

To get your store id, you may copy it from the store you have created on the [Playground](https://play.fga.dev). [Learn how ➡](https://docs.fga.dev/intro/playground#getting-store-id)

In the playground environment, you do not need to provide a client id and client secret.

### Calling the API

#### Write Authorization Model

> Note: To learn how to build your authorization model, check the Docs at https://docs.fga.dev/

> Note: The Auth0 FGA Playground, Dashboard and Documentation use a friendly syntax which gets translated to the API syntax seen below. Learn more about [the Auth0 FGA configuration language](https://docs.fga.dev/modeling/configuration-language).

```golang
body  := auth0fga.TypeDefinitions{TypeDefinitions: &[]auth0fga.TypeDefinition{
	{
		Type: "repo",
		Relations: map[string]auth0fga.Userset{
			"writer": {This: &map[string]interface{}{}},
			"reader": {Union: &auth0fga.Usersets{
				Child: &[]auth0fga.Userset{
					{This: &map[string]interface{}{}},
					{ComputedUserset: &auth0fga.ObjectRelation{
						Object:   auth0fga.PtrString(""),
						Relation: auth0fga.PtrString("writer"),
					}},
				},
			}},
		},
	},
}}
data, response, err := apiClient.Auth0FgaApi.WriteAuthorizationModel(context.Background()).Body(body).Execute()

fmt.Printf("%s", data.AuthorizationModelId) // 1uHxCSuTP0VKPYSnkq1pbb1jeZw
```

#### Read a Single Authorization Model

```golang
// Assuming `1uHxCSuTP0VKPYSnkq1pbb1jeZw` is an id of a single model
data, response, err := apiClient.Auth0FgaApi.ReadAuthorizationModel(context.Background(), "1uHxCSuTP0VKPYSnkq1pbb1jeZw").Execute()

// data = {"authorization_model":{"id":"1uHxCSuTP0VKPYSnkq1pbb1jeZw","type_definitions":[{"type":"repo","relations":{"writer":{"this":{}},"reader":{ ... }}}]}} // JSON

fmt.Printf("%s", data.AuthorizationModel.Id) // 1uHxCSuTP0VKPYSnkq1pbb1jeZw
```

#### Read Authorization Model IDs

```golang
data, response, err := apiClient.Auth0FgaApi.ReadAuthorizationModels(context.Background()).Execute()

// data = {"authorization_model_ids":["1uHxCSuTP0VKPYSnkq1pbb1jeZw","GtQpMohWezFmIbyXxVEocOCxxgq"]} // in JSON

fmt.Printf("%s", (*data.AuthorizationModelIds)[0]) // 1uHxCSuTP0VKPYSnkq1pbb1jeZw
```

#### Check
> Provide a tuple and ask the Auth0 FGA API to check for a relationship

```golang
body := auth0fga.CheckRequestParams{
	TupleKey: &auth0fga.TupleKey{
		User: auth0fga.PtrString("81684243-9356-4421-8fbf-a4f8d36aa31b"),
		Relation: auth0fga.PtrString("admin"),
		Object: auth0fga.PtrString("workspace:675bcac4-ad38-4fb1-a19a-94a5648c91d6"),
	},
}
data, response, err := apiClient.Auth0FgaApi.Check(context.Background()).Body(body).Execute()

// data = {"allowed":true,"resolution":""} // in JSON

fmt.Printf("%t", *data.Allowed) // True

```

#### Write Tuples

```golang
body := auth0fga.WriteRequestParams{
	Writes: &auth0fga.TupleKeys{
		TupleKeys: []auth0fga.TupleKey{
			{
				User: auth0fga.PtrString("81684243-9356-4421-8fbf-a4f8d36aa31b"),
				Relation: auth0fga.PtrString("admin"),
				Object: auth0fga.PtrString("workspace:675bcac4-ad38-4fb1-a19a-94a5648c91d6"),
			},
		},
	},
}
_, response, err := apiClient.Auth0FgaApi.Write(context.Background()).Body(body).Execute()

```

#### Delete Tuples

```golang
body := auth0fga.WriteRequestParams{
	Deletes: &auth0fga.TupleKeys{
		TupleKeys: []auth0fga.TupleKey{
			{
				User: auth0fga.PtrString("81684243-9356-4421-8fbf-a4f8d36aa31b"),
				Relation: auth0fga.PtrString("admin"),
				Object: auth0fga.PtrString("workspace:675bcac4-ad38-4fb1-a19a-94a5648c91d6"),
			},
		},
	},
}
_, response, err := apiClient.Auth0FgaApi.Write(context.Background()).Body(body).Execute()

```

#### Expand

```golang
body := auth0fga.ExpandRequestParams{
	TupleKey: &auth0fga.TupleKey{
		Relation: auth0fga.PtrString("admin"),
		Object: auth0fga.PtrString("workspace:675bcac4-ad38-4fb1-a19a-94a5648c91d6"),
	},
}
data, response, err := apiClient.Auth0FgaApi.Expand(context.Background()).Body(body).Execute()

// data = {"tree":{"root":{"name":"workspace:675bcac4-ad38-4fb1-a19a-94a5648c91d6#admin","leaf":{"users":{"users":["anne","beth"]}}}}} // JSON
```

#### Read

```golang
// Find if a relationship tuple stating that a certain user is an admin on a certain workspace
body := auth0fga.ReadRequestParams{
    TupleKey: &auth0fga.TupleKey{
        User:     auth0fga.PtrString("81684243-9356-4421-8fbf-a4f8d36aa31b"),
        Relation: auth0fga.PtrString("admin"),
        Object:   auth0fga.PtrString("workspace:675bcac4-ad38-4fb1-a19a-94a5648c91d6"),
    },
}

// Find all relationship tuples where a certain user has a relationship as any relation to a certain workspace
body := auth0fga.ReadRequestParams{
    TupleKey: &auth0fga.TupleKey{
        User:     auth0fga.PtrString("81684243-9356-4421-8fbf-a4f8d36aa31b"),
        Object:   auth0fga.PtrString("workspace:675bcac4-ad38-4fb1-a19a-94a5648c91d6"),
    },
}

// Find all relationship tuples where a certain user is an admin on any workspace
body := auth0fga.ReadRequestParams{
    TupleKey: &auth0fga.TupleKey{
        User:     auth0fga.PtrString("81684243-9356-4421-8fbf-a4f8d36aa31b"),
        Relation: auth0fga.PtrString("admin"),
        Object:   auth0fga.PtrString("workspace:"),
    },
}

// Find all relationship tuples where any user has a relationship as any relation with a particular workspace
body := auth0fga.ReadRequestParams{
    TupleKey: &auth0fga.TupleKey{
        Object:   auth0fga.PtrString("workspace:675bcac4-ad38-4fb1-a19a-94a5648c91d6"),
    },
}

data, response, err := apiClient.Auth0FgaApi.Read(context.Background()).Body(body).Execute()

// In all the above situations, the response will be of the form:
// data = {"tuples":[{"key":{"user":"...","relation":"...","object":"..."},"timestamp":"..."}]} // JSON
```


### API Endpoints

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*Auth0FgaApi* | [**Check**](docs/Auth0FgaApi.md#check) | **Post** /stores/{store_id}/check | Check whether a user is authorized to access an object
*Auth0FgaApi* | [**DeleteTokenIssuer**](docs/Auth0FgaApi.md#deletetokenissuer) | **Delete** /stores/{store_id}/settings/token-issuers/{id} | Remove 3rd party token issuer for Auth0 FGA read and write operations
*Auth0FgaApi* | [**Expand**](docs/Auth0FgaApi.md#expand) | **Post** /stores/{store_id}/expand | Expand all relationships in userset tree format, and following userset rewrite rules.  Useful to reason about and debug a certain relationship
*Auth0FgaApi* | [**Read**](docs/Auth0FgaApi.md#read) | **Post** /stores/{store_id}/read | Get tuples from the store that matches a query, without following userset rewrite rules
*Auth0FgaApi* | [**ReadAssertions**](docs/Auth0FgaApi.md#readassertions) | **Get** /stores/{store_id}/assertions/{authorization_model_id} | Read assertions for an authorization model ID
*Auth0FgaApi* | [**ReadAuthorizationModel**](docs/Auth0FgaApi.md#readauthorizationmodel) | **Get** /stores/{store_id}/authorization-models/{id} | Return a particular version of an authorization model
*Auth0FgaApi* | [**ReadAuthorizationModels**](docs/Auth0FgaApi.md#readauthorizationmodels) | **Get** /stores/{store_id}/authorization-models | Return all the authorization model IDs for a particular store
*Auth0FgaApi* | [**ReadSettings**](docs/Auth0FgaApi.md#readsettings) | **Get** /stores/{store_id}/settings | Return store settings, including the environment tag
*Auth0FgaApi* | [**Write**](docs/Auth0FgaApi.md#write) | **Post** /stores/{store_id}/write | Add or delete tuples from the store
*Auth0FgaApi* | [**WriteAssertions**](docs/Auth0FgaApi.md#writeassertions) | **Put** /stores/{store_id}/assertions/{authorization_model_id} | Upsert assertions for an authorization model ID
*Auth0FgaApi* | [**WriteAuthorizationModel**](docs/Auth0FgaApi.md#writeauthorizationmodel) | **Post** /stores/{store_id}/authorization-models | Create a new authorization model
*Auth0FgaApi* | [**WriteSettings**](docs/Auth0FgaApi.md#writesettings) | **Patch** /stores/{store_id}/settings | Update the environment tag for a store
*Auth0FgaApi* | [**WriteTokenIssuer**](docs/Auth0FgaApi.md#writetokenissuer) | **Post** /stores/{store_id}/settings/token-issuers | Add 3rd party token issuer for Auth0 FGA read and write operations


### Models

 - [Any](docs/Any.md)
 - [Assertion](docs/Assertion.md)
 - [AuthErrorCode](docs/AuthErrorCode.md)
 - [AuthenticationErrorMessageResponse](docs/AuthenticationErrorMessageResponse.md)
 - [AuthorizationModel](docs/AuthorizationModel.md)
 - [AuthorizationmodelDifference](docs/AuthorizationmodelDifference.md)
 - [AuthorizationmodelTupleToUserset](docs/AuthorizationmodelTupleToUserset.md)
 - [CheckRequestParams](docs/CheckRequestParams.md)
 - [CheckResponse](docs/CheckResponse.md)
 - [Computed](docs/Computed.md)
 - [Environment](docs/Environment.md)
 - [ErrorCode](docs/ErrorCode.md)
 - [ExpandRequestParams](docs/ExpandRequestParams.md)
 - [ExpandResponse](docs/ExpandResponse.md)
 - [InternalErrorCode](docs/InternalErrorCode.md)
 - [InternalErrorMessageResponse](docs/InternalErrorMessageResponse.md)
 - [Leaf](docs/Leaf.md)
 - [Node](docs/Node.md)
 - [Nodes](docs/Nodes.md)
 - [NotFoundErrorCode](docs/NotFoundErrorCode.md)
 - [ObjectRelation](docs/ObjectRelation.md)
 - [PathUnknownErrorMessageResponse](docs/PathUnknownErrorMessageResponse.md)
 - [ReadAssertionsResponse](docs/ReadAssertionsResponse.md)
 - [ReadAuthorizationModelResponse](docs/ReadAuthorizationModelResponse.md)
 - [ReadAuthorizationModelsResponse](docs/ReadAuthorizationModelsResponse.md)
 - [ReadRequestParams](docs/ReadRequestParams.md)
 - [ReadResponse](docs/ReadResponse.md)
 - [ReadSettingsResponse](docs/ReadSettingsResponse.md)
 - [ResourceExhaustedErrorCode](docs/ResourceExhaustedErrorCode.md)
 - [ResourceExhaustedErrorMessageResponse](docs/ResourceExhaustedErrorMessageResponse.md)
 - [Status](docs/Status.md)
 - [TokenIssuer](docs/TokenIssuer.md)
 - [Tuple](docs/Tuple.md)
 - [TupleKey](docs/TupleKey.md)
 - [TupleKeys](docs/TupleKeys.md)
 - [TypeDefinition](docs/TypeDefinition.md)
 - [TypeDefinitions](docs/TypeDefinitions.md)
 - [Users](docs/Users.md)
 - [Userset](docs/Userset.md)
 - [UsersetTree](docs/UsersetTree.md)
 - [UsersetTreeDifference](docs/UsersetTreeDifference.md)
 - [UsersetTreeTupleToUserset](docs/UsersetTreeTupleToUserset.md)
 - [Usersets](docs/Usersets.md)
 - [ValidationErrorMessageResponse](docs/ValidationErrorMessageResponse.md)
 - [WriteAssertionsRequestParams](docs/WriteAssertionsRequestParams.md)
 - [WriteAuthorizationModelResponse](docs/WriteAuthorizationModelResponse.md)
 - [WriteRequestParams](docs/WriteRequestParams.md)
 - [WriteSettingsRequestParams](docs/WriteSettingsRequestParams.md)
 - [WriteSettingsResponse](docs/WriteSettingsResponse.md)
 - [WriteTokenIssuersRequestParams](docs/WriteTokenIssuersRequestParams.md)
 - [WriteTokenIssuersResponse](docs/WriteTokenIssuersResponse.md)


## Contributing

### Issue Reporting

If you have found a bug or if you have a feature request, please report them at this repository [issues](https://github.com/auth0-lab/fga-go-sdk/issues) section. Please do not report security vulnerabilities on the public GitHub issue tracker. The [Responsible Disclosure Program](https://auth0.com/responsible-disclosure-policy) details the procedure for disclosing security issues.

For auth0 related questions/support please use the [Support Center](https://support.auth0.com).

### Pull Requests

Pull Requests are not currently open, please [raise an issue](https://github.com/auth0-lab/fga-go-sdk/issues) or contact a team member on https://discord.gg/8naAwJfWN6 if there is a feature you'd like us to implement.

## Author

[Auth0Lab](https://github.com/auth0-lab)

## License

This project is licensed under the MIT license. See the [LICENSE](https://github.com/auth0-lab/fga-go-sdk/blob/main/LICENSE) file for more info.

The code in this repo was auto generated by [OpenAPI Generator](https://github.com/OpenAPITools/openapi-generator) from a template based on the [go template](https://github.com/OpenAPITools/openapi-generator/tree/master/modules/openapi-generator/src/main/resources/go), licensed under the [Apache License 2.0](https://github.com/OpenAPITools/openapi-generator/blob/master/LICENSE).

This repo bundles some code from the [golang.org/x/oauth2](https://pkg.go.dev/golang.org/x/oauth2) package. You can find the code [here](./oauth2) and corresponding [BSD-3 License](./oauth2/LICENSE).
