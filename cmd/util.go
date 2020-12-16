package cmd

import (
	"fmt"
	"io"
	"io/ioutil"

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

func printResponseBodytoJSON(responseBody io.ReadCloser) error {
	bs, err := ioutil.ReadAll(responseBody)
	if err != nil {
		return err
	}

	fmt.Println(string(bs))
	return nil
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
