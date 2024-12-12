package main

import (
	"app/route"
	"log"
	"os"
	"time"

	"database/sql"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"

	"github.com/gofiber/fiber/v2"
)

func main() {
	err := godotenv.Load()
	// Set timeout configurations
	app := fiber.New(fiber.Config{
		ReadTimeout:  10 * time.Minute,
		WriteTimeout: 10 * time.Minute,
	})

	app.Use(func(c *fiber.Ctx) error {
		start := time.Now()
		err := c.Next()
		duration := time.Since(start)
		c.Append("Server-Timing", "app;dur="+duration.String())
		return err
	})

	// Database connection details
	connString := os.Getenv("DB_URL")

	// Connect to the database
	db, err := sql.Open("postgres", connString)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Define a route
	route.BaseRoutes(app, db)
	route.BaseRoutesTwo(app, db)
	route.SearchLimit(app, db)
	route.SearchLimitTwo(app, db)

	// Start the server
	log.Fatal(app.Listen(":1001"))
}
