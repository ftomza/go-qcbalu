package viperplus

import (
	"log"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type Viper struct {
	*viper.Viper
}

func New(name string) *Viper {
	config := viper.New()
	config.SetConfigName(name + "_config")
	config.SetConfigType("yaml")
	config.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			panic(err)
		}
	}
	go config.WatchConfig()
	config.OnConfigChange(func(e fsnotify.Event) {
		log.Println("config for " + name + " is change")
	})
	config.SetEnvPrefix(name)
	viperPlus := &Viper{config}
	viperPlus.SetX("env", "dev")
	viperPlus.SetX("debug", true)
	return viperPlus
}

func (v *Viper) Set(name string, def ...interface{}) error {
	if len(def) > 0 {
		v.SetDefault(name, def[0])
	}
	err := v.BindEnv(name)
	return err
}

func (v *Viper) SetX(name string, def ...interface{}) {
	err := v.Set(name, def...)
	if err != nil {
		panic(err)
	}
}

func (v *Viper) GetEnv() string {
	return v.GetString("env")
}

func (v *Viper) GetDebug() bool {
	return v.GetBool("debug")
}

func (v *Viper) IsProd() bool {
	return v.GetEnv() != "dev"
}
