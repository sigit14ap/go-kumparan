package service

import (
	"context"
	"encoding/json"
	"github.com/go-redis/redis/v7"
	"github.com/sigit14ap/go-kumparan/internal/domain"
	"github.com/sigit14ap/go-kumparan/internal/domain/dto"
	"github.com/sigit14ap/go-kumparan/internal/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type ArticlesService struct {
	repo        repository.Articles
	redisClient *redis.Client
}

func (service *ArticlesService) Get(ctx context.Context, query dto.SearchArticleDTO) ([]domain.Article, error) {
	return service.repo.Get(ctx, query)
}

func (service *ArticlesService) Find(ctx context.Context, articleID primitive.ObjectID) (domain.Article, error) {
	var article domain.Article
	cachedArticle, err := service.redisClient.Get(articleID.Hex()).Result()
	
	if err != nil {
		article, err = service.repo.Find(ctx, articleID)
	} else {
		err = json.Unmarshal([]byte(cachedArticle), &article)
	}

	return article, err
}

func (service *ArticlesService) Create(ctx context.Context, articleDTO dto.ArticleDTO) (domain.Article, error) {
	article, err := service.repo.Create(ctx, articleDTO)

	if err != nil {
		return article, err
	}

	result, err := service.CachingArticle(ctx, article.ID)

	return result, err
}

func (service *ArticlesService) CachingArticle(ctx context.Context, articleID primitive.ObjectID) (domain.Article, error) {

	article, err := service.repo.Find(ctx, articleID)

	if err != nil {
		return article, err
	}

	data, err := json.Marshal(&article)

	if err != nil {
		return article, err
	}

	err = service.redisClient.Set(article.ID.Hex(), data, time.Hour*24*7).Err()

	return article, err
}

func NewArticlesService(repo repository.Articles, redisClient *redis.Client) *ArticlesService {
	return &ArticlesService{
		repo:        repo,
		redisClient: redisClient,
	}
}
