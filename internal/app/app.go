package app

import (
	"context"
	"fmt"
	"github.com/sigit14ap/go-kumparan/internal/repository"
	"github.com/sigit14ap/go-kumparan/internal/service"
	"github.com/sigit14ap/go-kumparan/pkg/database/mongodb"
	"github.com/sigit14ap/go-kumparan/pkg/database/redis"
	"net/http"
	"time"

	"github.com/sigit14ap/go-kumparan/internal/config"
	delivery "github.com/sigit14ap/go-kumparan/internal/delivery/http"
	_ "github.com/sigit14ap/go-kumparan/pkg/logging"
	log "github.com/sirupsen/logrus"
)

func Run(configPath string) {
	log.Info("Application start ...")
	log.Info("Logger initialized ...")

	cfg := config.GetConfig(configPath)
	log.Info("Config created ...")

	mongoClient, err := mongodb.NewClient(context.Background(), cfg)

	if err != nil {
		log.Fatal(err)
	}

	log.Info("Mongodb connected ...")
	db := mongoClient.Database(cfg.DB.Database)

	redisClient, err := redis.NewClient(cfg)
	if err != nil {
		log.Fatal(err)
	}
	log.Info("Redis connected ...")

	repos := repository.NewRepositories(db)
	services := service.NewServices(service.Deps{
		Repos:       repos,
		RedisClient: redisClient,
	})

	handlers := delivery.NewHandler(services)
	log.Info("Services, repositories and handlers initialized")

	server := &http.Server{
		Handler:      handlers.Init(),
		Addr:         fmt.Sprintf("%s:%s", cfg.Listen.BindIP, cfg.Listen.Port),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Infof("Server started on  %s:%s", cfg.Listen.BindIP, cfg.Listen.Port)

	log.Fatal(server.ListenAndServe())
}
