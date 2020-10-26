package user

import (
	"fmt"

	jsoniter "github.com/json-iterator/go"
	"github.com/plant99/telegraphcl/pkg/util"
)

func CreateUser(user User) {

	userToBeCreated := createUser{
		ShortName:  user.ShortName,
		AuthorName: user.AuthorName,
		AuthorUrl:  user.AuthorUrl,
	}

	data, err := util.MakeRequest("createAccount", userToBeCreated)

	if err != nil {
		fmt.Println(err)
	}

	parser := jsoniter.ConfigFastest

	resultUser := new(createUserResponse)
	if err = parser.Unmarshal(data, resultUser); err != nil {
		fmt.Println("Couldn't handle api.telegra.ph response.")
	}

	// save access token
	util.StoreAccessToken(resultUser.AccessToken)
}

func ViewCurrentUserInfo() {
	accessToken, err := util.FetchAccessToken()
	if err != nil {
		fmt.Println("Error fetching access token, have you deleted ~/telegraph.token?")
	}
	viewInfoRequest := viewUserInfoRequest{
		AccessToken: accessToken,
		Fields:      []string{"short_name", "author_name", "author_url", "auth_url", "page_count"},
	}

	data, err := util.MakeRequest("getAccountInfo", viewInfoRequest)

	parser := jsoniter.ConfigFastest

	userInfo := new(viewUserInfoResponse)

	if err = parser.Unmarshal(data, userInfo); err != nil {
		fmt.Println("Couldn't handle api.telegra.ph response.")
	}

	// view user information
	fmt.Println("Short Name: ", userInfo.ShortName)
	fmt.Println("Author Name: ", userInfo.AuthorName)
	fmt.Println("Author URL: ", userInfo.AuthorUrl)
	fmt.Println("Page Count: ", userInfo.PageCount)
	fmt.Println("Auth URL", userInfo.AuthUrl)
}

func EditCurrentUserInfo(user User) {
	accessToken, err := util.FetchAccessToken()
	if err != nil {
		fmt.Println("Error fetching access token, have you deleted ~/telegraph.token?")
	}
	editUserInfo := editUserInfoRequest{
		AccessToken: accessToken,
		ShortName:   user.ShortName,
		AuthorName:  user.AuthorName,
		AuthorUrl:   user.AuthorUrl,
	}

	data, err := util.MakeRequest("editAccountInfo", editUserInfo)

	parser := jsoniter.ConfigFastest

	userInfo := new(editUserInfoResponse)

	if err = parser.Unmarshal(data, userInfo); err != nil {
		fmt.Println("Couldn't handle api.telegra.ph response.")
	}

	// view user information
	fmt.Println("Updated user account information!")
	fmt.Println("Short Name: ", userInfo.ShortName)
	fmt.Println("Author Name: ", userInfo.AuthorName)
	fmt.Println("Author URL: ", userInfo.AuthorUrl)
}
