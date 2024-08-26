package types

type (
	JsonResponse struct {
		Status  int         `json:"status"`
		Success bool        `json:"success"`
		Message string      `json:"message,omitempty"`
		Data    interface{} `json:"data,omitempty"`
	}

	DeploymentRequest struct {
		PlayerId   int    `validate:"required,min=1,max=2"`
		ShipType   string `json:"shipType" validate:"required"`
		Coordinate int    `json:"coordinate" validate:"required,min=1,max=100"`
		IsVertical bool   `json:"isVertical"`
	}

	FireRequest struct {
		PlayerId   int `validate:"required,min=1,max=2"`
		Coordinate int `json:"coordinate" validate:"required,min=1,max=100"`
	}

	FiredShot struct {
		Status   string `json:"status"`
		ShipType string `json:"shipType,omitempty"`
		Winner   bool   `json:"winner,omitempty"`
	}
)
