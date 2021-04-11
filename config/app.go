package config

import (
	"github.com/spf13/viper"
	"log"
)

type DevelopConfig struct {
	EnableCheckSignature bool
}

func GetDevelopConfig() (DevelopConfig, error) {
	viper.SetConfigName("config/app.config")
	viper.SetConfigType("toml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("获取配置文件出错!: %v", err)
		return DevelopConfig{}, err
	}
	var config DevelopConfig
	config.EnableCheckSignature = viper.GetBool("develop.enable_check_signature")
	return config, nil
}
