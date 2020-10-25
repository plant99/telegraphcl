package user

type User struct {
	// compulsory
	ShortName  string `json:"short_name"`
	AuthorName string `json:"author_name"`
	AuthorUrl  string `json:"author_url"`
	// optional
	AccessToken string `json:"access_token"`
	AuthUrl     string `json:"auth_url"`
	PageCount   int32  `json:"page_count"`
}

type createUser struct {
	ShortName  string `json:"short_name"`
	AuthorName string `json:"author_name"`
	AuthorUrl  string `json:"author_url"`
}

type createUserResponse struct {
	ShortName   string `json:"short_name"`
	AuthorName  string `json:"author_name"`
	AuthorUrl   string `json:"author_url"`
	AccessToken string `json:"access_token"`
	AuthUrl     string `json:"auth_url"`
}
