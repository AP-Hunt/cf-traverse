package testfixtures

import "fmt"

const V3ServicePlanGuid = "2b5fa5bf-eafa-4ae3-9727-e1f6bcc622ee"
const V3ServicePlanAlternateGuid = "8a138e97-42d6-46f6-be0f-5347d23b7e8e"

var V3ServicePlanPath = fmt.Sprintf("/v3/service_plans/%s", V3ServicePlanGuid)
var V3ServicePlansForOfferingPath = fmt.Sprintf("/v3/service_plans?per_page=5000&service_offering_guids=%s", V3ServiceOfferingGuid)

var V3ServicePlan = NewV3ServicePlan(V3ServicePlanGuid, V3ServiceOfferingGuid)
var V3ServicePlansForOfferingListing = fmt.Sprintf(`
{
  "pagination": {
    "total_results": 2,
    "total_pages": 1,
    "first": {
      "href": "https://api.example.org/v3/service_plans?page=1&per_page=5000&service_offering_guids=%[1]s"
    },
    "last": {
      "href": "https://api.example.org/v3/service_plans?page=2&per_page=5000&service_offering_guids=%[1]s"
    },
    "next": null,
    "previous": null
  },
  "resources": [
	%[2]s,
	%[3]s
  ]
}`,
V3ServiceOfferingGuid,
V3ServicePlan,
NewV3ServicePlan(V3ServicePlanAlternateGuid, V3ServiceOfferingGuid),
)

func NewV3ServicePlan(planGuid string, offeringGuid string) string  {
	return fmt.Sprintf(`
{
  "guid": "%[1]s",
  "name": "my_big_service_plan",
  "description": "Big",
  "visibility_type": "public",
  "available": true,
  "free": false,
  "costs": [
    {
      "currency": "USD",
      "amount": 199.99,
      "unit": "Monthly"
    }
  ],
  "created_at": "2019-11-28T13:44:02Z",
  "updated_at": "2019-11-28T13:44:02Z",
  "maintenance_info": {
    "version": "1.0.0+dev4",
    "description": "Database version 7.8.0"
  },
  "broker_catalog": {
    "id": "db730a8c-11e5-11ea-838a-0f4fff3b1cfb",
    "metadata": {
      "custom-key": "custom-information"
    },
    "maximum_polling_duration": null,
    "features": {
      "plan_updateable": true,
      "bindable": true
    }
  },
  "schemas": {
    "service_instance": {
      "create": {
        "parameters": {
          "$schema": "http://json-schema.org/draft-04/schema#",
          "type": "object",
          "properties": {
            "billing-account": {
              "description": "Billing account number used to charge use of shared fake server.",
              "type": "string"
            }
          }
        }
      },
      "update": {}
    },
    "service_binding": {
      "create": {}
    }
  },
  "relationships": {
    "service_offering": {
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
      "href": "https://api.example.org/v3/service_plans/%[1]s"
    },
    "service_offering": {
      "href": "https://api.example.org/v3/service_offerings/%[2]s"
    },
    "visibility": {
      "href": "https://api.example.org/v3/service_plans/%[1]s/visibility"
    }
  }
}
`,
		planGuid,
		offeringGuid)
}
