// Package resources contains utilities for azure sdk resource handling
package resources

import (
	"context"
	"fmt"

	"github.com/Azure/azure-sdk-for-go/profiles/latest/resources/mgmt/resources"
	"github.com/Azure/go-autorest/autorest"
)

// GetDeploymentsClient takes the azure authorizer and creates an ARM deployments client on the desired subscription
func GetDeploymentsClient(subscriptionID string, authorizer autorest.Authorizer) resources.DeploymentsClient {
	deployClient := resources.NewDeploymentsClient(subscriptionID)
	deployClient.Authorizer = authorizer
	return deployClient
}

// ValidateDeployment validates the template deployments and their
// parameters are correct and will produce a successful deployment.GetResource
func ValidateDeployment(ctx context.Context, client resources.DeploymentsClient, resourceGroupName, deploymentName string, deploymentMode string, template, params map[string]interface{}) (valid resources.DeploymentValidateResult, err error) {
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
		return valid, fmt.Errorf("cannot validate deployment: %v", err)
	}

	err = future.WaitForCompletionRef(ctx, client.Client)
	if err != nil {
		return valid, fmt.Errorf("cannot get the validate deployment future response: %v", err)
	}

	return future.Result(client)
}

// ValidateDeploymentAtSubscriptionScope validates the template deployments and their
// parameters are correct and will produce a successful deployment.GetResource (at subscription scope)
func ValidateDeploymentAtSubscriptionScope(ctx context.Context, client resources.DeploymentsClient, deploymentName string, deploymentMode string, template, params map[string]interface{}) (valid resources.DeploymentValidateResult, err error) {
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
		return valid, fmt.Errorf("cannot validate deployment: %v", err)
	}

	err = future.WaitForCompletionRef(ctx, client.Client)
	if err != nil {
		return valid, fmt.Errorf("cannot get the validate deployment future response: %v", err)
	}

	return future.Result(client)
}

// ValidateDeploymentAtManagementGroupScope validates the template deployments and their
// parameters are correct and will produce a successful deployment.GetResource (at management group scope)
func ValidateDeploymentAtManagementGroupScope(ctx context.Context, client resources.DeploymentsClient, managementGroupId string, deploymentName string, deploymentMode string, template, params map[string]interface{}) (valid resources.DeploymentValidateResult, err error) {
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
		return valid, fmt.Errorf("cannot validate deployment: %v", err)
	}

	err = future.WaitForCompletionRef(ctx, client.Client)
	if err != nil {
		return valid, fmt.Errorf("cannot get the validate deployment future response: %v", err)
	}

	return future.Result(client)
}

// CreateDeployment creates a template deployment using the
// referenced JSON files for the template and its parameters
func CreateDeployment(ctx context.Context, client resources.DeploymentsClient, resourceGroupName, deploymentName string, deploymentMode string, template, params map[string]interface{}) (de resources.DeploymentExtended, err error) {
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
		return de, fmt.Errorf("cannot create deployment: %v", err)
	}

	err = future.WaitForCompletionRef(ctx, client.Client)
	if err != nil {
		return de, fmt.Errorf("cannot get the create deployment future response: %v", err)
	}

	return future.Result(client)
}

// CreateDeploymentAtSubscriptionScope creates a template deployment using the
// referenced JSON files for the template and its parameters (at subscription scope)
func CreateDeploymentAtSubscriptionScope(ctx context.Context, client resources.DeploymentsClient, deploymentName, deploymentMode string, template, params map[string]interface{}) (de resources.DeploymentExtended, err error) {
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
		return de, fmt.Errorf("cannot create deployment: %v", err)
	}

	err = future.WaitForCompletionRef(ctx, client.Client)
	if err != nil {
		return de, fmt.Errorf("cannot get the create deployment future response: %v", err)
	}

	return future.Result(client)
}

// CreateDeploymentAtManagementGroupScope creates a template deployment using the
// referenced JSON files for the template and its parameters (at management group scope)
func CreateDeploymentAtManagementGroupScope(ctx context.Context, client resources.DeploymentsClient, managementGroupId string, deploymentName string, deploymentMode string, template, params map[string]interface{}) (valid resources.DeploymentExtended, err error) {
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
		return valid, fmt.Errorf("cannot validate deployment: %v", err)
	}

	err = future.WaitForCompletionRef(ctx, client.Client)
	if err != nil {
		return valid, fmt.Errorf("cannot get the validate deployment future response: %v", err)
	}

	return future.Result(client)
}