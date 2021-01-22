package testfixtures

import "fmt"

const V3SpaceGuid = "7ea72c2f-d0f1-4f4a-987e-7993807ab188"
const V3SpaceName = "a-space"

var V3SpacePath = fmt.Sprintf("/v3/spaces/%s", V3SpaceGuid)
var V3Space = fmt.Sprintf(`
{
	"guid": "%[1]s",
	"created_at": "2017-02-01T01:33:58Z",
	"updated_at": "2017-02-01T01:33:58Z",
	"name": "%[3]s",
	"relationships": {
	  "organization": {
		"data": {
		  "guid": "%[2]s"
		}
	  },
	  "quota": {
		"data": null
	  }
	},
	"links": {
	  "self": {
		"href": "https://api.example.org/v3/spaces/%[1]s"
	  },
	  "features": {
		"href": "https://api.example.org/v3/spaces/%[1]s/features"
	  },
	  "organization": {
		"href": "https://api.example.org/v3/organizations/%[2]s"
	  },
	  "apply_manifest": {
		"href": "https://api.example.org/v3/spaces/%[1]s/actions/apply_manifest",
		"method": "POST"
	  }
	},
	"metadata": {
	  "labels": {},
	  "annotations": {}
	}
  }
`,
	V3SpaceGuid,
	V3OrgGuid,
	V3SpaceName,
)
