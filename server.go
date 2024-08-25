package main

import (
	playerservice "battleship/internal/player-service"
	shipservice "battleship/internal/ship-service"
	"battleship/internal/types"
	"net/http"

	"github.com/go-playground/validator/v10"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	validate := validator.New(validator.WithRequiredStructEnabled())

	playerOne := playerservice.CreateNewPlayer(1)
	playerTwo := playerservice.CreateNewPlayer(2)

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

	e.POST("/player/deploy", func(c echo.Context) error {
		var request types.DeploymentRequest

		c.Bind(&request)

		err := validate.Struct(request)
		if err != nil {
			jsonResponse := types.JsonResponse{
				Status:  http.StatusBadRequest,
				Success: false,
				Message: err.Error(),
			}
			return c.JSON(http.StatusBadRequest, jsonResponse)
		}

		activePlayer := playerservice.GetPlayerFromId(&playerOne, &playerTwo, request.PlayerId)
		activePlayer, err = shipservice.DeployPlayerShip(request, activePlayer)
		if err != nil {
			jsonResponse := types.JsonResponse{
				Status:  http.StatusBadRequest,
				Success: false,
				Message: err.Error(),
			}
			return c.JSON(http.StatusBadRequest, jsonResponse)
		}

		jsonResponse := types.JsonResponse{
			Status:  http.StatusOK,
			Success: true,
			Data:    activePlayer,
			Message: request.ShipType + " sucessfully deployed",
		}
		return c.JSON(http.StatusOK, jsonResponse)
	})

	e.POST("/player/fire", func(c echo.Context) error {
		jsonResponse := types.JsonResponse{
			Status:  http.StatusOK,
			Success: true,
			Data:    types.Players{Players: []types.Player{playerOne, playerTwo}},
		}
		return c.JSON(http.StatusOK, jsonResponse)
	})

	e.Logger.Fatal(e.Start(":1323"))
}
