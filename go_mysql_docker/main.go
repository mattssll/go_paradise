package main

import (
	"log"
	"time"
	"fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func getDb() (*sql.DB, error) {
	fmt.Println("log: starting app, waiting 15 secs for database to start")
	time.Sleep(30 * time.Second)
	db, err := sql.Open("mysql", "root:example@tcp(mysqldb:3306)/testdb")
	if err != nil {
		return nil, err
	}
	//defer db.Close()  // making sure db will be closed
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	db.SetMaxOpenConns(20)
    db.SetMaxIdleConns(20)
    db.SetConnMaxLifetime(time.Minute * 5)
	fmt.Println("Successfully connected to mysql")
	return db, nil
}

func main() {
	db, err := getDb()
	if err != nil {
        panic(err.Error())
        return
    }
	insert, err := db.Query("INSERT INTO users (name) values ('Ted')")
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("Successfully added 1 person into 'users' table")
	defer insert.Close()

	// Querying database and returning the data row by row
	var (
		id int
		name string
	)

	rows, err := db.Query("select id, name from users")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&id, &name)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(id, name)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
	
}
