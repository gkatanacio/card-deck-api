package storage

import (
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	DbHost     string `envconfig:"DB_HOST"`
	DbPort     int    `envconfig:"DB_PORT" default:"5432"`
	DbName     string `envconfig:"DB_NAME"`
	DbUser     string `envconfig:"DB_USER"`
	DbPassword string `envconfig:"DB_PASSWORD"`
}

func NewConfig() *Config {
	c := &Config{}
	envconfig.MustProcess("", c)
	c.validate()

	return c
}

func (c *Config) validate() {
	for _, s := range []string{c.DbHost, c.DbName, c.DbUser, c.DbPassword} {
		if s == "" {
			panic("database details not properly configured")
		}
	}
}
