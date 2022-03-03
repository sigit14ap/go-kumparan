package tests

import (
	"context"
	"encoding/json"
	"github.com/restuwahyu13/go-supertest/supertest"
	"github.com/sigit14ap/go-kumparan/internal/config"
	http2 "github.com/sigit14ap/go-kumparan/internal/delivery/http"
	"github.com/sigit14ap/go-kumparan/internal/domain"
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

func TestCreate(t *testing.T) {
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

	payload := map[string]interface{}{
		"author": "kumparan",
		"title":  "Contoh artikel",
		"body":   "ini adalah isi konten artikel",
	}

	var article domain.Article

	test.Post("/api/v1/article")
	test.Send(payload)
	test.Set("Content-Type", "application/json")
	test.End(func(req *http.Request, rr *httptest.ResponseRecorder) {
		var response Response
		json.Unmarshal(rr.Body.Bytes(), &response)

		log.Infof("Test create article : %v",
			assert.Equal(t, "Success", response.Message))
		log.Infof("Hasil create : %v", response.Data)

		temp, err := json.Marshal(response.Data)

		if err != nil {
			log.Fatal(err)
		}

		err = json.Unmarshal([]byte(temp), &article)

		if err != nil {
			log.Fatal(err)
		}
	})

	articleID := article.ID.Hex()

	test.Get("/api/v1/article/" + articleID)
	test.Send(nil)
	test.End(func(req *http.Request, rr *httptest.ResponseRecorder) {
		var response Response
		json.Unmarshal(rr.Body.Bytes(), &response)

		log.Infof("Test detail article : %v",
			assert.Equal(t, "Success", response.Message))
		log.Infof("Hasil detail : %v", response.Data)
	})
}
