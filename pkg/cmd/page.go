package cmd

import (
	"github.com/plant99/telegraphcl/pkg/page"
	"github.com/spf13/cobra"
)

var pageCmd = &cobra.Command{
	Use:   "page",
	Short: "Manage your Telegra.ph pages",
}

var pageListCmd = &cobra.Command{
	Use:   "list",
	Short: "List your Telegra.ph pages",
	Run: func(cmd *cobra.Command, args []string) {
		page.ListPages()
	},
}

var pageViewsCmd = &cobra.Command{
	Use:   "views",
	Short: "Count views on your Telegra.ph page.",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		page.GetViews(args[0])
	},
}

func init() {
	pageCmd.AddCommand(pageListCmd)
	pageCmd.AddCommand(pageViewsCmd)
}
