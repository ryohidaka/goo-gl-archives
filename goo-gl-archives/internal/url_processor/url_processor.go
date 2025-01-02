package url_processor

import (
	"fmt"
	"goo-gl-archives/pkg/utils"
	"io"
	"net/http"
	"net/url"
	"time"

	"golang.org/x/net/html"
)

// Link represents the structure to hold the final URL information.
type Link struct {
	UID         string `gorm:"primaryKey;unique"`
	RedirectURL *string
	DomainName  *string
	PageTitle   *string
	HTTPStatus  *int
	CreatedAt   time.Time `gorm:"autoCreateTime"` // Automatically set to current time on insert
	UpdatedAt   time.Time `gorm:"autoUpdateTime"` // Automatically set to current time on update
	IsActive    bool
}

// TableName overrides the default table name used by GORM.
func (Link) TableName() string {
	return "links"
}

// ProcessRequest generates a random string, constructs the URL, and retrieves final URL information.
func ProcessRequest() (Link, error) {
	randomStr, err := utils.GenerateRandomString(5, 10)
	if err != nil {
		return Link{}, fmt.Errorf("failed to generate random string: %w", err)
	}

	urlStr := fmt.Sprintf("https://goo.gl/%s", randomStr)
	redirectURL, domain, title, statusCode, err := getRedirectURLInfo(urlStr)
	if err != nil {
		return Link{}, fmt.Errorf("failed to get URL info: %w", err)
	}

	// Check if domain is "goo.gl"
	if domain == "goo.gl" {
		return Link{
			UID:         randomStr,
			RedirectURL: nil,
			DomainName:  nil,
			PageTitle:   nil,
			HTTPStatus:  nil,
			IsActive:    false,
		}, nil
	}

	return Link{
		UID:         randomStr,
		RedirectURL: &redirectURL,
		DomainName:  &domain,
		PageTitle:   title,
		HTTPStatus:  &statusCode,
		IsActive:    true,
	}, nil
}

// getRedirectURLInfo performs an HTTP GET request to the provided URL and follows redirects.
func getRedirectURLInfo(urlStr string) (string, string, *string, int, error) {
	client := &http.Client{
		Timeout: 3 * time.Second, // Set timeout to 3 seconds
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return nil
		},
	}

	const maxRetries = 3
	for attempt := 1; attempt <= maxRetries; attempt++ {
		resp, err := client.Get(urlStr)
		if err != nil {
			if attempt == maxRetries {
				return "", "", nil, 0, fmt.Errorf("failed to get URL after %d attempts: %w", maxRetries, err)
			}
			continue // Retry on error
		}
		defer resp.Body.Close()

		redirectURL := resp.Request.URL.String()
		domain, err := extractDomain(redirectURL)
		if err != nil {
			return "", "", nil, 0, err
		}

		title, err := extractTitle(resp.Body)
		if err != nil {
			return "", "", nil, 0, err
		}

		return redirectURL, domain, title, resp.StatusCode, nil
	}

	return "", "", nil, 0, fmt.Errorf("failed to get URL: %s", urlStr)
}

// extractDomain parses the URL and extracts the domain.
func extractDomain(redirectURL string) (string, error) {
	parsedURL, err := url.Parse(redirectURL)
	if err != nil {
		return "", err
	}
	return parsedURL.Host, nil
}

// extractTitle extracts the page title from an HTML document.
// It returns a pointer to the title string or nil if no title is found.
func extractTitle(body io.Reader) (*string, error) {
	doc, err := html.Parse(body)
	if err != nil {
		return nil, err
	}
	return extractTitleFromNode(doc), nil
}

// extractTitleFromNode recursively searches for the <title> element in the HTML node tree.
func extractTitleFromNode(n *html.Node) *string {
	if n.Type == html.ElementNode && n.Data == "title" && n.FirstChild != nil {
		title := n.FirstChild.Data
		return &title
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if title := extractTitleFromNode(c); title != nil {
			return title
		}
	}
	return nil
}
