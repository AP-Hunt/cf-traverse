package cmd

import (
	"fmt"

	cliPlugin "code.cloudfoundry.org/cli/plugin"
	cfclient "github.com/cloudfoundry-community/go-cfclient"
	"github.com/spf13/cobra"
)

func NewServiceCommand(cliConnection cliPlugin.CliConnection) *cobra.Command {
	validTargetTypes := []string{
		"space",
	}
	return &cobra.Command{
		Use:     "service",
		Aliases: []string{"s"},
		Args:    cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			targetType := args[0]
			if !inSlice(validTargetTypes, targetType) {
				return fmt.Errorf("unknown relation '%s'", targetType)
			}

			identifier := args[1]

			switch targetType {
			case "space":
				return serviceToSpace(cliConnection, identifier)
			}

			return nil
		},
	}
}

func serviceToSpace(cliConnection cliPlugin.CliConnection, identifier string) error {
	endpoint, err := cliConnection.ApiEndpoint()
	if err != nil {
		return err
	}

	token, err := cliConnection.AccessToken()
	if err != nil {
		return err
	}

	cfg := cfclient.Config{
		ApiAddress: endpoint,
		Token:      token,
	}

	client, err := cfclient.NewClient(&cfg)
	if err != nil {
		return err
	}

	svcInstance, err := client.GetServiceInstanceByGuid(identifier)
	if err != nil {
		return err
	}

	resp, err := client.DoRequest(client.NewRequest("GET", fmt.Sprintf("/v3/spaces/%s", svcInstance.SpaceGuid)))
	if err != nil {
		return err
	}
	return printResponseBodytoJSON(resp.Body)
}
