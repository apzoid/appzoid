package config

import (
	"strings"

	"github.com/spf13/viper"
)

// ViperConfig viper configuration instance
var ViperConfig *viper.Viper

// LoadEnv instantiate viper instance
func LoadEnv() {
	ViperConfig = viper.New()
	ViperConfig.AddConfigPath("./")
	ViperConfig.SetConfigFile(".env")
	ViperConfig.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	_ = ViperConfig.ReadInConfig()
	ViperConfig.AutomaticEnv()
	ViperConfig.WatchConfig()
}
