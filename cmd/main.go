package main

import (
	"fmt"

	"app/controller"
)

func main() {

	users := controller.GenerateUser(10)
	for _, user := range users {
		fmt.Println(user)
	}
	fmt.Println()
	have := controller.FindPhoneNumbers("410", users)
	fmt.Println(have)
}