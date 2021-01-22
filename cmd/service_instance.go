package cmd

import (
	"fmt"
	"strings"

	cliPlugin "code.cloudfoundry.org/cli/plugin"
	cfclient "github.com/cloudfoundry-community/go-cfclient"
	"github.com/spf13/cobra"
)

func NewServiceInstancesCommand(cliConnection cliPlugin.CliConnection) *cobra.Command {

	rootCmd := &cobra.Command{
		Use:     "service_instance",
		Aliases: []string{"s_i"},
		Short: "Find relations of a service instance",
		SilenceUsage: true,
		TraverseChildren: true,
	}

	rootCmd.AddCommand(newServiceInstanceToSpaceCommand(cliConnection))
	rootCmd.AddCommand(newServiceInstanceToOrgCommand(cliConnection))
	rootCmd.AddCommand(newServiceInstanceToPlanCommand(cliConnection))
	rootCmd.AddCommand(newServiceInstanceToServiceOfferingCommand(cliConnection))
	rootCmd.AddCommand(newServiceInstanceOrgSpaceNameCommand(cliConnection))

	return rootCmd
}

func newServiceInstanceToSpaceCommand(cliConnection cliPlugin.CliConnection) *cobra.Command {
	return &cobra.Command{
		Use: "space",
		Args: cobra.ExactArgs(1),
		Short: "Find the space to which a service instance belongs",
		SilenceUsage: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			client, err := newClient(cliConnection)
			if err != nil {
				return err
			}

			identifier := args[0]

			if !isUUID(identifier) {
				identifier, err = serviceInstanceGuidFromName(client, identifier)
				if err != nil {
					return err
				}
			}

			space, err := serviceInstanceToSpace(client, identifier)
			if err != nil {
				return err
			}
			cmd.Print(string(space))
			return nil
		},
	}
}

func newServiceInstanceToOrgCommand(cliConnection cliPlugin.CliConnection) *cobra.Command {
	return &cobra.Command{
		Use: "org",
		Args: cobra.ExactArgs(1),
		Short: "Find the organization to which the service instance belongs",
		SilenceUsage: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			client, err := newClient(cliConnection)
			if err != nil {
				return err
			}

			identifier := args[0]

			if !isUUID(identifier) {
				identifier, err = serviceInstanceGuidFromName(client, identifier)
				if err != nil {
					return err
				}
			}

			org, err := serviceInstanceToOrg(client, identifier)
			if err != nil {
				return err
			}
			cmd.Print(string(org))
			return nil
		},
	}
}

func newServiceInstanceToPlanCommand(cliConnection cliPlugin.CliConnection) *cobra.Command {
	return &cobra.Command{
		Use: "plan",
		Args: cobra.ExactArgs(1),
		SilenceUsage: true,
		Short: "Find the service plan that the service instance is an instance of",
		RunE: func(cmd *cobra.Command, args []string) error {
			client, err := newClient(cliConnection)
			if err != nil {
				return err
			}

			identifier := args[0]

			if !isUUID(identifier) {
				identifier, err = serviceInstanceGuidFromName(client, identifier)
				if err != nil {
					return err
				}
			}

			plan, err := serviceInstanceToPlan(client, identifier)
			if err != nil {
				return err
			}
			cmd.Print(string(plan))
			return nil
		},
	}
}

func newServiceInstanceToServiceOfferingCommand(cliConnection cliPlugin.CliConnection) *cobra.Command {
	return &cobra.Command{
		Use: "service_offering",
		Args: cobra.ExactArgs(1),
		SilenceUsage: true,
		Short: "Find which service offering the service instance is an instance of",
		RunE: func(cmd *cobra.Command, args []string) error {
			client, err := newClient(cliConnection)
			if err != nil {
				return err
			}

			identifier := args[0]

			if !isUUID(identifier) {
				identifier, err = serviceInstanceGuidFromName(client, identifier)
				if err != nil {
					return err
				}
			}

			offering, err := serviceInstanceToServiceOffering(client, identifier)
			if err != nil {
				return err
			}
			cmd.Print(string(offering))
			return nil
		},
	}
}

func newServiceInstanceOrgSpaceNameCommand(cliConnection cliPlugin.CliConnection) *cobra.Command {
	var delimiter string
	cmd := &cobra.Command{
		Use: "org_space_name -d|--delimiter",
		Args: cobra.ExactArgs(1),
		Short: "Format the service instance name, organization name, and space name for the service instance. E.g. `org/space/instance_name`",
		SilenceUsage: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			client, err := newClient(cliConnection)
			if err != nil {
				return err
			}

			identifier := args[0]

			if !isUUID(identifier) {
				identifier, err = serviceInstanceGuidFromName(client, identifier)
				if err != nil {
					return err
				}
			}

			orgName, spaceName, serviceInstanceName, err := serviceInstanceToOrgSpaceInstanceName(client, identifier)
			if err != nil {
				return err
			}

			output := strings.Join([]string{orgName, spaceName, serviceInstanceName}, delimiter)
			cmd.Print(output)
			return nil
		},
	}

	cmd.Flags().StringVarP(&delimiter, "delimiter", "d", "", "Delimiter to separate the org, space, and service instance name")
	_ = cmd.MarkFlagRequired("delimiter")

	return cmd
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

func serviceInstanceToOrgSpaceInstanceName(client *cfclient.Client, identifier string) (string, string, string, error) {
	serviceInstanceWithIncludesPath := fmt.Sprintf("/v3/service_instances/%s?fields[space]=name&fields[space.organization]=name", identifier)
	serviceInstanceWithIncludesJSON, err := apiGetRequest(client, serviceInstanceWithIncludesPath)

	if err != nil {
		return "", "", "", err
	}

	serviceInstanceName, err := jsonPathString(serviceInstanceWithIncludesJSON, "$.name")
	if err != nil {
		return "", "", "", err
	}

	spaceName, err := jsonPathString(serviceInstanceWithIncludesJSON, "$.included.spaces[0].name")
	if err != nil {
		return "", "", "", err
	}

	orgName, err := jsonPathString(serviceInstanceWithIncludesJSON, "$.included.organizations[0].name")
	if err != nil {
		return "", "", "", err
	}

	return orgName, spaceName, serviceInstanceName, nil
}
