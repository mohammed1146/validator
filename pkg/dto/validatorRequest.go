package dto

import (
	"errors"
)

type State struct {
	GameID string `json:"gameId"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
	Score  int    `json:"score"`
	Fruit  Fruit  `json:"fruit"`
	Snake  Snake  `json:"snake"`
}

type Fruit struct {
	X int
	Y int
}

type Snake struct {
	X    int `json:"x"`
	Y    int `json:"y"`
	VelX int `json:"velX"` // X velocity of the snake (one of -1, 0, 1)
	VelY int `json:"velY"` // Y velocity of the snake (one of -1, 0, 1)
}

type ValidatorRequest struct {
	State State
	Tick  []Fruit
}

// IsValidTicks validate the snake moves with correct sequence.
func (vr ValidatorRequest) IsValidTicks() (bool, error) {
	var valid = true

	for index, tick := range vr.Tick {
		if index == 0 {
			continue
		}

		if !(tick.X-vr.Tick[index-1].X == 1 || tick.X == vr.Tick[index-1].X) || !(tick.Y-vr.Tick[index-1].Y == 1 || tick.Y == vr.Tick[index-1].Y) {
			valid = false
			return valid, errors.New("invalid ticks sequence")
		}
	}

	return valid, nil
}

// IsValidMoves validate the snake does not hit the walls or go out of bounds.
func (vr ValidatorRequest) IsValidMoves() (bool, error) {
	var valid = true

	for _, tick := range vr.Tick {
		if tick.X < 0 || tick.X >= vr.State.Width || tick.Y < 0 || tick.Y >= vr.State.Height {
			valid = false
			return valid, errors.New("game over. Snake went out of bounds or made invalid move")
		}
	}

	return valid, nil
}

// IsEatFruit is to check if the snake ate the fruit.
func (vr ValidatorRequest) IsEatFruit() (bool, error) {
	var equal = false

	for _, tick := range vr.Tick {
		if tick == vr.State.Fruit {
			equal = true
			return equal, nil
		}
	}

	return equal, errors.New("fruit not found, the ticks do not lead the snake to fruit position")
}
