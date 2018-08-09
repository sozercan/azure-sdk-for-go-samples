package resources

import (
	"context"
	"flag"
	"log"
	"os"
	"testing"
	"time"

	"github.com/Azure-Samples/azure-sdk-for-go-samples/internal/config"
	"github.com/Azure-Samples/azure-sdk-for-go-samples/internal/util"
)

func TestMain(m *testing.M) {
	err := setupEnvironment()
	if err != nil {
		log.Fatalf("could not set up environment: %v\n", err)
	}

	os.Exit(m.Run())
}

func setupEnvironment() error {
	err1 := config.ParseEnvironment()
	err2 := config.AddFlags()
	err3 := addLocalConfig()

	for _, err := range []error{err1, err2, err3} {
		if err != nil {
			return err
		}
	}

	flag.Parse()
	return nil
}

func addLocalConfig() error {
	return nil
}

func ExampleCreateManagedApp() {
	var groupName = config.GenerateGroupName("CreateManagedApp")
	config.SetGroupName(groupName)

	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(time.Hour*1))
	defer cancel()
	//defer Cleanup(ctx)

	_, err := CreateGroup(ctx, config.GroupName())
	if err != nil {
		util.PrintAndLog(err.Error())
	}

	applianceName := "<NAME>"
	managedResourceGroupID := "/subscriptions/<SUB_ID>/resourceGroups/<RG_NAME>"
	plan := Plan{
		Name:      "",
		Product:   "",
		Publisher: "",
		Version:   "0.0.1",
	}

	_, err = CreateManagedApp(ctx, config.GroupName(), applianceName, config.Location(), managedResourceGroupID, plan)
	if err != nil {
		util.PrintAndLog(err.Error())
	}
	util.PrintAndLog("created managed app")

	// Output:
	// created managed app
}
