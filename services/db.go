package services

import (
	"database/sql"
	"fmt"
)

const (
	Host     = "localhost"
	Port     = 3306
	DBName   = "myfirstdb"
	User     = "root"
	Password = "root"
)

func CreateDB() *sql.DB {
	var (
		err error
		db  *sql.DB
	)
	connString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", User, Password, Host, Port, DBName)
	db, err = sql.Open("mysql", connString)
	if err != nil {
		fmt.Println(err.Error())
	}

	return db
}
