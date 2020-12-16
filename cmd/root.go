package cmd

import (
	"github.com/spf13/cobra"
)

func NewRootCommand() *cobra.Command {
	return &cobra.Command{
		Use:     "traverse",
		Aliases: []string{"tr"},
		Run: func(cmd *cobra.Command, args []string) {
		},
	}
}
