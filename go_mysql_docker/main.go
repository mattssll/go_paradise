package main

import (
	"time"
	"fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)


func main() {
	fmt.Println("log: starting app, waiting 15 secs for database to start")
	time.Sleep(15 * time.Second)
	db, err := sql.Open("mysql", "root:example@tcp(mysqldb:3306)/testdb")
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("Successfully connected to mysql")
	defer db.Close()  // making sure db will be closed

	insert, err := db.Query("INSERT INTO people (name) values ('Ted')")
	if err != nil {
		panic(err.Error())
	}
	defer insert.Close()
	fmt.Println("Successfully added 1 person into people's table")
}
