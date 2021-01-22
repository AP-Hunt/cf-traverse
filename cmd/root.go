package cmd

import (
	"github.com/spf13/cobra"
	"os"
)

func NewRootCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "traverse",
		Aliases: []string{"tr"},
	}

	cmd.SetOut(os.Stdout)
	cmd.SetErr(os.Stderr)
	return cmd
}
