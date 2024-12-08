package config

type PasetoConfig struct {
	PrivateKey string `mapstructure:"private_key" json:"private_key"`
	Implicit   string `mapstructure:"implicit" json:"implicit"`
}

type UserServerConfig struct {
	Host       string       `mapstructure:"host" json:"host"`
	Port       int          `mapstructure:"port" json:"port"`
	PasetoInfo PasetoConfig `mapstructure:"paseto" json:"paseto"`
}
