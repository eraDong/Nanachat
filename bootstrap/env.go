package bootstrap

import (
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
	DBDriver string `mapstructure:"DB_DRIVER"`
	DBSource string `mapstructure:"DB_SOURCE"`
}

func LoadConfig(filePath string) (cfg Config, err error) {
	viper.AddConfigPath(filePath)
	viper.SetEnvPrefix("NANA_CHAT")
	viper.SetConfigName("app")
	viper.SetConfigType("env")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()
	if err = viper.ReadInConfig(); err != nil {
		return
	}
	if err = viper.Unmarshal(&cfg); err != nil {
		return
	}
	return
}
