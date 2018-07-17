package app

import (
	"fmt"

	"github.com/spf13/viper"
	"github.com/go-ozzo/ozzo-validation"
	"github.com/DecenterApps/CryptageCodeRedeem/model"
)

type appConfig struct {
	ErrorFile  string `mapstructure:"error_file"`
	Cards      string `mapstructure:"cards"`
	ServerPort int    `mapstructure:"server_port"`
	DSN        string `mapstructure:"dsn"`
}

type Cryptage struct {
	Cards                map[uint]map[uint]model.Card `json:"cards"`
	BoosterAddress       string                       `json:"boosterAddress"`
	CryptageAddress      string                       `json:"cryptageAddress"`
	SeleneanCardsAddress string                       `json:"seleneanCardsAddress"`
	CryptageStartBlock   uint64                       `json:"cryptageStartBlock"`
	HttpProvider         string                       `json:"httpProvider"`
	WsProvider           string                       `json:"wsProvider"`
	Levels               []Level                      `json:"levels"`
	CardsPerLevel        [][]uint                     `json:"cardsPerLevel"`
}

type Level struct {
	Level  uint `json:"level"`
	Exp    uint `json:"exp"`
	Change uint `json:"change"`
}

var Config appConfig

func (config appConfig) Validate() error {
	return validation.ValidateStruct(&config,
		validation.Field(&config.DSN, validation.Required),
	)
}

func LoadConfig(configPaths ...string) error {
	v := viper.New()
	v.SetConfigName("app")
	v.SetConfigType("yaml")
	v.SetEnvPrefix("restful")
	v.AutomaticEnv()
	v.SetDefault("error_file", "config/error.yaml")
	v.SetDefault("cards", "config/cards.yaml")
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
