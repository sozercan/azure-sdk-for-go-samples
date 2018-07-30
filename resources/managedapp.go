package resources

import (
	"context"
	"fmt"

	"github.com/Azure/go-autorest/autorest/to"

	"github.com/Azure-Samples/azure-sdk-for-go-samples/internal/config"
	"github.com/Azure-Samples/azure-sdk-for-go-samples/internal/iam"
	"github.com/Azure/azure-sdk-for-go/services/resources/mgmt/2016-09-01-preview/managedapplications"
)

func getManagedAppClient() managedapplications.AppliancesClient {
	managedAppClient := managedapplications.NewAppliancesClient(config.SubscriptionID())
	a, _ := iam.GetResourceManagementAuthorizer()
	managedAppClient.Authorizer = a
	managedAppClient.AddToUserAgent(config.UserAgent())
	return managedAppClient
}

// CreateManagedApp creates a managed app
func CreateManagedApp(ctx context.Context, resourceGroupName, applianceName, location, managedResourceGroupID, applianceDefinitionID string, parameters interface{}) (m managedapplications.Appliance, err error) {
	managedAppClient := getManagedAppClient()

	managedAppParameters := managedapplications.Appliance{
		Name:     &applianceName,
		Kind:     to.StringPtr("ServiceCatalog"),
		Location: &location,
		ApplianceProperties: &managedapplications.ApplianceProperties{
			ManagedResourceGroupID: &managedResourceGroupID,
			ApplianceDefinitionID:  &applianceDefinitionID,
			Parameters:             &parameters,
		},
	}

	future, err := managedAppClient.CreateOrUpdate(
		ctx,
		resourceGroupName,
		applianceName,
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
