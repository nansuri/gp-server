package config

import (
	"database/sql"
)

func Connect() *sql.DB {
	dbDriver := "mysql"
	dbUser := "gpuser"
	dbPass := "FtTY@ycvpf-tJ][*"
	dbConf := "192.168.2.105:3306"
	dbName := "multi_data"

	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@tcp("+dbConf+")/"+dbName)
	if err != nil {
		panic(err.Error())
	}
	return db
}
