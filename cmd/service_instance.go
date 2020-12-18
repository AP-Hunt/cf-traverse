package cmd

import (
	cliPlugin "code.cloudfoundry.org/cli/plugin"
	"fmt"
	cfclient "github.com/cloudfoundry-community/go-cfclient"
	"github.com/spf13/cobra"
)

func NewServiceCommand(cliConnection cliPlugin.CliConnection) *cobra.Command {

	return &cobra.Command{
		Use:     "service_instance",
		Aliases: []string{"s_i"},
		Args:    cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			client, err := newClient(cliConnection)
			if err != nil {
				return err
			}

			identifier := args[1]

			if !isUUID(identifier) {
				identifier, err = serviceInstanceGuidFromName(client, identifier)
			}

			targetType := args[0]
			switch targetType {
			case "space":
				space, err := serviceInstanceToSpace(client, identifier)
				if err != nil {
					return err
				}
				cmd.Print(string(space))
			case "org":
				org, err := serviceInstanceToOrg(client, identifier)
				if err != nil {
					return err
				}
				cmd.Print(string(org))
			case "plan":
				plan, err := serviceInstanceToPlan(client, identifier)
				if err != nil {
					return err
				}
				cmd.Print(string(plan))
			default:
				return fmt.Errorf("unknown relation '%s'", targetType)
			}

			return nil
		},
	}
}

func serviceInstanceGuidFromName(client *cfclient.Client, identifier string) (string, error) {
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

func serviceInstanceToSpace(client *cfclient.Client, identifier string) ([]byte, error) {

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

func serviceInstanceToOrg(client *cfclient.Client, identifier string) ([]byte, error) {
	space, err := serviceInstanceToSpace(client, identifier)
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

func serviceInstanceToPlan(client *cfclient.Client, identifier string) ([]byte, error) {
	svcInstance, err := apiGetRequest(client, fmt.Sprintf("/v3/service_instances/%s", identifier))
	if err != nil {
		return nil, err
	}

	planGUID, err := jsonPath(svcInstance, "$.relationships.service_plan.data.guid")
	if err != nil {
		return nil, err
	}

	planJSON, err := apiGetRequest(client, fmt.Sprintf("/v3/service_plans/%s", planGUID))
	if err != nil {
		return nil, err
	}

	return planJSON, nil
}
