package config

import (
	"fmt"
	"github.com/spf13/viper"
)

func NewViper() *viper.Viper {
	v := viper.New()
	v.SetConfigName("config")
	v.SetConfigType("env")
	v.AddConfigPath(".")
	v.AddConfigPath("./../")

	if err := v.ReadInConfig(); err != nil {
		panic(fmt.Sprintf("Error reading config file %v", err))
	}

	return v
}
