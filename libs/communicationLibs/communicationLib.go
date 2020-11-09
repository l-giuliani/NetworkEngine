package communicationLibs

import (
	"fmt"
	"github.com/l-giuliani/networkEngine/driver/drivermodel"
)

type CommunicationLib struct {
	CommunicationDriver drivermodel.CommunicationDriver
}

func (c *CommunicationLib) OnData(s string) {
	fmt.Println("do stuff")
}

func (c *CommunicationLib) SetCommDriver(commDriver drivermodel.CommunicationDriver){
	c.CommunicationDriver = commDriver
}