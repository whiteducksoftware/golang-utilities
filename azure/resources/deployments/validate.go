package deployments

import (
	"context"
	"fmt"

	"github.com/Azure/azure-sdk-for-go/profiles/latest/resources/mgmt/resources"
)

// Validate validates the template deployments and their
// parameters are correct and will produce a successful deployment.GetResource
func Validate(ctx context.Context, client resources.DeploymentsClient, resourceGroupName, deploymentName string, deploymentMode string, template, params map[string]interface{}) (resources.DeploymentValidateResult, error) {
	future, err := client.Validate(
		ctx,
		resourceGroupName,
		deploymentName,
		resources.Deployment{
			Properties: &resources.DeploymentProperties{
				Template:   template,
				Parameters: params,
				Mode:       resources.DeploymentMode(deploymentMode),
			},
		})

	if err != nil {
		return resources.DeploymentValidateResult{}, fmt.Errorf("cannot validate deployment: %v", err)
	}

	err = future.WaitForCompletionRef(ctx, client.Client)
	if err != nil {
		return resources.DeploymentValidateResult{}, fmt.Errorf("cannot get the validate deployment future response: %v", err)
	}

	return future.Result(client)
}

// ValidateAtSubscriptionScope validates the template deployments and their
// parameters are correct and will produce a successful deployment.GetResource (at subscription scope)
func ValidateAtSubscriptionScope(ctx context.Context, client resources.DeploymentsClient, deploymentName string, deploymentMode string, template, params map[string]interface{}) (resources.DeploymentValidateResult, error) {
	future, err := client.ValidateAtSubscriptionScope(
		ctx,
		deploymentName,
		resources.Deployment{
			Properties: &resources.DeploymentProperties{
				Template:   template,
				Parameters: params,
				Mode:       resources.DeploymentMode(deploymentMode),
			},
		})

	if err != nil {
		return resources.DeploymentValidateResult{}, fmt.Errorf("cannot validate deployment: %v", err)
	}

	err = future.WaitForCompletionRef(ctx, client.Client)
	if err != nil {
		return resources.DeploymentValidateResult{}, fmt.Errorf("cannot get the validate deployment future response: %v", err)
	}

	return future.Result(client)
}

// ValidateAtManagementGroupScope validates the template deployments and their
// parameters are correct and will produce a successful deployment.GetResource (at management group scope)
func ValidateAtManagementGroupScope(ctx context.Context, client resources.DeploymentsClient, managementGroupId string, deploymentName string, deploymentMode string, template, params map[string]interface{}) (resources.DeploymentValidateResult, error) {
	future, err := client.ValidateAtManagementGroupScope(
		ctx,
		managementGroupId,
		deploymentName,
		resources.ScopedDeployment{
			Properties: &resources.DeploymentProperties{
				Template:   template,
				Parameters: params,
				Mode:       resources.DeploymentMode(deploymentMode),
			},
		})

	if err != nil {
		return resources.DeploymentValidateResult{}, fmt.Errorf("cannot validate deployment: %v", err)
	}

	err = future.WaitForCompletionRef(ctx, client.Client)
	if err != nil {
		return resources.DeploymentValidateResult{}, fmt.Errorf("cannot get the validate deployment future response: %v", err)
	}

	return future.Result(client)
}
