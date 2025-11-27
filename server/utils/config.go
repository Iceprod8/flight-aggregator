package utils

import "github.com/spf13/viper"

type Config struct {
	JSERVER1_PORT string `mapstructure:"JSERVER1_PORT"`
	JSERVER1_FILE_PATH string `mapstructure:"JSERVER1_FILE_PATH"` 
	JSERVER1_NAME string `mapstructure:"JSERVER1_NAME"`
	JSERVER2_PORT string `mapstructure:"JSERVER2_PORT"`
	JSERVER2_NAME string `mapstructure:"JSERVER2_NAME"`
	JSERVER2_FILE_PATH string `mapstructure:"JSERVER2_FILE_PATH"`
	SERVER_PORT   string `mapstructure:"SERVER_PORT"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")
	viper.AutomaticEnv()
	err = viper.ReadInConfig()
	if err != nil {
		return
	}
	err = viper.Unmarshal(&config)
	return
}