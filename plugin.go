package main

import (
	"fmt"

	cliPlugin "code.cloudfoundry.org/cli/plugin"
	"github.com/AP-Hunt/cf-traverse/cmd"
)

type Plugin struct {
}

func NewPluin() *Plugin {
	return &Plugin{}
}

func (p *Plugin) Run(cliConnection cliPlugin.CliConnection, args []string) {
	root := cmd.NewRootCommand()
	root.AddCommand(cmd.NewServiceCommand(cliConnection))

	root.SetArgs(args[1:])
	err := root.Execute()
	if err != nil {
		fmt.Println(err.Error())
	}
}

func (p *Plugin) GetMetadata() cliPlugin.PluginMetadata {
	return cliPlugin.PluginMetadata{
		Name: "traverse",
		Version: cliPlugin.VersionType{
			Major: 0,
			Minor: 1,
			Build: 1,
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
