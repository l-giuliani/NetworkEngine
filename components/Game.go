package components

import "errors"

type Game struct {
	IdGame string
	NUsers uint32
	Turns Turns
	Users []User
	State map[string]State
}

func NewGame(idGame string, nusers uint32) *Game{
	game := new(Game)
	game.IdGame = idGame
	game.NUsers = nusers
	game.Users = make([]User, 0)
	game.State = make(map[string]State)
	return game
}

func (g *Game) Adduser(user User) {
	g.Users = append(g.Users, user)
}

func (g *Game) findUser(idUser string) {

}

func (g *Game) AddState(stateId string, state State) {
	g.State[stateId] = state
}

func (g *Game) GetState(stateId string) (State, error) {
	el, found := g.State[stateId]
	if found {
		return el, nil
	} else {
		return nil, errors.New("State Not Found")
	}
}