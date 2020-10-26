package cmd

import (
	"github.com/plant99/telegraphcl/pkg/page"
	"github.com/spf13/cobra"
)

var pageCmd = &cobra.Command{
	Use:   "page",
	Short: "Manage your Telegra.ph pages",
	Run: func(cmd *cobra.Command, args []string) {
		page.ListPages()
	},
}
