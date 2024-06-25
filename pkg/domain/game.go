package domain

import "github.com/mohammed1146/validator/pkg/dto"

type Game struct {
	GameID string
	Width  int
	Height int
	Score  int
	Fruit  Fruit
	Snake  Snake
}

// IGameRepository is the data layer for the game table.
type IGameRepository interface {
	NewGame(request dto.NewGameRequest) (*dto.GameResponse, error)
}
