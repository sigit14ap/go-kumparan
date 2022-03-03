package service

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/go-redis/redis/v7"
	"github.com/sigit14ap/go-kumparan/internal/domain"
	"github.com/sigit14ap/go-kumparan/internal/domain/dto"
	"github.com/sigit14ap/go-kumparan/internal/repository"
)

type Articles interface {
	Get(ctx context.Context, query dto.SearchArticleDTO) ([]domain.Article, error)
	Find(ctx context.Context, articleID primitive.ObjectID) (domain.Article, error)
	Create(ctx context.Context, article dto.ArticleDTO) (domain.Article, error)
}

type Services struct {
	Articles Articles
}

type Deps struct {
	Repos       *repository.Repositories
	Services    *Services
	RedisClient *redis.Client
}

func NewServices(deps Deps) *Services {
	articlesService := NewArticlesService(deps.Repos.Articles, deps.RedisClient)

	return &Services{
		Articles: articlesService,
	}
}
