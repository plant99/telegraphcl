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
