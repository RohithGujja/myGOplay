package main

import (
	"context"
	"fmt"

	"github.com/RohithGujja/myGOplay/services"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	ctx := context.Background()
	db := services.CreateDB()

	// db Check
	version := services.CheckVersion(db)
	fmt.Println("version: ", version)

	// query student table
	student := services.GetStudent(ctx, db, 1)
	fmt.Println("student id: ", student.Id)
	fmt.Println("student name: ", student.Name)

	defer db.Close()
}
