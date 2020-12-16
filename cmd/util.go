package cmd

import (
	"fmt"
	"io"
	"io/ioutil"
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
