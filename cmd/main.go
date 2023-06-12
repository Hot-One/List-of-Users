package main

import (
	"fmt"

	"app/config"
	"app/controller"
	"app/models"
)

func main() {

	cfg := config.Load()

	con := controller.NewController(&cfg)

	users := con.UserGenerate(100)
	con.Users = users 

	UserInfo :=  con.UserGetInfo()
	

	var (
		answer string
		dataLimit int
		page      int

	)
	fmt.Printf("1. List of Users\n2. Search by Phone numbers\n3. Search by Names or Letter\n4. Search by birthdays ")
	fmt.Scan(&answer)
	if answer == "1"{
		fmt.Printf("Input Data limit: ")
		fmt.Scan(&dataLimit)

		paginationCount := len(con.Users) / dataLimit
		fmt.Println("Pages count: ", paginationCount)

		for {
			fmt.Println("Input page:")
			fmt.Scan(&page)

			respUser := con.UserGetList(&models.UserGetListRequest{
				Offset: (page - 1) * dataLimit,
				Limit:  dataLimit,
			})

			for _, user := range respUser.Users {
				fmt.Println(user)
			}
		}
	}else if answer == "2"{
		var number string
		fmt.Printf("Enter the number: ")
		fmt.Scan(&number)

		have := controller.FindPhoneNumbers(number, UserInfo)
		for _, user := range have{
			fmt.Println(user)
		}
	}else if answer == "3"{
		var letters string
		fmt.Printf("Enter the letter: ")
		fmt.Scan(&letters)
		have := controller.SearchByName(letters, UserInfo)
		for _, user := range have{
			fmt.Println(user)
		}
	}else if answer == "4"{
		var from, to int
		fmt.Printf("Enter from: ")
		fmt.Scan(&from)
		fmt.Printf("Enter to: ")
		fmt.Scan(&to)
		have := controller.SearchByDate(from, to, UserInfo)
		for _, user := range have{
			fmt.Println(user)
		}
	}
}