package cmd

import (
	"github.com/spf13/cobra"
)

var TelegraphCommand = &cobra.Command{
	Use:   "telegraphcl",
	Short: "",
	Long:  "",
}

func Execute() {
	AddCommands()
	TelegraphCommand.Execute()
}

func AddCommands() {
	TelegraphCommand.AddCommand(versionCmd)
	TelegraphCommand.AddCommand(userCmd)
	TelegraphCommand.AddCommand(pageCmd)
}
