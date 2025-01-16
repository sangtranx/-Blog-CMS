package initialize

import (
	"Blog-CMS/common"
	"fmt"
	"github.com/spf13/viper"
)

func LoadConfig() {
	viper := viper.New()
	viper.SetConfigName("local")     // name of config file (without extension)
	viper.SetConfigType("yaml")      // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath("./config/") // path to look for the config file in

	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	if err := viper.Unmarshal(&common.Config); err != nil {
		fmt.Printf("Unable to decode configuration %v", err)
	}
}
