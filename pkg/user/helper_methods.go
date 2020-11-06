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

func GetCurrentUserNameAndURL() []string {
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

	returnValues := []string{userInfo.AuthorName, userInfo.AuthorUrl}
	return returnValues
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

func RevokeAccessToken() {
	accessToken, err := util.FetchAccessToken()
	if err != nil {
		fmt.Println("Error fetching access token, have you deleted ~/telegraph.token?")
	}

	revokeAccessToken := revokeAccessTokenRequest{
		AccessToken: accessToken,
	}

	data, err := util.MakeRequest("revokeAccessToken", revokeAccessToken)

	parser := jsoniter.ConfigFastest

	resultUser := new(revokeAccessTokenResponse)

	if err = parser.Unmarshal(data, resultUser); err != nil {
		fmt.Println("Couldn't handle api.telegra.ph response.")
	}

	// save access token
	util.StoreAccessToken(resultUser.AccessToken)

	fmt.Println("Revoked access token and created a new one!")
}
