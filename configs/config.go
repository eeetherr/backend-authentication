package configs

import (
	"log"

	"github.com/spf13/viper"
)

type JWTConfig struct {
	Secret string `mapstructure:"secret"`
}

type ConfigStruct struct {
	JWT JWTConfig `mapstructure:"jwt"`
}

var Config ConfigStruct

func LoadConfig() {
	viper.SetConfigName("local")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./configs") // or wherever your config is

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error reading config file: %v", err)
	}

	err = viper.Unmarshal(&Config)
	if err != nil {
		log.Fatalf("Unable to decode config into struct: %v", err)
	}
}
