package driver

import (
	"github.com/l-giuliani/networkEngine/driver/drivermodel"
	"github.com/l-giuliani/networkEngine/driver/socketio"
)

var commDriverMap = map[string](func()drivermodel.CommunicationDriver){
	"socketio": socketio.NewSocketIo,
}

func CommDriverGetter(driverType string) drivermodel.CommunicationDriver{
	x :=commDriverMap[driverType]()
	return x
}