package cnumber

import (
	"testing"

	"gotest.tools/assert"
)

func TestPositive(t *testing.T) {
	ok, _ := IsPositive[int](0)()
	assert.Assert(t, ok)

	ok, _ = IsPositive[int](1)()
	assert.Assert(t, ok)

	ok, _ = IsPositive[float32](0.0)()
	assert.Assert(t, ok)

	ok, _ = IsPositive[int](-1)()
	assert.Assert(t, !ok)
}

func TestStrictlyPositive(t *testing.T) {
	ok, _ := IsStrictlyPositive[int](0)()
	assert.Assert(t, !ok)

	ok, _ = IsStrictlyPositive[int](1)()
	assert.Assert(t, ok)

	ok, _ = IsStrictlyPositive[float32](0.0)()
	assert.Assert(t, !ok)

	ok, _ = IsStrictlyPositive[int](-1)()
	assert.Assert(t, !ok)
}

func TestNegative(t *testing.T) {
	ok, _ := IsNegative[int](0)()
	assert.Assert(t, ok)

	ok, _ = IsNegative[int](-1)()
	assert.Assert(t, ok)

	ok, _ = IsNegative[float64](0.1)()
	assert.Assert(t, !ok)
}

func TestStrictlyNegative(t *testing.T) {
	ok, _ := IsStrictlyNegative[int](0)()
	assert.Assert(t, !ok)

	ok, _ = IsStrictlyNegative[int](-1)()
	assert.Assert(t, ok)

	ok, _ = IsStrictlyNegative[float64](0.1)()
	assert.Assert(t, !ok)
}

func TestBetween(t *testing.T) {
	ok, _ := IsBetween(1, 0, 2)()
	assert.Assert(t, ok)

	ok, _ = IsBetween(-2, -2, 0)()
	assert.Assert(t, ok)

	ok, _ = IsBetween(0, -2, 0)()
	assert.Assert(t, ok)

	ok, _ = IsBetween(1, -2, 0)()
	assert.Assert(t, !ok)
}

func TestStrictlyBetween(t *testing.T) {
	ok, _ := IsStrictlyBetween(1, 0, 2)()
	assert.Assert(t, ok)

	ok, _ = IsStrictlyBetween(-2, -2, 0)()
	assert.Assert(t, !ok)

	ok, _ = IsStrictlyBetween(0, -2, 0)()
	assert.Assert(t, !ok)

	ok, _ = IsStrictlyBetween(1, -2, 0)()
	assert.Assert(t, !ok)
}

func TestLess(t *testing.T) {
	ok, _ := IsLess[float32](-1, 32)()
	assert.Assert(t, ok)

	ok, _ = IsLess[uint64](125, 125)()
	assert.Assert(t, ok)

	ok, _ = IsLess[uint](23, 0)()
	assert.Assert(t, !ok)
}

func TestStrictlyLess(t *testing.T) {
	ok, _ := IsStrictlyLess[float32](-1, 32)()
	assert.Assert(t, ok)

	ok, _ = IsStrictlyLess[uint64](125, 125)()
	assert.Assert(t, !ok)

	ok, _ = IsStrictlyLess[uint](23, 0)()
	assert.Assert(t, !ok)
}

func TestGreater(t *testing.T) {
	ok, _ := IsGreater[float32](12.34, 1.234)()
	assert.Assert(t, ok)

	ok, _ = IsGreater[uint64](125, 125)()
	assert.Assert(t, ok)

	ok, _ = IsGreater[uint](23, 32)()
	assert.Assert(t, !ok)
}

func TestStrictlyGreater(t *testing.T) {
	ok, _ := IsStrictlyGreater[float32](12.34, 1.234)()
	assert.Assert(t, ok)

	ok, _ = IsStrictlyGreater[uint64](125, 125)()
	assert.Assert(t, !ok)

	ok, _ = IsStrictlyGreater[uint](23, 32)()
	assert.Assert(t, !ok)
}
