/*
@Author: urmsone urmsone@163.com
@Date: 2/24/21 6:18 PM
@Name: config.go
*/
package utils

import (
	"github.com/spf13/viper"
)

var (
	cfgFile = "config"
	cfgDir  = "conf"
)

var v *viper.Viper

func ConfigInit(name, dir string, config interface{}) error {
	v = viper.New()
	if dir != "" {
		v.AddConfigPath(dir)
	} else {
		v.AddConfigPath(cfgDir)
	}
	if name != "" {
		v.SetConfigName(name)
	} else {
		v.SetConfigName(cfgFile)
	}
	v.SetConfigType("yaml")
	if err := v.ReadInConfig(); err != nil {
		return err
	}
	return nil
}

func GetConfig(c interface{}) error {
	if err := v.ReadInConfig(); err != nil {
		return err
	}
	if err := v.Unmarshal(c); err != nil {
		return err
	}
	return nil
}

func ConfigUtils() *viper.Viper {
	return v
}
