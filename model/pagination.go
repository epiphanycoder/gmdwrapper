package model

type Pagination interface {
	User | Post | Comment
}

type PaginatedResponse struct {
	PageSize int `json:"pageSize"`
	Limit    int `json:"limit"`
}
