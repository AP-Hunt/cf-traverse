package main

import (
	cliPlugin "code.cloudfoundry.org/cli/plugin"
)

func main() {
	cliPlugin.Start(NewPlugin())
}
