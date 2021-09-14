package config

import (
	"log"

	"github.com/spf13/viper"
)

type ServerConfig struct {
	DBConfig struct {
		Host     string
		User     string
		Name     string
		Password string
	} `mapstructure:"db_config"`
	StorageBucket  string `mapstructure:"storage_bucket"`
	FirebaseSecret string `mapstructure:"firebase_secret"`
	Mode           string
	Port           string
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
