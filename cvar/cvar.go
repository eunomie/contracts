package cvar

import (
	"reflect"

	"github.com/eunomie/contracts"
)

func IsNil(v any) contracts.Check {
	return func() (bool, string) {
		return isNil(v), "variable should be nil"
	}
}

func IsNotNil(v any) contracts.Check {
	return func() (bool, string) {
		return !isNil(v), "variable should not be nil"
	}
}

func isNil(v any) bool {
	if v == nil {
		return true
	}
	if reflect.ValueOf(v) == reflect.Zero(reflect.TypeOf(v)) {
		return true
	}
	return false
}
