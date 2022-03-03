package tests

import (
	"context"
	"encoding/json"
	"github.com/restuwahyu13/go-supertest/supertest"
	"github.com/sigit14ap/go-kumparan/internal/config"
	http2 "github.com/sigit14ap/go-kumparan/internal/delivery/http"
	"github.com/sigit14ap/go-kumparan/internal/repository"
	"github.com/sigit14ap/go-kumparan/internal/service"
	"github.com/sigit14ap/go-kumparan/pkg/database/mongodb"
	"github.com/sigit14ap/go-kumparan/pkg/database/redis"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

type Response struct {
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

func TestArticle(t *testing.T) {
	cfg := config.GetConfig("../config/config.yml")
	redisClient, err := redis.NewClient(cfg)

	if err != nil {
		log.Fatal(err)
	}

	mongoClient, err := mongodb.NewClient(context.Background(), cfg)

	if err != nil {
		log.Fatal(err)
	}

	db := mongoClient.Database(cfg.DB.Database)

	repos := repository.NewRepositories(db)
	services := service.NewServices(service.Deps{
		Repos:       repos,
		RedisClient: redisClient,
	})

	handlers := http2.NewHandler(services)

	var router = handlers.Init()

	test := supertest.NewSuperTest(router, t)

	test.Get("/api/v1/article")
	test.Send(nil)
	test.Set("Content-Type", "application/json")
	test.End(func(req *http.Request, rr *httptest.ResponseRecorder) {
		var response Response
		json.Unmarshal(rr.Body.Bytes(), &response)

		log.Infof("Test get article : %v",
			assert.Equal(t, "Success", response.Message))
	})

}
