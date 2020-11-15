package components

import "errors"

type Game struct {
	IdGame string
	NUsers uint32
	CurrentNUsers uint32
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
	game.Turns.CurrentUserTurn = 0
	return game
}

func (g *Game) Adduser(user User) {
	g.Users = append(g.Users, user)
}

func (g *Game) FindUser(idUser string) (*User, int){
	for index, user := range g.Users {
		if user.IdUser == idUser {
			return &user, index
		}
	}
	return nil, 9999
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

func (g* Game) GetUserTurn() uint32{
	return g.Turns.CurrentUserTurn
}

func (g* Game) SetNextTurn(){
	g.Turns.CurrentUserTurn = (g.Turns.CurrentUserTurn + 1) % g.NUsers
}