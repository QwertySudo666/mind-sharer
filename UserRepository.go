package main

import (
	"database/sql"
	"fmt"
	//"github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func fetchUsers() ([]user, error) {
	var users []user

	rows, err := db.Query("SELECT * FROM users")
	if err != nil {
		return nil, fmt.Errorf("%q", err)
	}
	defer rows.Close()
	for rows.Next() {
		var currentUser user
		if err := rows.Scan(&currentUser.Id, &currentUser.Email, &currentUser.Username, &currentUser.Password); err != nil {
			return nil, fmt.Errorf("%q", err)
		}
		users = append(users, currentUser)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("%q", err)
	}
	fmt.Println(users)
	return users, nil
}

func fetchUserById(id uuid.UUID) (user, error) {
	var existedUser user
	row := db.QueryRow("SELECT * FROM users WHERE id = ?", id)
	if err := row.Scan(&existedUser.Id, &existedUser.Email, &existedUser.Username, &existedUser.Password); err != nil {
		if err == sql.ErrNoRows {
			return existedUser, err
		}
		//return existedUser, err
	}
	return existedUser, nil
}

func persistUser(newUser user) (user, error) {
	dsn := "root:localhost@tcp(127.0.0.1:3306)/mind_sharer?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	db.Create(&newUser)
	//if err != nil {
	//	return &user{}, fmt.Errorf("%v", err)
	//}
	return newUser, err
}
