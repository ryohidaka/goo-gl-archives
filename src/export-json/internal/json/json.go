package json

import (
	"encoding/json"
	"export-json/pkg/models"
	"os"
)

// ExportToJSON writes the given links data to a minified JSON file.
func ExportToJSON(filename string, links []models.Link) error {
	// Marshal the data into a compact JSON format
	data, err := json.Marshal(links)
	if err != nil {
		return err
	}

	// Write the minified JSON to the file
	return os.WriteFile(filename, data, 0644)
}
