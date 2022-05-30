package config

import "github.com/spf13/viper"

type Config struct {
	PORT            string `mapstructure:"PORT"`
	NATS_CLUSTER    string `mapstructure:"NATS_CLUSTER"`
	SPACES_ENDPOINT string `mapstructure:"SPACES_ENDPOINT"`
	SPACES_TOKEN    string `mapstructure:"SPACES_TOKEN"`
	TOKEN_SECRET    string `mapstructure:"TOKEN_SECRET"`
	GOOGLE_CLIENTID string `mapstructure:"GOOGLE_CLIENTID"`
	DB_NAME         string `mapstructure:"DB_NAME"`
	DB_USER         string `mapstructure:"DB_USER"`
	DB_PW           string `mapstructure:"DB_PW"`
	DB_PORT         uint16 `mapstructure:"DB_PORT"`
	DB_HOST         string `mapstructure:"DB_HOST"`
}

func LoadConfig() (config Config, err error) {
	viper.AddConfigPath("./config/envs")
	viper.SetConfigName("dev")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
