package main

import (
	playerservice "battleship/internal/player-service"
	"battleship/internal/types"
	"net/http"

	"github.com/labstack/echo"
)

type JsonResponse struct {
	Status  int         `json:"status"`
	Success bool        `json:"success"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

func main() {
	e := echo.New()

	playerOne := playerservice.CreateNewPlayer(1)
	playerTwo := playerservice.CreateNewPlayer(2)

	e.GET("/players", func(c echo.Context) error {
		jsonResponse := JsonResponse{
			Status:  http.StatusOK,
			Success: true,
			Data:    types.Players{Players: []types.Player{playerOne, playerTwo}},
		}
		return c.JSON(http.StatusOK, jsonResponse)
	})
	e.Logger.Fatal(e.Start(":1323"))
}
