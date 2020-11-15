package dto

import (
	"github.com/l-giuliani/networkEngine/components"
)

type GameEndpointDto struct{
	StateIds []string
	Endpoint func(game *components.Game, data string)
}
