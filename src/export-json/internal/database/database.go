package database

import (
	"export-json/pkg/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// InitializeDatabase opens a SQLite database connection using GORM.
func InitializeDatabase(dbName string) (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(dbName), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// Automatically migrate the database schema for the Link model
	if err := db.AutoMigrate(&models.Link{}); err != nil {
		return nil, err
	}

	return db, nil
}

// GetActiveLinks retrieves links where is_active == 1 and http_status == 200,
// and sorts them by uid in ascending order.
func GetActiveLinks(db *gorm.DB) ([]models.Link, error) {
	var links []models.Link
	if err := db.Where("is_active = ? AND http_status = ?", 1, 200).Order("uid ASC").Find(&links).Error; err != nil {
		return nil, err
	}
	return links, nil
}
