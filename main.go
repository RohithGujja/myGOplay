package main

import (
	"fmt"

	"github.com/RohithGujja/myGOplay/services"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db := services.CreateDB()

	// db Check
	version := services.CheckVersion(db)

	fmt.Println("version: ", version)
}
