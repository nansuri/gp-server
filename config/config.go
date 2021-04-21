package config

import (
	"database/sql"
)

func Connect() *sql.DB {
	dbDriver := "mysql"
	dbUser := "gpuser"
	dbPass := "FtTY@ycvpf-tJ][*"
	dbConf := "2.tcp.ngrok.io:12904"
	dbName := "multi_data"

	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@tcp("+dbConf+")/"+dbName)
	if err != nil {
		panic(err.Error())
	}
	return db
}
