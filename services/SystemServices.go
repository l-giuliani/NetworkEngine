package services

import (
	"github.com/l-giuliani/networkEngine/components"
	"github.com/l-giuliani/networkEngine/context"
	//"networkEngine/driver"
	"github.com/l-giuliani/networkEngine/driver"
	"github.com/l-giuliani/networkEngine/libs"
	"github.com/l-giuliani/networkEngine/libs/communicationLibs"
	"github.com/l-giuliani/networkEngine/libs/communicationLibs/commInterfaces"
)

func initGamePool(){
	gamePool := components.NewGamePool()
	context.SetGamePool(gamePool)
}

func InitSystem() {
	conf, err := libs.GetConfig()
	if err != nil {
		panic(err)
	}
	comDriverType := conf.CommDriver.DriverType

	commDriver := driver.CommDriverGetter(comDriverType)
	commDriver.Init(conf)

	var commLibInterface commInterfaces.CommSubscriber
	commLib := new(communicationLibs.CommunicationLib)
	commLib.SetCommDriver(commDriver)
	//context.SetCommLib(commLib)
	commLibInterface = commLib
	commDriver.Subscribe(commLibInterface)

	initGamePool()
}
