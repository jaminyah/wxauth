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

func (store *DataStore) GetUser(email string) models.User {
	user := models.User{}

	return user
}

func (store *DataStore) ReadUsers() []models.User {
	users := []models.User{}

	return users
}
