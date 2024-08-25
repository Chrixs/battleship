package types

type (
	JsonResponse struct {
		Status  int         `json:"status"`
		Success bool        `json:"success"`
		Message string      `json:"message,omitempty"`
		Data    interface{} `json:"data,omitempty"`
	}

	DeploymentRequest struct {
		PlayerId   int    `json:"playerId" validate:"required,min=1,max=2"`
		ShipType   string `json:"shipType" validate:"required"`
		GridNumber int    `json:"gridNumber" validate:"required,min=1,max=100"`
		IsVertical bool   `json:"isVertical"`
	}
)
