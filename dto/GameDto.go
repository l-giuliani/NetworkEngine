package dto

import (
	"github.com/l-giuliani/networkEngine/components"
)

type GameDto struct {
	Users *[]components.User
	State *map[string]components.State
}
