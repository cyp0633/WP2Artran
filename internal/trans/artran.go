package trans

// Specs at https://artalk.js.org/guide/transfer.html
type Artran struct {
	ID            int    `json:"id"`
	RID           int    `json:"rid"`
	Content       string `json:"content"`
	UA            string `json:"ua"`
	IP            string `json:"ip"`
	CreatedAt     string `json:"created_at"`
	UpdatedAt     string `json:"updated_at"`
	IsCollapsed   bool   `json:"is_collapsed"`
	IsPending     bool   `json:"is_pending"`
	VoteUp        int    `json:"vote_up"`
	VoteDown      int    `json:"vote_down"`
	Nickname      string `json:"nick"`
	Email         string `json:"email"`
	Link          string `json:"link"`
	Password      string `json:"password"`
	BadgeName     string `json:"badge_name"`
	BadgeColor    string `json:"badge_color"`
	PageKey       string `json:"page_key"`
	PageTitle     string `json:"page_title"`
	PageAdminOnly string `json:"page_admin_only"`
	SiteName      string `json:"site_name"`
	SiteURLs      string `json:"site_urls"`
}
