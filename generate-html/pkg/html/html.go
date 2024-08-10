package html

import (
	"fmt"
	"generate-html/internal/database"
	"html/template"
	"io/ioutil"
	"path/filepath"
	"strings"
	"unicode/utf8"
)

// TemplateFuncs defines custom template functions for use in HTML templates.
var TemplateFuncs = template.FuncMap{
	"sanitize": sanitizeUTF8,
	"escape":   escapeHTMLSpecialChars,
	"addOne":   func(i int) int { return i + 1 },
}

// sanitizeUTF8 ensures that a string is valid UTF-8.
// If the string contains invalid UTF-8 sequences, it replaces the string with an empty string.
func sanitizeUTF8(str string) string {
	if !utf8.ValidString(str) {
		return ""
	}
	return str
}

// escapeHTMLSpecialChars escapes special characters in HTML, such as angle brackets.
func escapeHTMLSpecialChars(str string) string {
	str = strings.ReplaceAll(str, "&", "&amp;")
	str = strings.ReplaceAll(str, "<", "&lt;")
	str = strings.ReplaceAll(str, ">", "&gt;")
	str = strings.ReplaceAll(str, `"`, "&quot;")
	return str
}

// LoadTemplateFromFile loads an HTML template from the specified file path.
// It parses the template and returns it along with any error encountered.
//
// Parameters:
// - templatePath: The path to the HTML template file.
//
// Returns:
// - A parsed template and any error encountered during the process.
func LoadTemplateFromFile(templatePath string) (*template.Template, error) {
	// Read the content of the template file
	content, err := ioutil.ReadFile(filepath.Clean(templatePath))
	if err != nil {
		return nil, err
	}

	// Parse the template with custom functions
	tmpl, err := template.New("table").Funcs(TemplateFuncs).Parse(string(content))
	if err != nil {
		return nil, err
	}

	return tmpl, nil
}

// GenerateTable generates an HTML table from a slice of database.Link structs using a template.
// Each row in the table corresponds to a link's data, formatted with specific columns.
//
// Parameters:
// - links: A slice of database.Link structs containing link data.
//
// Returns:
// - A string representing an HTML table with the link data or an error message if the template fails.
func GenerateTable(links []database.Link) string {
	// Path to the template file
	templatePath := "templates/html_template.html"

	// Load and parse the template
	tmpl, err := LoadTemplateFromFile(templatePath)
	if err != nil {
		return fmt.Sprintf("Error loading template: %s", err)
	}

	var sb strings.Builder
	if err := tmpl.Execute(&sb, map[string]interface{}{"Links": links}); err != nil {
		return fmt.Sprintf("Error executing template: %s", err)
	}

	return sb.String()
}
