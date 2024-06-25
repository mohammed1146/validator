package service

import (
	"github.com/mohammed1146/validator/pkg/domain"
	"github.com/mohammed1146/validator/pkg/dto"
	"math/rand"
)

// IGameService interface
type IGameService interface {
	CreateNewGame(request dto.NewGameRequest) (*dto.GameResponse, error)
}

// GameService is to inject all dependencies for gameService.
type GameService struct {
	gameRepository domain.IGameRepository
}

func NewGameService(gameRepository domain.IGameRepository) *GameService {
	return &GameService{gameRepository: gameRepository}
}

// CreateNewGame is responsible for creating new game.
func (g *GameService) CreateNewGame(request dto.NewGameRequest) (*dto.GameResponse, error) {
	game, err := g.gameRepository.NewGame(request)
	if err != nil {
		return nil, err
	}

	// Generate food position.
	fruit := generateFoodPosition(request.X, request.Y)

	// Set fruit
	game.Fruit.X = fruit.X
	game.Fruit.Y = fruit.Y

	return game, nil
}

// generateFoodPosition is to generate new position for the fruit.
func generateFoodPosition(width int, height int) *dto.Fruit {
	return &dto.Fruit{rand.Intn(width), rand.Intn(height)}
}
