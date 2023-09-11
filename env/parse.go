package env

import (
	"github.com/spf13/viper"
	"os"
	"reflect"
	"strings"
)

var (
	defaultTAG = "mapstructure"
)

func Unmarshal(confPath string, obj any) error {
	SetViperDefaults(obj)

	viper.SetConfigFile(confPath)

	_ = viper.ReadInConfig()

	viper.AutomaticEnv()
	err := viper.Unmarshal(obj)
	return err
}

func SetDefaultTag(tag string) {
	defaultTAG = tag
}

func SetOsENV(obj any) map[string]interface{} {
	m := map[string]interface{}{}
	v := reflect.Indirect(reflect.ValueOf(obj))
	fields := reflect.VisibleFields(v.Type())

	var fieldTag string
	var tagName string

	for _, field := range fields {
		if field.Anonymous || !field.IsExported() {
			continue
		}

		fieldTag = field.Tag.Get(defaultTAG)
		if fieldTag == "" {
			continue
		}

		tagName = strings.SplitN(fieldTag, ",", 2)[0]

		m[tagName] = os.Getenv(tagName)
	}
	return m
}

func SetViperDefaults(obj any) {
	v := reflect.Indirect(reflect.ValueOf(obj))
	fields := reflect.VisibleFields(v.Type())

	var fieldTag string
	var tagName string

	for _, field := range fields {
		if field.Anonymous || !field.IsExported() {
			continue
		}

		fieldTag = field.Tag.Get(defaultTAG)
		if fieldTag == "" {
			continue
		}

		tagName = strings.SplitN(fieldTag, ",", 2)[0]

		viper.SetDefault(tagName, os.Getenv(tagName))
	}
}
