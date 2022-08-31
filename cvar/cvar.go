package cvar

import (
	"fmt"
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

func IsInsideEnum[T any](v T, collection ...T) contracts.Check {
	return IsInsideEnumFunc(collection...)(v)
}

func IsInsideEnumFunc[T any](collection ...T) func(T) contracts.Check {
	return func(v T) contracts.Check {
		return func() (bool, string) {
			for _, item := range collection {
				if reflect.DeepEqual(item, v) {
					return true, ""
				}
			}
			return false, fmt.Sprintf("variable should match the enum")
		}
	}
}
