package database

import (
	"database/sql"
	"fmt"
	"log"
	"wxauth/datatype"
)

type DbHandle struct {
	DB *sql.DB
}

func CreateTable(db *sql.DB) (*DbHandle, error) {

	stmt, err := db.Prepare(`
		CREATE TABLE IF NOT EXISTS "users" (
			"ID" INTEGER UNIQUE,
			"Email" TEXT UNIQUE,
			"Token" TEXT,
			"Role" TEXT,
			"Services" TEXT,
			PRIMARY KEY ("ID" AUTOINCREMENT)
		);
	`)

	if err != nil {
		log.Fatal("Database not created.")
	}

	_, err = stmt.Exec()

	return &DbHandle{
		DB: db,
	}, err
}

func (handle *DbHandle) InsertUser(user datatype.UserDataModel) error {

	sql, err := handle.DB.Prepare(`
		INSERT INTO "users" (Email, Token, Role, Services) values (?, ?, ?, ?)
	`)
	if err != nil {
		fmt.Println("Insert user error")
	}

	trans, err := handle.DB.Begin()
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

	return err
}

func (handle *DbHandle) UpdateMail(user datatype.UserDataModel, newEmail string) error {

	sql, err := handle.DB.Prepare(`
		UPDATE users SET Email=? WHERE ID=?
	`)
	if err != nil {
		fmt.Println("Update email error.")
	}

	trans, err := handle.DB.Begin()
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

	return err
}

func (handle *DbHandle) UpdateToken(user datatype.UserDataModel, newToken string) error {

	sql, err := handle.DB.Prepare(`
		UPDATE users SET Token=? WHERE ID=?
	`)
	if err != nil {
		fmt.Println("Update token error.")
	}

	trans, err := handle.DB.Begin()
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

	return err
}

func (handle *DbHandle) UpdateServices(user datatype.UserDataModel, services string) error {

	sql, err := handle.DB.Prepare(`
		UPDATE users SET Services=? WHERE ID=?
	`)
	if err != nil {
		fmt.Println("Update services error.")
	}

	trans, err := handle.DB.Begin()
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

	return err
}

func (handle *DbHandle) DeleteUser(userEmail string) error {

	sql, err := handle.DB.Prepare(`
		DELETE FROM users WHERE Email=?
	`)
	if err != nil {
		fmt.Println("Update services error.")
	}

	trans, err := handle.DB.Begin()
	if err != nil {
		fmt.Println(err)
	}

	_, err = trans.Stmt(sql).Exec(userEmail)
	if err != nil {
		fmt.Println("Doing rollback")
		trans.Rollback()
	} else {
		trans.Commit()
	}

	return err
}

func (handle *DbHandle) GetUser(userEmail string) (datatype.UserDataModel, error) {

	user := datatype.UserDataModel{}
	rows, err := handle.DB.Query(`
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
		user = datatype.UserDataModel{
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
	return user, err
}

func (handle *DbHandle) ReadUsers() ([]datatype.UserDataModel, error) {

	userList := []datatype.UserDataModel{}
	rows, err := handle.DB.Query(`
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
		userDm := datatype.UserDataModel{
			ID:       id,
			Email:    email,
			Token:    token,
			UserRole: role,
			Services: services,
		}
		userList = append(userList, userDm)
	}
	rows.Close()

	return userList, err
}
