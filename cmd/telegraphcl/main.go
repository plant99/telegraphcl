package main

import (
	"fmt"

	"github.com/plant99/telegraphcl/pkg/user"
)

func main() {
	fmt.Println("Microblog with telegra.ph from your friendly neighbourhood terminal :)")
	user.PrintSomethingAboutUser("something for the user")
}
