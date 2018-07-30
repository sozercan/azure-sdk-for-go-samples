package resources

import (
	"context"
	"time"

	"github.com/Azure-Samples/azure-sdk-for-go-samples/internal/config"
	"github.com/Azure-Samples/azure-sdk-for-go-samples/internal/util"
)

func ExampleCreateManagedApp() {
	var groupName = config.GenerateGroupName("CreateManagedApp")
	config.SetGroupName(groupName)

	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(time.Hour*1))
	defer cancel()
	defer Cleanup(ctx)

	_, err := CreateGroup(ctx, config.GroupName())
	if err != nil {
		util.PrintAndLog(err.Error())
	}

	applianceName := "sertac-managedApp-test"
	managedResourceGroupID := "/subscriptions//resourceGroups/sertac-managedApp-test"
	applianceDefinitionID := "/subscriptions//resourceGroups/sertac-managedk8s-masters/providers/Microsoft.Solutions/applicationDefinitions/ManagedACSEngineMasters"
	parameters := ""

	_, err = CreateManagedApp(ctx, config.GroupName(), applianceName, config.Location(), managedResourceGroupID, applianceDefinitionID, parameters)
	if err != nil {
		util.PrintAndLog(err.Error())
	}
	util.PrintAndLog("created managed app")

	// Output:
	// created managed app
}
