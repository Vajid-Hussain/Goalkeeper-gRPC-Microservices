package config

import "github.com/spf13/viper"

type Config struct {
	Port        string `mapstructure:"PORT"`
	DBUrl       string `mapstructure:"DB_URL"`
	JwtSecret   string `mapstructure:"JWT_SECRET_KEY"`
	AppPasskey  string `mapstructure:"SECRETKEY"`
	FromEmail   string `mapstructure:"FROM"`
	MailSvcPort string `mapstructure:"MAILSVCPORT"`
}

func LoadConfig() (c Config, err error) {
	viper.AddConfigPath("./")
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
