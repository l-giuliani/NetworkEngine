package commInterfaces

type CommSubscriber interface {
	OnData(jwt string, endpointId string, data string)
}
