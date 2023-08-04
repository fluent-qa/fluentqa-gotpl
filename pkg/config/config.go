package config

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"os"
)

var (
	ConfigMapping        = map[string]*viper.Viper{}
	DEFFAULT_CONFIG_FILE = "app.toml"
)

func NewDefaultConfig() *viper.Viper {
	v := viper.New()
	v.SetConfigFile(DEFFAULT_CONFIG_FILE)

	if err := v.ReadInConfig(); err != nil {
		fmt.Printf("couldn't load config: %s", err)
		os.Exit(1)
	}
	ConfigMapping["DEFAULT"] = v
	watchConfigChanges(v, &GenericConfig{})
	return v
}

func watchConfigChanges(v *viper.Viper, config any) {
	v.WatchConfig()
	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)
		fmt.Println(v.AllKeys())
		err := v.Unmarshal(config)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(config)
	})
}

func NewNamedConfig(configFile string, configName string, config any) any {
	v := NewViper(configFile)
	ConfigMapping[configName] = v
	err := v.Unmarshal(config)
	if err != nil {
		return nil
	}
	fmt.Println(err)
	watchConfigChanges(v, config)
	return config
}

func NewViper(configFile string) *viper.Viper {
	v := viper.New()
	v.SetConfigFile(configFile)

	if err := v.ReadInConfig(); err != nil {
		fmt.Printf("couldn't load config: %s", err)
		os.Exit(1)
	}
	return v
}

func GetConfigByViperName(name string) *viper.Viper {
	return ConfigMapping[name]
}

func NewConfig(configFile string, configName string) *GenericConfig {
	v := NewViper(configFile)
	g, config, done := setNamedConfig(configName, v)
	if done {
		return config
	}
	return g
}

func setNamedConfig(configName string, v *viper.Viper) (*GenericConfig, *GenericConfig, bool) {
	ConfigMapping[configName] = v
	g := &GenericConfig{}
	err := v.Unmarshal(9)
	if err != nil {
		return nil, nil, true
	}
	fmt.Println(err)
	watchConfigChanges(v, g)
	return g, nil, false
}
