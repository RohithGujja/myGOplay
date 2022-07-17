package services

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/RohithGujja/myGOplay/model"
)

const studentQuery = "SELECT NAME FROM STUDENT where ID = ?"

func CheckVersion(db *sql.DB) string {
	var (
		err     error
		version string
	)

	err = db.QueryRow("SELECT VERSION()").Scan(&version)
	if err != nil {
		fmt.Println(err.Error())
	}

	return version
}

func GetStudent(ctx context.Context, db *sql.DB, id int) *model.Student {
	var (
		err  error
		name string
	)

	err = db.QueryRowContext(
		ctx,
		studentQuery,
		id,
	).Scan(&name)
	if err != nil {
		fmt.Println(err.Error())
	}

	return &model.Student{
		Id:   id,
		Name: name,
	}
}
