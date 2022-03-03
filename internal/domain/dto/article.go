package dto

import "time"

type SearchArticleDTO struct {
	Author string
	Search string
}

type ArticleInput struct {
	Author string `json:"author" validate:"required"`
	Title  string `json:"title" validate:"required"`
	Body   string `json:"body" validate:"required"`
}

type ArticleDTO struct {
	Author  string    `json:"author" bson:"author"`
	Title   string    `json:"title" bson:"title"`
	Body    string    `json:"body" bson:"body"`
	Created time.Time `json:"created" bson:"created"`
}
