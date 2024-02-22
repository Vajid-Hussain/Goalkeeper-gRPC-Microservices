package config

import "github.com/spf13/viper"

type Config struct {
	Port      string `mapstructure:"PORT"`
	DBUrl     string `mapstructure:"DB_URL"`
	JwtSecret string `mapstructure:"JWT_SECRET_KEY"`
}

func LoadConfig() (c Config, err error) {
	viper.AddConfigPath("./pkg/config/envs/")
	viper.SetConfigName("dev")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	viper.Unmarshal(&c)
	return
}
