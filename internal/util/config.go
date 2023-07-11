package util

import "github.com/spf13/viper"

type Config struct {
	KitchenServerAddress string `mapstructure:"KITCHEN_SERVER_ADDRESS"`
	WaiterServerAddress  string `mapstructure:"WAITER_SERVER_ADDRESS"`
	KafkaServerAddress   string `mapstructure:"KAFKA_ADVERTISED_LISTENERS"`
}

// LoadConfig reads configuration from file or environment variables.
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
