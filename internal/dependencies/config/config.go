package config

import (
	"github.com/cynx-io/cynx-core/src/configuration"
)

var Config *AppConfig

type AppConfig struct {
	Elastic struct {
		Url   string `json:"url"`
		Level string `json:"level"`
	} `json:"elastic"`
	Xendit struct {
		ApiKey     string `mapstructure:"api_key"`
		WebhookKey string `mapstructure:"webhook_key"`
	} `mapstructure:"xendit"`
	Hermes struct {
		Url string `mapstructure:"url"`
	} `mapstructure:"hermes"`
	Ananke struct {
		Url string `mapstructure:"url"`
	} `mapstructure:"ananke"`
	App struct {
		Name    string `mapstructure:"name"`
		Address string `mapstructure:"address"`
		Key     string `mapstructure:"key"`
		Port    int    `mapstructure:"port"`
		Debug   bool   `mapstructure:"debug"`
	} `mapstructure:"app"`
	Database struct {
		Host        string `mapstructure:"host"`
		Database    string `mapstructure:"database"`
		Username    string `mapstructure:"username"`
		Password    string `mapstructure:"password"`
		Dialect     string `mapstructure:"dialect"`
		AutoMigrate bool   `mapstructure:"autoMigrate"`
		Pool        struct {
			Max     int `mapstructure:"max"`
			Min     int `mapstructure:"min"`
			Acquire int `mapstructure:"acquire"`
			Idle    int `mapstructure:"idle"`
		} `mapstructure:"pool"`
		Port int `mapstructure:"port"`
	} `mapstructure:"database"`
}

func Init() {

	Config = &AppConfig{}
	err := configuration.InitConfig("config.json", Config)
	if err != nil {
		panic("failed to initialize config: " + err.Error())
	}
}
