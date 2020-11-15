package dto

type GameEndpointDto struct{
	StateIds []string
	Endpoint func(game *GameDto, data string)
}
