package config

type Config struct {
	CommDriver CommDriver
}

type CommDriver struct {
	DriverType string
	Params CommDriverParams
}

type CommDriverParams struct {
	ServerPort uint16
}
