package drivermodel

import (
	"github.com/l-giuliani/networkEngine/libs/communicationLibs/commInterfaces"
	"github.com/l-giuliani/networkEngine/config"
)

type CommunicationDriver interface {
	Init(config *config.Config)
	Write()
	Read()
	Subscribe(subscriber commInterfaces.CommSubscriber) error
}
