package utils

import "github.com/spf13/viper"

func InitConfig(configName string) error {
	viper.AddConfigPath("configs")
	viper.SetConfigName(configName)
	return viper.ReadInConfig()
}
