package networkEngine

import (
	"github.com/l-giuliani/networkEngine/components"
)

var initialUserState components.State
var initialGameState map[string]components.State

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

func RegisterGameEndPoint(endpointId string, stateIds []string, endpoint (func())){

}