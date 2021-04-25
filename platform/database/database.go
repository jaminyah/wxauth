package database

import (
	"database/sql"
	"fmt"
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

	sql, err := store.DB.Prepare(`
		INSERT INTO "users" (Email, Token, Role, Services) values (?, ?, ?, ?)
	`)
	if err != nil {
		fmt.Println("Insert user error")
	}

	trans, err := store.DB.Begin()
	if err != nil {
		fmt.Println(err)
	}

	_, err = trans.Stmt(sql).Exec(user.Email, user.Token, user.UserRole, user.Services)
	if err != nil {
		fmt.Println("Doing rollback")
		trans.Rollback()
	} else {
		trans.Commit()
	}
}

func (store *DataStore) UpdateMail(user models.UserDataModel, newEmail string) {

	sql, err := store.DB.Prepare(`
		UPDATE users SET Email=? WHERE ID=?
	`)
	if err != nil {
		fmt.Println("Update email error.")
	}

	trans, err := store.DB.Begin()
	if err != nil {
		fmt.Println(err)
	}

	_, err = trans.Stmt(sql).Exec(newEmail, user.ID)
	if err != nil {
		fmt.Println("Doing rollback")
		trans.Rollback()
	} else {
		trans.Commit()
	}
}

func (store *DataStore) UpdateToken(user models.UserDataModel, newToken string) {

	sql, err := store.DB.Prepare(`
		UPDATE users SET Token=? WHERE ID=?
	`)
	if err != nil {
		fmt.Println("Update token error.")
	}

	trans, err := store.DB.Begin()
	if err != nil {
		fmt.Println(err)
	}

	_, err = trans.Stmt(sql).Exec(newToken, user.ID)
	if err != nil {
		fmt.Println("Doing rollback")
		trans.Rollback()
	} else {
		trans.Commit()
	}
}

func (store *DataStore) UpdateServices(user models.UserDataModel, services string) {

	sql, err := store.DB.Prepare(`
		UPDATE users SET Services=? WHERE ID=?
	`)
	if err != nil {
		fmt.Println("Update services error.")
	}

	trans, err := store.DB.Begin()
	if err != nil {
		fmt.Println(err)
	}

	_, err = trans.Stmt(sql).Exec(services, user.ID)
	if err != nil {
		fmt.Println("Doing rollback")
		trans.Rollback()
	} else {
		trans.Commit()
	}
}

func (store *DataStore) DeleteUser(user models.UserDataModel) {

	sql, err := store.DB.Prepare(`
		DELETE FROM users WHERE ID=?
	`)
	if err != nil {
		fmt.Println("Update services error.")
	}

	trans, err := store.DB.Begin()
	if err != nil {
		fmt.Println(err)
	}

	_, err = trans.Stmt(sql).Exec(user.ID)
	if err != nil {
		fmt.Println("Doing rollback")
		trans.Rollback()
	} else {
		trans.Commit()
	}
}

func (store *DataStore) GetUser(userEmail string) models.UserDataModel {

	user := models.UserDataModel{}
	rows, err := store.DB.Query(`
		SELECT * FROM users
	`)

	if err != nil {
		fmt.Println("Query users failed.")
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
	rows.Close()
	return user
}

func (store *DataStore) ReadUsers() []models.UserDataModel {

	userList := []models.UserDataModel{}
	rows, err := store.DB.Query(`
		SELECT * FROM users
	`)

	if err != nil {
		fmt.Println("Query users failed.")
	}
	var id int
	var email string
	var token string
	var role string
	var services string

	if rows.Next() {
		rows.Scan(&id, &email, &token, &role, &services)
		userDm := models.UserDataModel{
			ID:       id,
			Email:    email,
			Token:    token,
			UserRole: role,
			Services: services,
		}
		userList = append(userList, userDm)
	}
	rows.Close()
	return userList
}
