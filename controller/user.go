package controller

import (
	"strconv"
	"strings"
	"log"

	"github.com/bxcodec/faker/v4"

	"app/models"
)

func (c *Controller) UserGenerate(count int) []*models.User {
	var users []*models.User
	for count > 0 {
		users = append(users, &models.User{
			Name:        faker.Name(),
			PhoneNumber: faker.Phonenumber(),
			Data:		 faker.Date(),
		})
		count--
	}

	return users
}

func (c *Controller) UserGetInfo() []*models.User{
	return c.Users
}

func (c *Controller) UserGetList(req *models.UserGetListRequest) *models.UserGetListResponse {

	log.Printf("UserGetList req: %+v\n", req)

	var (
		response = &models.UserGetListResponse{}
		offset   = c.Cfg.DefaultOffset
		limit    = c.Cfg.DefaultLimit
	)

	if req.Offset > 0 {
		offset = req.Offset
	}

	if req.Limit > 0 {
		limit = req.Limit
	}

	response.Count = len(c.Users)
	if len(c.Users) < limit+offset {
		response.Users = c.Users[offset:]
	} else {
		response.Users = c.Users[offset : offset+limit]
	}

	return response
}

func SearchByName(Name string, users []*models.User) []*models.User{
	var have_user = []*models.User{}
	for _, user := range users{
		if strings.Contains(user.Name, Name){
			have_user = append(have_user, user)
		}
	}
	return have_user
}

func FindPhoneNumbers(PhoneNumber string, users []*models.User) []*models.User{
	var have_user = []*models.User{}
	for _, user := range users{
		if strings.Contains(user.PhoneNumber, PhoneNumber){
			have_user = append(have_user, user)
		}
	}
	return have_user
}



func SearchByDate(from, to int, users []*models.User) []*models.User{
	var (
		// fromToInt = strconv.Atoi(from)
		// toToint = strconv.Atoi(to)
		have_user = []*models.User{}
	)
	for year := from; year <= to; year++{
		yearTostr := strconv.Itoa(year)
		for _, user := range users{
			if strings.Contains(user.Data, yearTostr){
				have_user = append(have_user, user)
			}
		}
	}
	return have_user
	// var have_user = []*models.User{}
	// for _, user := range users{
	// 	if strings.Contains(user.Data, from){
	// 		have_user = append(have_user, user)
	// 	}
	// 	if strings.Contains(user.Data, to){
	// 		have_user = append(have_user, user)
	// 	}
	// }
	// return have_user
}