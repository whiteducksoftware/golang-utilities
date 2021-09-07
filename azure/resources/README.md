# Azure / Resources
Installation: `go get github.com/whiteducksoftware/golang-utilities/azure/resources`

## Types & Functions
```go
package deployments

// GetClient takes the azure authorizer and creates an ARM deployments client on the desired subscription
func GetClient(subscriptionID string, authorizer autorest.Authorizer) resources.DeploymentsClient {}

// Validate validates the template deployments and their parameters are correct and will produce a successful deployment.GetResource
func Validate(ctx context.Context, client resources.DeploymentsClient, resourceGroupName, deploymentName string, deploymentMode string, template, params map[string]interface{}) (resources.DeploymentValidateResult, error) {}

// ValidateAtSubscriptionScope validates the template deployments and their parameters are correct and will produce a successful deployment.GetResource (at subscription scope)
func ValidateAtSubscriptionScope(ctx context.Context, client resources.DeploymentsClient, deploymentName string, deploymentMode string, template, params map[string]interface{}) (resources.DeploymentValidateResult, error) {}

// ValidateAtManagementGroupScope validates the template deployments and their parameters are correct and will produce a successful deployment.GetResource (at management group scope)
func ValidateAtManagementGroupScope(ctx context.Context, client resources.DeploymentsClient, managementGroupId string, deploymentName string, deploymentMode string, template, params map[string]interface{}) (resources.DeploymentValidateResult, error) {}

// Create creates a template deployment using the referenced JSON files for the template and its parameters
func Create(ctx context.Context, client resources.DeploymentsClient, resourceGroupName, deploymentName string, deploymentMode string, template, params map[string]interface{}) (resources.DeploymentExtended, error) {}

// CreatAtSubscriptionScope creates a template deployment using the referenced JSON files for the template and its parameters (at subscription scope)
func CreatAtSubscriptionScope(ctx context.Context, client resources.DeploymentsClient, deploymentName, deploymentMode string, template, params map[string]interface{}) (resources.DeploymentExtended, error) {}

// CreateAtManagementGroupScope creates a template deployment using the referenced JSON files for the template and its parameters (at management group scope)
func CreateAtManagementGroupScope(ctx context.Context, client resources.DeploymentsClient, managementGroupId string, deploymentName string, deploymentMode string, template, params map[string]interface{}) (resources.DeploymentExtended, error) {}
```