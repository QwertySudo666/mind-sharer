package main

import (
	"database/sql"
	"fmt"
	"github.com/google/uuid"
	"strconv"
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

func persistUser(user user) (persistedUUID uuid.UUID, err error) {
	result, err := db.Exec("INSERT INTO users (Id, Email, Username, Password) VALUES (?, ?, ?, ?)",
		user.Id, user.Email, user.Username, user.Password)

	a, _ := result.LastInsertId()
	b, _ := result.RowsAffected()

	fmt.Println("%v", a)
	fmt.Println("%v", b)
	if err != nil {
		return persistedUUID, fmt.Errorf("%v", err)
	}
	id, err := result.LastInsertId()
	persistedUUID = uuid.MustParse(strconv.FormatInt(id, 10))
	if err != nil {
		return persistedUUID, fmt.Errorf("%v", err)
	}
	return persistedUUID, nil
}
