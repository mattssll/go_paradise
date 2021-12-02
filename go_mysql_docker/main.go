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
	db, err := sql.Open("mysql", "root:example@tcp(mysqldb:3306)/testdb")
	if err != nil {
		return nil, err
	}
	//defer db.Close()  // making sure db will be closed
	
	for i:=0 ; i<=10; i++ {
		err = db.Ping()
		if err != nil {
			fmt.Println("Connection failed, trying again in 5 secs, will retry 10 times")
			time.Sleep(5 * time.Second)
			continue
		}
		if i==10{
			panic(err.Error()) 
		}
		fmt.Println("Connection successfull")
		break // leave loop if successfully connected
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

	// create table
	createTable, err := db.Query("CREATE TABLE IF NOT EXISTS users (id int primary key auto_increment, name text)")
	if err != nil {
		panic(err.Error())
	}
	defer createTable.Close()
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
