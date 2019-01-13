package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var (
	dbHost = "127.0.0.1"
	dbProtocol = fmt.Sprintf("tcp(%s:3306)", dbHost)
	dbUser = "root"
	dbPass = "root"
	dbName = "graphql"
)

func Store(user *User) error {
	con := fmt.Sprintf("%s:%s@%s/%s?parseTime=true", dbUser, dbPass, dbProtocol, dbName)
	sql, err := sql.Open("mysql", con)
	if err != nil {
		return err
	}
	defer sql.Close()
	_, err = sql.Exec("insert into users (user_id, user_name, description, photo_url, email) values (?,?,?,?,?)",
		user.UserId, user.UserName, user.Description, user.PhotoURL, user.Email)
	if err != nil {
		return err
	}
	return nil
}

func FindUserById(id string) (*User, error) {
	con := fmt.Sprintf("%s:%s@%s/%s?parseTime=true", dbUser, dbPass, dbProtocol, dbName)
	sql, err := sql.Open("mysql", con)
	if err != nil {
		return nil, err
	}
	defer sql.Close()
	var user User
	err = sql.QueryRow("select user_id, user_name from users where user_id = ?", id).Scan(&(user.UserId), &(user.UserName))
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func UserList() (*[]User, error) {
	con := fmt.Sprintf("%s:%s@%s/%s?parseTime=true", dbUser, dbPass, dbProtocol, dbName)
	sql, err := sql.Open("mysql", con)
	if err != nil {
		return nil, err
	}
	defer sql.Close()

	var users []User
	rows, err := sql.Query("select user_id, user_name from users")
	for rows.Next() {
		var user User
		if err := rows.Scan(&(user.UserId), &(user.UserName)); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return &users, nil
}
