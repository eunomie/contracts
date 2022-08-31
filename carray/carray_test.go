package carray

import (
	"testing"

	"github.com/eunomie/contracts/cnumber"

	"gotest.tools/assert"
)

func TestNotEmpty(t *testing.T) {
	ok, _ := IsNotEmpty[string]([]string{""})()
	assert.Assert(t, ok)
	ok, _ = IsEmpty[string]([]string{""})()
	assert.Assert(t, !ok)

	ok, _ = IsNotEmpty[int]([]int{1, 2, 3})()
	assert.Assert(t, ok)
	ok, _ = IsEmpty[int]([]int{1, 2, 3})()
	assert.Assert(t, !ok)
}

func TestEmpty(t *testing.T) {
	ok, _ := IsEmpty[string]([]string{})()
	assert.Assert(t, ok)
	ok, _ = IsNotEmpty[string]([]string{})()
	assert.Assert(t, !ok)

	ok, _ = IsEmpty[int]([]int{})()
	assert.Assert(t, ok)
	ok, _ = IsNotEmpty[int]([]int{})()
	assert.Assert(t, !ok)

	ok, _ = IsEmpty[interface{}]([]interface{}{})()
	assert.Assert(t, ok)
	ok, _ = IsNotEmpty[interface{}]([]interface{}{})()
	assert.Assert(t, !ok)

	ok, _ = IsEmpty[func()]([]func(){})()
	assert.Assert(t, ok)
	ok, _ = IsNotEmpty[func()]([]func(){})()
	assert.Assert(t, !ok)

	ok, _ = IsEmpty[string]([]string(nil))()
	assert.Assert(t, ok)
	ok, _ = IsNotEmpty[string]([]string(nil))()
	assert.Assert(t, !ok)
}

func TestOfSize(t *testing.T) {
	ok, _ := OfSize[int]([]int{1, 2, 3, 4}, 4)()
	assert.Assert(t, ok)

	ok, _ = OfSize[int]([]int{1, 2, 3, 4}, 3)()
	assert.Assert(t, !ok)
}

func TestWithEachElement(t *testing.T) {
	ok, _ := WithEachElement[int]([]int{1, 2, 1, 2, 3, 1, 2}, cnumber.IsBetweenFunc(1, 3))()
	assert.Assert(t, ok)

	ok, _ = WithEachElement[int]([]int{1, 2, 1, 2, 3, -1, 2}, cnumber.IsPositiveFunc[int]())()
	assert.Assert(t, !ok)
}
