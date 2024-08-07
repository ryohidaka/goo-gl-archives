package database

import (
	"log"

	"goo-gl-archives/internal/url_processor"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// InitializeDatabase opens a SQLite database connection using GORM and performs migrations.
func InitializeDatabase(dbName string) (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(dbName), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	if err := db.AutoMigrate(&url_processor.Link{}); err != nil {
		return nil, err
	}

	return db, nil
}

// StoreLinks saves or updates links in the database.
// It logs valid links and ignores those with empty UID or RedirectURL.
func StoreLinks(db *gorm.DB, links []url_processor.Link, logger *log.Logger) error {
	for _, link := range links {
		// Skip links with empty UID or RedirectURL
		if link.UID == "" || link.RedirectURL == "" {
			continue
		}

		// Ensure PageTitle is not nil
		if link.PageTitle == nil {
			link.PageTitle = new(string) // or set a default value
		}

		// Log details of valid links
		logger.Printf("UID: %s | Redirect URL: %s | Domain: %s | Page Title: %s | HTTP Status Code: %d",
			link.UID, link.RedirectURL, link.DomainName, *link.PageTitle, link.HTTPStatus)

		// Insert or update link in the database
		if err := db.Clauses(clause.OnConflict{
			UpdateAll: true, // Update all columns on conflict
		}).Create(&link).Error; err != nil {
			logger.Printf("Error saving link with UID %s: %v", link.UID, err)
			return err
		}
	}

	return nil
}
