package models

type Post struct {
	Base
	Title  string `json:"title" db:"title"`
	Body   string `json:"body" db:"body"`
	UserID int64  `json:"user_id" db:"user_id"`
}

type GetAllPostsResponse struct {
	Meta Meta    `json:"meta"`
	Data []*Post `json:"data"`
}

type PostRequest struct {
	Title  string `json:"title" db:"title"`
	Body   string `json:"body" db:"body"`
	UserID int    `json:"user_id" db:"user_id"`
}
