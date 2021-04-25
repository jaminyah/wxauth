package main

import (
	"database/sql"
	"fmt"
	"os"
	"wxauth/models"
	"wxauth/platform/database"

	_ "github.com/mattn/go-sqlite3"
)

func main() {

	os.Remove("./auth.db")

	db, err := sql.Open("sqlite3", "./auth.db")
	if err != nil {
		fmt.Println("Cannot create database.")
	}
	defer db.Close()

	userDB := database.Init(db)

	userDm1 := models.UserDataModel{
		Email:    "jan@mail.com",
		Token:    "2468",
		UserRole: "Client",
		Services: "Notify, WebRTC",
	}

	userDB.InsertUser(userDm1)

	userDm2 := models.UserDataModel{
		Email:    "der@mail.com",
		Token:    "72019",
		UserRole: "Client",
		Services: "Notify",
	}

	userDB.InsertUser(userDm2)

	userDms := userDB.ReadUsers()
	for _, user := range userDms {
		fmt.Printf("%v, %s, %s, %s, %s\n", user.ID, user.Email, user.Token, user.UserRole, user.Services)
	}

	userDm := userDB.GetUser("der@mail.com")
	fmt.Printf("%v, %s, %s, %s, %s\n", userDm.ID, userDm.Email, userDm.Token, userDm.UserRole, userDm.Services)
	fmt.Println()
	/*
		fmt.Println(models.Client, models.Admin, models.Notify, models.WebRTC, models.UserLogs)

		user := models.User{
			Email:    "jan@mail.com",
			Password: "2468",
			UserRole: models.Client,
			Services: []models.Service{models.Notify, models.WebRTC},
		}

		fmt.Println(user.Email)
		for _, service := range user.Services {
			fmt.Printf("Registered for: %s\n", service)
		}

		fmt.Printf("User role: %s\n", user.UserRole)
	*/
}
