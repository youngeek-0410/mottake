package config

import (
	"log"

	"github.com/spf13/viper"
)

type ServerConfig struct {
	DbConfig struct {
		Host     string
		User     string
		Name     string
		Password string
	} `mapstructure:"db_config"`
	FirebaseSecret   string `mapstructure:"firebase_secret"`
	Mode             string
	Port             string
	GeocoderClientID string `mapstructure:"geocoder_client_id"`
}

const (
	ModeDevelop    = "DEVELOP"
	ModeProduction = "PRODUCTION"
	ModeDebug      = "DEBUG"
)

var Config *ServerConfig

func Init(configFile string) {
	v := viper.New()
	v.SetConfigType("yaml")
	v.SetConfigName(configFile)
	v.AddConfigPath(".")
	if err := v.ReadInConfig(); err != nil {
		log.Fatal(err)
	}
	Config = &ServerConfig{}
	if err := v.Unmarshal(Config); err != nil {
		log.Fatal(err)
	}
}
