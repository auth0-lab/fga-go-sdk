/**
 * Go SDK for Auth0 Fine Grained Authorization (FGA)
 *
 * API version: 0.1
 * Website: <https://fga.dev>
 * Documentation: <https://docs.fga.dev>
 * Support: <https://discord.gg/8naAwJfWN6>
 * License: [MIT](https://github.com/auth0-lab/fga-go-sdk/blob/main/LICENSE)
 *
 * NOTE: This file was auto generated by [OpenAPI Generator](https://openapi-generator.tech). DO NOT EDIT.
 */

package auth0fga

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/jarcoal/httpmock"
	"net/http"
	"testing"
)

type TestDefinition struct {
	Name           string
	JsonResponse   string
	ResponseStatus int
	Method         string
	RequestPath    string
}

func TestAuth0FgaApiConfiguration(t *testing.T) {
	t.Run("Providing no store id should error", func(t *testing.T) {
		_, err := NewConfiguration(UserConfiguration{})

		if err == nil {
			t.Errorf("Expected an error")
			return
		}
	})

	t.Run("Providing no environment should not error", func(t *testing.T) {
		_, err := NewConfiguration(UserConfiguration{
			StoreId:      "6c181474-aaa1-4df7-8929-6e7b3a992754",
			ClientId:     "some-id",
			ClientSecret: "some-secret",
		})

		if err != nil {
			t.Errorf("%v", err)
			return
		}
	})

	t.Run("Providing invalid environment should error", func(t *testing.T) {
		_, err := NewConfiguration(UserConfiguration{
			StoreId:     "6c181474-aaa1-4df7-8929-6e7b3a992754",
			Environment: "does-not-exist",
		})

		if err == nil {
			t.Errorf("Expected an error")
			return
		}
	})

	t.Run("Providing no client id and secret when they are required should error", func(t *testing.T) {
		_, err := NewConfiguration(UserConfiguration{
			StoreId:     "6c181474-aaa1-4df7-8929-6e7b3a992754",
			Environment: "staging",
		})

		if err == nil {
			t.Errorf("Expected an error: client id and seceret required")
			return
		}

		_, err = NewConfiguration(UserConfiguration{
			StoreId:     "6c181474-aaa1-4df7-8929-6e7b3a992754",
			Environment: "staging",
			ClientId:    "some-id",
		})

		if err == nil {
			t.Errorf("Expected an error: client secret required")
			return
		}

		_, err = NewConfiguration(UserConfiguration{
			StoreId:      "6c181474-aaa1-4df7-8929-6e7b3a992754",
			Environment:  "staging",
			ClientSecret: "some-secret",
		})

		if err == nil {
			t.Errorf("Expected an error: client id required")
			return
		}

		_, err = NewConfiguration(UserConfiguration{
			StoreId:      "6c181474-aaa1-4df7-8929-6e7b3a992754",
			Environment:  "staging",
			ClientId:     "some-id",
			ClientSecret: "some-secret",
		})

		if err != nil {
			t.Errorf("Unexpected error: %v", err)
			return
		}
	})

	t.Run("Providing no client id and secret when they are not required should not error", func(t *testing.T) {
		_, err := NewConfiguration(UserConfiguration{
			StoreId:     "6c181474-aaa1-4df7-8929-6e7b3a992754",
			Environment: "playground",
		})

		if err != nil {
			t.Errorf("%v", err)
			return
		}
	})

	t.Run("should issue a network call to get the token at the first request if client id is provided", func(t *testing.T) {
		configuration, err := NewConfiguration(UserConfiguration{
			StoreId:      "6c181474-aaa1-4df7-8929-6e7b3a992754",
			Environment:  "staging",
			ClientId:     "some-id",
			ClientSecret: "some-secret",
		})
		if err != nil {
			t.Errorf("%v", err)
			return
		}

		apiClient := NewAPIClient(configuration)

		httpmock.Activate()
		defer httpmock.DeactivateAndReset()
		httpmock.RegisterResponder("GET", fmt.Sprintf("%s://%s/stores/%s/authorization-models", configuration.ApiScheme, configuration.ApiHost, configuration.StoreId),
			func(req *http.Request) (*http.Response, error) {
				resp, err := httpmock.NewJsonResponse(200, ReadAuthorizationModelsResponse{AuthorizationModels: &[]AuthorizationModel{
					{
						Id:              PtrString("1uHxCSuTP0VKPYSnkq1pbb1jeZw"),
						TypeDefinitions: &[]TypeDefinition{},
					},
					{
						Id:              PtrString("GtQpMohWezFmIbyXxVEocOCxxgq"),
						TypeDefinitions: &[]TypeDefinition{},
					},
				}})
				if err != nil {
					return httpmock.NewStringResponse(500, ""), nil
				}
				return resp, nil
			},
		)

		httpmock.RegisterResponder("POST", fmt.Sprintf("https://%s/oauth/token", configuration.ApiTokenIssuer),
			func(req *http.Request) (*http.Response, error) {
				resp, err := httpmock.NewJsonResponse(200, struct {
					AccessToken string `json:"access_token"`
				}{AccessToken: "abcde"})
				if err != nil {
					return httpmock.NewStringResponse(500, ""), nil
				}
				return resp, nil
			},
		)

		if _, _, err = apiClient.Auth0FgaApi.ReadAuthorizationModels(context.Background()).Execute(); err != nil {
			t.Errorf("%v", err)
			return
		}

		info := httpmock.GetCallCountInfo()
		numCalls := info[fmt.Sprintf("POST https://%s/oauth/token", configuration.ApiTokenIssuer)]
		if numCalls != 1 {
			t.Errorf("Expected call to get access token to be made exactly once, saw: %d", numCalls)
			return
		}
		numCalls = info[fmt.Sprintf("GET %s://%s/stores/%s/authorization-models", configuration.ApiScheme, configuration.ApiHost, configuration.StoreId)]
		if numCalls != 1 {
			t.Errorf("Expected call to get authorization models to be made exactly once, saw: %d", numCalls)
			return
		}
	})

	t.Run("should not issue a network call to get the token at the first request if the clientId is not provided", func(t *testing.T) {
		configuration, err := NewConfiguration(UserConfiguration{
			Environment: "playground",
			StoreId:     "6c181474-aaa1-4df7-8929-6e7b3a992754",
		})
		if err != nil {
			t.Errorf("%v", err)
			return
		}
		configuration.ApiHost = "api.fga.example"

		apiClient := NewAPIClient(configuration)

		httpmock.Activate()
		defer httpmock.DeactivateAndReset()
		httpmock.RegisterResponder("GET", fmt.Sprintf("%s://%s/stores/%s/authorization-models", configuration.ApiScheme, configuration.ApiHost, configuration.StoreId),
			func(req *http.Request) (*http.Response, error) {
				resp, err := httpmock.NewJsonResponse(200, ReadAuthorizationModelsResponse{AuthorizationModels: &[]AuthorizationModel{
					{
						Id:              PtrString("1uHxCSuTP0VKPYSnkq1pbb1jeZw"),
						TypeDefinitions: &[]TypeDefinition{},
					},
					{
						Id:              PtrString("GtQpMohWezFmIbyXxVEocOCxxgq"),
						TypeDefinitions: &[]TypeDefinition{},
					},
				}})
				if err != nil {
					return httpmock.NewStringResponse(500, ""), nil
				}
				return resp, nil
			},
		)
		if _, _, err = apiClient.Auth0FgaApi.ReadAuthorizationModels(context.Background()).Execute(); err != nil {
			t.Errorf("%v", err)
			return
		}

		info := httpmock.GetCallCountInfo()
		numCalls := info[fmt.Sprintf("POST https://%s/oauth/token", configuration.ApiTokenIssuer)]
		if numCalls != 0 {
			t.Errorf("Unexpected call to get access token made. Expected 0, saw: %d", numCalls)
			return
		}
		numCalls = info[fmt.Sprintf("GET %s://%s/stores/%s/authorization-models", configuration.ApiScheme, configuration.ApiHost, configuration.StoreId)]
		if numCalls != 1 {
			t.Errorf("Expected call to get authorization models to be made exactly once, saw: %d", numCalls)
			return
		}
	})
}

