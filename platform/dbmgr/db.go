package dbmgr

import (
	"database/sql"
	"fmt"
	"log"
	"sync"
	"wxauth/datatype"
)

const (
	driver = "sqlite3"
	source = "./auth.db"
)

type DbHandle struct {
	DB *sql.DB
}

var handle *DbHandle
var once sync.Once
var conn *sql.DB

func GetInstance() *DbHandle {
	once.Do(func() {
		conn, err := sql.Open(driver, source)
		if err != nil {
			log.Fatal("cannot create database: ", err)
		}
		//handle = &DbHandle{DB: conn}
		handle, err = createTable(conn)
		// TODO - defer conn.Close()
	})
	return handle
}

func createTable(db *sql.DB) (*DbHandle, error) {

	stmt, err := db.Prepare(`
		CREATE TABLE IF NOT EXISTS "users" (
			"ID" INTEGER UNIQUE,
			"Email" TEXT UNIQUE,
			"PassRSA" TEXT,
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
		INSERT INTO "users" (Email, PassRSA, Role, Services) values (?, ?, ?, ?)
	`)
	if err != nil {
		fmt.Println("Insert user error")
	}

	trans, err := handle.DB.Begin()
	if err != nil {
		fmt.Println(err)
	}

	_, err = trans.Stmt(sql).Exec(user.Email, user.PassRSA, user.UserRole, user.Services)
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
	var passrsa string
	var role string
	var services string

	for rows.Next() {

		rows.Scan(&id, &email, &passrsa, &role, &services)
		user = datatype.UserDataModel{
			ID:       id,
			Email:    email,
			PassRSA:  passrsa,
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
	var passrsa string
	var role string
	var services string

	for rows.Next() {
		rows.Scan(&id, &email, &passrsa, &role, &services)
		userDm := datatype.UserDataModel{
			ID:       id,
			Email:    email,
			PassRSA:  passrsa,
			UserRole: role,
			Services: services,
		}
		userList = append(userList, userDm)
	}
	rows.Close()

	return userList, err
}

func CloseConn() {
	conn.Close()
}
