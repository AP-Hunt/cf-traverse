package testfixtures

import "fmt"

const V3ServiceInstanceGuid = "e9aca2a7-f4f9-431e-924c-664ddc236604"
const V3ServiceInstanceAlternateGuid = "0f52de7f-b7f6-435c-9f9d-f0bca1a6f874"
const V3ServiceInstanceName = "a-service-instance"

var V3ServiceInstancePath = fmt.Sprintf("/v3/service_instances/%s", V3ServiceInstanceGuid)
var V3ServiceInstanceWithOrgAndSpaceNamePath = fmt.Sprintf("/v3/service_instances/%s?fields[space]=name&fields[space.organization]=name", V3ServiceInstanceGuid)
var V3ServiceInstanceByNameListingPath = fmt.Sprintf("/v3/service_instances?names=%s", V3ServiceInstanceName)
var V3ServiceInstancesBySinglePlanListingPath = fmt.Sprintf("/v3/service_instances?per_page=5000&service_plan_guids=%s", V3ServicePlanGuid)
var V3ServiceInstancesByMultiplePlanListingPath = fmt.Sprintf("/v3/service_instances?per_page=5000&service_plan_guids=%s,%s", V3ServicePlanGuid, V3ServicePlanAlternateGuid)

var V3ServiceInstance = NewV3ServiceInstance(V3ServiceInstanceGuid, V3ServiceInstanceName, V3SpaceGuid, V3ServicePlanGuid)
var V3ServiceInstanceWithOrgAndSpaceName = NewV3ServiceInstanceWithOrgAndSpaceName(V3ServiceInstanceGuid, V3ServiceInstanceName, V3SpaceGuid, V3ServicePlanGuid, V3OrgName, V3SpaceName)

var V3ServiceInstancesByNameListing = fmt.Sprintf(`
{
  "pagination": {
    "total_results": 1,
    "total_pages": 1,
    "first": {
      "href": "https://api.example.org/v3/service_instances?page=1&per_page=5000&names=%[2]s"
    },
    "last": {
      "href": "https://api.example.org/v3/service_instances?page=1&per_page=5000&names=%[2]s"
    },
    "next": null,
    "previous": null
  },
  "resources": [
	%[1]s
  ]
}
`,
	V3ServiceInstance,
	V3ServiceInstanceName)

var V3ServiceInstancesBySinglePlanListing = fmt.Sprintf(`
{
	"pagination": {
	  "total_results": 1,
	  "total_pages": 1,
	  "first": {
		"href": "https://api.example.org/v3/service_instances?page=1&per_page=50&service_plan_guids=%[3]s"
	  },
	  "last": {
		"href": "https://api.example.org/v3/service_instances?page=1&per_page=50&service_plan_guids=%[3]s"
	  },
	  "next": null,
	  "previous": null
	},
	"resources": [
	  %[1]s,
	  %[2]s
	]
  }
`,
	V3ServiceInstance,
	NewV3ServiceInstance(V3ServiceInstanceAlternateGuid, "b-service-instance", V3SpaceGuid, V3ServicePlanGuid),
	V3ServicePlanGuid)

var V3ServiceInstancesByMultiplePlanListing = fmt.Sprintf(`
{
	"pagination": {
	  "total_results": 1,
	  "total_pages": 1,
	  "first": {
		"href": "https://api.example.org/v3/service_instances?page=1&per_page=50&service_plan_guids=%[3]s,%[4]s"
	  },
	  "last": {
		"href": "https://api.example.org/v3/service_instances?page=1&per_page=50&service_plan_guids=%[3]s,%[4]s"
	  },
	  "next": null,
	  "previous": null
	},
	"resources": [
	  %[1]s,
	  %[2]s
	]
  }
`,
	V3ServiceInstance,
	NewV3ServiceInstance(V3ServiceInstanceAlternateGuid, "b-service-instance", V3SpaceGuid, V3ServicePlanAlternateGuid),
	V3ServicePlanGuid,
	V3ServicePlanAlternateGuid)


func NewV3ServiceInstance(instanceGuid string, instanceName string, spaceGuid string, servicePlanGuid string) string {
	return fmt.Sprintf(`
	{
		"guid": "%[1]s",
		"created_at": "2020-03-10T15:49:29Z",
		"updated_at": "2020-03-10T15:49:29Z",
		"name": "%[3]s",
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
				"guid": "%[4]s"
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
			"href": "https://api.example.org/v3/service_plans/%[4]s"
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
		instanceGuid,
		spaceGuid,
		instanceName,
		servicePlanGuid,
	)
}

func NewV3ServiceInstanceWithOrgAndSpaceName(
	instanceGuid string,
	instanceName string,
	spaceGuid string,
	servicePlanGuid string,
	orgName string,
	spaceName string,
) string  {
	return fmt.Sprintf(`
	{
		"guid": "%[1]s",
		"created_at": "2020-03-10T15:49:29Z",
		"updated_at": "2020-03-10T15:49:29Z",
		"name": "%[3]s",
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
				"guid": "%[4]s"
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
			"href": "https://api.example.org/v3/service_plans/%[4]s"
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
		},
		"included": {
			"spaces": [
				{ "name": "%[5]s" }
			],
			"organizations": [
				{ "name": "%[6]s" }
			]
		}
		}
	`,
	instanceGuid,
	spaceGuid,
	instanceName,
	servicePlanGuid,
	spaceName,
	orgName,
	)
}
