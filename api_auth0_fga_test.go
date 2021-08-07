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
	t.Run("Providing no tenant should error", func(t *testing.T) {
		_, err := NewConfiguration(UserConfiguration{
			Environment: "playground",
		})

		if err == nil {
			t.Errorf("Expected an error")
			return
		}
	})

	t.Run("Providing no environment should error", func(t *testing.T) {
		_, err := NewConfiguration(UserConfiguration{
			StoreId: "6c181474-aaa1-4df7-8929-6e7b3a992754",
		})

		if err == nil {
			t.Errorf("Expected an error")
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
		configuration.Host = "api.fga.auth0.example"

		apiClient := NewAPIClient(configuration)

		httpmock.Activate()
		defer httpmock.DeactivateAndReset()
		httpmock.RegisterResponder("GET", fmt.Sprintf("%s://%s/%s/authorization-models", configuration.Scheme, configuration.Host, configuration.StoreId),
			func(req *http.Request) (*http.Response, error) {
				resp, err := httpmock.NewJsonResponse(200, Auth0FgaReadAuthorizationModelsResponse{AuthorizationModelIds: &[]string{
					"1uHxCSuTP0VKPYSnkq1pbb1jeZw",
					"GtQpMohWezFmIbyXxVEocOCxxgq",
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
		numCalls = info[fmt.Sprintf("GET %s://%s/%s/authorization-models", configuration.Scheme, configuration.Host, configuration.StoreId)]
		if numCalls != 1 {
			t.Errorf("Expected call to get authorization models to be made exactly once, saw: %d", numCalls)
			return
		}
	})

	t.Run("should not issue a network call to get the token at the first request if the clientId is not provided", func(t *testing.T) {
		configuration, err := NewConfiguration(UserConfiguration{
			StoreId:     "6c181474-aaa1-4df7-8929-6e7b3a992754",
			Environment: "playground",
		})
		if err != nil {
			t.Errorf("%v", err)
			return
		}
		configuration.Host = "api.fga.auth0.example"

		apiClient := NewAPIClient(configuration)

		httpmock.Activate()
		defer httpmock.DeactivateAndReset()
		httpmock.RegisterResponder("GET", fmt.Sprintf("%s://%s/%s/authorization-models", configuration.Scheme, configuration.Host, configuration.StoreId),
			func(req *http.Request) (*http.Response, error) {
				resp, err := httpmock.NewJsonResponse(200, Auth0FgaReadAuthorizationModelsResponse{AuthorizationModelIds: &[]string{
					"1uHxCSuTP0VKPYSnkq1pbb1jeZw",
					"GtQpMohWezFmIbyXxVEocOCxxgq",
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
		numCalls = info[fmt.Sprintf("GET %s://%s/%s/authorization-models", configuration.Scheme, configuration.Host, configuration.StoreId)]
		if numCalls != 1 {
			t.Errorf("Expected call to get authorization models to be made exactly once, saw: %d", numCalls)
			return
		}
	})
}

func TestAuth0FgaApi(t *testing.T) {
	configuration, err := NewConfiguration(UserConfiguration{
		StoreId:     "6c181474-aaa1-4df7-8929-6e7b3a992754",
		Environment: "playground",
	})
	if err != nil {
		t.Errorf("%v", err)
		return
	}

	//Ensure we never talk to a live host
	configuration.Host = "api.fga.auth0.example"

	apiClient := NewAPIClient(configuration)

	t.Run("ReadAuthorizationModels", func(t *testing.T) {
		test := TestDefinition{
			Name:           "ReadAuthorizationModels",
			JsonResponse:   `{"authorization_model_ids":["1uHxCSuTP0VKPYSnkq1pbb1jeZw"]}`,
			ResponseStatus: 200,
			Method:         "GET",
			RequestPath:    "authorization-models",
		}

		var expectedResponse Auth0FgaReadAuthorizationModelsResponse
		if err := json.Unmarshal([]byte(test.JsonResponse), &expectedResponse); err != nil {
			t.Errorf("%v", err)
			return
		}

		httpmock.Activate()
		defer httpmock.DeactivateAndReset()
		httpmock.RegisterResponder(test.Method, fmt.Sprintf("%s://%s/%s/%s", configuration.Scheme, configuration.Host, configuration.StoreId, test.RequestPath),
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
			t.Errorf("%v", err)
			return
		}

		if response.StatusCode != test.ResponseStatus {
			t.Errorf("Auth0Fga%v().Execute() = %v, want %v", test.Name, response.StatusCode, test.ResponseStatus)
			return
		}

		responseJson, err := got.MarshalJSON()
		if err != nil {
			t.Errorf("%v", err)
			return
		}

		if len(*got.AuthorizationModelIds) != 1 {
			t.Errorf("%v", err)
			return
		}

		if (*got.AuthorizationModelIds)[0] != (*expectedResponse.AuthorizationModelIds)[0] {
			t.Errorf("Auth0Fga%v().Execute() = %v, want %v", test.Name, string(responseJson), test.JsonResponse)
			return
		}
	})

	t.Run("WriteAuthorizationModel", func(t *testing.T) {
		test := TestDefinition{
			Name:           "WriteAuthorizationModel",
			JsonResponse:   `{"id":"1uHxCSuTP0VKPYSnkq1pbb1jeZw"}`,
			ResponseStatus: 200,
			Method:         "POST",
			RequestPath:    "authorization-models",
		}
		requestBody := AuthorizationmodelTypeDefinitions{
			TypeDefinitions: &[]AuthorizationmodelTypeDefinition{{
				Type: "github-repo",
				Relations: map[string]AuthorizationmodelUserset{
					"repo_writer": {
						This: &map[string]interface{}{},
					},
					"repo_reader": {Union: &AuthorizationmodelUsersets{
						Child: &[]AuthorizationmodelUserset{
							{This: &map[string]interface{}{}},
							{ComputedUserset: &AuthorizationmodelObjectRelation{
								Object:   PtrString(""),
								Relation: PtrString("repo_writer"),
							}},
						},
					}},
				},
			}},
		}

		var expectedResponse Auth0FgaWriteAuthorizationModelResponse
		if err := json.Unmarshal([]byte(test.JsonResponse), &expectedResponse); err != nil {
			t.Errorf("%v", err)
			return
		}

		httpmock.Activate()
		defer httpmock.DeactivateAndReset()
		httpmock.RegisterResponder(test.Method, fmt.Sprintf("%s://%s/%s/%s", configuration.Scheme, configuration.Host, configuration.StoreId, test.RequestPath),
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
			t.Errorf("%v", err)
			return
		}

		if response.StatusCode != test.ResponseStatus {
			t.Errorf("Auth0Fga%v().Execute() = %v, want %v", test.Name, response.StatusCode, test.ResponseStatus)
			return
		}

		_, err = got.MarshalJSON()
		if err != nil {
			t.Errorf("%v", err)
			return
		}
	})

	t.Run("ReadAuthorizationModel", func(t *testing.T) {
		test := TestDefinition{
			Name:           "ReadAuthorizationModel",
			JsonResponse:   `{"authorization_model":{"id":"1uHxCSuTP0VKPYSnkq1pbb1jeZw", "type_definitions":[{"type":"github-repo", "relations":{"repo_reader":{"this":{}}}}]}}`,
			ResponseStatus: 200,
			Method:         "GET",
			RequestPath:    "authorization-models",
		}

		var expectedResponse Auth0FgaReadAuthorizationModelResponse
		if err := json.Unmarshal([]byte(test.JsonResponse), &expectedResponse); err != nil {
			t.Errorf("%v", err)
			return
		}
		modelId := *(*expectedResponse.AuthorizationModel).Id

		httpmock.Activate()
		defer httpmock.DeactivateAndReset()
		httpmock.RegisterResponder(test.Method, fmt.Sprintf("%s://%s/%s/%s/%s", configuration.Scheme, configuration.Host, configuration.StoreId, test.RequestPath, modelId),
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
			t.Errorf("%v", err)
			return
		}

		if response.StatusCode != test.ResponseStatus {
			t.Errorf("Auth0Fga%v().Execute() = %v, want %v", test.Name, response.StatusCode, test.ResponseStatus)
			return
		}

		responseJson, err := got.MarshalJSON()
		if err != nil {
			t.Errorf("%v", err)
			return
		}

		if *(*got.AuthorizationModel).Id != modelId {
			t.Errorf("Auth0Fga%v().Execute() = %v, want %v", test.Name, string(responseJson), test.JsonResponse)
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
		requestBody := Auth0FgaCheckRequestParams{
			TupleKey: &Auth0FgaTupleKey{
				User:     PtrString("anne@auth0.com"),
				Relation: PtrString("repo_reader"),
				Object:   PtrString("github-repo:auth0/express-jwt"),
			},
		}

		var expectedResponse Auth0FgaCheckResponse
		if err := json.Unmarshal([]byte(test.JsonResponse), &expectedResponse); err != nil {
			t.Errorf("%v", err)
			return
		}

		httpmock.Activate()
		defer httpmock.DeactivateAndReset()
		httpmock.RegisterResponder(test.Method, fmt.Sprintf("%s://%s/%s/%s", configuration.Scheme, configuration.Host, configuration.StoreId, test.RequestPath),
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
			t.Errorf("%v", err)
			return
		}

		if response.StatusCode != test.ResponseStatus {
			t.Errorf("Auth0Fga%v().Execute() = %v, want %v", test.Name, response.StatusCode, test.ResponseStatus)
			return
		}

		responseJson, err := got.MarshalJSON()
		if err != nil {
			t.Errorf("%v", err)
			return
		}

		if *got.Allowed != *expectedResponse.Allowed {
			t.Errorf("Auth0Fga%v().Execute() = %v, want %v", test.Name, string(responseJson), test.JsonResponse)
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
		requestBody := Auth0FgaWriteRequestParams{
			Writes: &Auth0FgaTupleKeys{
				TupleKeys: []Auth0FgaTupleKey{{
					User:     PtrString("anne@auth0.com"),
					Relation: PtrString("repo_reader"),
					Object:   PtrString("github-repo:auth0/express-jwt"),
				}},
			},
		}

		var expectedResponse map[string]interface{}
		if err := json.Unmarshal([]byte(test.JsonResponse), &expectedResponse); err != nil {
			t.Errorf("%v", err)
			return
		}

		httpmock.Activate()
		defer httpmock.DeactivateAndReset()
		httpmock.RegisterResponder(test.Method, fmt.Sprintf("%s://%s/%s/%s", configuration.Scheme, configuration.Host, configuration.StoreId, test.RequestPath),
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
			t.Errorf("%v", err)
			return
		}

		if response.StatusCode != test.ResponseStatus {
			t.Errorf("Auth0Fga%v().Execute() = %v, want %v", test.Name, response.StatusCode, test.ResponseStatus)
			return
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

		requestBody := Auth0FgaWriteRequestParams{
			Deletes: &Auth0FgaTupleKeys{
				TupleKeys: []Auth0FgaTupleKey{{
					User:     PtrString("anne@auth0.com"),
					Relation: PtrString("repo_reader"),
					Object:   PtrString("github-repo:auth0/express-jwt"),
				}},
			},
		}

		var expectedResponse map[string]interface{}
		if err := json.Unmarshal([]byte(test.JsonResponse), &expectedResponse); err != nil {
			t.Errorf("%v", err)
			return
		}

		httpmock.Activate()
		defer httpmock.DeactivateAndReset()
		httpmock.RegisterResponder(test.Method, fmt.Sprintf("%s://%s/%s/%s", configuration.Scheme, configuration.Host, configuration.StoreId, test.RequestPath),
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
			t.Errorf("%v", err)
			return
		}

		if response.StatusCode != test.ResponseStatus {
			t.Errorf("Auth0Fga%v().Execute() = %v, want %v", test.Name, response.StatusCode, test.ResponseStatus)
			return
		}
	})

	t.Run("Expand", func(t *testing.T) {
		test := TestDefinition{
			Name:           "Expand",
			JsonResponse:   `{"tree":{"root":{"name":"github-repo:auth0/express-jwt#repo_reader","union":{"nodes":[{"name": "github-repo:auth0/express-jwt#repo_reader","leaf":{"users":{"users":["anne@auth0.com"]}}}]}}}}`,
			ResponseStatus: 200,
			Method:         "POST",
			RequestPath:    "expand",
		}

		requestBody := Auth0FgaExpandRequestParams{
			TupleKey: &Auth0FgaTupleKey{
				Relation: PtrString("repo_reader"),
				Object:   PtrString("github-repo:auth0/express-jwt"),
			},
		}

		var expectedResponse Auth0FgaExpandResponse
		if err := json.Unmarshal([]byte(test.JsonResponse), &expectedResponse); err != nil {
			t.Errorf("%v", err)
			return
		}

		httpmock.Activate()
		defer httpmock.DeactivateAndReset()
		httpmock.RegisterResponder(test.Method, fmt.Sprintf("%s://%s/%s/%s", configuration.Scheme, configuration.Host, configuration.StoreId, test.RequestPath),
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
			t.Errorf("%v", err)
			return
		}

		if response.StatusCode != test.ResponseStatus {
			t.Errorf("Auth0Fga%v().Execute() = %v, want %v", test.Name, response.StatusCode, test.ResponseStatus)
			return
		}

		_, err = got.MarshalJSON()
		if err != nil {
			t.Errorf("%v", err)
			return
		}
	})

	t.Run("Read", func(t *testing.T) {
		test := TestDefinition{
			Name:           "Read",
			JsonResponse:   `{"tuples":[{"key":{"user":"anne@auth0.com","relation":"repo_reader","object":"github-repo:auth0/express-jwt"},"timestamp": "2000-01-01T00:00:00Z"}]}`,
			ResponseStatus: 200,
			Method:         "POST",
			RequestPath:    "read",
		}

		requestBody := Auth0FgaReadRequestParams{
			TupleKey: &Auth0FgaTupleKey{
				User:     PtrString("anne@auth0.com"),
				Relation: PtrString("repo_reader"),
				Object:   PtrString("github-repo:auth0/express-jwt"),
			},
		}

		var expectedResponse Auth0FgaReadResponse
		if err := json.Unmarshal([]byte(test.JsonResponse), &expectedResponse); err != nil {
			t.Errorf("%v", err)
			return
		}

		httpmock.Activate()
		defer httpmock.DeactivateAndReset()
		httpmock.RegisterResponder(test.Method, fmt.Sprintf("%s://%s/%s/%s", configuration.Scheme, configuration.Host, configuration.StoreId, test.RequestPath),
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
			t.Errorf("%v", err)
			return
		}

		if response.StatusCode != test.ResponseStatus {
			t.Errorf("Auth0Fga%v().Execute() = %v, want %v", test.Name, response.StatusCode, test.ResponseStatus)
			return
		}

		responseJson, err := got.MarshalJSON()
		if err != nil {
			t.Errorf("%v", err)
			return
		}

		if len(*got.Tuples) != len(*expectedResponse.Tuples) {
			t.Errorf("Auth0Fga%v().Execute() = %v, want %v", test.Name, string(responseJson), test.JsonResponse)
			return
		}
	})
}
