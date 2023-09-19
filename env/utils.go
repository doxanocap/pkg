package env

import (
	"reflect"
)

var (
	// priority = 1 -> os env vars important
	// priority = 2 -> env file vars more important
	priority = 1

	defaultTAG = "mapstructure"
)

func SetDefaultTag(tag string) {
	defaultTAG = tag
}

func SetPriorityENV() {
	priority = 2
}

func removeBrackets(str string) string {
	if len(str) < 3 {
		return str
	}

	if str[0] == '"' {
		str = str[1:]

	}
	ln := len(str)
	if str[ln-1] == '"' {
		str = str[:ln-1]
	}
	return str
}

func isIntegerType(kind reflect.Kind) bool {
	return kind == reflect.Int ||
		kind == reflect.Int8 ||
		kind == reflect.Int32 ||
		kind == reflect.Int64 ||
		kind == reflect.Uint ||
		kind == reflect.Uint8 ||
		kind == reflect.Uint16 ||
		kind == reflect.Uint32 ||
		kind == reflect.Uint64
}
