package main

import (
	"export-json/internal/database"
	"export-json/internal/json"
	"log"
)

func main() {
	// Initialize the database connection
	db, err := database.InitializeDatabase("../../db/archives.db")
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	// Retrieve active links (is_active == 1)
	links, err := database.GetActiveLinks(db)
	if err != nil {
		log.Fatalf("Failed to retrieve active links: %v", err)
	}

	// Export the retrieved links to a JSON file
	if err := json.ExportToJSON("../../json/archives.json", links); err != nil {
		log.Fatalf("Failed to export to JSON: %v", err)
	}

	log.Println("JSON export completed successfully.")
}
