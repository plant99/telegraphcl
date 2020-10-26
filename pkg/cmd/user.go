package cmd

import (
	"fmt"

	"github.com/plant99/telegraphcl/pkg/user"
	"github.com/spf13/cobra"
)

var userCmd = &cobra.Command{
	Use:   "user",
	Short: "Operations related to Telegraph account management",
}

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create an user",
	Run: func(cmd *cobra.Command, args []string) {
		// Taking input from user
		fmt.Print("Enter your username: ")
		var shortName string
		fmt.Scanln(&shortName)
		fmt.Print("Enter your fullname: ")
		var fullName string
		fmt.Scanln(&fullName)
		fmt.Print("Enter an url to be displayed on your articles: ")
		var authorUrl string
		fmt.Scanln(&authorUrl)
		userNew := user.User{
			ShortName:   shortName,
			AuthorName:  fullName,
			AuthorUrl:   authorUrl,
			AccessToken: "",
			AuthUrl:     "",
			PageCount:   0,
		}
		user.CreateUser(userNew)
	},
}

var viewCmd = &cobra.Command{
	Use:   "view",
	Short: "View current user information",
	Run: func(cmd *cobra.Command, args []string) {
		user.ViewCurrentUserInfo()
	},
}

var editCmd = &cobra.Command{
	Use:   "edit",
	Short: "Edit current user information",
	Run: func(cmd *cobra.Command, args []string) {
		// Taking input from user
		fmt.Print("Enter your username: ")
		var shortName string
		fmt.Scanln(&shortName)
		fmt.Print("Enter your fullname: ")
		var fullName string
		fmt.Scanln(&fullName)
		fmt.Print("Enter an url to be displayed on your articles: ")
		var authorUrl string
		fmt.Scanln(&authorUrl)
		userUpdate := user.User{
			ShortName:   shortName,
			AuthorName:  fullName,
			AuthorUrl:   authorUrl,
			AccessToken: "",
			AuthUrl:     "",
			PageCount:   0,
		}
		user.EditCurrentUserInfo(userUpdate)
	},
}

func init() {
	userCmd.AddCommand(createCmd)
	userCmd.AddCommand(viewCmd)
	userCmd.AddCommand(editCmd)
}
