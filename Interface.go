package networkEngine

import (
	"github.com/l-giuliani/networkEngine/components"
	"github.com/l-giuliani/networkEngine/dto"
)

var initialUserState components.State
var initialGameState map[string]components.State = make(map[string]components.State)

var gameEndpoints map[string]dto.GameEndpointDto = make(map[string]dto.GameEndpointDto)

func SetInitialUserState(state components.State){
	initialUserState = state
}

func GetInitialUserState() components.State {
	return initialUserState
}

func SetInitialGameState(gameState string, state components.State){
	initialGameState[gameState] = state
}

func GetInitialGameStates() map[string]components.State{
	return initialGameState
}

func RegisterGameEndPoint(endpointId string, gameEndpoint dto.GameEndpointDto){
	gameEndpoints[endpointId] = gameEndpoint
}

func GetGameEndpoints() map[string]dto.GameEndpointDto {
	return gameEndpoints
}