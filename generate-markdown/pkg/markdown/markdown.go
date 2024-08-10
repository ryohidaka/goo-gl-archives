package markdown

import (
	"fmt"
	"generate-markdown/internal/database"
	"strings"
)

// GenerateTable generates a markdown table from a slice of database.Link structs.
// Each row in the table corresponds to a link's data, formatted with specific columns.
//
// Parameters:
// - links: A slice of database.Link structs containing link data.
//
// Returns:
// - A string representing a markdown table with the link data.
func GenerateTable(links []database.Link) string {
	var sb strings.Builder

	// Write the header row of the markdown table
	sb.WriteString("| Index | UID | Original URL | Redirect URL | Domain Name | Page Title | HTTP Status |\n")
	sb.WriteString("|---|---|---|---|---|---|---|\n")

	// Iterate over the links and add each link's data as a row in the table
	for i, link := range links {
		pageTitle := ""

		// Check if PageTitle is not nil and use its value if available
		if link.PageTitle != nil {
			pageTitle = *link.PageTitle
		}

		// Write a formatted row with the link data
		sb.WriteString(fmt.Sprintf("| %d | `%s` | https://goo.gl/%s | %s | %s | %s | %d |\n",
			i+1, link.UID, link.UID, link.RedirectURL, link.DomainName, pageTitle, link.HTTPStatus))
	}

	// Return the generated markdown table as a string
	return sb.String()
}
