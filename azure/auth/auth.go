/*
Copyright (c) 2021 white duck Gesellschaft für Softwareentwicklung mbH

This code is licensed under MIT license (see LICENSE for details)
*/
package auth

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/adal"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/azure/auth"
	"github.com/Azure/go-autorest/autorest/azure/cli"
)

// SDKAuth represents an Azure AD Service Principal
type SDKAuth struct {
	ClientID       string `json:"clientId"`
	ClientSecret   string `json:"clientSecret"`
	SubscriptionID string `json:"subscriptionId"`
	TenantID       string `json:"tenantId"`
	ARMEndpointURL string `json:"resourceManagerEndpointUrl"`
	ADEndpointURL  string `json:"activeDirectoryEndpointUrl"`
}

// FromString fills the struct with the information from the input json string
func (auth *SDKAuth) FromString(credentials string) error {
	err := json.Unmarshal([]byte(credentials), auth)
	if err != nil {
		return fmt.Errorf("failed to parse the credentials passed, marshal error: %s", err)
	}

	return nil
}

// GetResourceManagerAuthorizer builds an autorest.Authorizer for the Azure Resource Manager using the given credentials
func (auth SDKAuth) GetResourceManagerAuthorizer() (autorest.Authorizer, error) {
	// If the Active Directory Endpoint is not set, fallback to the default public cloud endpoint
	if len(auth.ADEndpointURL) == 0 {
		auth.ADEndpointURL = azure.PublicCloud.ActiveDirectoryEndpoint
	}

	oauthConfig, err := adal.NewOAuthConfig(auth.ADEndpointURL, auth.TenantID)
	if err != nil {
		return nil, err
	}

	// If the Resource Manager Endpoint is not set, fallback to the default public cloud endpoint
	if len(auth.ARMEndpointURL) == 0 {
		auth.ARMEndpointURL = azure.PublicCloud.ResourceManagerEndpoint
	}

	token, err := adal.NewServicePrincipalToken(*oauthConfig, auth.ClientID, auth.ClientSecret, auth.ARMEndpointURL)
	if err != nil {
		return nil, err
	}

	// Create authorizer from the bearer token
	var authorizer autorest.Authorizer
	authorizer = autorest.NewBearerAuthorizer(token)

	return authorizer, nil
}

// GetSdkAuthFromString builds from the cmd flags a ServicePrincipal
// Deprecated: Use SDKAuth.FromString instead.
func GetSdkAuthFromString(credentials string) (SDKAuth, error) {
	var sdkAuth SDKAuth
	if err := sdkAuth.FromString(credentials); err != nil {
		return SDKAuth{}, err
	}

	return sdkAuth, nil
}

// GetArmAuthorizerFromSdkAuth creates an ARM authorizer from an Sp
// Deprecated: Use SDKAuth.GetResourceManagerAuthorizer instead.
func GetArmAuthorizerFromSdkAuth(auth SDKAuth) (autorest.Authorizer, error) {
	return auth.GetResourceManagerAuthorizer()
}

// GetArmAuthorizerFromSdkAuthJSON creats am ARM authorizer from the passed sdk auth file
func GetArmAuthorizerFromSdkAuthJSON(path string, resourceManagerEndpoint string) (autorest.Authorizer, error) {
	var authorizer autorest.Authorizer

	// Manipulate the AZURE_AUTH_LOCATION var at runtime
	os.Setenv("AZURE_AUTH_LOCATION", path)
	defer os.Unsetenv("AZURE_AUTH_LOCATION")

	authorizer, err := auth.NewAuthorizerFromFile(resourceManagerEndpoint)
	return authorizer, err
}

// GetArmAuthorizerFromSdkAuthJSONString creates an ARM authorizer from the sdk auth credentials
func GetArmAuthorizerFromSdkAuthJSONString(credentials string, resourceManagerEndpoint string) (autorest.Authorizer, error) {
	var authorizer autorest.Authorizer

	// create a temporary file, as the sdk credentials need to be read from a file
	tmpFile, err := ioutil.TempFile(os.TempDir(), "azure-sdk-auth-")
	if err != nil {
		return authorizer, fmt.Errorf("Cannot create temporary sdk auth file: %s", err)
	}
	defer os.Remove(tmpFile.Name())

	text := []byte(credentials)
	if _, err = tmpFile.Write(text); err != nil {
		return authorizer, fmt.Errorf("Failed to write to temporary sdk auth file: %s", err)
	}
	tmpFile.Close()

	// Manipulate the AZURE_AUTH_LOCATION var at runtime
	os.Setenv("AZURE_AUTH_LOCATION", tmpFile.Name())
	defer os.Unsetenv("AZURE_AUTH_LOCATION")

	authorizer, err = auth.NewAuthorizerFromFile(resourceManagerEndpoint)

	return authorizer, err
}

// GetArmAuthorizerFromEnvironment creates an ARM authorizer from a MSI (e.g. AAD Pod Identity)
func GetArmAuthorizerFromEnvironment() (autorest.Authorizer, error) {
	var authorizer autorest.Authorizer
	authorizer, err := auth.NewAuthorizerFromEnvironment()

	return authorizer, err
}

// GetArmAuthorizerFromCLI creates an ARM authorizer from the local azure cli
func GetArmAuthorizerFromCLI(params cli.GetAccessTokenParams) (autorest.Authorizer, error) {
	token, err := cli.GetTokenFromCLIWithParams(params)
	if err != nil {
		return nil, err
	}

	adalToken, err := token.ToADALToken()
	if err != nil {
		return nil, err
	}

	return autorest.NewBearerAuthorizer(&adalToken), nil
}