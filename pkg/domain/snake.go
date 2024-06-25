package domain

type Snake struct {
	X    int `json:"x"`
	Y    int `json:"y"`
	VelX int `json:"velX"` // X velocity of the snake (one of -1, 0, 1)
	VelY int `json:"velY"` // Y velocity of the snake (one of -1, 0, 1)
}
