package Cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var version string

func VersionCmd() *cobra.Command {
	return &cobra.Command{
		Use:     "version",
		Short:   "Print the Jnote-server version.",
		Long:    "Print the Jnote-server version.",
		Args:    cobra.NoArgs,
		GroupID: rootGroupID,
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				fmt.Println("Jnote-server version: " + version)
				return
			}
			cmd.Help()
		},
	}
}
