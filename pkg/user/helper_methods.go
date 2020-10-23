package user

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func CreateUser(user User) {
	// create URL
	createUserURL := url.URL{
		Scheme: "https",
		Host:   "api.telegra.ph",
		Path:   "/createAccount",
	}
	createUserURLQuery := createUserURL.Query()
	createUserURLQuery.Set("short_name", user.ShortName)
	createUserURLQuery.Set("author_name", user.AuthorName)
	createUserURLQuery.Set("author_url", user.AuthorUrl)

	createUserURL.RawQuery = createUserURLQuery.Encode()
	createUserURLString := createUserURL.String()

	// make request
	res, err := http.Get(createUserURLString)

	if err != nil {
		fmt.Println("Error while making request for Create Account")
	}

	// read all response body
	data, _ := ioutil.ReadAll(res.Body)

	// close response body
	res.Body.Close()

	// print `data` as a string
	fmt.Printf("%s\n", data)

}
