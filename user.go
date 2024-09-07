package main

import (
	"log"
)

func GetAllUsers() []User {
	var users []User
	result := BotDB.Table("users").Find(&users)
	if result.Error != nil {
		log.Fatal(result.Error)
	}
	return users
}

func GetUserByUsername(username string) (User, error) {
	var user User
	result := BotDB.Where("username = ?", username).First(&user)
	if result.Error != nil {
		return user, result.Error
	}
	return user, nil
}
