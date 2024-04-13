package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"

	"github.com/iriskin77/testgo/configs"
	"github.com/iriskin77/testgo/pkg/logging"
	"github.com/iriskin77/testgo/server"
	"github.com/joho/godotenv"
)

// @title Orders API
// @version 1.0
// @description This is a sample service for managing orders
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email soberkoder@gmail.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8000
// @BasePath /
func main() {

	if err := RunServer(); err != nil {
		log.Fatal(err)
	}
	os.Exit(0)

}

func RunServer() error {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
		return err
	}

	// Initialize logger
	logging.InitLogger()

	logger := logging.GetLogger()

	logger.Info("logger has been initialized")

	// Initialize configs
	config := configs.NewConfig()

	// Initialize server, db, routing
	ctx := context.Background()
	srv, err := server.NewHttpServer(ctx, logger, config.Postgres, config.Bindaddr)

	if err != nil {
		logger.Fatal("Failed to start server", err.Error())
		return err
	}

	if err = srv.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
		logger.Fatal("HTTP server ListenAndServe", err.Error())
		return err
	}

	logger.Info("starting API Server")

	return nil
}
