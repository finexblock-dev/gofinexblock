package main

import (
	"github.com/finexblock-dev/gofinexblock/cmd/backoffice/internal"
	"github.com/finexblock-dev/gofinexblock/pkg/safety"
	"github.com/joho/godotenv"
	"log"

	_ "github.com/finexblock-dev/gofinexblock/cmd/backoffice/docs"
)

// @title						Finexblock backoffice API Documentation
// @version					1.0
// @description				Finexblock backoffice API Documentation
// @securityDefinitions.apikey	BearerAuth
// @type						apiKey
// @in							header
// @name						Authorization
// @termsOfService				http://swagger.io/terms/
// @contact.name				API Support
// @contact.email				fiber@swagger.io
// @license.name				Apache 2.0
// @license.url				http://www.apache.org/licenses/LICENSE-2.0.html
// @host						localhost:8000
// @BasePath					/
func main() {
	log.Println("Launch finexblock backoffice server")

	if err := godotenv.Load(".env"); err != nil {
		log.Fatalf("failed to load env file : %v", err)
	}

	safety.GracefullyStopBootstrap(internal.Bootstrap)
}