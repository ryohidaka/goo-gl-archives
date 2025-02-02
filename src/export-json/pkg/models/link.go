package models

// Link represents a record in the links table.
type Link struct {
	UID         string `json:"uid"`
	RedirectURL string `json:"redirect_url"`
	DomainName  string `json:"domain_name"`
	PageTitle   string `json:"page_title"`
	HTTPStatus  int    `json:"http_status"`
	IsActive    int    `json:"is_active"`
}
