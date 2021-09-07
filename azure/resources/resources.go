/*
Copyright (c) 2021 white duck Gesellschaft für Softwareentwicklung mbH

This code is licensed under MIT license (see LICENSE for details)
*/

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
func ValidateDeployment(ctx context.Context, deployClient resources.DeploymentsClient, resourceGroupName, deploymentName string, deploymentMode string, template, params map[string]interface{}) (valid resources.DeploymentValidateResult, err error) {
	future, err := deployClient.Validate(
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

	err = future.WaitForCompletionRef(ctx, deployClient.Client)
	if err != nil {
		return valid, fmt.Errorf("cannot get the validate deployment future response: %v", err)
	}

	return future.Result(deployClient)
}

// ValidateDeploymentAtSubscriptionScope validates the template deployments and their
// parameters are correct and will produce a successful deployment.GetResource (at subscription scope)
func ValidateDeploymentAtSubscriptionScope(ctx context.Context, deployClient resources.DeploymentsClient, deploymentName string, deploymentMode string, template, params map[string]interface{}) (valid resources.DeploymentValidateResult, err error) {
	future, err := deployClient.ValidateAtSubscriptionScope(
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

	err = future.WaitForCompletionRef(ctx, deployClient.Client)
	if err != nil {
		return valid, fmt.Errorf("cannot get the validate deployment future response: %v", err)
	}

	return future.Result(deployClient)
}

// CreateDeployment creates a template deployment using the
// referenced JSON files for the template and its parameters
func CreateDeployment(ctx context.Context, deployClient resources.DeploymentsClient, resourceGroupName, deploymentName string, deploymentMode string, template, params map[string]interface{}) (de resources.DeploymentExtended, err error) {
	future, err := deployClient.CreateOrUpdate(
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

	err = future.WaitForCompletionRef(ctx, deployClient.Client)
	if err != nil {
		return de, fmt.Errorf("cannot get the create deployment future response: %v", err)
	}

	return future.Result(deployClient)
}

// CreateDeploymentAtSubscriptionScope creates a template deployment using the
// referenced JSON files for the template and its parameters (at subscription scope)
func CreateDeploymentAtSubscriptionScope(ctx context.Context, deployClient resources.DeploymentsClient, deploymentName, deploymentMode string, template, params map[string]interface{}) (de resources.DeploymentExtended, err error) {
	future, err := deployClient.CreateOrUpdateAtSubscriptionScope(
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

	err = future.WaitForCompletionRef(ctx, deployClient.Client)
	if err != nil {
		return de, fmt.Errorf("cannot get the create deployment future response: %v", err)
	}

	return future.Result(deployClient)
}
