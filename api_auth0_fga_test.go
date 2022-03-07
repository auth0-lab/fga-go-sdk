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
		httpmock.RegisterResponder("GET", fmt.Sprintf("%s://%s/stores/%s/authorization-models", configuration.Scheme, configuration.Host, configuration.StoreId),
			func(req *http.Request) (*http.Response, error) {
				resp, err := httpmock.NewJsonResponse(200, ReadAuthorizationModelsResponse{AuthorizationModelIds: &[]string{
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
		numCalls = info[fmt.Sprintf("GET %s://%s/stores/%s/authorization-models", configuration.Scheme, configuration.Host, configuration.StoreId)]
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
		httpmock.RegisterResponder("GET", fmt.Sprintf("%s://%s/stores/%s/authorization-models", configuration.Scheme, configuration.Host, configuration.StoreId),
			func(req *http.Request) (*http.Response, error) {
				resp, err := httpmock.NewJsonResponse(200, ReadAuthorizationModelsResponse{AuthorizationModelIds: &[]string{
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
		numCalls = info[fmt.Sprintf("GET %s://%s/stores/%s/authorization-models", configuration.Scheme, configuration.Host, configuration.StoreId)]
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

		var expectedResponse ReadAuthorizationModelsResponse
		if err := json.Unmarshal([]byte(test.JsonResponse), &expectedResponse); err != nil {
			t.Errorf("%v", err)
			return
		}

		httpmock.Activate()
		defer httpmock.DeactivateAndReset()
		httpmock.RegisterResponder(test.Method, fmt.Sprintf("%s://%s/stores/%s/%s", configuration.Scheme, configuration.Host, configuration.StoreId, test.RequestPath),
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
			JsonResponse:   `{"authorization_model_id":"1uHxCSuTP0VKPYSnkq1pbb1jeZw"}`,
			ResponseStatus: 200,
			Method:         "POST",
			RequestPath:    "authorization-models",
		}
		requestBody := TypeDefinitions{
			TypeDefinitions: &[]TypeDefinition{{
				Type: "github-repo",
				Relations: map[string]Userset{
					"repo_writer": {
						This: &map[string]interface{}{},
					},
					"repo_reader": {Union: &Usersets{
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
			t.Errorf("%v", err)
			return
		}

		httpmock.Activate()
		defer httpmock.DeactivateAndReset()
		httpmock.RegisterResponder(test.Method, fmt.Sprintf("%s://%s/stores/%s/%s", configuration.Scheme, configuration.Host, configuration.StoreId, test.RequestPath),
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

		var expectedResponse ReadAuthorizationModelResponse
		if err := json.Unmarshal([]byte(test.JsonResponse), &expectedResponse); err != nil {
			t.Errorf("%v", err)
			return
		}
		modelId := *(*expectedResponse.AuthorizationModel).Id

		httpmock.Activate()
		defer httpmock.DeactivateAndReset()
		httpmock.RegisterResponder(test.Method, fmt.Sprintf("%s://%s/stores/%s/%s/%s", configuration.Scheme, configuration.Host, configuration.StoreId, test.RequestPath, modelId),
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
		requestBody := CheckRequestParams{
			TupleKey: &TupleKey{
				User:     PtrString("anne@auth0.com"),
				Relation: PtrString("repo_reader"),
				Object:   PtrString("github-repo:auth0/express-jwt"),
			},
		}

		var expectedResponse CheckResponse
		if err := json.Unmarshal([]byte(test.JsonResponse), &expectedResponse); err != nil {
			t.Errorf("%v", err)
			return
		}

		httpmock.Activate()
		defer httpmock.DeactivateAndReset()
		httpmock.RegisterResponder(test.Method, fmt.Sprintf("%s://%s/stores/%s/%s", configuration.Scheme, configuration.Host, configuration.StoreId, test.RequestPath),
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
		requestBody := WriteRequestParams{
			Writes: &TupleKeys{
				TupleKeys: []TupleKey{{
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
		httpmock.RegisterResponder(test.Method, fmt.Sprintf("%s://%s/stores/%s/%s", configuration.Scheme, configuration.Host, configuration.StoreId, test.RequestPath),
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

		requestBody := WriteRequestParams{
			Deletes: &TupleKeys{
				TupleKeys: []TupleKey{{
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
		httpmock.RegisterResponder(test.Method, fmt.Sprintf("%s://%s/stores/%s/%s", configuration.Scheme, configuration.Host, configuration.StoreId, test.RequestPath),
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

		requestBody := ExpandRequestParams{
			TupleKey: &TupleKey{
				Relation: PtrString("repo_reader"),
				Object:   PtrString("github-repo:auth0/express-jwt"),
			},
		}

		var expectedResponse ExpandResponse
		if err := json.Unmarshal([]byte(test.JsonResponse), &expectedResponse); err != nil {
			t.Errorf("%v", err)
			return
		}

		httpmock.Activate()
		defer httpmock.DeactivateAndReset()
		httpmock.RegisterResponder(test.Method, fmt.Sprintf("%s://%s/stores/%s/%s", configuration.Scheme, configuration.Host, configuration.StoreId, test.RequestPath),
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

		requestBody := ReadRequestParams{
			TupleKey: &TupleKey{
				User:     PtrString("anne@auth0.com"),
				Relation: PtrString("repo_reader"),
				Object:   PtrString("github-repo:auth0/express-jwt"),
			},
		}

		var expectedResponse ReadResponse
		if err := json.Unmarshal([]byte(test.JsonResponse), &expectedResponse); err != nil {
			t.Errorf("%v", err)
			return
		}

		httpmock.Activate()
		defer httpmock.DeactivateAndReset()
		httpmock.RegisterResponder(test.Method, fmt.Sprintf("%s://%s/stores/%s/%s", configuration.Scheme, configuration.Host, configuration.StoreId, test.RequestPath),
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

	t.Run("Check with 400 error", func(t *testing.T) {
		test := TestDefinition{
			Name:           "Check",
			JsonResponse:   `{"allowed":true, "resolution":""}`,
			ResponseStatus: 400,
			Method:         "POST",
			RequestPath:    "check",
		}
		requestBody := CheckRequestParams{
			TupleKey: &TupleKey{
				User:     PtrString("anne@auth0.com"),
				Relation: PtrString("repo_reader"),
				Object:   PtrString("github-repo:auth0/express-jwt"),
			},
		}

		var expectedResponse CheckResponse
		if err := json.Unmarshal([]byte(test.JsonResponse), &expectedResponse); err != nil {
			t.Errorf("%v", err)
			return
		}

		httpmock.Activate()
		defer httpmock.DeactivateAndReset()
		httpmock.RegisterResponder(test.Method, fmt.Sprintf("%s://%s/stores/%s/%s", configuration.Scheme, configuration.Host, configuration.StoreId, test.RequestPath),
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
			t.Errorf("Expected error with 400 request but there is none")
			return
		}
		validationError, ok := err.(Auth0FgaApiValidationError)
		if !ok {
			t.Errorf("Expected validation Error but type is incorrect %v", err)
			return
		}
		// Do some basic validation of the error itself

		if validationError.StoreId() != configuration.StoreId {
			t.Errorf("Expected store id to be %s but actual %s", configuration.StoreId, validationError.StoreId())
			return
		}

		if validationError.EndpointCategory() != "Check" {
			t.Errorf("Expected category to be Check but actual %s", validationError.EndpointCategory())
			return
		}

		if validationError.RequestMethod() != "POST" {
			t.Errorf("Expected category to be POST but actual %s", validationError.RequestMethod())
			return
		}

		if validationError.ResponseStatusCode() != 400 {
			t.Errorf("Expected status code to be 400 but actual %d", validationError.ResponseStatusCode())
			return
		}

		if validationError.ResponseCode() != VALIDATION_ERROR {
			t.Errorf("Expected response code to be VALIDATION_ERROR but actual %s", validationError.ResponseCode())
			return
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
		requestBody := CheckRequestParams{
			TupleKey: &TupleKey{
				User:     PtrString("anne@auth0.com"),
				Relation: PtrString("repo_reader"),
				Object:   PtrString("github-repo:auth0/express-jwt"),
			},
		}

		var expectedResponse CheckResponse
		if err := json.Unmarshal([]byte(test.JsonResponse), &expectedResponse); err != nil {
			t.Errorf("%v", err)
			return
		}

		httpmock.Activate()
		defer httpmock.DeactivateAndReset()
		httpmock.RegisterResponder(test.Method, fmt.Sprintf("%s://%s/stores/%s/%s", configuration.Scheme, configuration.Host, configuration.StoreId, test.RequestPath),
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
			t.Errorf("Expected error with 401 request but there is none")
			return
		}
		authenticationError, ok := err.(Auth0FgaApiAuthenticationError)
		if !ok {
			t.Errorf("Expected authentication Error but type is incorrect %v", err)
			return
		}
		// Do some basic validation of the error itself

		if authenticationError.StoreId() != configuration.StoreId {
			t.Errorf("Expected store id to be %s but actual %s", configuration.StoreId, authenticationError.StoreId())
			return
		}

		if authenticationError.EndpointCategory() != "Check" {
			t.Errorf("Expected category to be Check but actual %s", authenticationError.EndpointCategory())
			return
		}

		if authenticationError.ResponseStatusCode() != 401 {
			t.Errorf("Expected status code to be 401 but actual %d", authenticationError.ResponseStatusCode())
			return
		}

		if authenticationError.ResponseCode() != AUTH_FAILURE {
			t.Errorf("Expected response code to be AUTH_FAILURE but actual %s", authenticationError.ResponseCode())
			return
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
		requestBody := CheckRequestParams{
			TupleKey: &TupleKey{
				User:     PtrString("anne@auth0.com"),
				Relation: PtrString("repo_reader"),
				Object:   PtrString("github-repo:auth0/express-jwt"),
			},
		}

		var expectedResponse CheckResponse
		if err := json.Unmarshal([]byte(test.JsonResponse), &expectedResponse); err != nil {
			t.Errorf("%v", err)
			return
		}

		httpmock.Activate()
		defer httpmock.DeactivateAndReset()
		httpmock.RegisterResponder(test.Method, fmt.Sprintf("%s://%s/stores/%s/%s", configuration.Scheme, configuration.Host, configuration.StoreId, test.RequestPath),
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
			t.Errorf("Expected error with 404 request but there is none")
			return
		}
		notFoundError, ok := err.(Auth0FgaApiNotFoundError)
		if !ok {
			t.Errorf("Expected not found Error but type is incorrect %v", err)
			return
		}
		// Do some basic validation of the error itself

		if notFoundError.StoreId() != configuration.StoreId {
			t.Errorf("Expected store id to be %s but actual %s", configuration.StoreId, notFoundError.StoreId())
			return
		}

		if notFoundError.EndpointCategory() != "Check" {
			t.Errorf("Expected category to be Check but actual %s", notFoundError.EndpointCategory())
			return
		}

		if notFoundError.RequestMethod() != "POST" {
			t.Errorf("Expected category to be POST but actual %s", notFoundError.RequestMethod())
			return
		}

		if notFoundError.ResponseStatusCode() != 404 {
			t.Errorf("Expected status code to be 404 but actual %d", notFoundError.ResponseStatusCode())
			return
		}

		if notFoundError.ResponseCode() != UNDEFINED_ENDPOINT {
			t.Errorf("Expected response code to be UNDEFINED_ENDPOINT but actual %s", notFoundError.ResponseCode())
			return
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
		requestBody := CheckRequestParams{
			TupleKey: &TupleKey{
				User:     PtrString("anne@auth0.com"),
				Relation: PtrString("repo_reader"),
				Object:   PtrString("github-repo:auth0/express-jwt"),
			},
		}

		var expectedResponse CheckResponse
		if err := json.Unmarshal([]byte(test.JsonResponse), &expectedResponse); err != nil {
			t.Errorf("%v", err)
			return
		}

		httpmock.Activate()
		defer httpmock.DeactivateAndReset()
		httpmock.RegisterResponder(test.Method, fmt.Sprintf("%s://%s/stores/%s/%s", configuration.Scheme, configuration.Host, configuration.StoreId, test.RequestPath),
			func(req *http.Request) (*http.Response, error) {
				errObj := ErrorResponse{
					Code:    "rate_limit_exceeded",
					Message: "Foo",
				}
				return httpmock.NewJsonResponse(429, errObj)
			},
		)

		updatedConfiguration, err := NewConfiguration(UserConfiguration{
			StoreId:     "6c181474-aaa1-4df7-8929-6e7b3a992754",
			Environment: "playground",
			RetryParams: &RetryParams{
				MaxRetry:    3,
				MinWaitInMs: 5,
			},
		})
		if err != nil {
			t.Errorf("%v", err)
			return
		}

		//Ensure we never talk to a live host
		updatedConfiguration.Host = "api.fga.auth0.example"

		updatedApiClient := NewAPIClient(updatedConfiguration)

		_, _, err = updatedApiClient.Auth0FgaApi.Check(context.Background()).Body(requestBody).Execute()
		if err == nil {
			t.Errorf("Expected error with 429 request but there is none")
			return
		}
		rateLimitError, ok := err.(Auth0FgaApiRateLimitError)
		if !ok {
			t.Errorf("Expected authentication Error but type is incorrect %v", err)
			return
		}
		// Do some basic validation of the error itself

		if rateLimitError.StoreId() != configuration.StoreId {
			t.Errorf("Expected store id to be %s but actual %s", configuration.StoreId, rateLimitError.StoreId())
			return
		}

		if rateLimitError.EndpointCategory() != "Check" {
			t.Errorf("Expected category to be Check but actual %s", rateLimitError.EndpointCategory())
			return
		}

		if rateLimitError.ResponseStatusCode() != 429 {
			t.Errorf("Expected status code to be 429 but actual %d", rateLimitError.ResponseStatusCode())
			return
		}

		if rateLimitError.ResponseCode() != RATE_LIMIT_EXCEEDED {
			t.Errorf("Expected response code to be RATE_LIMIT_EXCEEDED but actual %s", rateLimitError.ResponseCode())
			return
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
		requestBody := CheckRequestParams{
			TupleKey: &TupleKey{
				User:     PtrString("anne@auth0.com"),
				Relation: PtrString("repo_reader"),
				Object:   PtrString("github-repo:auth0/express-jwt"),
			},
		}

		var expectedResponse CheckResponse
		if err := json.Unmarshal([]byte(test.JsonResponse), &expectedResponse); err != nil {
			t.Errorf("%v", err)
			return
		}

		httpmock.Activate()
		defer httpmock.DeactivateAndReset()
		firstMock := httpmock.NewStringResponder(429, "")
		secondMock, _ := httpmock.NewJsonResponder(200, expectedResponse)
		httpmock.RegisterResponder(test.Method, fmt.Sprintf("%s://%s/stores/%s/%s", configuration.Scheme, configuration.Host, configuration.StoreId, test.RequestPath),
			firstMock.Then(firstMock).Then(firstMock).Then(secondMock),
		)
		updatedConfiguration, err := NewConfiguration(UserConfiguration{
			StoreId:     "6c181474-aaa1-4df7-8929-6e7b3a992754",
			Environment: "playground",
			RetryParams: &RetryParams{
				MaxRetry:    2,
				MinWaitInMs: 5,
			},
		})
		if err != nil {
			t.Errorf("%v", err)
			return
		}

		//Ensure we never talk to a live host
		updatedConfiguration.Host = "api.fga.auth0.example"

		updatedApiClient := NewAPIClient(updatedConfiguration)

		got, response, err := updatedApiClient.Auth0FgaApi.Check(context.Background()).Body(requestBody).Execute()

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

	t.Run("Check with 500 error", func(t *testing.T) {
		test := TestDefinition{
			Name:           "Check",
			JsonResponse:   `{"allowed":true, "resolution":""}`,
			ResponseStatus: 500,
			Method:         "POST",
			RequestPath:    "check",
		}
		requestBody := CheckRequestParams{
			TupleKey: &TupleKey{
				User:     PtrString("anne@auth0.com"),
				Relation: PtrString("repo_reader"),
				Object:   PtrString("github-repo:auth0/express-jwt"),
			},
		}

		var expectedResponse CheckResponse
		if err := json.Unmarshal([]byte(test.JsonResponse), &expectedResponse); err != nil {
			t.Errorf("%v", err)
			return
		}

		httpmock.Activate()
		defer httpmock.DeactivateAndReset()
		httpmock.RegisterResponder(test.Method, fmt.Sprintf("%s://%s/stores/%s/%s", configuration.Scheme, configuration.Host, configuration.StoreId, test.RequestPath),
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
			t.Errorf("Expected error with 500 request but there is none")
			return
		}
		internalError, ok := err.(Auth0FgaApiInternalError)
		if !ok {
			t.Errorf("Expected internal Error but type is incorrect %v", err)
			return
		}
		// Do some basic validation of the error itself

		if internalError.StoreId() != configuration.StoreId {
			t.Errorf("Expected store id to be %s but actual %s", configuration.StoreId, internalError.StoreId())
			return
		}

		if internalError.EndpointCategory() != "Check" {
			t.Errorf("Expected category to be Check but actual %s", internalError.EndpointCategory())
			return
		}

		if internalError.RequestMethod() != "POST" {
			t.Errorf("Expected category to be POST but actual %s", internalError.RequestMethod())
			return
		}

		if internalError.ResponseStatusCode() != 500 {
			t.Errorf("Expected status code to be 500 but actual %d", internalError.ResponseStatusCode())
			return
		}

		if internalError.ResponseCode() != INTERNAL_ERROR {
			t.Errorf("Expected response code to be INTERNAL_ERROR but actual %s", internalError.ResponseCode())
			return
		}
	})
}
