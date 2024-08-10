package markdown

import (
	"fmt"
	"generate-markdown/internal/database"
	"strings"
	"unicode/utf8"
)

// sanitizeUTF8 ensures that a string is valid UTF-8.
// If the string contains invalid UTF-8 sequences, it converts it to a valid UTF-8 string.
func sanitizeUTF8(str string) string {
	if !utf8.ValidString(str) {
		return string([]rune(str))
	}
	return str
}

// escapeMarkdownSpecialChars escapes special characters in Markdown, such as pipes.
func escapeMarkdownSpecialChars(str string) string {
	str = strings.ReplaceAll(str, "|", "\\|")
	str = strings.ReplaceAll(str, "`", "\\`")
	str = strings.ReplaceAll(str, "[", "\\[")
	str = strings.ReplaceAll(str, "]", "\\]")
	str = strings.ReplaceAll(str, "!", "\\!")
	return str
}

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
			// Ensure the page title is valid UTF-8 and escape special characters
			pageTitle = sanitizeUTF8(*link.PageTitle)
			pageTitle = escapeMarkdownSpecialChars(pageTitle)
		}

		// Escape other special characters in the link data
		uid := escapeMarkdownSpecialChars(link.UID)
		redirectURL := escapeMarkdownSpecialChars(link.RedirectURL)
		domainName := escapeMarkdownSpecialChars(link.DomainName)

		// Write a formatted row with the link data
		sb.WriteString(fmt.Sprintf("| %d | `%s` | https://goo.gl/%s | ![%s](%s) | `%s` | `%s` | `%d` |\n",
			i+1, uid, uid, redirectURL, redirectURL, domainName, pageTitle, link.HTTPStatus))
	}

	// Return the generated markdown table as a string
	return sb.String()
}
