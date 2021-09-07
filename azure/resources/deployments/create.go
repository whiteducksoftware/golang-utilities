package deployments

import (
	"context"
	"fmt"

	"github.com/Azure/azure-sdk-for-go/profiles/latest/resources/mgmt/resources"
)

// Create creates a template deployment using the
// referenced JSON files for the template and its parameters
func Create(ctx context.Context, client resources.DeploymentsClient, resourceGroupName, deploymentName string, deploymentMode string, template, params map[string]interface{}) (resources.DeploymentExtended, error) {
	future, err := client.CreateOrUpdate(
		ctx,
		resourceGroupName,
		deploymentName,
		resources.Deployment{
			Properties: &resources.DeploymentProperties{
				Template:   template,
				Parameters: params,
				Mode:       resources.DeploymentMode(deploymentMode),
			},
		},
	)
	if err != nil {
		return resources.DeploymentExtended{}, fmt.Errorf("cannot create deployment: %v", err)
	}

	err = future.WaitForCompletionRef(ctx, client.Client)
	if err != nil {
		return resources.DeploymentExtended{}, fmt.Errorf("cannot get the create deployment future response: %v", err)
	}

	return future.Result(client)
}

// CreateAtSubscriptionScope creates a template deployment using the
// referenced JSON files for the template and its parameters (at subscription scope)
func CreateAtSubscriptionScope(ctx context.Context, client resources.DeploymentsClient, deploymentName, deploymentMode string, template, params map[string]interface{}) (resources.DeploymentExtended, error) {
	future, err := client.CreateOrUpdateAtSubscriptionScope(
		ctx,
		deploymentName,
		resources.Deployment{
			Properties: &resources.DeploymentProperties{
				Template:   template,
				Parameters: params,
				Mode:       resources.DeploymentMode(deploymentMode),
			},
		},
	)

	if err != nil {
		return resources.DeploymentExtended{}, fmt.Errorf("cannot create deployment: %v", err)
	}

	err = future.WaitForCompletionRef(ctx, client.Client)
	if err != nil {
		return resources.DeploymentExtended{}, fmt.Errorf("cannot get the create deployment future response: %v", err)
	}

	return future.Result(client)
}

// CreateAtManagementGroupScope creates a template deployment using the
// referenced JSON files for the template and its parameters (at management group scope)
func CreateAtManagementGroupScope(ctx context.Context, client resources.DeploymentsClient, managementGroupId string, deploymentName string, deploymentMode string, template, params map[string]interface{}) (resources.DeploymentExtended, error) {
	future, err := client.CreateOrUpdateAtManagementGroupScope(
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
		return resources.DeploymentExtended{}, fmt.Errorf("cannot validate deployment: %v", err)
	}

	err = future.WaitForCompletionRef(ctx, client.Client)
	if err != nil {
		return resources.DeploymentExtended{}, fmt.Errorf("cannot get the validate deployment future response: %v", err)
	}

	return future.Result(client)
}
