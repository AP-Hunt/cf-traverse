package cmd

import (
	"fmt"

	cliPlugin "code.cloudfoundry.org/cli/plugin"
	cfclient "github.com/cloudfoundry-community/go-cfclient"
	"github.com/spf13/cobra"
)

func NewServicePlansCommand(cliConnection cliPlugin.CliConnection) *cobra.Command {
	return &cobra.Command{
		Use:     "service_plan",
		Aliases: []string{"s_p"},
		Args:    cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			client, err := newClient(cliConnection)
			if err != nil {
				return err
			}

			identifier := args[1]
			targetType := args[0]

			switch targetType {
			case "instances_of":
				instances, err := serviceInstancesFromPlanGuid(client, identifier)
				if err != nil {
					return err
				}
				cmd.Print(string(instances))

			default:
				return fmt.Errorf("unknown relation '%s'", targetType)
			}

			return nil
		},
	}
}

func serviceInstancesFromPlanGuid(client *cfclient.Client, identifier string) ([]byte, error) {
	listing, err := apiGetRequest(client, fmt.Sprintf("/v3/service_instances?per_page=5000&service_plan_guids=%s", identifier))

	if err != nil {
		return nil, err
	}

	return listing, nil
}
