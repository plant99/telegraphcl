package cmd

import (
	"fmt"

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
	Short: "Count views on your Telegra.ph page. Arguments: <path>",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		page.GetViews(args[0])
	},
}

var pageCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create Page from a Markdown file. Arguments: <markdown-path> <title>",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		page.CreatePage(args[0], args[1])
	},
}

var pageGetCmd = &cobra.Command{
	Use:   "get",
	Short: "Get page with Telegra.ph path. Arguments: <path>",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		page.GetPage(args[0])
	},
}

var pageEditCmd = &cobra.Command{
	Use:   "edit",
	Short: "Edit page with Telegra.ph path. Arguments: <path> <markdown-path>",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		// Taking input from user
		fmt.Print("Enter the updated title of this page: ")
		var title string
		fmt.Scanln(&title)
		page.EditPage(args[0], args[1], title)
	},
}

func init() {
	pageCmd.AddCommand(pageListCmd)
	pageCmd.AddCommand(pageViewsCmd)
	pageCmd.AddCommand(pageCreateCmd)
	pageCmd.AddCommand(pageGetCmd)
	pageCmd.AddCommand(pageEditCmd)
}
