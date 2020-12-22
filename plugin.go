package main

import (
	"fmt"
	"strconv"

	"github.com/AP-Hunt/cf-traverse/version"

	cliPlugin "code.cloudfoundry.org/cli/plugin"
	"github.com/AP-Hunt/cf-traverse/cmd"
)

type Plugin struct {
}

func NewPlugin() *Plugin {
	return &Plugin{}
}

func (p *Plugin) Run(cliConnection cliPlugin.CliConnection, args []string) {
	root := cmd.NewRootCommand()
	root.AddCommand(cmd.NewServiceInstancesCommand(cliConnection))
	root.AddCommand(cmd.NewServicePlansCommand(cliConnection))
	root.AddCommand(cmd.NewServiceOfferingsCommand(cliConnection))

	root.SetArgs(args[1:])
	err := root.Execute()
	if err != nil {
		fmt.Println(err.Error())
	}
}

func (p *Plugin) GetMetadata() cliPlugin.PluginMetadata {
	majorVersionInt, err := strconv.Atoi(version.MAJOR_VERSION)
	if err != nil {
		panic(fmt.Sprintf("cannot convert major version '%s' to an integer", version.MAJOR_VERSION))
	}

	minorVersionInt, err := strconv.Atoi(version.MINOR_VERSION)
	if err != nil {
		panic(fmt.Sprintf("cannot convert minor version '%s' to an integer", version.MINOR_VERSION))
	}

	patchVersionInt, err := strconv.Atoi(version.PATCH_VERSION)
	if err != nil {
		panic(fmt.Sprintf("cannot convert patch version '%s' to an integer", version.PATCH_VERSION))
	}

	return cliPlugin.PluginMetadata{
		Name: "traverse",
		Version: cliPlugin.VersionType{
			Major: majorVersionInt,
			Minor: minorVersionInt,
			Build: patchVersionInt,
		},
		MinCliVersion: cliPlugin.VersionType{
			Major: 7,
			Minor: 0,
			Build: 0,
		},
		Commands: []cliPlugin.Command{
			{
				Name:     "traverse",
				HelpText: "Traverse the CF API to find relations beween entities",
				UsageDetails: cliPlugin.Usage{
					Usage: "traverse\n    [source-entity-type] [destination-entity-type] [source-entity-identifier]",
				},
			},
		},
	}
}
