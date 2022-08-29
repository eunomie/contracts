package carray

import (
	"fmt"

	"github.com/eunomie/contracts"
)

func IsNotEmpty[T any](arr []T) contracts.Check {
	return func() (bool, string) {
		return len(arr) > 0, "array should not be empty"
	}
}

func IsEmpty[T any](arr []T) contracts.Check {
	return func() (bool, string) {
		lenArr := len(arr)
		return lenArr == 0, fmt.Sprintf("array should be empty but is of size %v", lenArr)
	}
}

func OfSize[T any](arr []T, size int) contracts.Check {
	return func() (bool, string) {
		lenArr := len(arr)
		return lenArr == size, fmt.Sprintf("array should be of size %v but is of size %v", size, lenArr)
	}
}

func WithEachElement[T any](arr []T, checks ...func(T) contracts.Check) contracts.Check {
	return func() (bool, string) {
		for i, el := range arr {
			for _, check := range checks {
				if ok, errMsg := check(el)(); !ok {
					return false, fmt.Sprintf("array element %v is not valid: %v", i, errMsg)
				}
			}
		}
		return true, ""
	}
}
