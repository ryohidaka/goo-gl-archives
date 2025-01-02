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
		// Log details of valid links, using safeString to handle nil pointers
		logger.Printf("UID: %s | Redirect URL: %s | Domain: %s | Page Title: %s | HTTP Status Code: %d",
			link.UID, safeString(link.RedirectURL), safeString(link.DomainName), safeString(link.PageTitle), link.HTTPStatus)

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

func safeString(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}
