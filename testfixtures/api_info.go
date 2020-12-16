package testfixtures

var V2Info = `
{
	"name": "vcap",
	"build": "2222",
	"support": "http://support.cloudfoundry.com",
	"version": 2,
	"description": "Cloud Foundry sponsored by Pivotal",
	"authorization_endpoint": "http://localhost:8080/uaa",
	"token_endpoint": "http://localhost:8080/uaa",
	"min_cli_version": null,
	"min_recommended_cli_version": null,
	"api_version": "2.158.0",
	"app_ssh_endpoint": "ssh.system.domain.example.com:2222",
	"app_ssh_host_key_fingerprint": "47:0d:d1:c8:c3:3d:0a:36:d1:49:2f:f2:90:27:31:d0",
	"app_ssh_oauth_client": null,
	"routing_endpoint": "http://localhost:3000",
	"logging_endpoint": "ws://loggregator.vcap.me:80"
  }
`

var V3Info = `
{
	"build": "afa73e3fe",
	"cli_version": {
	  "minimum": "6.22.0",
	  "recommended": "latest"
	},
	"custom": {
	  "arbitrary": "stuff"
	},
	"description": "Put your apps here!",
	"name": "Cloud Foundry",
	"version": 123,
	"links": {
	  "self": { "href": "http://api.example.com/v3/info" } ,
	  "support": { "href": "http://support.example.com" }
	}
  }
`
