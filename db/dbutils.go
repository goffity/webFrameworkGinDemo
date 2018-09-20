package db

import (
	"database/sql"
	"fmt"
)

func getDatabaseConnection(username, password, host, port, databaseName string) (*sql.DB, error) {
	var connectionString = ""

	if host == "" && port == "" {
		connectionString = fmt.Sprintf("%s:%s/%s", username, password, databaseName)
	} else {
		connectionString = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", username, password, host, port, databaseName)
	}

	return sql.Open("mysql", connectionString)
}
