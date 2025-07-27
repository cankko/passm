package repository

import (
	"passm/internal/db"
	"passm/internal/model"
)

func GetEntries() ([]model.Entry, error) {
	rows, err := db.DB.Query("SELECT * FROM entries ORDER BY source COLLATE NOCASE ASC")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var entries []model.Entry
	for rows.Next() {
		var entry model.Entry
		if err := rows.Scan(&entry.ID, &entry.Source, &entry.User, &entry.Password); err != nil {
			return nil, err
		}
		entries = append(entries, entry)
	}

	return entries, nil
}

func CreateEntry(source, user, encryptPassword string) error {
	_, err := db.DB.Exec("INSERT INTO entries(user, source, password) VALUES(?,?,?)", user, source, encryptPassword)
	if err != nil {
		return err
	}

	return nil
}

func DeleteEntry(id int) error {
	_, err := db.DB.Exec("DELETE FROM entries WHERE id = ?", id)
	if err != nil {
		return err
	}

	return nil
}

func UpdateEntry(id int, source, user, encryptPassword string) error {
	_, err := db.DB.Exec("UPDATE entries SET user = ?, source = ?, password = ? WHERE id = ?", user, source, encryptPassword, id)
	if err != nil {
		return err
	}

	return nil
}
