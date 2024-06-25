package service

import (
	"github.com/mohammed1146/validator/pkg/domain"
	"github.com/mohammed1146/validator/pkg/dto"
)

type IValidatorService interface {
	ValidateGame(request dto.ValidatorRequest) (*dto.GameResponse, error)
}

type ValidatorService struct {
	gameRepository domain.IGameRepository
}

func NewValidatorService(gameRepository domain.IGameRepository) *ValidatorService {
	return &ValidatorService{gameRepository: gameRepository}
}

// ValidateGame is responsible for running series of checks.
func (v *ValidatorService) ValidateGame(request dto.ValidatorRequest) (*dto.GameResponse, error) {
	// Get the game from db.
	// Generate new position for the fruit
	fruit := generateFoodPosition(request.State.Width, request.State.Height)

	// Send back the response.
	return &dto.GameResponse{
		GameID: request.State.GameID,
		Width:  request.State.Width,
		Height: request.State.Height,
		Score:  request.State.Score + 1,
		Fruit: dto.Fruit{
			X: fruit.X,
			Y: fruit.Y,
		},
		Snake: request.State.Snake,
	}, nil
}
