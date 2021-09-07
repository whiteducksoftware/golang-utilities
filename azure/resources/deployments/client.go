// Package resources contains utilities for azure sdk resource handling
package deployments

import (
	"github.com/Azure/azure-sdk-for-go/profiles/latest/resources/mgmt/resources"
	"github.com/Azure/go-autorest/autorest"
)

// GetClient takes the azure authorizer and creates an ARM deployments client on the desired subscription
func GetClient(subscriptionID string, authorizer autorest.Authorizer) resources.DeploymentsClient {
	deployClient := resources.NewDeploymentsClient(subscriptionID)
	deployClient.Authorizer = authorizer

	return deployClient
}

// GetClientWithBaseUri takes the azure authorizer and creates an ARM deployments client on the desired subscription
func GetClientWithBaseUri(baseUri string, subscriptionID string, authorizer autorest.Authorizer) resources.DeploymentsClient {
	deployClient := resources.NewDeploymentsClientWithBaseURI(baseUri, subscriptionID)
	deployClient.Authorizer = authorizer

	return deployClient
}
