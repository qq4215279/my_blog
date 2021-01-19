package utils

import "errors"

func StreamMapToString(list []interface{}, f interface{}) ([]string, error) {
	switch f.(type) {
	case func(interface{}) string:
		res := make([]string, 0)
		for _, e := range list {
			res = append(res, f.(func(interface{}) string)(e))
		}
		return res, nil
	default:
		return nil, errors.New("f is not a function")
	}
}
