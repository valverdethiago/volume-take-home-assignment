package config

import (
	"github.com/spf13/viper"
)

type AppConfig struct {
	Port         string `mapstructure:"PORT"`
	ReadTimeout  int    `mapstructure:"READ_TIMEOUT"`
	WriteTimeout int    `mapstructure:"WRITE_TIMEOUT"`
}

// LoadConfig loads config from env
func LoadConfig(path, file string) (AppConfig, error) {
	var config AppConfig
	viper.AddConfigPath(path)
	viper.SetConfigName(file)
	viper.SetConfigType("env")
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		return config, err
	}
	err = viper.Unmarshal(&config)
	return config, err
}
