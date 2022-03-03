package app

import (
	"fmt"
	"github.com/sigit14ap/go-kumparan/internal/domain"
	"net/http"
	"time"

	"github.com/sigit14ap/go-kumparan/internal/config"
	delivery "github.com/sigit14ap/go-kumparan/internal/delivery/http"
	_ "github.com/sigit14ap/go-kumparan/pkg/logging"
	log "github.com/sirupsen/logrus"
)

func Run(configPath string, command domain.Command) {
	log.Info("Application start ...")
	log.Info("Logger initialized ...")

	cfg := config.GetConfig(configPath)
	log.Info("Config created ...")

	handlers := delivery.NewHandler()
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
