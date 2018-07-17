package app

import (
	"fmt"

	"github.com/spf13/viper"
	"github.com/go-ozzo/ozzo-validation"
)

var Config appConfig

type appConfig struct {
	ErrorFile          string `mapstructure:"error_file"`
	ServerPort         int    `mapstructure:"server_port"`
	DSN                string `mapstructure:"dsn"`
	JWTSigningMethod   string `mapstructure:"jwt_signing_method"`
	JWTSigningKey      string `mapstructure:"jwt_signing_key"`
	JWTVerificationKey string `mapstructure:"jwt_verification_key"`
}

func (config appConfig) Validate() error {
	return validation.ValidateStruct(&config,
		validation.Field(&config.DSN, validation.Required),
		validation.Field(&config.JWTSigningKey, validation.Required),
		validation.Field(&config.JWTVerificationKey, validation.Required),
	)
}

func LoadConfig(configPaths ...string) error {
	v := viper.New()
	v.SetConfigName("app")
	v.SetConfigType("yaml")
	v.SetEnvPrefix("restful")
	v.AutomaticEnv()
	v.SetDefault("error_file", "config/error.yaml")
	v.SetDefault("server_port", 8088)
	v.SetDefault("jwt_signing_method", "HS256")
	for _, path := range configPaths {
		v.AddConfigPath(path)
	}
	if err := v.ReadInConfig(); err != nil {
		return fmt.Errorf("Failed to read the configuration file: %s", err)
	}
	if err := v.Unmarshal(&Config); err != nil {
		return err
	}
	return Config.Validate()
}
