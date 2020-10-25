package user

import (
	"fmt"

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

	// print `data` as a string
	fmt.Printf("%s\n", data)

}
