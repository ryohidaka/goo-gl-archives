package main

import (
	"fmt"
	"log"
	"os"

	"generate-markdown/internal/database"
	"generate-markdown/pkg/markdown"
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

	// Generate a Markdown table from the retrieved data
	mdTable := markdown.GenerateTable(links)

	// Create or overwrite the index.md file in the docs directory
	file, err := os.Create("../docs/index.md")
	if err != nil {
		log.Fatalf("Failed to create the file: %v", err)
	}
	defer file.Close()

	// Write the generated Markdown table to the index.md file
	if _, err := file.WriteString(mdTable); err != nil {
		log.Fatalf("Failed to write to the file: %v", err)
	}

	fmt.Println("index.md has been created.")
}
