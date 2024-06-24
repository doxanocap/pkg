package env

import (
	"bufio"
	"fmt"
	"github.com/doxanocap/pkg/errs"
	"os"
	"reflect"
	"strconv"
	"strings"
)

func LoadFile(filePath string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return errs.Wrap("env: open file: ", err)
	}

	scanner := bufio.NewScanner(file)
	for i := 0; scanner.Scan(); i++ {
		line := scanner.Text()
		if line == "" {
			continue
		}

		vars := strings.Split(line, "=")
		if len(vars) < 2 {
			return fmt.Errorf("scanning line %d: invalid format", i)
		}
		key := vars[0]
		value := removeBrackets(line[len(key)+1:])

		temp := os.Getenv(key)
		if temp != "" && priority == 1 {
			continue
		}

		err := os.Setenv(key, value)
		if err != nil {
			return fmt.Errorf("scanning line %d: %v", i, err)
		}
	}

	if err := scanner.Err(); err != nil {
		return errs.Wrap("env: scanner: ", err)
	}

	return nil
}

func Unmarshal(cfg interface{}) error {
	baseValue := reflect.ValueOf(&cfg)
	v := reflect.Indirect(baseValue.Elem().Elem())
	if !v.CanAddr() {
		return fmt.Errorf("value cannot be addressed")
	}

	fields := reflect.VisibleFields(v.Type())
	var (
		fieldTag   string
		fieldValue reflect.Value
		kind       reflect.Kind
		key        string
		value      string
	)

	for _, field := range fields {
		if field.Anonymous || !field.IsExported() {
			continue
		}

		fieldTag = field.Tag.Get("env")
		if fieldTag == "" {
			continue
		}

		fieldValue = v.FieldByIndex(field.Index)
		if !fieldValue.IsValid() {
			return fmt.Errorf("unmarshal: value is not valid: %s", field.Name)
		}

		if !fieldValue.CanAddr() {
			return fmt.Errorf("unmarshal: value cannot be addressed: %s", field.Name)
		}

		if !fieldValue.CanSet() {
			return fmt.Errorf("unmarshal: value cannot be changed: %s", fieldValue)
		}

		key = strings.SplitN(fieldTag, ",", 2)[0]
		value = os.Getenv(key)

		if value == "" {
			return fmt.Errorf("unmarshal: env value not found by key: %s", key)
		}

		kind = field.Type.Kind()

		switch {
		case isIntegerType(kind):
			n, err := strconv.Atoi(value)
			if err != nil {
				return fmt.Errorf("unmarshal: parsed value by key: %s: %v", key, err)
			}
			fieldValue.SetInt(int64(n))
		case kind == reflect.String:
			fieldValue.SetString(value)
		case kind == reflect.Bool:
			fieldValue.SetBool(value == "true")
		default:
			return fmt.Errorf("this type is not valid: %s", kind.String())
		}
	}

	return nil
}
