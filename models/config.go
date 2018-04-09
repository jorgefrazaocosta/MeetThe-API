package models

type TomlConfig struct {
	Owner             OwnerInfo
	Server            Server
	DB                Database `toml:"database"`
	SendGrid          Sendgrid
	DefaultProperties Defaultproperties `toml:"default_properties"`
}

type Server struct {
	Port string
}

type OwnerInfo struct {
	Name string
}

type Database struct {
	Server   string
	Port     int
	Name     string
	User     string
	Password string
}

type Sendgrid struct {
	APIKEY string `toml:"api_key"`
}

type Defaultproperties struct {
	Bundle string
}
