package context

import (
	"github.com/l-giuliani/networkEngine/libs/communicationLibs"
	"github.com/l-giuliani/networkEngine/components"
)


var gamePool *components.GamePool
var commLib *communicationLibs.CommunicationLib

func SetGamePool(pool *components.GamePool){
	gamePool = pool
}

func GetGamePool() *components.GamePool{
	return gamePool
}

func SetCommLib(lib *communicationLibs.CommunicationLib) {
	commLib = lib
}

func GetCommLib() *communicationLibs.CommunicationLib {
	return commLib
}
