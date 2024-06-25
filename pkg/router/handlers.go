package router

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/mohammed1146/validator/pkg/dto"
	"github.com/mohammed1146/validator/pkg/httputils"
	"github.com/mohammed1146/validator/pkg/service"
)

// GameHandler is for handling dependencies.
type GameHandler struct {
	gameService      service.IGameService
	validatorService service.IValidatorService
}

// newGame is for creating new game.
func (gh GameHandler) newGame(w http.ResponseWriter, r *http.Request) {
	// Get the query parameters.
	x, err := strconv.Atoi(r.URL.Query().Get("width"))
	if err != nil {
		fmt.Println("error converting string to int")
		return
	}

	var y int
	y, err = strconv.Atoi(r.URL.Query().Get("height"))
	if err != nil {
		fmt.Println("error converting string to int")
		return
	}

	newGameRequest := dto.NewGameRequest{
		X: x,
		Y: y,
	}

	// Call the service
	game, err := gh.gameService.CreateNewGame(newGameRequest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	httputils.WriteResponse(w, http.StatusOK, game)
}

// validateGame is for validate the game.
func (gh GameHandler) validateGame(w http.ResponseWriter, r *http.Request) {
	var request dto.ValidatorRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		httputils.WriteHandlerError(r.Header, httputils.NewBadRequestError("malformed request"), w)
		return
	}

	//TODO
	// Validate the request and make sure all fields required is there.

	// Validate the snake ticks are valid sequence of ticks.
	_, err = request.IsValidTicks()
	if err != nil {
		httputils.WriteHandlerError(r.Header, httputils.NewGameOverError(err), w)
		return
	}

	// Validate the snake does not hit the walls.
	_, err = request.IsValidMoves()
	if err != nil {
		httputils.WriteHandlerError(r.Header, httputils.NewGameOverError(err), w)
		return
	}

	// Validate the ticks if lead to fruit position.
	_, err = request.IsEatFruit()
	if err != nil {
		httputils.WriteHandlerError(r.Header, httputils.NewNotFoundError(err.Error()), w)
		return
	}

	// Send the request to service.
	state, err := gh.validatorService.ValidateGame(request)
	if err != nil {
		httputils.WriteHandlerError(r.Header, httputils.NewGameOverError(err), w)
		return
	}

	// Send back the response.
	httputils.WriteResponse(w, http.StatusOK, state)
}
