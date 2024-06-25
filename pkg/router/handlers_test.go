package router

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/mohammed1146/validator/pkg/domain"
	"github.com/mohammed1146/validator/pkg/dto"
	"github.com/mohammed1146/validator/pkg/httputils"
	"github.com/mohammed1146/validator/pkg/service"
)

// Mocks
type mockGameRepository struct {
	mock.Mock
}

func newMockGameRepository() *mockGameRepository {
	return &mockGameRepository{}
}

func (c *mockGameRepository) NewGame(request dto.NewGameRequest) (*dto.GameResponse, error) {
	args := c.Called(request)
	return args.Get(0).(*dto.GameResponse), args.Error(1)
}

func TestValidateGame(t *testing.T) {
	tests := []struct {
		desc               string
		input              string
		postParam          map[string]interface{}
		gameRepository     *mockGameRepository
		expectedResult     dto.GameResponse
		expectedErr        *httputils.Error
		expectedStatusCode int
	}{
		{
			desc: "case 1: test getting game with updated score.",
			input: `{
			   "state": {
						"gameId": "xxxx01",
						"width": 20,
						"height": 20,
						"score": 0,
						"fruit": {
							"x": 1,
							"y": 3
						},
						"snake": {
							"x": 1,
							"y": 3,
							"velX": 0,
							"velY": 0
						}
			   },
				"tick" : [
					{
						"x": 0,
						"y": 1
					},
					{
						"x": 0,
						"y": 2
					},
					 {
						"x": 1,
						"y": 2
					},
					{
						"x": 1,
						"y": 3
					}
				]
			}`,
			gameRepository: func() *mockGameRepository {
				game := dto.GameResponse{
					GameID: "xxxx01",
					Width:  20,
					Height: 20,
					Score:  0,
					Fruit: dto.Fruit{
						X: 0,
						Y: 1},
					Snake: dto.Snake{
						1,
						3,
						0,
						0,
					},
				}

				a := newMockGameRepository()
				a.
					On("NewGame", mock.Anything).
					Return(&game, nil)
				return a
			}(),
			expectedResult: dto.GameResponse{
				GameID: "xxxx01",
				Width:  20,
				Height: 20,
				Score:  1,
				Fruit:  dto.Fruit{0, 1},
				Snake: dto.Snake{
					1,
					3,
					0,
					0,
				},
			},
			expectedErr:        nil,
			expectedStatusCode: http.StatusOK,
		},
		{
			desc: "case 2: test getting 418 because of invalid ticks sequence",
			input: `{
			   "state": {
						"gameId": "xxxx01",
						"width": 20,
						"height": 20,
						"score": 0,
						"fruit": {
							"x": 1,
							"y": 3
						},
						"snake": {
							"x": 1,
							"y": 3,
							"velX": 0,
							"velY": 0
						}
			   },
				"tick" : [
					{
						"x": 0,
						"y": 1
					},
					{
						"x": 0,
						"y": 2
					},
					 {
						"x": 1,
						"y": 2
					},
					{
						"x": 21,
						"y": 33
					}
				]
			}`,
			gameRepository:     newMockGameRepository(),
			expectedResult:     dto.GameResponse{},
			expectedErr:        httputils.NewGameOverError(errors.New("invalid ticks sequence")),
			expectedStatusCode: http.StatusTeapot,
		},
		{
			desc: "case 3: test getting 418 because of snake go out of bounds",
			input: `{
			   "state": {
						"gameId": "xxxx01",
						"width": 2,
						"height": 2,
						"score": 0,
						"fruit": {
							"x": 1,
							"y": 3
						},
						"snake": {
							"x": 1,
							"y": 3,
							"velX": 0,
							"velY": 0
						}
			   },
				"tick" : [
					{
						"x": 0,
						"y": 1
					},
					{
						"x": 0,
						"y": 2
					},
					 {
						"x": 1,
						"y": 2
					},
					{
						"x": 1,
						"y": 3
					}
				]
			}`,
			gameRepository:     newMockGameRepository(),
			expectedResult:     dto.GameResponse{},
			expectedErr:        httputils.NewGameOverError(errors.New("game over. Snake went out of bounds or made invalid move")),
			expectedStatusCode: http.StatusTeapot,
		},
		{
			desc: "case 4: test getting 404 because of snake does not eat the fruit",
			input: `{
			   "state": {
						"gameId": "xxxx01",
						"width": 3,
						"height": 3,
						"score": 0,
						"fruit": {
							"x": 0,
							"y": 3
						},
						"snake": {
							"x": 1,
							"y": 3,
							"velX": 0,
							"velY": 0
						}
			   },
				"tick" : [
					{
						"x": 0,
						"y": 1
					},
					{
						"x": 0,
						"y": 2
					},
					 {
						"x": 1,
						"y": 2
					},
					{
						"x": 2,
						"y": 2
					}
				]
			}`,
			gameRepository:     newMockGameRepository(),
			expectedResult:     dto.GameResponse{},
			expectedErr:        httputils.NewNotFoundError("fruit not found, the ticks do not lead the snake to fruit position"),
			expectedStatusCode: http.StatusNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			// given
			req := httptest.NewRequest(http.MethodPost, "/validate", strings.NewReader(tt.input))
			rr := httptest.NewRecorder()

			// Bootstrap the application dependencies.
			gameHandler := GameHandler{
				gameService:      service.NewGameService(domain.NewStateRepositoryStub()),
				validatorService: service.NewValidatorService(domain.NewStateRepositoryStub()),
			}

			// when
			gameHandler.validateGame(rr, req)

			// then
			res := rr.Result()
			defer res.Body.Close()

			assert.Equal(t, tt.expectedStatusCode, res.StatusCode)
			bdy, err := io.ReadAll(res.Body)
			assert.NoError(t, err)

			// check the error
			if tt.expectedStatusCode != http.StatusOK {
				var actualErr httputils.Error
				err = json.Unmarshal(bdy, &actualErr)
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedErr, &actualErr)
			} else {
				var actualResult dto.GameResponse
				err = json.Unmarshal(bdy, &actualResult)
				assert.NoError(t, err)

				// TODO revamp this part.
				// Special case for match new fruit generation
				tt.expectedResult.Fruit.X = actualResult.Fruit.X
				tt.expectedResult.Fruit.Y = actualResult.Fruit.Y

				assert.Equal(t, tt.expectedResult, actualResult)
			}
		})
	}
}
