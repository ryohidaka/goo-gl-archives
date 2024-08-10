package main

import (
	"fmt"
	"log"
	"os"

	"generate-html/internal/database"
	"generate-html/pkg/html"
)

func main() {
	// Establish a connection to the database
	db, err := database.Connect()
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	// Retrieve the records from the database
	links, err := database.GetLinks(db)
	if err != nil {
		log.Fatalf("Failed to fetch records: %v", err)
	}

	// Generate a HTML table from the retrieved data
	htmltable := html.GenerateTable(links)

	// Create or overwrite the index.html file in the docs directory
	file, err := os.Create("../docs/index.html")
	if err != nil {
		log.Fatalf("Failed to create the file: %v", err)
	}
	defer file.Close()

	// Write the generated HTML table to the index.html file
	if _, err := file.WriteString(htmltable); err != nil {
		log.Fatalf("Failed to write to the file: %v", err)
	}

	fmt.Println("index.html has been created.")
}
