package libs

import (
	"encoding/json"
	"github.com/l-giuliani/networkEngine/config"
	"io/ioutil"
	"os"
)

var c *config.Config = nil

func GetConfig() (*config.Config, error){
	if c == nil {
		conf, err := ioutil.ReadFile(os.Args[1])
		if err != nil {
			return nil, err
		}
		c = new(config.Config)
		err = json.Unmarshal(conf, c)
		return c, err
	} else {
		return c, nil
	}

}
