package config

import (
	"errors"
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	VaultSvcPort string `mapstructure:"VAULT_SVC_PORT"`
	MailSvcPort string `mapstructure:"PORT"`
	AppPasskey string `mapstructure:"SECRETKEY"`
	FromEmail  string `mapstructure:"FROM"`
}

func LoadConfig() (*Config, error) {

	config := Config{}

	viper.AddConfigPath("./")
	viper.SetConfigName("dev")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		return nil, fmt.Errorf("env is not loaded lead to error %v ",err)
	}

	err= viper.Unmarshal(&config)
	if err!=nil{
		return nil, errors.New("env unmarshel lead to error")
	}
	return &config, nil
}
