package service

import (
	"context"
	"github.com/sigit14ap/go-kumparan/internal/domain"
	"github.com/sigit14ap/go-kumparan/internal/domain/dto"
	"github.com/sigit14ap/go-kumparan/internal/repository"
)

type ArticlesService struct {
	repo repository.Articles
}

func (service *ArticlesService) Get(ctx context.Context, query dto.SearchArticleDTO) ([]domain.Article, error) {
	return service.repo.Get(ctx, query)
}

func (service *ArticlesService) Create(ctx context.Context, article dto.ArticleDTO) (domain.Article, error) {
	return service.repo.Create(ctx, article)
}

func NewArticlesService(repo repository.Articles) *ArticlesService {
	return &ArticlesService{
		repo: repo,
	}
}
