package resources

import (
	"context"
	"fmt"

	"github.com/Azure/go-autorest/autorest/to"

	"github.com/Azure-Samples/azure-sdk-for-go-samples/internal/config"
	"github.com/Azure-Samples/azure-sdk-for-go-samples/internal/iam"
	"github.com/Azure/azure-sdk-for-go/services/resources/mgmt/2018-06-01/managedapplications"
)

func getManagedAppClient() managedapplications.ApplicationsClient {
	managedAppClient := managedapplications.NewApplicationsClient(config.SubscriptionID())
	a, _ := iam.GetResourceManagementAuthorizer()
	managedAppClient.Authorizer = a
	managedAppClient.AddToUserAgent(config.UserAgent())
	return managedAppClient
}

type Plan struct {
	Name          string
	Product       string
	PromotionCode string
	Publisher     string
	Version       string
}

// CreateManagedApp creates a managed app
func CreateManagedApp(ctx context.Context, resourceGroupName, applicationName, location, managedResourceGroupID string, plan Plan) (m managedapplications.Application, err error) {
	managedAppClient := getManagedAppClient()

	managedAppParameters := managedapplications.Application{
		Name:     &applicationName,
		Kind:     to.StringPtr("MarketPlace"),
		Location: &location,
		Plan: &managedapplications.Plan{
			Name:          &plan.Name,
			Product:       &plan.Product,
			PromotionCode: &plan.PromotionCode,
			Publisher:     &plan.Publisher,
			Version:       &plan.Version,
		},
		ApplicationProperties: &managedapplications.ApplicationProperties{
			ManagedResourceGroupID: &managedResourceGroupID,
		},
	}

	future, err := managedAppClient.CreateOrUpdate(
		ctx,
		resourceGroupName,
		applicationName,
		managedAppParameters,
	)
	if err != nil {
		return m, fmt.Errorf("cannot create appliance: %v", err)
	}

	err = future.WaitForCompletion(ctx, managedAppClient.Client)
	if err != nil {
		return m, fmt.Errorf("cannot get the create appliance future response: %v", err)
	}

	return future.Result(managedAppClient)
}
