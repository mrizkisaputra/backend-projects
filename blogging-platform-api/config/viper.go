package config

import (
	"fmt"
	"github.com/spf13/viper"
)

func NewViper() *viper.Viper {
	v := viper.New()
	v.SetConfigType("env")
	v.SetConfigName("config")
	v.AddConfigPath(".")
	v.AddConfigPath("./../")

	if err := v.ReadInConfig(); err != nil {
		panic(fmt.Sprintf("Fatal error config file:: %s", err.Error()))
	}

	return v
}
