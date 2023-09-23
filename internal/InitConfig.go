package internal

import (
	"github.com/spf13/viper"
	"github.com/yanguiyuan/yuan/pkg/config"
)

var GLOBAL_CONFIG *viper.Viper = config.NewConfig()
