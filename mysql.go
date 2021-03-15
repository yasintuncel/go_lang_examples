package main

/*
go get -u github.com/go-sql-driver/mysql
go run mysql.go
*/
import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	fmt.Println("Hello, World!")
	db, err := sql.Open("mysql", "golang:golang@tcp(127.0.0.1:3306)/golang") // "user:password@tcp(127.0.0.1:3306)/dbName"
	if err != nil {
		panic(err)
	}

	defer db.Close()
	insert, err := db.Query("INSERT INTO user (id, userName, userPassword) VALUES (1,'yasintuncel', '123' );")

	if err != nil {
		panic(err.Error())
	}

	defer insert.Close()
}
