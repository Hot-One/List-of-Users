package controller

import (
	// "fmt"
	"strings"

	"github.com/bxcodec/faker/v4"

	"app/models"
)

func GenerateUser(count int) []models.User {
	var users []models.User
	for count >= 0 {
		users = append(users, models.User{
			Name:        faker.Name(),
			PhoneNumber: faker.Phonenumber(),
		})
		count--
	}

	return users
}

func FindPhoneNumbers(PhoneNumber string, users []models.User) []models.User{
	var have_user = []models.User{}
	for _, user := range users{
		if strings.Contains(user.PhoneNumber, PhoneNumber){
			have_user = append(have_user, user)
		}
	}
	return have_user
}