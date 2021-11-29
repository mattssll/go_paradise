package main

import (
	"os"
	"fmt"
	"github.com/gofiber/fiber"
	"github.com/mattssll/go-fiber-gorm-api-rest/book"
	"github.com/mattssll/go-fiber-gorm-api-rest/database"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func helloWorld (c *fiber.Ctx) {
	c.Send("Hello World")
}

func setupRoutes(app *fiber.App) {
	app.Get("api/v1/books", book.GetBooks)
	app.Get("api/v1/book/:id", book.GetBook)
	app.Post("api/v1/book", book.NewBook)
	app.Delete("api/v1/book/:id", book.DeleteBook)
}

func initDatabase() {
	var err error
	database.DBConn, err = gorm.Open("sqlite3", "books.db")
	if err != nil {
		panic("Failed to connect to database")
	}
	fmt.Println("Database connection successfully opened")

	database.DBConn.AutoMigrate(&book.Book{})
	fmt.Println("Database Migrated")
}

func main() {
	app := fiber.New()
	initDatabase()
	defer database.DBConn.Close()   // make sure connections will be closed
	setupRoutes(app)
	if(os.Getenv("PORT") != "") {  // env var coming from docker-compose.yml
		fmt.Println("log: starting webserver in port ", os.Getenv("PORT"))
		app.Listen(os.Getenv("PORT"))	
	} else {
		fmt.Println("log: starting webserver in port 3001")
		app.Listen(3001)
	}
}