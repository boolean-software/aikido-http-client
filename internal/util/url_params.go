package util

import (
	"fmt"
	"net/url"
	"reflect"
)

func BuildURLParams(params interface{}) (string, error) {
	v := reflect.ValueOf(params)
	if v.Kind() != reflect.Struct {
		return "", fmt.Errorf("params must be a struct")
	}

	values := url.Values{}

	for i := 0; i < v.NumField(); i++ {
		field := v.Type().Field(i)
		urlTag := field.Tag.Get("url")
		if urlTag == "" {
			continue
		}
		value := fmt.Sprintf("%v", v.Field(i).Interface())

		values.Set(urlTag, value)
	}

	return values.Encode(), nil
}
