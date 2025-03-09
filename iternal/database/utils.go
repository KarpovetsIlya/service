package database

import (
	"database/sql"
	"fmt"
)

func TableExists(db *sql.DB, tableName string) bool {
	var name string

	query := fmt.Sprintf(`SELECT to_regclass('%s')`, tableName) // Формируем запрос
	err := db.QueryRow(query).Scan(&name)

	if err != nil || name == "" {
		return false
	}
	return true
}
