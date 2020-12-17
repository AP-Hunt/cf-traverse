package cmd

import (
	cliPlugin "code.cloudfoundry.org/cli/plugin"
	"fmt"
	"io/ioutil"
	"regexp"
	"strings"

	"github.com/cloudfoundry-community/go-cfclient"
	"github.com/spyzhov/ajson"
)

func inSlice(xs []string, y string) bool {
	for _, x := range xs {
		if x == y {
			return true
		}
	}

	return false
}

func apiGetRequest(client *cfclient.Client, path string) ([]byte, error) {
	resp, err := client.DoRequest(client.NewRequest("GET", path))
	if err != nil {
		return nil, err
	}

	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return respBytes, nil
}

func jsonPath(json []byte, path string) (string, error) {
	root, err := ajson.Unmarshal(json)
	if err != nil {
		return "", err
	}

	nodes, err := root.JSONPath(path)
	if err != nil {
		return "", err
	}

	if len(nodes) == 0 {
		return "", fmt.Errorf("JSON path '%s' returned no nodes", path)
	}

	if len(nodes) > 1 {
		return "", fmt.Errorf("cannot get value at path '%s' because it returned >1 node", path)
	}

	nodeVal, err := nodes[0].Value()
	if err != nil {
		return "", err
	}
	return nodeVal.(string), nil
}

func newClient(cliConnection cliPlugin.CliConnection) (*cfclient.Client, error) {
	endpoint, err := cliConnection.ApiEndpoint()
	if err != nil {
		return nil, err
	}

	token, err := cliConnection.AccessToken()
	if err != nil {
		return nil, err
	}

	cfg := cfclient.Config{
		ApiAddress: endpoint,
		Token:      strings.Replace(token, "bearer ", "", -1),
	}

	client, err := cfclient.NewClient(&cfg)
	if err != nil {
		return nil, err
	}
	return client, nil
}

func isUUID(uuid string) bool {
	r := regexp.MustCompile("^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$")
	return r.MatchString(uuid)
}
