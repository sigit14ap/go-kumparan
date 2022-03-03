package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/sigit14ap/go-kumparan/internal/domain"
	"github.com/sigit14ap/go-kumparan/internal/domain/dto"
	"go.mongodb.org/mongo-driver/mongo"
)

type Articles interface {
	Get(ctx context.Context, query dto.SearchArticleDTO) ([]domain.Article, error)
	Find(ctx context.Context, articleID primitive.ObjectID) (domain.Article, error)
	Create(ctx context.Context, article dto.ArticleDTO) (domain.Article, error)
}

type Repositories struct {
	Articles Articles
}

func NewRepositories(db *mongo.Database) *Repositories {
	return &Repositories{
		Articles: NewArticlesRepo(db),
	}
}
