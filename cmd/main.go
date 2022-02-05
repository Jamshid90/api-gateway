package main

import (
	"log"
	"time"

	"github.com/Jamshid90/api-getawey/api"
	configpkg "github.com/Jamshid90/api-getawey/internal/config"
	"github.com/Jamshid90/api-getawey/internal/http/server"
	loggerpkg "github.com/Jamshid90/api-getawey/internal/logger"
	"github.com/Jamshid90/api-getawey/services"
	"go.uber.org/zap"
)

func main() {
	var (
		contextTimeout time.Duration
	)

	// initialization config
	config := configpkg.New()

	// initialization logger
	logger, err := loggerpkg.New(config.LogLevel, config.Environment)
	if err != nil {
		log.Fatal(err)
	}
	defer logger.Sync()

	// context timeout initialization
	contextTimeout, err = time.ParseDuration(config.Context.Timeout)
	if err != nil {
		log.Fatalf("Error during parse duration for context timeout : %v\n", err)
	}

	// services initialization
	service, err := services.New(config)
	if err != nil {
		log.Fatal(err)
	}

	// api handler initialization
	apiHandler := api.New(api.Option{
		Logger:         logger,
		Service:        service,
		ContextTimeout: contextTimeout,
	})

	logger.Info("Listen: ", zap.String("address", config.Server.Host+config.Server.Port))
	log.Fatal(server.NewServer(config, apiHandler).Run())

}
