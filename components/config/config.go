package config

import (
	"fmt"

	"github.com/BurntSushi/toml"
)

type tomlConfig struct {
	Owner             ownerInfo
	Server            server
	DB                database `toml:"database"`
	SendGrid          sendgrid
	DefaultProperties defaultproperties `toml:"default_properties"`
}

type server struct {
	Port string
}

type ownerInfo struct {
	Name string
}

type database struct {
	Server   string
	Port     int
	Name     string
	User     string
	Password string
}

type sendgrid struct {
	APIKEY string `toml:"api_key"`
}

type defaultproperties struct {
	Bundle string
}

var Data tomlConfig

func init() {

	if _, err := toml.DecodeFile("config/config.dev.toml", &Data); err != nil {

		fmt.Println(err)
		return

	}

}
