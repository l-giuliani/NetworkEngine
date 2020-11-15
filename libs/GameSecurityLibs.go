package libs

import (
	"github.com/l-giuliani/networkEngine/components"
)

func IsUserInGame(idUser string, game *components.Game) bool{
	user, _ := game.FindUser(idUser)
	if user != nil {
		return true
	} else {
		return false
	}
}

func IsUserTurn(idUser string, game *components.Game) bool {
	_, index := game.FindUser(idUser)
	return (uint32(index) == game.GetUserTurn())
}

func PerformSecurityGameCheck(idUser string, game *components.Game) bool {
	return IsUserInGame(idUser, game) && IsUserTurn(idUser, game)
}
