package main

import (
	"os"
	"fmt"
	"time"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {
	conStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s sslmode=disable",
		os.Getenv("POSTGRES_HOST"), os.Getenv("POSTGRES_PORT"),
		os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASSWORD"))
	db, err := gorm.Open("postgres", conStr)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	database := db.DB()
	err = database.Ping()
	fmt.Println("Waiting 10 seconds for database to be initialized")
	time.Sleep(10 * time.Second)
	if err != nil {
		panic(err.Error())
	} 
	fmt.Println("Database was initialized successfully")
}