package dto

type GameRequestDto struct {
	IdGame string `json:"idGame" binding:"required"`
}
