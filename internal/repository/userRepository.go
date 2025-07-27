package repository

import (
	"passm/internal/db"
	"passm/internal/model"
)

func LoadUsers() ([]model.User, error) {
	rows, err := db.DB.Query("SELECT * FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []model.User
	for rows.Next() {
		var user model.User
		if err := rows.Scan(&user.ID, &user.MainPassword); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func FinMainPasswordById(userID int) (string, error) {
	row, err := db.DB.Query("SELECT main_password FROM users WHERE id = ?", userID)
	if err != nil {
		return "", err
	}
	defer row.Close()

	var entry model.Entry
	err = row.Scan(&entry.Password)
	if err != nil {
		return "", err
	}

	return entry.Password, nil
}
