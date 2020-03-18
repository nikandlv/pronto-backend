package drivers

import (
	"errors"
	"fmt"
	"github.com/spf13/viper"
)

type viperConfiguration struct {
}

func NewViperConfiguration() viperConfiguration {
	config := viperConfiguration{}
	if err := config.Boot(); err != nil {
		panic(err)
	}
	return config
}

func (v viperConfiguration) Boot() error {
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath(".")

	readErr := viper.ReadInConfig()
	if readErr != nil {
		return readErr
	}
	return nil
}

func (v viperConfiguration) Get(key string) (interface{}, error) {
	val := viper.Get(key)
	if val == nil {
		return nil, errors.New(fmt.Sprint("configuration Entry:%v not found", key))
	}
	return val, nil
}
