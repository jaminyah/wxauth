package main

import (
	"fmt"
	"wxauth/models"
)

func main() {
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
}
