package database

import (
	"database/sql"
	"fmt"
	"log"
	"wxauth/models"
)

type DataStore struct {
	DB *sql.DB
}

func Init(db *sql.DB) *DataStore {

	stmt, err := db.Prepare(`
		CREATE TABLE IF NOT EXISTS "users" (
			"ID" INTEGER UNIQUE,
			"Email" TEXT,
			"Token" TEXT,
			"Role" TEXT,
			"Services" TEXT,
			PRIMARY KEY ("ID" AUTOINCREMENT)
		);
	`)

	if err != nil {
		fmt.Println("Database not created.")
	}

	stmt.Exec()
	return &DataStore{
		DB: db,
	}
}

func (store *DataStore) InsertUser(user models.UserDataModel) {
	stmt, err := store.DB.Prepare(`
		INSERT INTO "users" (Email, Token, Role, Services) values (?, ?, ?, ?)
	`)
	if err != nil {
		log.Fatal("Insert user error")
	}
	stmt.Exec(user.Email, user.Token, user.UserRole, user.Services)
}

func (store *DataStore) UpdateUser(user models.User) {

}

func (store *DataStore) DeleteUser(user models.User) {

}

func (store *DataStore) GetUser(userEmail string) models.UserDataModel {

	user := models.UserDataModel{}
	rows, err := store.DB.Query(`
		SELECT * FROM users
	`)

	if err != nil {
		log.Fatal("Query users failed.")
	}
	var id int
	var email string
	var token string
	var role string
	var services string

	for rows.Next() {

		rows.Scan(&id, &email, &token, &role, &services)
		user = models.UserDataModel{
			ID:       id,
			Email:    email,
			Token:    token,
			UserRole: role,
			Services: services,
		}

		if user.Email == userEmail {
			break
		}
	}
	return user
}

func (store *DataStore) ReadUsers() []models.UserDataModel {

	userList := []models.UserDataModel{}
	rows, err := store.DB.Query(`
		SELECT * FROM users
	`)

	if err != nil {
		log.Fatal("Query users failed.")
	}
	var id int
	var email string
	var token string
	var role string
	var services string

	for rows.Next() {
		fmt.Println("Rows has next.")

		rows.Scan(&id, &email, &token, &role, &services)
		userDm := models.UserDataModel{
			ID:       id,
			Email:    email,
			Token:    token,
			UserRole: role,
			Services: services,
		}
		fmt.Printf("user email: %s\n", userDm.Email)
		userList = append(userList, userDm)
	}
	return userList
}
