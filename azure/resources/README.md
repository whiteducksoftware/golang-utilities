# Azure / Resources
Installation: `go get github.com/whiteducksoftware/golang-utilities/azure/resources`

## Types & Functions
```go
package resources

// GetDeploymentsClient takes the azure authorizer and creates an ARM deployments client on the desired subscription
func GetDeploymentsClient(subscriptionID string, authorizer autorest.Authorizer) resources.DeploymentsClient {}

// ValidateDeployment validates the template deployments and their parameters are correct and will produce a successful deployment.GetResource
func ValidateDeployment(ctx context.Context, deployClient resources.DeploymentsClient, resourceGroupName, deploymentName string, deploymentMode string, template, params map[string]interface{}) (valid resources.DeploymentValidateResult, err error) {}

// ValidateDeploymentAtSubscriptionScope validates the template deployments and their parameters are correct and will produce a successful deployment.GetResource (at subscription scope)
func ValidateDeploymentAtSubscriptionScope(ctx context.Context, deployClient resources.DeploymentsClient, deploymentName string, deploymentMode string, template, params map[string]interface{}) (valid resources.DeploymentValidateResult, err error) {}

// CreateDeployment creates a template deployment using the referenced JSON files for the template and its parameters
func CreateDeployment(ctx context.Context, deployClient resources.DeploymentsClient, resourceGroupName, deploymentName string, deploymentMode string, template, params map[string]interface{}) (de resources.DeploymentExtended, err error) {}

// CreateDeploymentAtSubscriptionScope creates a template deployment using the referenced JSON files for the template and its parameters (at subscription scope)
func CreateDeploymentAtSubscriptionScope(ctx context.Context, deployClient resources.DeploymentsClient, deploymentName, deploymentMode string, template, params map[string]interface{}) (de resources.DeploymentExtended, err error) {}
```