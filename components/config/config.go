package config

import (
	"fmt"

	"github.com/BurntSushi/toml"

	m "api.meet.the/models"
)

var Data m.TomlConfig

func init() {

	if _, err := toml.DecodeFile("config/config.dev.toml", &Data); err != nil {

		fmt.Println(err)
		return

	}

}
