package cmd

import (
	"bufio"
	"fmt"
	"os"

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
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter your username: ")
		shortName, _ := reader.ReadString('\n')
		fmt.Print("Enter your fullname: ")
		fullName, _ := reader.ReadString('\n')
		fmt.Print("Enter an url to be displayed on your articles: ")
		authorUrl, _ := reader.ReadString('\n')
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
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter your username: ")
		shortName, _ := reader.ReadString('\n')
		fmt.Print("Enter your fullname: ")
		fullName, _ := reader.ReadString('\n')
		fmt.Print("Enter an url to be displayed on your articles: ")
		authorUrl, _ := reader.ReadString('\n')
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

var revokeCmd = &cobra.Command{
	Use:   "revoke",
	Short: "Revoke current access token, and regenerate access token.",
	Run: func(cmd *cobra.Command, args []string) {
		user.RevokeAccessToken()
	},
}

func init() {
	userCmd.AddCommand(createCmd)
	userCmd.AddCommand(viewCmd)
	userCmd.AddCommand(editCmd)
	userCmd.AddCommand(revokeCmd)
}
