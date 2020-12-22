package testfixtures

import "fmt"

const V3ServiceOfferingGuid = "38c435ac-87a5-4c19-9794-1ab3e99bd97c"
const V3ServiceOfferingName = "a-service-offering"

var V3ServiceOfferingPath = fmt.Sprintf("/v3/service_offerings/%s", V3ServiceOfferingGuid)
var V3ServiceOfferingByNamePath = fmt.Sprintf("/v3/service_offerings?names=%s", V3ServiceOfferingName)

var V3ServiceOffering = fmt.Sprintf(`
{
  "guid": "%[1]s",
  "name": "%[2]s",
  "description": "Provides my service",
  "available": true,
  "tags": ["relational", "caching"],
  "requires": [],
  "created_at": "2019-11-28T13:44:02Z",
  "updated_at": "2019-11-28T13:44:02Z",
  "shareable": true,
  "documentation_url": "https://some-documentation-link.io",
  "broker_catalog": {
    "id": "db730a8c-11e5-11ea-838a-0f4fff3b1cfb",
    "metadata": {
      "shareable": true
    },
    "features": {
      "plan_updateable": true,
      "bindable": true,
      "instances_retrievable": true,
      "bindings_retrievable": true,
      "allow_context_updates": false
    }
  },
  "relationships": {
    "service_broker": {
      "data": {
        "guid": "13c60e38-11e7-11ea-9106-33ee3c5bd4d7"
      }
    }
  },
  "metadata": {
    "labels": {},
    "annotations": {}
  },
  "links": {
    "self": {
      "href": "https://api.example.org/v3/service_offerings/%[1]s"
    },
    "service_plans": {
      "href": "https://api.example.org/v3/service_plans?service_offering_guids=%[1]s"
    },
    "service_broker": {
      "href": "https://api.example.org/v3/service_brokers/13c60e38-11e7-11ea-9106-33ee3c5bd4d7"
    }
  }
}
`,
V3ServiceOfferingGuid,
V3ServiceOfferingName)

var V3ServiceOfferingByNameListing = fmt.Sprintf(`
{
  "pagination": {
    "total_results": 1,
    "total_pages": 1,
    "first": {
      "href": "https://api.example.org/v3/service_offerings?page=1&per_page=1&names=%[2]s"
    },
    "last": {
      "href": "https://api.example.org/v3/service_offerings?page=1&per_page=1&names=%[2]s"
    },
    "next": null,
    "previous": null
  },
  "resources": [
	%[1]s
  ]
}
`,
V3ServiceOffering,
V3ServiceOfferingName)
