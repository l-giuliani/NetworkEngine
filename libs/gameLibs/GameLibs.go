package gameLibs

import (
	"errors"
	"github.com/l-giuliani/networkEngine/components"
	"github.com/l-giuliani/networkEngine/context"
)

func CreateGame(idGame string, nusers uint32) (*components.Game, error){
	gamePool := context.GetGamePool()
	_, err := gamePool.GetGame(idGame)
	if err == nil {
		return nil, errors.New("Game already exist")
	}
	game := components.NewGame(idGame, nusers)
	gamePool.AddGame(idGame, game)
	return game, nil
}

func AddUserToGame(idGame string, idUser string) error{
	gamePool := context.GetGamePool()
	game, err := gamePool.GetGame(idGame)
	if err != nil {
		return errors.New("Game not found")
	}
	user := components.User{IdUser:idUser}
	//TODO set default state
	game.Adduser(user)
	return nil
}

func AddGameState(idGame string, stateId string, state components.State) error{
	gamePool := context.GetGamePool()
	game, err := gamePool.GetGame(idGame)
	if err != nil {
		return errors.New("Game not found")
	}
	game.AddState(stateId, state)
	return nil
}

func GenerateTurns(){

}
