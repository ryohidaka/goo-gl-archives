package html

import (
	"fmt"
	"generate-html/internal/database"
	"html/template"
	"strings"
	"unicode/utf8"
)

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

// GenerateTable generates an HTML table from a slice of database.Link structs using a template.
// Each row in the table corresponds to a link's data, formatted with specific columns.
//
// Parameters:
// - links: A slice of database.Link structs containing link data.
//
// Returns:
// - A string representing an HTML table with the link data.
func GenerateTable(links []database.Link) string {
	const tableTemplate = `
<table>
	<thead>
		<tr>
			<th>Index</th>
			<th>UID</th>
			<th>Original URL</th>
			<th>Redirect URL</th>
			<th>Domain Name</th>
			<th>Page Title</th>
			<th>HTTP Status</th>
		</tr>
	</thead>
	<tbody>
		{{- range $i, $link := .Links }}
		<tr>
			<td>{{ $i | addOne }}</td>
			<td><code>{{ $link.UID | sanitize | escape }}</code></td>
			<td><a href="https://goo.gl/{{ $link.UID | sanitize | escape }}" target="_blank" rel="noopener noreferrer">https://goo.gl/{{ $link.UID | sanitize | escape }}</a></td>
			<td><a href="{{ $link.RedirectURL | sanitize | escape }}" target="_blank" rel="noopener noreferrer">{{ $link.RedirectURL | sanitize | escape }}</a></td>
			<td><code>{{ $link.DomainName | sanitize | escape }}</code></td>
			<td><code>{{ if $link.PageTitle }}{{ $link.PageTitle | sanitize | escape }}{{ end }}</code></td>
			<td><code>{{ $link.HTTPStatus }}</code></td>
		</tr>
		{{- end }}
	</tbody>
</table>
`

	// Define custom template functions
	funcMap := template.FuncMap{
		"sanitize": sanitizeUTF8,
		"escape":   escapeHTMLSpecialChars,
		"addOne":   func(i int) int { return i + 1 },
	}

	// Parse and execute the template
	tmpl, err := template.New("table").Funcs(funcMap).Parse(tableTemplate)
	if err != nil {
		return fmt.Sprintf("Error parsing template: %s", err)
	}

	var sb strings.Builder
	if err := tmpl.Execute(&sb, map[string]interface{}{"Links": links}); err != nil {
		return fmt.Sprintf("Error executing template: %s", err)
	}

	return sb.String()
}
