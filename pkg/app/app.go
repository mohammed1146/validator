package app

import (
	"log"
	"net/http"

	// Application packages.
	"github.com/mohammed1146/validator/pkg/logger"
	"github.com/mohammed1146/validator/pkg/router"
)

func Start() {
	// Load routes.
	routes := router.New()

	// Start the application.
	logger.Info("app start on port 8081")
	log.Fatal(http.ListenAndServe(":8081", routes))
}
