package cmd

import (
	cliPlugin "code.cloudfoundry.org/cli/plugin"
	"fmt"
	cfclient "github.com/cloudfoundry-community/go-cfclient"
	"github.com/spf13/cobra"
)

func NewServiceCommand(cliConnection cliPlugin.CliConnection) *cobra.Command {
	validTargetTypes := []string{
		"space",
		"org",
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

			client, err := newClient(cliConnection)
			if err != nil {
				return err
			}

			identifier := args[1]

			if !isUUID(identifier) {
				identifier, err = serviceGuidFromName(client, identifier)
			}

			switch targetType {
			case "space":
				space, err := serviceToSpace(client, identifier)
				if err != nil {
					return err
				}
				cmd.Print(string(space))
			case "org":
				org, err := serviceToOrg(client, identifier)
				if err != nil {
					return err
				}
				cmd.Print(string(org))
			}

			return nil
		},
	}
}

func serviceGuidFromName(client *cfclient.Client, identifier string) (string, error) {
	listing, err := apiGetRequest(client, fmt.Sprintf("/v3/service_instances?names=%s", identifier))
	if err != nil {
		return "" ,err
	}

	guid, err := jsonPath(listing, "$.resources[0].guid")
	if err != nil {
		return "", err
	}

	return guid, nil
}

func serviceToSpace(client *cfclient.Client, identifier string) ([]byte, error) {

	svcInstance, err := apiGetRequest(client, fmt.Sprintf("/v3/service_instances/%s", identifier))
	if err != nil {
		return nil, err
	}

	spaceGUID, err := jsonPath(svcInstance, "$.relationships.space.data.guid")
	if err != nil {
		return nil, err
	}

	spaceJSON, err := apiGetRequest(client, fmt.Sprintf("/v3/spaces/%s", spaceGUID))
	if err != nil {
		return nil, err
	}

	return spaceJSON, nil
}

func serviceToOrg(client *cfclient.Client, identifier string) ([]byte, error) {
	space, err := serviceToSpace(client, identifier)
	if err != nil {
		return nil, err
	}

	orgGUID, err := jsonPath(space, "$.relationships.organization.data.guid")
	if err != nil {
		return nil, err
	}

	org, err := apiGetRequest(client, fmt.Sprintf("/v3/organizations/%s", orgGUID))
	if err != nil {
		return nil, err
	}

	return org, nil
}
