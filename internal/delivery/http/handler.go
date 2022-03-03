package http

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/sigit14ap/go-kumparan/internal/delivery/http/v1"
	"github.com/sigit14ap/go-kumparan/internal/service"
)

type Handler struct {
	services *service.Services
}

func NewHandler(services *service.Services) *Handler {
	return &Handler{
		services: services,
	}
}

func (h *Handler) Init() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()

	//h.InitAPI(router)

	handlerV1 := v1.NewHandler(h.services)

	api := router.Group("/api/v1")
	{
		api.GET("/article", handlerV1.GetArticle)
		api.POST("/article", handlerV1.CreateArticle)
		api.GET("/article/:id", handlerV1.DetailArticle)
	}

	return router
}

func (h *Handler) InitAPI(router *gin.Engine) {
	handlerV1 := v1.NewHandler(h.services)
	api := router.Group("/api")
	{
		handlerV1.Init(api)
	}
}
