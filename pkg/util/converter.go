package util

import (
	"fmt"
	"reflect"
	"strings"
	"time"
)

func MapToStructGeneric(target, source interface{}, tag string, mustfit bool) error {
	targetType := reflect.TypeOf(target).Elem()
	if targetType.Kind() != reflect.Struct {
		return fmt.Errorf("%v is not ptr of struct", reflect.TypeOf(target))
	}

	sourceType := reflect.TypeOf(source).Elem()
	if sourceType.Kind() != reflect.Map {
		return fmt.Errorf("%v is not ptr of map", reflect.TypeOf(source))
	}

	sourceVals := reflect.ValueOf(source).Elem()
	if len(sourceVals.MapKeys()) < 1 {
		return fmt.Errorf("source is empty map")
	}

	if mustfit && len(sourceVals.MapKeys()) != targetType.NumField() {
		return fmt.Errorf("mustfit option is true but length of source and target field not same")
	}

	if sourceVals.MapKeys()[0].Kind() != reflect.String {
		return fmt.Errorf("source map key isn't string")
	}

	targetVals := reflect.ValueOf(target).Elem()

	for i := 0; i < targetType.NumField(); i++ {
		key := targetType.Field(i).Tag.Get(tag)

		keys := strings.Split(key, ",")
		if len(keys) > 1 {
			key = strings.Trim(keys[0], " ")
		}

		if key == "" || key == "-" {
			continue
		}

		if !targetVals.Field(i).CanSet() {
			if mustfit {
				return fmt.Errorf("mustfit option is true but there is field that can't set")
			}
			continue
		}

		value := sourceVals.MapIndex(reflect.ValueOf(key))
		if !value.IsValid() {
			//key doesn't exist
			if mustfit {
				return fmt.Errorf("mustfit option is true but there is field that invalid")
			}
			continue
		}

		if value.Type() != targetVals.Field(i).Type() {
			return fmt.Errorf("map type and field must be same")
		}

		iValue := value.Interface()

		//parse time convert
		if targetVals.Field(i).Type() == reflect.TypeOf(time.Now()) && reflect.TypeOf(iValue).Kind() == reflect.String {
			vStr, _ := iValue.(string)
			if tm, err := time.Parse("20060102150405", vStr); err == nil {
				targetVals.Field(i).Set(reflect.ValueOf(tm))
				continue
			} else if tm, err := time.Parse("20060102", vStr); err == nil {
				targetVals.Field(i).Set(reflect.ValueOf(tm))
				continue
			}
		}

		//set to respective type
		targetVals.Field(i).Set(reflect.ValueOf(iValue))
	}
	return nil
}
