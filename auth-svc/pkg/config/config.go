package config

import (
	"crypto/tls"
	"fmt"

	"github.com/spf13/viper"
	"google.golang.org/grpc/credentials"
)

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

func LoadTlsCredential() (credentials.TransportCredentials, error) {
	serverCert, err := tls.LoadX509KeyPair("./../cert/server-key.pem", "./../cert/server-req.pem")
	if err != nil {
		fmt.Println("err at tls fetching the server sertificates")
	}

	config := &tls.Config{
		Certificates: []tls.Certificate{serverCert},
		ClientAuth:   tls.NoClientCert,
	}

	return credentials.NewTLS(config), nil
}
