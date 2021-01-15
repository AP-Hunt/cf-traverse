package cmd

import (
	cliPlugin "code.cloudfoundry.org/cli/plugin"
	"fmt"
	"github.com/cloudfoundry-community/go-cfclient"
	"github.com/spf13/cobra"
	"strings"
)

func NewServiceOfferingsCommand(cliConnection cliPlugin.CliConnection) *cobra.Command {
	return &cobra.Command{
		Use: "service_offering",
		Aliases: []string{"s_o"},
		Args: cobra.ExactArgs(2),
		SilenceUsage: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			client, err := newClient(cliConnection)
			if err != nil {
				return err
			}

			identifier := args[1]
			targetType := args[0]

			if !isUUID(identifier) {
				identifier, err = serviceOfferingGuidFromName(client, identifier)
				if err != nil {
					return err
				}
			}

			switch targetType {
			default:
				return fmt.Errorf("unknown relation '%s'", targetType)
			case "instances_of":
				instances, err := serviceInstancesFromOfferingGuid(client, identifier)
				if err != nil {
					return err
				}
				cmd.Print(string(instances))
			}

			return nil
		},
	}
}

func serviceOfferingGuidFromName(client *cfclient.Client, identifier string) (string, error) {
	listing, err := apiGetRequest(client, fmt.Sprintf("/v3/service_offerings?names=%s", identifier))
	if err != nil {
		return "", err
	}

	guid, err := jsonPathString(listing, "$.resources[0].guid")
	if err != nil {
		return "", err
	}

	return guid, nil
}

func serviceInstancesFromOfferingGuid(client *cfclient.Client, identifier string) ([]byte, error) {
	servicePlans, err := plansFromOfferingGuid(client, identifier)
	if err != nil {
		return nil, err
	}

	servicePlanGuids, err := jsonPathStringSlice(servicePlans, "$.resources[*].guid")
	if err != nil {
		return nil, err
	}

	planCsv := strings.Join(servicePlanGuids, ",")
	instancesPath := fmt.Sprintf("/v3/service_instances?per_page=5000&service_plan_guids=%s", planCsv)
	instances, err := apiGetRequest(client, instancesPath)
	if err != nil {
		return nil, err
	}

	return instances, nil
}

func plansFromOfferingGuid(client *cfclient.Client, identifier string) ([]byte, error) {
	plans, err := apiGetRequest(client, fmt.Sprintf("/v3/service_plans?per_page=5000&service_offering_guids=%s", identifier))
	if err != nil {
		return nil, err
	}
	return plans, nil
}
