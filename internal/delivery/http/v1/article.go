package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"github.com/sigit14ap/go-kumparan/internal/delivery/http/services"
	"github.com/sigit14ap/go-kumparan/internal/domain/dto"
	"net/http"
	"time"
)

func (h *Handler) initArticleRoutes(api *gin.RouterGroup) {
	article := api.Group("/article")
	{
		article.GET("/", h.getArticle)
		article.POST("/", h.createArticle)
	}
}

func (h *Handler) getArticle(context *gin.Context) {
	var input dto.SearchArticleDTO

	authorQuery, ok := context.GetQuery("author")

	if ok {
		input.Author = authorQuery
	}

	searchQuery, ok := context.GetQuery("search")

	if ok {
		input.Search = searchQuery
	}

	data, err := h.services.Articles.Get(context, input)

	if err != nil {
		services.ErrorResponse(context, http.StatusInternalServerError, err.Error())
		return
	}

	services.SuccessResponse(context, data)
}

func (h *Handler) createArticle(context *gin.Context) {
	var input dto.ArticleInput
	_ = context.ShouldBindJSON(&input)

	err := validate.Struct(input)
	if err != nil {
		services.ErrorValidationResponse(context, err)
		return
	}

	articleDTO := dto.ArticleDTO{}
	copier.Copy(&articleDTO, &input)

	articleDTO.Created = time.Now()

	data, err := h.services.Articles.Create(context, articleDTO)

	if err != nil {
		services.ErrorResponse(context, http.StatusInternalServerError, err.Error())
		return
	}

	services.CreatedResponse(context, data)
}
