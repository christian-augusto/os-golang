package operational_system

import (
	"github.com/spf13/viper"
)

func loadEnvironmentVars() {
	viper.AddConfigPath("./operational_system")

	viper.SetConfigName("config")
	viper.SetConfigType("env")

	viper.ReadInConfig()
}
