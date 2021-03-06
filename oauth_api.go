/* 
 * Phone.com API
 *
 * This is a Phone.com api Swagger definition
 *
 * OpenAPI spec version: 1.0.0
 * Contact: apisupport@phone.com
 * Generated by: https://github.com/swagger-api/swagger-codegen.git
 */

package swagger

import (
	"net/url"
	"strings"
	"encoding/json"
)

type OauthApi struct {
	Configuration *Configuration
}

func NewOauthApi() *OauthApi {
	configuration := NewConfiguration()
	return &OauthApi{
		Configuration: configuration,
	}
}

func NewOauthApiWithBasePath(basePath string) *OauthApi {
	configuration := NewConfiguration()
	configuration.BasePath = basePath

	return &OauthApi{
		Configuration: configuration,
	}
}

/**
 * To create an access token via the /oauth/access-token API, an API user may choose any one of the grant types it supports: Authorization Code Grant, Client Credential Grant, Password Credential Grant or Refresh Token Grant.
 * To create an access token via the /oauth/access-token API, an API user may choose any one of the grant types it supports: Authorization Code Grant, Client Credential Grant, Password Credential Grant or Refresh Token Grant. For Authorization Code Grant, the input parameter &#39;code&#39; is generated via the Create Authorization API. NOTE: The Create Access Token API now accepts requests in query string format as well as JSON body format. See OAuth for more details on how to obtain client id and client secret to create an access token at real time.
 *
 * @param data Oauth data
 * @return *OauthAccessToken
 */
func (a OauthApi) CreateOauthAccessToken(data CreateOauthParams) (*OauthAccessToken, *APIResponse, error) {

	var localVarHttpMethod = strings.ToUpper("Post")
	// create path and map variables
	localVarPath := a.Configuration.BasePath + "/oauth/access-token"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make(map[string]string)
	var localVarPostBody interface{}
	var localVarFileName string
	var localVarFileBytes []byte
	// authentication '(apiKey)' required
	// set key with prefix in header
	localVarHeaderParams["Authorization"] = a.Configuration.GetAPIKeyWithPrefix("Authorization")
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		localVarHeaderParams[key] = a.Configuration.DefaultHeader[key]
	}

	clearEmptyParams(localVarQueryParams)

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	localVarPostBody = &data
	var successPayload = new(OauthAccessToken)
	localVarHttpResponse, err := a.Configuration.APIClient.CallAPI(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)

	var localVarURL, _ = url.Parse(localVarPath)
	localVarURL.RawQuery = localVarQueryParams.Encode()
	var localVarAPIResponse = &APIResponse{Operation: "CreateOauthAccessToken", Method: localVarHttpMethod, RequestURL: localVarURL.String()}
	if localVarHttpResponse != nil {
		localVarAPIResponse.Response = localVarHttpResponse.RawResponse
		localVarAPIResponse.Payload = localVarHttpResponse.Body()
	}

	if err != nil {
		return successPayload, localVarAPIResponse, err
	}
	err = json.Unmarshal(localVarHttpResponse.Body(), &successPayload)
	return successPayload, localVarAPIResponse, err
}

/**
 * Create Authorization Code or Access Token.
 * Create Authorization Code or Access Token. The /oauth/authorization API supports Authorization Grant and Implicit Grant. In Authorization Grant, this API returns a code (response_type&#x3D;code) for clients implemented in a browser using a scripting language such as JavaScript. Users may then use the code via the Create Access Token API to create an access token. The Implicit Grant is a simplified authorization code flow. In the implicit flow, instead of issuing the client an authorization code, the client is issued an access token (response_type&#x3D;token) directly. See OAuth for more details on how to obtain client id and client secret to create authorization code access token at real time.
 *
 * @param clientId Client ID
 * @param responseType &#39;token&#39; for Implicit Grant; &#39;code&#39; for Authorization Code Grant
 * @param scope account-owner, extension-user and/or methods:ALL, separated by space (%20)
 * @param redirectUri The URL where the response data of this API is redirected to
 * @return void
 */
func (a OauthApi) CreateOauthAuthorization(clientId string, responseType string, scope string, redirectUri string) (*APIResponse, error) {

	var localVarHttpMethod = strings.ToUpper("Get")
	// create path and map variables
	localVarPath := a.Configuration.BasePath + "/oauth/authorization"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make(map[string]string)
	var localVarPostBody interface{}
	var localVarFileName string
	var localVarFileBytes []byte
	// authentication '(apiKey)' required
	// set key with prefix in header
	localVarHeaderParams["Authorization"] = a.Configuration.GetAPIKeyWithPrefix("Authorization")
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		localVarHeaderParams[key] = a.Configuration.DefaultHeader[key]
	}
	localVarQueryParams.Add("client_id", a.Configuration.APIClient.ParameterToString(clientId, ""))
	localVarQueryParams.Add("response_type", a.Configuration.APIClient.ParameterToString(responseType, ""))
	localVarQueryParams.Add("scope", a.Configuration.APIClient.ParameterToString(scope, ""))
	localVarQueryParams.Add("redirect_uri", a.Configuration.APIClient.ParameterToString(redirectUri, ""))

	clearEmptyParams(localVarQueryParams)

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	localVarHttpResponse, err := a.Configuration.APIClient.CallAPI(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)

	var localVarURL, _ = url.Parse(localVarPath)
	localVarURL.RawQuery = localVarQueryParams.Encode()
	var localVarAPIResponse = &APIResponse{Operation: "CreateOauthAuthorization", Method: localVarHttpMethod, RequestURL: localVarURL.String()}
	if localVarHttpResponse != nil {
		localVarAPIResponse.Response = localVarHttpResponse.RawResponse
		localVarAPIResponse.Payload = localVarHttpResponse.Body()
	}

	if err != nil {
		return localVarAPIResponse, err
	}
	return localVarAPIResponse, err
}

/**
 * Retrieve details of an access token, such as scope, expiration and extension ID.
 * Retrieve details of an access token, such as scope, expiration and extension ID. Voip ID will be returned as well if scope contains account-owner scope.
 *
 * @return *GetOauthAccessToken
 */
func (a OauthApi) GetOauthAccessToken() (*GetOauthAccessToken, *APIResponse, error) {

	var localVarHttpMethod = strings.ToUpper("Get")
	// create path and map variables
	localVarPath := a.Configuration.BasePath + "/oauth/access-token"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make(map[string]string)
	var localVarPostBody interface{}
	var localVarFileName string
	var localVarFileBytes []byte
	// authentication '(apiKey)' required
	// set key with prefix in header
	localVarHeaderParams["Authorization"] = a.Configuration.GetAPIKeyWithPrefix("Authorization")
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		localVarHeaderParams[key] = a.Configuration.DefaultHeader[key]
	}

	clearEmptyParams(localVarQueryParams)

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	var successPayload = new(GetOauthAccessToken)
	localVarHttpResponse, err := a.Configuration.APIClient.CallAPI(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)

	var localVarURL, _ = url.Parse(localVarPath)
	localVarURL.RawQuery = localVarQueryParams.Encode()
	var localVarAPIResponse = &APIResponse{Operation: "GetOauthAccessToken", Method: localVarHttpMethod, RequestURL: localVarURL.String()}
	if localVarHttpResponse != nil {
		localVarAPIResponse.Response = localVarHttpResponse.RawResponse
		localVarAPIResponse.Payload = localVarHttpResponse.Body()
	}

	if err != nil {
		return successPayload, localVarAPIResponse, err
	}
	err = json.Unmarshal(localVarHttpResponse.Body(), &successPayload)
	return successPayload, localVarAPIResponse, err
}

