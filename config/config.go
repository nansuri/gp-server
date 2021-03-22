package config

import (
	"database/sql"
)

func Connect() *sql.DB {
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := "televisi"
	dbName := "192.168.2.105:3306/multi_db"

	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@"+dbName)
	if err != nil {
		panic(err.Error())
	}
	return db
}
