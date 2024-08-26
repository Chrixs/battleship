package main

import (
	playerservice "battleship/internal/player-service"
	shipservice "battleship/internal/ship-service"
	"battleship/internal/types"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"

	"github.com/labstack/echo"
)

type Test struct {
	ID string `param:"id"`
}

func main() {
	e := echo.New()
	validate := validator.New(validator.WithRequiredStructEnabled())

	playerOne := playerservice.CreateNewPlayer(1)
	playerTwo := playerservice.CreateNewPlayer(2)

	e.GET("/", func(c echo.Context) error {
		jsonResponse := types.JsonResponse{
			Status:  http.StatusOK,
			Success: true,
			Data:    "Hello World",
		}
		return c.JSON(http.StatusOK, jsonResponse)
	})

	e.GET("/players", func(c echo.Context) error {
		jsonResponse := types.JsonResponse{
			Status:  http.StatusOK,
			Success: true,
			Data:    types.Players{Players: []types.Player{playerOne, playerTwo}},
		}
		return c.JSON(http.StatusOK, jsonResponse)
	})

	e.GET("/reset", func(c echo.Context) error {
		playerOne = playerservice.CreateNewPlayer(1)
		playerTwo = playerservice.CreateNewPlayer(2)
		jsonResponse := types.JsonResponse{
			Status:  http.StatusOK,
			Success: true,
			Data:    types.Players{Players: []types.Player{playerOne, playerTwo}},
		}
		return c.JSON(http.StatusOK, jsonResponse)
	})

	e.PUT("/player/:id/deploy", func(c echo.Context) error {
		var request types.DeploymentRequest
		var err error

		id := c.Param("id")
		request.PlayerId, err = strconv.Atoi(id)
		if err != nil {
			return c.JSON(http.StatusBadRequest, errorResponse(http.StatusBadRequest, err))
		}

		c.Bind(&request)

		err = validate.Struct(request)
		if err != nil {
			return c.JSON(http.StatusBadRequest, errorResponse(http.StatusBadRequest, err))
		}

		activePlayer, _ := playerservice.GetPlayersFromId(&playerOne, &playerTwo, request.PlayerId)
		deployedShip, err := shipservice.DeployPlayerShip(request, activePlayer)
		if err != nil {
			return c.JSON(http.StatusBadRequest, errorResponse(http.StatusBadRequest, err))
		}

		jsonResponse := types.JsonResponse{
			Status:  http.StatusOK,
			Success: true,
			Data:    deployedShip,
			Message: request.ShipType + " sucessfully deployed",
		}
		return c.JSON(http.StatusOK, jsonResponse)
	})

	e.PUT("/player/:id/fire", func(c echo.Context) error {
		var request types.FireRequest
		var err error

		id := c.Param("id")
		request.PlayerId, err = strconv.Atoi(id)
		if err != nil {
			return c.JSON(http.StatusBadRequest, errorResponse(http.StatusBadRequest, err))
		}

		c.Bind(&request)

		err = validate.Struct(request)
		if err != nil {
			return c.JSON(http.StatusBadRequest, errorResponse(http.StatusBadRequest, err))
		}

		attackingPlayer, defendingPlayer := playerservice.GetPlayersFromId(&playerOne, &playerTwo, request.PlayerId)
		firedShot, err := playerservice.Fire(request.Coordinate, attackingPlayer, defendingPlayer)
		if err != nil {
			return c.JSON(http.StatusBadRequest, errorResponse(http.StatusBadRequest, err))
		}

		jsonResponse := types.JsonResponse{
			Status:  http.StatusOK,
			Success: true,
			Data:    firedShot,
		}
		return c.JSON(http.StatusOK, jsonResponse)
	})

	e.Logger.Fatal(e.Start(":1323"))
}

func errorResponse(status int, err error) types.JsonResponse {
	return types.JsonResponse{
		Status:  status,
		Success: false,
		Message: err.Error(),
	}
}
