package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Article struct {
	ID      primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Author  string             `json:"author"`
	Title   string             `json:"title"`
	Body    string             `json:"body"`
	Created time.Time          `json:"created"`
}
