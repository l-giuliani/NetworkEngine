package communicationLibs

import (
	"github.com/l-giuliani/networkEngine"
	"github.com/l-giuliani/networkEngine/driver/drivermodel"
	"github.com/l-giuliani/networkEngine/libs"
)

type CommunicationLib struct {
	CommunicationDriver drivermodel.CommunicationDriver
}

func (c *CommunicationLib) OnData(jwt string, endpointId string, data string) {
	idUser, idGame, err := libs.ParseGameToken(jwt)
	if err != nil {
		return
	}

	gamePool := libs.GetGamePool()
	game, err :=gamePool.GetGame(idGame)
	if err != nil {
		return
	}
	if !libs.PerformSecurityGameCheck(idUser, game){
		return
	}

	gameEndpoints := networkEngine.GetGameEndpoints()
	gameEndpoint, found := gameEndpoints[endpointId]
	if !found {
		return
	}
	gameEndpoint.Endpoint(game, data)

}

func (c *CommunicationLib) SetCommDriver(commDriver drivermodel.CommunicationDriver){
	c.CommunicationDriver = commDriver
}