package operational_system

import "github.com/spf13/viper"

func showLogs() bool {
	return viper.Get(logsEnvName) == booleanEnvValue
}
