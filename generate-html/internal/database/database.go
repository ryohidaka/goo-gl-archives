package database

import (
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// Link represents the structure to hold the final URL information.
// This struct is used to map the 'links' table in the SQLite database.
type Link struct {
	UID         string    // Unique identifier for the link
	RedirectURL string    // The URL to which the original URL redirects
	DomainName  string    // The domain name of the redirect URL
	PageTitle   *string   // The title of the page, if available
	HTTPStatus  int       // The HTTP status code of the redirect URL
	CreatedAt   time.Time // Timestamp when the record was created
	UpdatedAt   time.Time // Timestamp when the record was last updated
}

// Connect establishes a connection to the SQLite database and performs
// automatic migration to ensure the 'links' table exists.
//
// Returns:
// - A pointer to the gorm.DB instance representing the database connection.
// - An error if there is an issue connecting to the database.
func Connect() (*gorm.DB, error) {
	// Connect to the SQLite database located at '../db/archives.db'
	db, err := gorm.Open(sqlite.Open("../db/archives.db"), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// Perform automatic migration to create the 'links' table if it doesn't exist
	db.AutoMigrate(&Link{})
	return db, nil
}

// GetLinks retrieves all records from the 'links' table in the SQLite database.
//
// Parameters:
// - db: A pointer to the gorm.DB instance representing the database connection.
//
// Returns:
// - A slice of Link structs containing all records from the 'links' table.
// - An error if there is an issue retrieving the records.
func GetLinks(db *gorm.DB) ([]Link, error) {
	var links []Link

	// Execute the query to find all records in the 'links' table
	if err := db.Find(&links).Error; err != nil {
		return nil, err
	}

	// Return the retrieved links
	return links, nil
}
