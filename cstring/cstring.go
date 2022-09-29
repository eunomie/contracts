package cstring

import (
	"fmt"
	"regexp"

	"github.com/eunomie/contracts"
)

func IsNotEmpty(s string) contracts.Check {
	return func() (bool, string) {
		return len(s) > 0, "string should not be empty"
	}
}

func IsEmpty(s string) contracts.Check {
	return func() (bool, string) {
		return len(s) == 0, fmt.Sprintf("string %q should be empty", s)
	}
}

func OfSize(s string, size int) contracts.Check {
	return func() (bool, string) {
		return len(s) == size, fmt.Sprintf("string should be of size %v but is of size %v", size, len(s))
	}
}

func ShouldMatch(s string, r *regexp.Regexp) contracts.Check {
	return func() (bool, string) {
		return r.MatchString(s), fmt.Sprintf("string should match %v", r.String())
	}
}
