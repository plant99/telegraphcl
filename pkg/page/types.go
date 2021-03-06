package page

type GetPageListRequest struct {
	AccessToken string `json:"access_token"`
	Offset      int32  `json:"offset"`
	Limit       int32  `json:"limit"`
}

// following types were copied as it is from https://github.com/toby3d/telegraph/blob/master/types.go
type PageList struct {
	TotalCount int    `json:"total_count"`
	Pages      []Page `json:"pages"`
}

// Page represents a page on Telegraph.
type Page struct {
	Path        string `json:"path"`
	URL         string `json:"url"`
	Title       string `json:"title"`
	Description string `json:"description"`
	AuthorName  string `json:"author_name,omitempty"`
	AuthorURL   string `json:"author_url,omitempty"`
	ImageURL    string `json:"image_url,omitempty"`
	Content     []Node `json:"content,omitempty"`
	Views       int    `json:"views"`
	CanEdit     bool   `json:"can_edit,omitempty"`
}

type Node interface{}

type NodeElement struct {
	Tag      string            `json:"tag"`
	Attrs    map[string]string `json:"attrs,omitempty"`
	Children []Node            `json:"children,omitempty"`
}

type PageViews struct {
	Views int `json:"views"`
}

type createPageRequest struct {
	// user related fields
	AccessToken string `json:"access_token"`
	AuthorName  string `json:"author_name"`
	AuthorUrl   string `json:"author_url"`
	// page related fields
	Title         string `json:"title"`
	Content       []Node `json:"content"`
	ReturnContent bool   `json:"return_content"`
}

type editPageRequest struct {
	// user related fields
	AccessToken string `json:"access_token"`
	AuthorName  string `json:"author_name"`
	AuthorUrl   string `json:"author_url"`
	// page related fields
	Title         string `json:"title"`
	Content       []Node `json:"content"`
	ReturnContent bool   `json:"return_content"`
	Path          string `json:"path"`
}
