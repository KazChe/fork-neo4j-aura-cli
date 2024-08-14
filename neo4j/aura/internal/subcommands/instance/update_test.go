package instance_test

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/neo4j/cli/neo4j/aura/internal/test/testutils"
)

func TestUpdateMemory(t *testing.T) {
	helper := testutils.NewAuraTestHelper(t)
	defer helper.Close()

	instanceId := "2f49c2b3"

	mockHandler := helper.NewRequestHandlerMock(fmt.Sprintf("/v1/instances/%s", instanceId), http.StatusAccepted, `{
		"data": {
			"id": "2f49c2b3",
			"name": "Production",
			"status": "updating",
			"connection_url": "YOUR_CONNECTION_URL",
			"tenant_id": "YOUR_TENANT_ID",
			"cloud_provider": "gcp",
			"memory": "8GB",
			"region": "europe-west1",
			"type": "enterprise-db"
		}
	}`)

	helper.ExecuteCommand(fmt.Sprintf("instance update %s --memory 8GB", instanceId))

	mockHandler.AssertCalledTimes(1)
	mockHandler.AssertCalledWithMethod(http.MethodPatch)
	mockHandler.AssertCalledWithBody(`{"memory":"8GB"}`)

	helper.AssertOutJson(`{
		"data": {
			"id": "2f49c2b3",
			"name": "Production",
			"status": "updating",
			"connection_url": "YOUR_CONNECTION_URL",
			"tenant_id": "YOUR_TENANT_ID",
			"cloud_provider": "gcp",
			"memory": "8GB",
			"region": "europe-west1",
			"type": "enterprise-db"
		}
	}
	`)
}

func TestUpdateName(t *testing.T) {
	helper := testutils.NewAuraTestHelper(t)
	defer helper.Close()

	instanceId := "2f49c2b3"

	mockHandler := helper.NewRequestHandlerMock(fmt.Sprintf("/v1/instances/%s", instanceId), http.StatusOK, `{
		"data": {
			"id": "2f49c2b3",
			"name": "New Name",
			"status": "updating",
			"connection_url": "YOUR_CONNECTION_URL",
			"tenant_id": "YOUR_TENANT_ID",
			"cloud_provider": "gcp",
			"memory": "4GB",
			"region": "europe-west1",
			"type": "enterprise-db"
		}
	}`)

	helper.ExecuteCommand(fmt.Sprintf(`instance update %s --name "New Name"`, instanceId))

	mockHandler.AssertCalledTimes(1)
	mockHandler.AssertCalledWithMethod(http.MethodPatch)
	mockHandler.AssertCalledWithBody(`{"name":"New Name"}`)

	helper.AssertOutJson(`{
		"data": {
			"id": "2f49c2b3",
			"name": "New Name",
			"status": "updating",
			"connection_url": "YOUR_CONNECTION_URL",
			"tenant_id": "YOUR_TENANT_ID",
			"cloud_provider": "gcp",
			"memory": "4GB",
			"region": "europe-west1",
			"type": "enterprise-db"
		}
	}
	`)
}

func TestUpdateMemoryAndName(t *testing.T) {
	helper := testutils.NewAuraTestHelper(t)
	defer helper.Close()

	instanceId := "2f49c2b3"

	mockHandler := helper.NewRequestHandlerMock(fmt.Sprintf("/v1/instances/%s", instanceId), http.StatusAccepted, `{
		"data": {
			"id": "2f49c2b3",
			"name": "New Name",
			"status": "updating",
			"connection_url": "YOUR_CONNECTION_URL",
			"tenant_id": "YOUR_TENANT_ID",
			"cloud_provider": "gcp",
			"memory": "8GB",
			"region": "europe-west1",
			"type": "enterprise-db"
		}
	}`)

	helper.ExecuteCommand(fmt.Sprintf(`instance update %s --name "New Name" --memory 8GB`, instanceId))

	mockHandler.AssertCalledTimes(1)
	mockHandler.AssertCalledWithMethod(http.MethodPatch)
	mockHandler.AssertCalledWithBody(`{"memory":"8GB","name":"New Name"}`)

	helper.AssertOutJson(`{
		"data": {
			"id": "2f49c2b3",
			"name": "New Name",
			"status": "updating",
			"connection_url": "YOUR_CONNECTION_URL",
			"tenant_id": "YOUR_TENANT_ID",
			"cloud_provider": "gcp",
			"memory": "8GB",
			"region": "europe-west1",
			"type": "enterprise-db"
		}
	}
	`)
}

func TestUpdateErrorsWithNoFlags(t *testing.T) {
	helper := testutils.NewAuraTestHelper(t)
	defer helper.Close()

	instanceId := "2f49c2b3"

	mockHandler := helper.NewRequestHandlerMock(fmt.Sprintf("/v1/instances/%s", instanceId), http.StatusAccepted, "")

	helper.ExecuteCommand(fmt.Sprintf(`instance update %s`, instanceId))

	mockHandler.AssertCalledTimes(0)

	helper.AssertErr(`Error: at least one of the flags in the group [memory name] is required
`)
	helper.AssertOut(`Usage:
  aura instance update [flags]

Flags:
  -h, --help            help for update
      --memory string   The size of the instance memory in GB.
      --name string     The name of the instance (any UTF-8 characters with no trailing or leading whitespace).

Global Flags:
      --auth-url string   
      --base-url string   
      --output string

`)
}
