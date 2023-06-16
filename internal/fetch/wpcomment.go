package fetch

// Specs at https://developer.wordpress.org/rest-api/reference/comments/#list-comments
type WPComment struct {
	ID         int    `json:"id"`
	Post       int    `json:"post"`
	Parent     int    `json:"parent"`
	Author     int    `json:"author"`
	AuthorName string `json:"author_name"`
	AuthorURL  string `json:"author_url"`
	Date       string `json:"date"`
	DateGmt    string `json:"date_gmt"`
	Content    struct {
		Rendered string `json:"rendered"`
	} `json:"content"`
	Link             string `json:"link"`
	Status           string `json:"status"`
	Type             string `json:"type"`
	AuthorAvatarURLs struct {
		Size24 string `json:"24"`
		Size48 string `json:"48"`
		Size96 string `json:"96"`
	} `json:"author_avatar_urls"`
	Meta []struct {
		Key   string `json:"key"`
		Value string `json:"value"`
	} `json:"meta"`
	Links struct {
		Self []struct {
			Href string `json:"href"`
		} `json:"self"`
		Collection []struct {
			Href string `json:"href"`
		} `json:"collection"`
		Up []struct {
			Embeddable bool   `json:"embeddable"`
			Href       string `json:"href"`
			PostType   string `json:"post_type"`
		} `json:"up"`
		InReplyTo []struct {
			Embeddable bool   `json:"embeddable"`
			Href       string `json:"href"`
		} `json:"in-reply-to"`
	} `json:"_links"`
}
