package main

import (
	"database/sql"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

// Initialize MySQL connection
func initDB() {
	var err error
	db, err = sql.Open("mysql", "root:root123@tcp(127.0.0.1:3306)/time_api")
	if err != nil {
		log.Fatal("Error connecting to the database: ", err)
	}

	// Test the database connection
	if err := db.Ping(); err != nil {
		log.Fatal("Error pinging the database: ", err)
	}
	log.Println("Successfully connected to the database.")
}

// Log the current time to the MySQL database
func logTimeToDatabase(currentTime string) error {
	query := "INSERT INTO time_log (timestamp) VALUES (?)"
	_, err := db.Exec(query, currentTime)
	if err != nil {
		log.Printf("Database log failed: %v", err)
		return err
	}
	log.Printf("Successfully logged time: %s", currentTime)
	return nil
}

// Handler for the /current-time API endpoint
func getCurrentTime(c *gin.Context) {
	// Set the time zone to Toronto
	location, err := time.LoadLocation("America/Toronto")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to load Toronto timezone"})
		return
	}

	// Get the current time in Toronto
	currentTime := time.Now().In(location).Format(time.RFC3339)

	// Log the time to the database
	if err := logTimeToDatabase(currentTime); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database log failed"})
		return
	}

	// Return the current time in JSON format
	c.JSON(http.StatusOK, gin.H{"current_time": currentTime})
}

func main() {
	// Initialize the database
	initDB()

	// Set up the Gin router
	r := gin.Default()

	// Define the /current-time API endpoint
	r.GET("/current-time", getCurrentTime)

	// Run the server on port 8080
	if err := r.Run(":8080"); err != nil {
		log.Fatal("Error starting the server: ", err)
	}
}
