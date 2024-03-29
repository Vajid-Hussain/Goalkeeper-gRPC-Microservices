package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Port       string `mapstructure:"PORT"`
	Auth_port  string `mapstructure:"AUTH_SVC_PORT"`
	Vault_port string `mapstructure:"VAULT_SVC_PORT"`
}

func LoadConfig() (c *Config, err error) {
	viper.AddConfigPath("./")
	viper.SetConfigName("dev")
	viper.SetConfigType("env")

	viper.AutomaticEnv()
	err = viper.ReadInConfig()
	if err != nil {
		return nil, fmt.Errorf("error for getting env file %v", err)
	}

	viper.Unmarshal(&c)
	return
}
