package domain

import (
	"github.com/mohammed1146/validator/pkg/dto"
)

type StateRepositoryStub struct {
	game dto.GameResponse
}

func NewStateRepositoryStub() *StateRepositoryStub {
	game := dto.GameResponse{
		GameID: "xxxx01",
		Width:  20,
		Height: 20,
		Score:  0,
		Fruit:  dto.Fruit{0, 1},
		Snake: dto.Snake{
			0,
			0,
			0,
			0,
		},
	}

	return &StateRepositoryStub{game}
}

func (s StateRepositoryStub) NewGame(request dto.NewGameRequest) (*dto.GameResponse, error) {
	return &s.game, nil
}
