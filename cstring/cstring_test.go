package cstring

import (
	"regexp"
	"testing"

	"gotest.tools/assert"
)

func TestNotEmpty(t *testing.T) {
	ok, _ := IsNotEmpty("plop")()
	assert.Assert(t, ok)
	ok, _ = IsEmpty("plop")()
	assert.Assert(t, !ok)
}

func TestEmpty(t *testing.T) {
	ok, _ := IsEmpty("")()
	assert.Assert(t, ok)
	ok, _ = IsNotEmpty("")()
	assert.Assert(t, !ok)
}

func TestOfSize(t *testing.T) {
	ok, _ := OfSize("1234", 4)()
	assert.Assert(t, ok)

	ok, _ = OfSize("1234", 3)()
	assert.Assert(t, !ok)
}

func TestShouldMatch(t *testing.T) {
	ok, _ := ShouldMatch("234 a", regexp.MustCompile("[0-9]{2,3} [a-z]"))()
	assert.Assert(t, ok)

	ok, _ = ShouldMatch("2 a", regexp.MustCompile("[0-9]{2,3} [a-z]"))()
	assert.Assert(t, !ok)

	ok, _ = ShouldMatch("2 A", regexp.MustCompile("[0-9]{2,3} [a-z]"))()
	assert.Assert(t, !ok)

	ok, _ = ShouldMatch("1234 aA", regexp.MustCompile("[0-9]{2,3} [a-z]"))()
	assert.Assert(t, ok)

	ok, _ = ShouldMatch("1234 aA", regexp.MustCompile("^[0-9]{2,3} [a-z]$"))()
	assert.Assert(t, !ok)
}
