package cmd

import (
	"fmt"

	cliPlugin "code.cloudfoundry.org/cli/plugin"
	cfclient "github.com/cloudfoundry-community/go-cfclient"
	"github.com/spf13/cobra"
)

func NewServiceInstancesCommand(cliConnection cliPlugin.CliConnection) *cobra.Command {

	return &cobra.Command{
		Use:     "service_instance",
		Aliases: []string{"s_i"},
		Args:    cobra.ExactArgs(2),
		SilenceUsage: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			client, err := newClient(cliConnection)
			if err != nil {
				return err
			}

			identifier := args[1]

			if !isUUID(identifier) {
				identifier, err = serviceInstanceGuidFromName(client, identifier)
				if err != nil {
					return err
				}
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
			case "service_offering":
				offering, err := serviceInstanceToServiceOffering(client, identifier)
				if err != nil {
					return err
				}
				cmd.Print(string(offering))
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
		return "", err
	}

	guid, err := jsonPathString(listing, "$.resources[0].guid")
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

	spaceGUID, err := jsonPathString(svcInstance, "$.relationships.space.data.guid")
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

	orgGUID, err := jsonPathString(space, "$.relationships.organization.data.guid")
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

	planGUID, err := jsonPathString(svcInstance, "$.relationships.service_plan.data.guid")
	if err != nil {
		return nil, err
	}

	planJSON, err := apiGetRequest(client, fmt.Sprintf("/v3/service_plans/%s", planGUID))
	if err != nil {
		return nil, err
	}

	return planJSON, nil
}

func serviceInstanceToServiceOffering(client *cfclient.Client, identifier string) ([]byte, error) {
	plan, err := serviceInstanceToPlan(client, identifier)

	if err != nil {
		return nil, err
	}

	offeringGUID, err := jsonPathString(plan, "$.relationships.service_offering.data.guid")
	if err != nil {
		return nil, err
	}

	offeringJSON, err := apiGetRequest(client, fmt.Sprintf("/v3/service_offerings/%s", offeringGUID))
	if err != nil {
		return nil, err
	}

	return offeringJSON, nil
}
