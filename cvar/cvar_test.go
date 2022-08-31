package cvar

import (
	"fmt"
	"testing"

	"gotest.tools/assert"
)

func TestNil(t *testing.T) {
	testCases := []any{
		nil,
		interface{}(nil),
		([]string)(nil),
		(func())(nil),
	}

	for i, tc := range testCases {
		i, tc := i, tc
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			ok, _ := IsNil(tc)()
			assert.Assert(t, ok)

			ok, _ = IsNotNil(tc)()
			assert.Assert(t, !ok)
		})
	}
}

func TestNotNil(t *testing.T) {
	testCases := []any{
		0,
		0.0,
		"",
		[]interface{}{},
		func() {},
	}

	for i, tc := range testCases {
		i, tc := i, tc
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			ok, _ := IsNotNil(tc)()
			assert.Assert(t, ok)

			ok, _ = IsNil(tc)()
			assert.Assert(t, !ok)
		})
	}
}

func TestInsideEnum(t *testing.T) {
	ok, _ := IsInsideEnum(1, 1, 2, 3, 4, 5)()
	assert.Assert(t, ok)

	ok, _ = IsInsideEnum("ipsum", "lorem", "ipsum", "dolor", "sit", "amet")()
	assert.Assert(t, ok)

	ok, _ = IsInsideEnum(1, 2, 3, 4, 5)()
	assert.Assert(t, !ok)

	ok, _ = IsInsideEnum("ipsum", "lorem", "", "dolor", "sit", "amet")()
	assert.Assert(t, !ok)

	type s struct {
		i int
		s string
		a []int
	}

	ok, _ = IsInsideEnum(s{1, "1", []int{1, 2, 3}},
		s{0, "0", []int{1, 2, 3}},
		s{1, "1", []int{1, 2, 3}},
		s{2, "2", []int{1, 2, 3}})()
	assert.Assert(t, ok)

	ok, _ = IsInsideEnum(s{1, "1", []int{1, 2, 3}},
		s{0, "0", []int{1, 2, 3}},
		s{1, "1", []int{1, 2}},
		s{2, "2", []int{1, 2, 3}})()
	assert.Assert(t, !ok)

	ok, _ = IsInsideEnum(&s{1, "1", []int{1, 2, 3}},
		&s{0, "0", []int{1, 2, 3}},
		&s{1, "1", []int{1, 2, 3}},
		&s{2, "2", []int{1, 2, 3}})()
	assert.Assert(t, ok)

	ok, _ = IsInsideEnum(&s{1, "1", []int{1, 2, 3}},
		&s{0, "0", []int{1, 2, 3}},
		&s{1, "1", []int{1, 2, 3, 4}},
		&s{2, "2", []int{1, 2, 3}})()
	assert.Assert(t, !ok)
}
