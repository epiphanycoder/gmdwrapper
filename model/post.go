package model

type Post struct {
	ID          string   `json:"id"`
	Text        string   `json:"text"`
	Image       string   `json:"image"`
	Likes       int      `json:"likes"`
	Link        string   `json:"link"`
	Tags        []string `json:"tags"`
	PublishDate string   `json:"publishDate"`
}
