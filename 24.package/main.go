package main

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/syedxshah/golang/auth"
	"github.com/syedxshah/golang/user"
)

func main() {

	auth.Loginwithcredentials("Ansar", "Ali")

	fmt.Println(auth.Getsession())

	u := user.Users{
		Email: "hello@gmail.com",
		Name:  "Ansar",
	}

	u.Name = "Syed Ansar Ali"

	fmt.Println(u.Name, u.Email)

	// command to download the pacakge : go get github.com/fatih/color

	color.Green("Uname Is := %v \n\n\n", u.Name)
	color.Red("Email Is := %v ", u.Email)

	// another command is go mod tidy to fix and auto download the package if its was used in the code
}
