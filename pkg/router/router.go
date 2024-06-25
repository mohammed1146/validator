package router

import (
	"net/http"

	"github.com/gorilla/mux"

	// Application.
	"github.com/mohammed1146/validator/pkg/domain"
	"github.com/mohammed1146/validator/pkg/service"
)

// New is responsible for bootstrapping the application and routing.
func New() http.Handler {
	// Initialize router.
	router := mux.NewRouter().StrictSlash(true)

	// Bootstrap the application dependencies.
	gameHandler := GameHandler{
		gameService:      service.NewGameService(domain.NewStateRepositoryStub()),
		validatorService: service.NewValidatorService(domain.NewStateRepositoryStub()),
	}

	// Application routes.
	// Create new game.
	router.HandleFunc("/new", gameHandler.newGame).Methods(http.MethodGet)

	// Validate game.
	router.HandleFunc("/validate", gameHandler.validateGame).Methods(http.MethodPost)

	return router
}
