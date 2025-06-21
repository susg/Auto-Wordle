package config

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"github.com/susg/autowordle/utils"
)

type Config struct {
	BaseWordsPath        string            `mapstructure:"baseWordsPath"`
	Colours              map[string]string `mapstructure:"colours"`
	MandatoryFile        string            `mapstructure:"mandatoryFile"`
	SupportedWordLengths []int             `mapstructure:"supportedWordLengths"`
	FileChunkSize        int               `mapstructure:"fileChunkSize"`
	WordsBatchSize       int               `mapstructure:"wordsBatchSize"`
}

var AppConfig Config

func init() {
	viperCfg := viper.New()

	viperCfg.SetConfigName("config")
	viperCfg.SetConfigType("yaml")

	projectRoot := utils.FindProjectRoot()

	if flag.Lookup("test.v") != nil || strings.HasSuffix(os.Args[0], ".test") {
		viperCfg.AddConfigPath(filepath.Join(projectRoot, "config/test"))
	} else {
		viperCfg.AddConfigPath(filepath.Join(projectRoot, "config/prod"))
	}

	err := viperCfg.ReadInConfig()
	if err != nil {
		panic(err)
	}

	err = viperCfg.Unmarshal(&AppConfig)
	if err != nil {
		panic(err)
	}

	viperCfg.WatchConfig()
	viperCfg.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)
		if err = viperCfg.Unmarshal(&AppConfig); err != nil {
			fmt.Println(err)
		}
	})
}

func GetConfig() Config {
	return AppConfig
}
