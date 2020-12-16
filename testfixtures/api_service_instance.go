package testfixtures

import "fmt"

const V3ServiceInstanceGuid = "e9aca2a7-f4f9-431e-924c-664ddc236604"

var V3ServiceInstancePath = fmt.Sprintf("/v3/service_instances/%s", V3ServiceInstanceGuid)
var V3ServiceInstance = fmt.Sprintf(`
{
	"guid": "%[1]s",
	"created_at": "2020-03-10T15:49:29Z",
	"updated_at": "2020-03-10T15:49:29Z",
	"name": "my-managed-instance",
	"tags": [],
	"type": "managed",
	"maintenance_info": {
	  "version": "1.0.0"
	},
	"upgrade_available": false,
	"dashboard_url": "https://service-broker.example.org/dashboard",
	"last_operation": {
	  "type": "create",
	  "state": "succeeded",
	  "description": "Operation succeeded",
	  "updated_at": "2020-03-10T15:49:32Z",
	  "created_at": "2020-03-10T15:49:29Z"
	},
	"relationships": {
	  "service_plan": {
		"data": {
		  "guid": "5358d122-638e-11ea-afca-bf6e756684ac"
		}
	  },
	  "space": {
		"data": {
		  "guid": "%[2]s"
		}
	  }
	},
	"metadata": {
	  "labels": {},
	  "annotations": {}
	},
	"links": {
	  "self": {
		"href": "https://api.example.org/v3/service_instances/%[1]s"
	  },
	  "service_plan": {
		"href": "https://api.example.org/v3/service_plans/5358d122-638e-11ea-afca-bf6e756684ac"
	  },
	  "space": {
		"href": "https://api.example.org/v3/spaces/%[2]s"
	  },
	  "parameters": {
		"href": "https://api.example.org/v3/service_instances/%[1]s/parameters"
	  },
	  "service_credential_bindings": {
		"href": "https://api.example.org/v3/service_credential_bindings?service_instance_guids=%[1]s"
	  },
	  "service_route_bindings": {
		"href": "https://api.example.org/v3/service_route_bindings?service_instance_guids=%[1]s"
	  }
	}
  }
`,
	V3ServiceInstanceGuid,
	V3SpaceGuid,
)
