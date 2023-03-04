package model

import (
	"errors"
)

// user represents data about a record user.
type UserModel struct {
	ID        string `json:"id" binding:"required" example:"string" maxLength:"15"`
	Email     string `json:"email" binding:"required" example:"string" maxLength:"255"`
	FirstName string `json:"firstname" binding:"required" example:"string" maxLength:"255"`
	LastName  string `json:"lastname" binding:"required" example:"string" maxLength:"255"`
	JobTitle  string `json:"jobtitle" binding:"required" example:"string" maxLength:"255"`
}

var users = []UserModel{
	{ID: "1", Email: "chatchai.w@netpoleons.com", FirstName: "Chatchai", LastName: "Wongdetsakul", JobTitle: "DevSecOps Engineer"},
	{ID: "2", Email: "natchapong.b@netpoleons.com", FirstName: "Natchapong", LastName: "Buretes", JobTitle: "iSec and Network Engineer"},
	{ID: "3", Email: "chananya.k@netpoleons.com", FirstName: "Chananya", LastName: "Krudnim", JobTitle: "iSec and Network Engineer"},
}

func GetUsers() ([]UserModel, error) {
	return users, nil
}

func GetUserByID(id string) (UserModel, error) {
	// Loop over the list of users, looking for
	// an user whose ID value matches the parameter.
	var user UserModel
	for _, a := range users {
		if a.ID == id {
			user = a
			return user, nil
		}
	}
	return user, errors.New("User not found")
}
