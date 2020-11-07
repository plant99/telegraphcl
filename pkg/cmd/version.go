package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var version string

func init() {
	version = "v0.0.1-beta"
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of telegraphcl",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("telegraphcl", version)
	},
}
