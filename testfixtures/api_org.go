package testfixtures

import "fmt"

const V3OrgGuid = "13200e2e-cb27-4c3e-b9fe-393d20370677"

var V3OrgPath = fmt.Sprintf("/v3/organizations/%s", V3OrgGuid)
var V3Org = fmt.Sprintf(`
{
	"guid": "%[1]s",
	"created_at": "2017-02-01T01:33:58Z",
	"updated_at": "2017-02-01T01:33:58Z",
	"name": "my-organization",
	"suspended": false,
	"relationships": {
	  "quota": {
		"data": {
		  "guid": "b7887f5c-34bb-40c5-9778-577572e4fb2d"
		}
	  }
	},
	"links": {
	  "self": {
		"href": "https://api.example.org/v3/organizations/%[1]s"
	  },
	  "domains": {
		"href": "https://api.example.org/v3/organizations/%[1]s/domains"
	  },
	  "default_domain": {
		"href": "https://api.example.org/v3/organizations/%[1]s/domains/default"
	  },
	  "quota": {
		"href": "https://api.example.org/v3/organization_quotas/%[1]s"
	  }
	},
	"metadata": {
	  "labels": {},
	  "annotations": {}
	}
  }
`,
	V3OrgGuid)
