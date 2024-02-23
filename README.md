# Go SDK for Auth0 Fine Grained Authorization (FGA)

[![Go Reference](https://pkg.go.dev/badge/github.com/auth0-lab/fga-go-sdk.svg)](https://pkg.go.dev/github.com/auth0-lab/fga-go-sdk)
[![Release](https://img.shields.io/github/v/release/auth0-lab/fga-go-sdk?sort=semver&color=green)](https://github.com/auth0-lab/fga-go-sdk/releases)
[![License](https://img.shields.io/badge/License-MIT-blue.svg)](./LICENSE)

This is an autogenerated Go SDK for Auth0 Fine Grained Authorization (FGA). It provides a wrapper around the [Auth0 Fine Grained Authorization API](https://docs.fga.dev/api/service).

**This SDK is considered deprecated**.

## Table of Contents

- [About Auth0 Fine Grained Authorization (FGA)](#about)
- [Resources](#resources)
- [Contributing](#contributing)
- [License](#license)

## About

[Okta Fine Grained Authorization (FGA)](https://fga.dev) is designed to make it easy for application builders to model their permission layer, and to add and integrate fine-grained authorization into their applications. Okta Fine Grained Authorization (FGA)’s design is optimized for reliability and low latency at a high scale.

**DEPRECATION WARNING**: This project is no longer maintained. We recommend using the [OpenFGA Go SDK](https://github.com/openfga/go-sdk) with the following configuration instead of this SDK:

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

For US1 (Production US) environment, use the following values:
- API URL: `https://api.us1.fga.dev`
- Credential Method: ClientCredentials
- API Token Issuer: `fga.us.auth0.com`
- API Audience: `https://api.us1.fga.dev/`

You can get the rest of the necessary variables from the FGA Dashboard. See [here](https://docs.fga.dev/intro/dashboard#create-api-credentials).

## Resources

- [Okta Fine Grained Authorization (FGA) Documentation](https://docs.fga.dev)
- [Okta Fine Grained Authorization (FGA) API Documentation](https://docs.fga.dev/api/service)
- [Zanzibar Academy](https://zanzibar.academy)
- [Google's Zanzibar Paper (2019)](https://research.google/pubs/pub48190/)


## Contributing

This repo is deprecated and no longer accepting contributions.

## Author

[Okta FGA](https://github.com/auth0-lab)

## License

This project is licensed under the MIT license. See the [LICENSE](https://github.com/auth0-lab/fga-go-sdk/blob/main/LICENSE) file for more info.

The code in this repo was auto generated by [OpenAPI Generator](https://github.com/OpenAPITools/openapi-generator) from a template based on the [go template](https://github.com/OpenAPITools/openapi-generator/tree/master/modules/openapi-generator/src/main/resources/go), licensed under the [Apache License 2.0](https://github.com/OpenAPITools/openapi-generator/blob/master/LICENSE).

This repo bundles some code from the [golang.org/x/oauth2](https://pkg.go.dev/golang.org/x/oauth2) package. You can find the code [here](./oauth2) and corresponding [BSD-3 License](./oauth2/LICENSE).

