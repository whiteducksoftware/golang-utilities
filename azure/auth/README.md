# Azure / Auth
Installation: `go get github.com/whiteducksoftware/golang-utilities/azure/auth`

## Types & Functions
```go
package auth

// SDKAuth represents an Azure AD Service Principal
type SDKAuth struct {}

// FromString fills the struct with the information from the input json string
func (auth *SDKAuth) FromString(credentials string) error {}

// GetResourceManagerAuthorizer builds an autorest.Authorizer for the Azure Resource Manager using the given credentials
func (auth *SDKAuth) GetResourceManagerAuthorizer() (autorest.Authorizer, error) {}

// GetSdkAuthFromString builds from the cmd flags a ServicePrincipal
// Deprecated: Use SDKAuth.FromString instead.
func GetSdkAuthFromString(credentials string) (SDKAuth, error) {}

// GetArmAuthorizerFromSdkAuth creates an ARM authorizer from an Sp
// Deprecated: Use SDKAuth.GetResourceManagerAuthorizer instead.
func GetArmAuthorizerFromSdkAuth(auth SDKAuth) (autorest.Authorizer, error) {}

// GetArmAuthorizerFromSdkAuthJSON creats am ARM authorizer from the passed sdk auth file
func GetArmAuthorizerFromSdkAuthJSON(path string, resourceManagerEndpoint string) (autorest.Authorizer, error) {}

// GetArmAuthorizerFromSdkAuthJSONString creates an ARM authorizer from the sdk auth credentials
func GetArmAuthorizerFromSdkAuthJSONString(credentials string, resourceManagerEndpoint string) (autorest.Authorizer, error) {}

// GetArmAuthorizerFromEnvironment creates an ARM authorizer from a MSI (e.g. AAD Pod Identity)
func GetArmAuthorizerFromEnvironment() (autorest.Authorizer, error) {}

// GetArmAuthorizerFromCLI creates an ARM authorizer from the local azure cli
func GetArmAuthorizerFromCLI(params cli.GetAccessTokenParams) (autorest.Authorizer, error) {}
```