func TestAuth0FgaApi(t *testing.T) {
	configuration, err := NewConfiguration(UserConfiguration{
		Environment: "playground",
		StoreId:     "6c181474-aaa1-4df7-8929-6e7b3a992754",
	})
	if err != nil {
		t.Fatalf("%v", err)
	}

	apiClient := NewAPIClient(configuration)

	t.Run("ReadAuthorizationModels", func(t *testing.T) {
		test := TestDefinition{
			Name:           "ReadAuthorizationModels",
			JsonResponse:   `{"authorization_models":[{"id":"1uHxCSuTP0VKPYSnkq1pbb1jeZw","type_definitions":[]}]}`,
			ResponseStatus: 200,
			Method:         "GET",
			RequestPath:    "authorization-models",
		}

		var expectedResponse ReadAuthorizationModelsResponse
		if err := json.Unmarshal([]byte(test.JsonResponse), &expectedResponse); err != nil {
			t.Fatalf("%v", err)
		}

		httpmock.Activate()
		defer httpmock.DeactivateAndReset()
		httpmock.RegisterResponder(test.Method, fmt.Sprintf("%s://%s/stores/%s/%s", configuration.ApiScheme, configuration.ApiHost, configuration.StoreId, test.RequestPath),
			func(req *http.Request) (*http.Response, error) {
				resp, err := httpmock.NewJsonResponse(test.ResponseStatus, expectedResponse)
				if err != nil {
					return httpmock.NewStringResponse(500, ""), nil
				}
				return resp, nil
			},
		)

		got, response, err := apiClient.Auth0FgaApi.ReadAuthorizationModels(context.Background()).Execute()
		if err != nil {
			t.Fatalf("%v", err)
		}

		if response.StatusCode != test.ResponseStatus {
			t.Fatalf("Auth0Fga%v().Execute() = %v, want %v", test.Name, response.StatusCode, test.ResponseStatus)
		}

		if len(*got.AuthorizationModels) != 1 {
			t.Fatalf("%v", err)
		}

		if *((*got.AuthorizationModels)[0].Id) != *((*expectedResponse.AuthorizationModels)[0].Id) {
			t.Fatalf("Auth0Fga%v().Execute() = %v, want %v", test.Name, *((*got.AuthorizationModels)[0].Id), *((*expectedResponse.AuthorizationModels)[0].Id))
		}
	})

	t.Run("WriteAuthorizationModel", func(t *testing.T) {
		test := TestDefinition{
			Name:           "WriteAuthorizationModel",
			JsonResponse:   `{"authorization_model_id":"1uHxCSuTP0VKPYSnkq1pbb1jeZw"}`,
			ResponseStatus: 200,
			Method:         "POST",
			RequestPath:    "authorization-models",
		}
		requestBody := WriteAuthorizationModelRequest{
			TypeDefinitions: []TypeDefinition{{
				Type: "github-repo",
				Relations: &map[string]Userset{
					"repo_writer": {
						This: &map[string]interface{}{},
					},
					"viewer": {Union: &Usersets{
						Child: &[]Userset{
							{This: &map[string]interface{}{}},
							{ComputedUserset: &ObjectRelation{
								Object:   PtrString(""),
								Relation: PtrString("repo_writer"),
							}},
						},
					}},
				},
			}},
		}

		var expectedResponse WriteAuthorizationModelResponse
		if err := json.Unmarshal([]byte(test.JsonResponse), &expectedResponse); err != nil {
			t.Fatalf("%v", err)
		}

		httpmock.Activate()
		defer httpmock.DeactivateAndReset()
		httpmock.RegisterResponder(test.Method, fmt.Sprintf("%s://%s/stores/%s/%s", configuration.ApiScheme, configuration.ApiHost, configuration.StoreId, test.RequestPath),
			func(req *http.Request) (*http.Response, error) {
				resp, err := httpmock.NewJsonResponse(test.ResponseStatus, expectedResponse)
				if err != nil {
					return httpmock.NewStringResponse(500, ""), nil
				}
				return resp, nil
			},
		)
		got, response, err := apiClient.Auth0FgaApi.WriteAuthorizationModel(context.Background()).Body(requestBody).Execute()
		if err != nil {
			t.Fatalf("%v", err)
		}

		if response.StatusCode != test.ResponseStatus {
			t.Fatalf("Auth0Fga%v().Execute() = %v, want %v", test.Name, response.StatusCode, test.ResponseStatus)
		}

		_, err = got.MarshalJSON()
		if err != nil {
			t.Fatalf("%v", err)
		}
	})

	t.Run("ReadAuthorizationModel", func(t *testing.T) {
		test := TestDefinition{
			Name:           "ReadAuthorizationModel",
			JsonResponse:   `{"authorization_model":{"id":"1uHxCSuTP0VKPYSnkq1pbb1jeZw", "type_definitions":[{"type":"github-repo", "relations":{"viewer":{"this":{}}}}]}}`,
			ResponseStatus: 200,
			Method:         "GET",
			RequestPath:    "authorization-models",
		}

		var expectedResponse ReadAuthorizationModelResponse
		if err := json.Unmarshal([]byte(test.JsonResponse), &expectedResponse); err != nil {
			t.Fatalf("%v", err)
		}
		modelId := *(*expectedResponse.AuthorizationModel).Id

		httpmock.Activate()
		defer httpmock.DeactivateAndReset()
		httpmock.RegisterResponder(test.Method, fmt.Sprintf("%s://%s/stores/%s/%s/%s", configuration.ApiScheme, configuration.ApiHost, configuration.StoreId, test.RequestPath, modelId),
			func(req *http.Request) (*http.Response, error) {
				resp, err := httpmock.NewJsonResponse(test.ResponseStatus, expectedResponse)
				if err != nil {
					return httpmock.NewStringResponse(500, ""), nil
				}
				return resp, nil
			},
		)
		got, response, err := apiClient.Auth0FgaApi.ReadAuthorizationModel(context.Background(), modelId).Execute()
		if err != nil {
			t.Fatalf("%v", err)
		}

		if response.StatusCode != test.ResponseStatus {
			t.Fatalf("Auth0Fga%v().Execute() = %v, want %v", test.Name, response.StatusCode, test.ResponseStatus)
		}

		responseJson, err := got.MarshalJSON()
		if err != nil {
			t.Fatalf("%v", err)
		}

		if *(*got.AuthorizationModel).Id != modelId {
			t.Fatalf("Auth0Fga%v().Execute() = %v, want %v", test.Name, string(responseJson), test.JsonResponse)
		}
	})

	t.Run("Check", func(t *testing.T) {
		test := TestDefinition{
			Name:           "Check",
			JsonResponse:   `{"allowed":true, "resolution":""}`,
			ResponseStatus: 200,
			Method:         "POST",
			RequestPath:    "check",
		}
		requestBody := CheckRequest{
			TupleKey: TupleKey{
				User:     PtrString("user:81684243-9356-4421-8fbf-a4f8d36aa31b"),
				Relation: PtrString("viewer"),
				Object:   PtrString("document:roadmap"),
			},
		}

		var expectedResponse CheckResponse
		if err := json.Unmarshal([]byte(test.JsonResponse), &expectedResponse); err != nil {
			t.Fatalf("%v", err)
		}

		httpmock.Activate()
		defer httpmock.DeactivateAndReset()
		httpmock.RegisterResponder(test.Method, fmt.Sprintf("%s://%s/stores/%s/%s", configuration.ApiScheme, configuration.ApiHost, configuration.StoreId, test.RequestPath),
			func(req *http.Request) (*http.Response, error) {
				resp, err := httpmock.NewJsonResponse(test.ResponseStatus, expectedResponse)
				if err != nil {
					return httpmock.NewStringResponse(500, ""), nil
				}
				return resp, nil
			},
		)
		got, response, err := apiClient.Auth0FgaApi.Check(context.Background()).Body(requestBody).Execute()
		if err != nil {
			t.Fatalf("%v", err)
		}

		if response.StatusCode != test.ResponseStatus {
			t.Fatalf("Auth0Fga%v().Execute() = %v, want %v", test.Name, response.StatusCode, test.ResponseStatus)
		}

		responseJson, err := got.MarshalJSON()
		if err != nil {
			t.Fatalf("%v", err)
		}

		if *got.Allowed != *expectedResponse.Allowed {
			t.Fatalf("Auth0Fga%v().Execute() = %v, want %v", test.Name, string(responseJson), test.JsonResponse)
		}
	})

	t.Run("Write (Write Tuple)", func(t *testing.T) {
		test := TestDefinition{
			Name:           "Write",
			JsonResponse:   `{}`,
			ResponseStatus: 200,
			Method:         "POST",
			RequestPath:    "write",
		}
		requestBody := WriteRequest{
			Writes: &TupleKeys{
				TupleKeys: []TupleKey{{
					User:     PtrString("user:81684243-9356-4421-8fbf-a4f8d36aa31b"),
					Relation: PtrString("viewer"),
					Object:   PtrString("document:roadmap"),
				}},
			},
		}

		var expectedResponse map[string]interface{}
		if err := json.Unmarshal([]byte(test.JsonResponse), &expectedResponse); err != nil {
			t.Fatalf("%v", err)
		}

		httpmock.Activate()
		defer httpmock.DeactivateAndReset()
		httpmock.RegisterResponder(test.Method, fmt.Sprintf("%s://%s/stores/%s/%s", configuration.ApiScheme, configuration.ApiHost, configuration.StoreId, test.RequestPath),
			func(req *http.Request) (*http.Response, error) {
				resp, err := httpmock.NewJsonResponse(test.ResponseStatus, expectedResponse)
				if err != nil {
					return httpmock.NewStringResponse(500, ""), nil
				}
				return resp, nil
			},
		)
		_, response, err := apiClient.Auth0FgaApi.Write(context.Background()).Body(requestBody).Execute()
		if err != nil {
			t.Fatalf("%v", err)
		}

		if response.StatusCode != test.ResponseStatus {
			t.Fatalf("Auth0Fga%v().Execute() = %v, want %v", test.Name, response.StatusCode, test.ResponseStatus)
		}
	})

	t.Run("Write (Delete Tuple)", func(t *testing.T) {
		test := TestDefinition{
			Name:           "Write",
			JsonResponse:   `{}`,
			ResponseStatus: 200,
			Method:         "POST",
			RequestPath:    "write",
		}

		requestBody := WriteRequest{
			Deletes: &TupleKeys{
				TupleKeys: []TupleKey{{
					User:     PtrString("user:81684243-9356-4421-8fbf-a4f8d36aa31b"),
					Relation: PtrString("viewer"),
					Object:   PtrString("document:roadmap"),
				}},
			},
		}

		var expectedResponse map[string]interface{}
		if err := json.Unmarshal([]byte(test.JsonResponse), &expectedResponse); err != nil {
			t.Fatalf("%v", err)
		}

		httpmock.Activate()
		defer httpmock.DeactivateAndReset()
		httpmock.RegisterResponder(test.Method, fmt.Sprintf("%s://%s/stores/%s/%s", configuration.ApiScheme, configuration.ApiHost, configuration.StoreId, test.RequestPath),
			func(req *http.Request) (*http.Response, error) {
				resp, err := httpmock.NewJsonResponse(test.ResponseStatus, expectedResponse)
				if err != nil {
					return httpmock.NewStringResponse(500, ""), nil
				}
				return resp, nil
			},
		)
		_, response, err := apiClient.Auth0FgaApi.Write(context.Background()).Body(requestBody).Execute()
		if err != nil {
			t.Fatalf("%v", err)
		}

		if response.StatusCode != test.ResponseStatus {
			t.Fatalf("Auth0Fga%v().Execute() = %v, want %v", test.Name, response.StatusCode, test.ResponseStatus)
		}
	})

	t.Run("Expand", func(t *testing.T) {
		test := TestDefinition{
			Name:           "Expand",
			JsonResponse:   `{"tree":{"root":{"name":"document:roadmap#viewer","union":{"nodes":[{"name": "document:roadmap#viewer","leaf":{"users":{"users":["user:81684243-9356-4421-8fbf-a4f8d36aa31b"]}}}]}}}}`,
			ResponseStatus: 200,
			Method:         "POST",
			RequestPath:    "expand",
		}

		requestBody := ExpandRequest{
			TupleKey: TupleKey{
				Relation: PtrString("viewer"),
				Object:   PtrString("document:roadmap"),
			},
		}

		var expectedResponse ExpandResponse
		if err := json.Unmarshal([]byte(test.JsonResponse), &expectedResponse); err != nil {
			t.Fatalf("%v", err)
		}

		httpmock.Activate()
		defer httpmock.DeactivateAndReset()
		httpmock.RegisterResponder(test.Method, fmt.Sprintf("%s://%s/stores/%s/%s", configuration.ApiScheme, configuration.ApiHost, configuration.StoreId, test.RequestPath),
			func(req *http.Request) (*http.Response, error) {
				resp, err := httpmock.NewJsonResponse(test.ResponseStatus, expectedResponse)
				if err != nil {
					return httpmock.NewStringResponse(500, ""), nil
				}
				return resp, nil
			},
		)
		got, response, err := apiClient.Auth0FgaApi.Expand(context.Background()).Body(requestBody).Execute()
		if err != nil {
			t.Fatalf("%v", err)
		}

		if response.StatusCode != test.ResponseStatus {
			t.Fatalf("Auth0Fga%v().Execute() = %v, want %v", test.Name, response.StatusCode, test.ResponseStatus)
		}

		_, err = got.MarshalJSON()
		if err != nil {
			t.Fatalf("%v", err)
		}
	})

	t.Run("Read", func(t *testing.T) {
		test := TestDefinition{
			Name:           "Read",
			JsonResponse:   `{"tuples":[{"key":{"user":"user:81684243-9356-4421-8fbf-a4f8d36aa31b","relation":"viewer","object":"document:roadmap"},"timestamp": "2000-01-01T00:00:00Z"}]}`,
			ResponseStatus: 200,
			Method:         "POST",
			RequestPath:    "read",
		}

		requestBody := ReadRequest{
			TupleKey: &TupleKey{
				User:     PtrString("user:81684243-9356-4421-8fbf-a4f8d36aa31b"),
				Relation: PtrString("viewer"),
				Object:   PtrString("document:roadmap"),
			},
		}

		var expectedResponse ReadResponse
		if err := json.Unmarshal([]byte(test.JsonResponse), &expectedResponse); err != nil {
			t.Fatalf("%v", err)
		}

		httpmock.Activate()
		defer httpmock.DeactivateAndReset()
		httpmock.RegisterResponder(test.Method, fmt.Sprintf("%s://%s/stores/%s/%s", configuration.ApiScheme, configuration.ApiHost, configuration.StoreId, test.RequestPath),
			func(req *http.Request) (*http.Response, error) {
				resp, err := httpmock.NewJsonResponse(test.ResponseStatus, expectedResponse)
				if err != nil {
					return httpmock.NewStringResponse(500, ""), nil
				}
				return resp, nil
			},
		)
		got, response, err := apiClient.Auth0FgaApi.Read(context.Background()).Body(requestBody).Execute()
		if err != nil {
			t.Fatalf("%v", err)
		}

		if response.StatusCode != test.ResponseStatus {
			t.Fatalf("Auth0Fga%v().Execute() = %v, want %v", test.Name, response.StatusCode, test.ResponseStatus)
		}

		responseJson, err := got.MarshalJSON()
		if err != nil {
			t.Fatalf("%v", err)
		}

		if len(*got.Tuples) != len(*expectedResponse.Tuples) {
			t.Fatalf("Auth0Fga%v().Execute() = %v, want %v", test.Name, string(responseJson), test.JsonResponse)
		}
	})

	t.Run("ReadChanges", func(t *testing.T) {
		test := TestDefinition{
			Name:           "ReadChanges",
			JsonResponse:   `{"changes":[{"tuple_key":{"user":"user:81684243-9356-4421-8fbf-a4f8d36aa31b","relation":"viewer","object":"document:roadmap"},"operation":"TUPLE_OPERATION_WRITE","timestamp": "2000-01-01T00:00:00Z"}],"continuation_token":"eyJwayI6IkxBVEVTVF9OU0NPTkZJR19hdXRoMHN0b3JlIiwic2siOiIxem1qbXF3MWZLZExTcUoyN01MdTdqTjh0cWgifQ=="}`,
			ResponseStatus: 200,
			Method:         "GET",
			RequestPath:    "changes",
		}

		var expectedResponse ReadChangesResponse
		if err := json.Unmarshal([]byte(test.JsonResponse), &expectedResponse); err != nil {
			t.Fatalf("%v", err)
		}

		httpmock.Activate()
		defer httpmock.DeactivateAndReset()
		httpmock.RegisterResponder(test.Method, fmt.Sprintf("%s://%s/stores/%s/%s", configuration.ApiScheme, configuration.ApiHost, configuration.StoreId, test.RequestPath),
			func(req *http.Request) (*http.Response, error) {
				resp, err := httpmock.NewJsonResponse(test.ResponseStatus, expectedResponse)
				if err != nil {
					return httpmock.NewStringResponse(500, ""), nil
				}
				return resp, nil
			},
		)
		got, response, err := apiClient.Auth0FgaApi.ReadChanges(context.Background()).
			Type_("repo").
			PageSize(25).
			ContinuationToken("eyJwayI6IkxBVEVTVF9OU0NPTkZJR19hdXRoMHN0b3JlIiwic2siOiIxem1qbXF3MWZLZExTcUoyN01MdTdqTjh0cWgifQ==").
			Execute()
		if err != nil {
			t.Fatalf("%v", err)
		}

		if response.StatusCode != test.ResponseStatus {
			t.Fatalf("Auth0Fga%v().Execute() = %v, want %v", test.Name, response.StatusCode, test.ResponseStatus)
		}

		responseJson, err := got.MarshalJSON()
		if err != nil {
			t.Fatalf("%v", err)
		}

		if len(*got.Changes) != len(*expectedResponse.Changes) {
			t.Fatalf("Auth0Fga%v().Execute() = %v, want %v", test.Name, string(responseJson), test.JsonResponse)
		}
	})

	t.Run("ListObjects", func(t *testing.T) {
		test := TestDefinition{
			Name:           "ListObjects",
			JsonResponse:   `{"objects":["document:roadmap"]}`,
			ResponseStatus: 200,
			Method:         "POST",
			RequestPath:    "list-objects",
		}

		requestBody := ListObjectsRequest{
			AuthorizationModelId: PtrString("01GAHCE4YVKPQEKZQHT2R89MQV"),
			User:                 "user:81684243-9356-4421-8fbf-a4f8d36aa31b",
			Relation:             "can_read",
			Type:                 "document",
			ContextualTuples: &ContextualTupleKeys{
				TupleKeys: []TupleKey{{
					User:     PtrString("user:81684243-9356-4421-8fbf-a4f8d36aa31b"),
					Relation: PtrString("editor"),
					Object:   PtrString("folder:product"),
				}, {
					User:     PtrString("folder:product"),
					Relation: PtrString("parent"),
					Object:   PtrString("document:roadmap"),
				}},
			},
		}

		var expectedResponse ListObjectsResponse
		if err := json.Unmarshal([]byte(test.JsonResponse), &expectedResponse); err != nil {
			t.Fatalf("%v", err)
		}

		httpmock.Activate()
		defer httpmock.DeactivateAndReset()
		httpmock.RegisterResponder(test.Method, fmt.Sprintf("%s://%s/stores/%s/%s", configuration.ApiScheme, configuration.ApiHost, configuration.StoreId, test.RequestPath),
			func(req *http.Request) (*http.Response, error) {
				resp, err := httpmock.NewJsonResponse(test.ResponseStatus, expectedResponse)
				if err != nil {
					return httpmock.NewStringResponse(500, ""), nil
				}
				return resp, nil
			},
		)
		got, response, err := apiClient.Auth0FgaApi.ListObjects(context.Background()).
			Body(requestBody).
			Execute()
		if err != nil {
			t.Fatalf("%v", err)
		}

		if response.StatusCode != test.ResponseStatus {
			t.Fatalf("Auth0Fga%v().Execute() = %v, want %v", test.Name, response.StatusCode, test.ResponseStatus)
		}

		responseJson, err := got.MarshalJSON()
		if err != nil {
			t.Fatalf("%v", err)
		}

		if len(*got.Objects) != len(*expectedResponse.Objects) || (*got.Objects)[0] != (*expectedResponse.Objects)[0] {
			t.Fatalf("Auth0Fga%v().Execute() = %v, want %v", test.Name, string(responseJson), test.JsonResponse)
		}
	})

	t.Run("Check with 400 error", func(t *testing.T) {
		test := TestDefinition{
			Name:           "Check",
			JsonResponse:   `{"allowed":true, "resolution":""}`,
			ResponseStatus: 400,
			Method:         "POST",
			RequestPath:    "check",
		}
		requestBody := CheckRequest{
			TupleKey: TupleKey{
				User:     PtrString("user:81684243-9356-4421-8fbf-a4f8d36aa31b"),
				Relation: PtrString("viewer"),
				Object:   PtrString("document:roadmap"),
			},
		}

		var expectedResponse CheckResponse
		if err := json.Unmarshal([]byte(test.JsonResponse), &expectedResponse); err != nil {
			t.Fatalf("%v", err)
		}

		httpmock.Activate()
		defer httpmock.DeactivateAndReset()
		httpmock.RegisterResponder(test.Method, fmt.Sprintf("%s://%s/stores/%s/%s", configuration.ApiScheme, configuration.ApiHost, configuration.StoreId, test.RequestPath),
			func(req *http.Request) (*http.Response, error) {
				errObj := ErrorResponse{
					Code:    "validation_error",
					Message: "Foo",
				}
				return httpmock.NewJsonResponse(400, errObj)
			},
		)
		_, _, err := apiClient.Auth0FgaApi.Check(context.Background()).Body(requestBody).Execute()
		if err == nil {
			t.Fatalf("Expected error with 400 request but there is none")
		}
		validationError, ok := err.(FgaApiValidationError)
		if !ok {
			t.Fatalf("Expected validation Error but type is incorrect %v", err)
		}
		// Do some basic validation of the error itself

		if validationError.StoreId() != configuration.StoreId {
			t.Fatalf("Expected store id to be %s but actual %s", configuration.StoreId, validationError.StoreId())
		}

		if validationError.EndpointCategory() != "Check" {
			t.Fatalf("Expected category to be Check but actual %s", validationError.EndpointCategory())
		}

		if validationError.RequestMethod() != "POST" {
			t.Fatalf("Expected category to be POST but actual %s", validationError.RequestMethod())
		}

		if validationError.ResponseStatusCode() != 400 {
			t.Fatalf("Expected status code to be 400 but actual %d", validationError.ResponseStatusCode())
		}

		if validationError.ResponseCode() != VALIDATION_ERROR {
			t.Fatalf("Expected response code to be VALIDATION_ERROR but actual %s", validationError.ResponseCode())
		}
	})

	t.Run("Check with 401 error", func(t *testing.T) {
		test := TestDefinition{
			Name:           "Check",
			JsonResponse:   `{"allowed":true, "resolution":""}`,
			ResponseStatus: 401,
			Method:         "POST",
			RequestPath:    "check",
		}
		requestBody := CheckRequest{
			TupleKey: TupleKey{
				User:     PtrString("user:81684243-9356-4421-8fbf-a4f8d36aa31b"),
				Relation: PtrString("viewer"),
				Object:   PtrString("document:roadmap"),
			},
		}

		var expectedResponse CheckResponse
		if err := json.Unmarshal([]byte(test.JsonResponse), &expectedResponse); err != nil {
			t.Fatalf("%v", err)
		}

		httpmock.Activate()
		defer httpmock.DeactivateAndReset()
		httpmock.RegisterResponder(test.Method, fmt.Sprintf("%s://%s/stores/%s/%s", configuration.ApiScheme, configuration.ApiHost, configuration.StoreId, test.RequestPath),
			func(req *http.Request) (*http.Response, error) {
				errObj := ErrorResponse{
					Code:    "auth_failure",
					Message: "Foo",
				}
				return httpmock.NewJsonResponse(401, errObj)
			},
		)
		_, _, err := apiClient.Auth0FgaApi.Check(context.Background()).Body(requestBody).Execute()
		if err == nil {
			t.Fatalf("Expected error with 401 request but there is none")
		}
		authenticationError, ok := err.(FgaApiAuthenticationError)
		if !ok {
			t.Fatalf("Expected authentication Error but type is incorrect %v", err)
		}
		// Do some basic validation of the error itself

		if authenticationError.StoreId() != configuration.StoreId {
			t.Fatalf("Expected store id to be %s but actual %s", configuration.StoreId, authenticationError.StoreId())
		}

		if authenticationError.EndpointCategory() != "Check" {
			t.Fatalf("Expected category to be Check but actual %s", authenticationError.EndpointCategory())
		}

		if authenticationError.ResponseStatusCode() != 401 {
			t.Fatalf("Expected status code to be 401 but actual %d", authenticationError.ResponseStatusCode())
		}

	})

	t.Run("Check with 404 error", func(t *testing.T) {
		test := TestDefinition{
			Name:           "Check",
			JsonResponse:   `{"allowed":true, "resolution":""}`,
			ResponseStatus: 404,
			Method:         "POST",
			RequestPath:    "check",
		}
		requestBody := CheckRequest{
			TupleKey: TupleKey{
				User:     PtrString("user:81684243-9356-4421-8fbf-a4f8d36aa31b"),
				Relation: PtrString("viewer"),
				Object:   PtrString("document:roadmap"),
			},
		}

		var expectedResponse CheckResponse
		if err := json.Unmarshal([]byte(test.JsonResponse), &expectedResponse); err != nil {
			t.Fatalf("%v", err)
		}

		httpmock.Activate()
		defer httpmock.DeactivateAndReset()
		httpmock.RegisterResponder(test.Method, fmt.Sprintf("%s://%s/stores/%s/%s", configuration.ApiScheme, configuration.ApiHost, configuration.StoreId, test.RequestPath),
			func(req *http.Request) (*http.Response, error) {
				errObj := ErrorResponse{
					Code:    "undefined_endpoint",
					Message: "Foo",
				}
				return httpmock.NewJsonResponse(404, errObj)
			},
		)
		_, _, err := apiClient.Auth0FgaApi.Check(context.Background()).Body(requestBody).Execute()
		if err == nil {
			t.Fatalf("Expected error with 404 request but there is none")
		}
		notFoundError, ok := err.(FgaApiNotFoundError)
		if !ok {
			t.Fatalf("Expected not found Error but type is incorrect %v", err)
		}
		// Do some basic validation of the error itself

		if notFoundError.StoreId() != configuration.StoreId {
			t.Fatalf("Expected store id to be %s but actual %s", configuration.StoreId, notFoundError.StoreId())
		}

		if notFoundError.EndpointCategory() != "Check" {
			t.Fatalf("Expected category to be Check but actual %s", notFoundError.EndpointCategory())
		}

		if notFoundError.RequestMethod() != "POST" {
			t.Fatalf("Expected category to be POST but actual %s", notFoundError.RequestMethod())
		}

		if notFoundError.ResponseStatusCode() != 404 {
			t.Fatalf("Expected status code to be 404 but actual %d", notFoundError.ResponseStatusCode())
		}

		if notFoundError.ResponseCode() != UNDEFINED_ENDPOINT {
			t.Fatalf("Expected response code to be UNDEFINED_ENDPOINT but actual %s", notFoundError.ResponseCode())
		}
	})

	t.Run("Check with 429 error", func(t *testing.T) {
		test := TestDefinition{
			Name:           "Check",
			JsonResponse:   `{"allowed":true, "resolution":""}`,
			ResponseStatus: 429,
			Method:         "POST",
			RequestPath:    "check",
		}
		requestBody := CheckRequest{
			TupleKey: TupleKey{
				User:     PtrString("user:81684243-9356-4421-8fbf-a4f8d36aa31b"),
				Relation: PtrString("viewer"),
				Object:   PtrString("document:roadmap"),
			},
		}

		var expectedResponse CheckResponse
		if err := json.Unmarshal([]byte(test.JsonResponse), &expectedResponse); err != nil {
			t.Fatalf("%v", err)
		}

		httpmock.Activate()
		defer httpmock.DeactivateAndReset()
		httpmock.RegisterResponder(test.Method, fmt.Sprintf("%s://%s/stores/%s/%s", configuration.ApiScheme, configuration.ApiHost, configuration.StoreId, test.RequestPath),
			func(req *http.Request) (*http.Response, error) {
				errObj := ErrorResponse{
					Code:    "rate_limit_exceeded",
					Message: "Foo",
				}
				return httpmock.NewJsonResponse(429, errObj)
			},
		)

		updatedConfiguration, err := NewConfiguration(UserConfiguration{
			Environment: "playground",
			StoreId:     "6c181474-aaa1-4df7-8929-6e7b3a992754",
			RetryParams: &RetryParams{
				MaxRetry:    3,
				MinWaitInMs: 5,
			},
		})
		if err != nil {
			t.Fatalf("%v", err)
		}

		updatedApiClient := NewAPIClient(updatedConfiguration)

		_, _, err = updatedApiClient.Auth0FgaApi.Check(context.Background()).Body(requestBody).Execute()
		if err == nil {
			t.Fatalf("Expected error with 429 request but there is none")
		}
		rateLimitError, ok := err.(FgaApiRateLimitExceededError)
		if !ok {
			t.Fatalf("Expected rate limit exceeded Error but type is incorrect %v", err)
		}
		// Do some basic validation of the error itself

		if rateLimitError.StoreId() != configuration.StoreId {
			t.Fatalf("Expected store id to be %s but actual %s", configuration.StoreId, rateLimitError.StoreId())
		}

		if rateLimitError.EndpointCategory() != "Check" {
			t.Fatalf("Expected category to be Check but actual %s", rateLimitError.EndpointCategory())
		}

		if rateLimitError.ResponseStatusCode() != 429 {
			t.Fatalf("Expected status code to be 429 but actual %d", rateLimitError.ResponseStatusCode())
		}

	})

	t.Run("Check with initial 429 but eventually resolved", func(t *testing.T) {
		test := TestDefinition{
			Name:           "Check",
			JsonResponse:   `{"allowed":true, "resolution":""}`,
			ResponseStatus: 200,
			Method:         "POST",
			RequestPath:    "check",
		}
		requestBody := CheckRequest{
			TupleKey: TupleKey{
				User:     PtrString("user:81684243-9356-4421-8fbf-a4f8d36aa31b"),
				Relation: PtrString("viewer"),
				Object:   PtrString("document:roadmap"),
			},
		}

		var expectedResponse CheckResponse
		if err := json.Unmarshal([]byte(test.JsonResponse), &expectedResponse); err != nil {
			t.Fatalf("%v", err)
		}

		httpmock.Activate()
		defer httpmock.DeactivateAndReset()
		firstMock := httpmock.NewStringResponder(429, "")
		secondMock, _ := httpmock.NewJsonResponder(200, expectedResponse)
		httpmock.RegisterResponder(test.Method, fmt.Sprintf("%s://%s/stores/%s/%s", configuration.ApiScheme, configuration.ApiHost, configuration.StoreId, test.RequestPath),
			firstMock.Then(firstMock).Then(firstMock).Then(secondMock),
		)
		updatedConfiguration, err := NewConfiguration(UserConfiguration{
			Environment: "playground",
			StoreId:     "6c181474-aaa1-4df7-8929-6e7b3a992754",
			RetryParams: &RetryParams{
				MaxRetry:    2,
				MinWaitInMs: 5,
			},
		})
		if err != nil {
			t.Fatalf("%v", err)
		}

		updatedApiClient := NewAPIClient(updatedConfiguration)

		got, response, err := updatedApiClient.Auth0FgaApi.Check(context.Background()).Body(requestBody).Execute()

		if err != nil {
			t.Fatalf("%v", err)
		}

		if response.StatusCode != test.ResponseStatus {
			t.Fatalf("Auth0Fga%v().Execute() = %v, want %v", test.Name, response.StatusCode, test.ResponseStatus)
		}

		responseJson, err := got.MarshalJSON()
		if err != nil {
			t.Fatalf("%v", err)
		}

		if *got.Allowed != *expectedResponse.Allowed {
			t.Fatalf("Auth0Fga%v().Execute() = %v, want %v", test.Name, string(responseJson), test.JsonResponse)
		}
	})

	t.Run("Check with 500 error", func(t *testing.T) {
		test := TestDefinition{
			Name:           "Check",
			JsonResponse:   `{"allowed":true, "resolution":""}`,
			ResponseStatus: 500,
			Method:         "POST",
			RequestPath:    "check",
		}
		requestBody := CheckRequest{
			TupleKey: TupleKey{
				User:     PtrString("user:81684243-9356-4421-8fbf-a4f8d36aa31b"),
				Relation: PtrString("viewer"),
				Object:   PtrString("document:roadmap"),
			},
		}

		var expectedResponse CheckResponse
		if err := json.Unmarshal([]byte(test.JsonResponse), &expectedResponse); err != nil {
			t.Fatalf("%v", err)
		}

		httpmock.Activate()
		defer httpmock.DeactivateAndReset()
		httpmock.RegisterResponder(test.Method, fmt.Sprintf("%s://%s/stores/%s/%s", configuration.ApiScheme, configuration.ApiHost, configuration.StoreId, test.RequestPath),
			func(req *http.Request) (*http.Response, error) {
				errObj := ErrorResponse{
					Code:    "internal_error",
					Message: "Foo",
				}
				return httpmock.NewJsonResponse(500, errObj)
			},
		)
		_, _, err := apiClient.Auth0FgaApi.Check(context.Background()).Body(requestBody).Execute()
		if err == nil {
			t.Fatalf("Expected error with 500 request but there is none")
		}
		internalError, ok := err.(FgaApiInternalError)
		if !ok {
			t.Fatalf("Expected internal Error but type is incorrect %v", err)
		}
		// Do some basic validation of the error itself

		if internalError.StoreId() != configuration.StoreId {
			t.Fatalf("Expected store id to be %s but actual %s", configuration.StoreId, internalError.StoreId())
		}

		if internalError.EndpointCategory() != "Check" {
			t.Fatalf("Expected category to be Check but actual %s", internalError.EndpointCategory())
		}

		if internalError.RequestMethod() != "POST" {
			t.Fatalf("Expected category to be POST but actual %s", internalError.RequestMethod())
		}

		if internalError.ResponseStatusCode() != 500 {
			t.Fatalf("Expected status code to be 500 but actual %d", internalError.ResponseStatusCode())
		}

		if internalError.ResponseCode() != INTERNAL_ERROR {
			t.Fatalf("Expected response code to be INTERNAL_ERROR but actual %s", internalError.ResponseCode())
		}
	})
}