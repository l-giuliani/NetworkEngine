package components

import "errors"

type GamePool struct {
	Games map[string]*Game
}

func NewGamePool() *GamePool{
	gamePool := new(GamePool)
	gamePool.Games = make(map[string]*Game)
	return gamePool
}

func (g *GamePool) AddGame(idGame string, game *Game){
	g.Games[idGame] = game
}

func (g *GamePool) GetGame(idGame string)(*Game,error){
	el, found := g.Games[idGame]
	if found {
		return el, nil
	} else {
		return nil, errors.New("Game Not Found")
	}
}