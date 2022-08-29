package contracts_test

import (
	"fmt"
	"testing"

	"github.com/eunomie/contracts"
	"github.com/eunomie/contracts/carray"
	"github.com/eunomie/contracts/cnumber"
	"github.com/eunomie/contracts/ctype"
)

func TestRequires_Number(t *testing.T) {
	contracts.Requires(
		cnumber.IsStrictlyPositive(0.1),
		cnumber.IsPositive(12.3),
		cnumber.IsPositive(0),
		cnumber.IsNegative(0),
		cnumber.IsNegative(-1.2),
		cnumber.IsStrictlyNegative(-2),
		cnumber.IsBetween(1, 1, 3),
		cnumber.IsBetween(2, 1, 3),
		cnumber.IsBetween(3, 1, 3),
		cnumber.IsStrictlyBetween(0.0, -0.1, 0.1),
		cnumber.IsLessOrEqual(2.1, 2.1),
		cnumber.IsLessOrEqual(2.1, 2.2),
		cnumber.IsLess(1.2, 1.21),
		cnumber.IsGreater(2, 1),
		cnumber.IsGreaterOrEqual(2, 2),
		cnumber.IsGreaterOrEqual(2.2, 2.1),
	)
}

func TestRequires_Array(t *testing.T) {
	contracts.Requires(
		carray.IsNotEmpty([]string{""}),
		carray.IsEmpty([]struct{}{}),
		carray.IsEmpty(([]int)(nil)),
		carray.OfSize([]interface{}{1, false, nil}, 3),
		carray.WithEachElement([]int{-1, 0, 1, 2}, cnumber.IsBetweenFunc(-1, 2)),
		carray.WithEachElement([]int{0, 1, 2, 0}, cnumber.IsPositiveFunc[int]()),
	)
}

func TestRequiresPanic_Number(t *testing.T) {
	testRequiresPanic(t, []contracts.Check{
		cnumber.IsStrictlyPositive(0),
		cnumber.IsStrictlyPositive(-1.1),
		cnumber.IsPositive(-1),
		cnumber.IsNegative(0.1),
		cnumber.IsStrictlyNegative(0.1),
		cnumber.IsStrictlyNegative(0.0),
		cnumber.IsBetween(0, 1, 2),
		cnumber.IsBetween(3, 1, 2),
		cnumber.IsStrictlyBetween(0, 0, 1),
		cnumber.IsStrictlyBetween(1.1, 0.0, 1.1),
		cnumber.IsLessOrEqual(2, 1),
		cnumber.IsLess(1, 1),
		cnumber.IsLess(2.1, 2.0),
		cnumber.IsGreater(1, 2),
		cnumber.IsGreater(2, 2),
		cnumber.IsGreaterOrEqual(2.1, 2.2),
	})
}

func TestRequires_Type(t *testing.T) {
	type fn func()
	var (
		nilFn    fn = nil
		notNilFn fn = func() {}
	)

	type st struct{}
	var (
		nilStruct    *st = nil
		notNilStruct *st = &st{}
	)

	contracts.Requires(
		ctype.IsNil(nilFn),
		ctype.IsNotNil(notNilFn),
		ctype.IsNotNil(""),
		ctype.IsNotNil(0),
		ctype.IsNil(nilStruct),
		ctype.IsNotNil(notNilStruct),
	)
}

func TestRequiresPanic_Type(t *testing.T) {
	type fn func()
	var (
		nilFn    fn = nil
		notNilFn fn = func() {}
	)

	type st struct{}
	var (
		nilStruct    *st = nil
		notNilStruct *st = &st{}
	)

	testRequiresPanic(t, []contracts.Check{
		ctype.IsNil(notNilFn),
		ctype.IsNotNil(nilFn),
		ctype.IsNil(notNilStruct),
		ctype.IsNotNil(nilStruct),
	})
}

func testRequiresPanic(t *testing.T, testCases []contracts.Check) {
	for i, tc := range testCases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			defer func() {
				if r := recover(); r == nil {
					t.Error("should have panic")
				}
			}()
			contracts.Requires(tc)
		})
	}
}

func TestEnsure_Number(t *testing.T) {
	func() {
		defer contracts.Ensure(
			cnumber.IsStrictlyPositive(0.1),
			cnumber.IsPositive(12.3),
			cnumber.IsPositive(0),
			cnumber.IsNegative(0),
			cnumber.IsNegative(-1.2),
			cnumber.IsStrictlyNegative(-2),
			cnumber.IsBetween(1, 1, 3),
			cnumber.IsBetween(2, 1, 3),
			cnumber.IsBetween(3, 1, 3),
			cnumber.IsStrictlyBetween(0.0, -0.1, 0.1),
			cnumber.IsLessOrEqual(2.1, 2.1),
			cnumber.IsLessOrEqual(2.1, 2.2),
			cnumber.IsLess(1.2, 1.21),
			cnumber.IsGreater(2, 1),
			cnumber.IsGreaterOrEqual(2, 2),
			cnumber.IsGreaterOrEqual(2.2, 2.1),
		)
	}()
}

func TestEnsure_Array(t *testing.T) {
	defer contracts.Ensure(
		carray.IsNotEmpty([]string{""}),
		carray.IsEmpty([]struct{}{}),
		carray.IsEmpty(([]int)(nil)),
		carray.OfSize([]interface{}{1, false, nil}, 3),
		carray.WithEachElement([]int{-1, 0, 1, 2}, cnumber.IsBetweenFunc(-1, 2)),
		carray.WithEachElement([]int{0, 1, 2, 0}, cnumber.IsPositiveFunc[int]()),
	)
}

func TestEnsurePanic_Number(t *testing.T) {
	testEnsurePanic(t, []contracts.Check{
		cnumber.IsStrictlyPositive(0),
		cnumber.IsStrictlyPositive(-1.1),
		cnumber.IsPositive(-1),
		cnumber.IsNegative(0.1),
		cnumber.IsStrictlyNegative(0.1),
		cnumber.IsStrictlyNegative(0.0),
		cnumber.IsBetween(0, 1, 2),
		cnumber.IsBetween(3, 1, 2),
		cnumber.IsStrictlyBetween(0, 0, 1),
		cnumber.IsStrictlyBetween(1.1, 0.0, 1.1),
		cnumber.IsLessOrEqual(2, 1),
		cnumber.IsLess(1, 1),
		cnumber.IsLess(2.1, 2.0),
		cnumber.IsGreater(1, 2),
		cnumber.IsGreater(2, 2),
		cnumber.IsGreaterOrEqual(2.1, 2.2),
	})
}

func TestEnsure_Type(t *testing.T) {
	type fn func()
	var (
		nilFn    fn = nil
		notNilFn fn = func() {}
	)

	type st struct{}
	var (
		nilStruct    *st = nil
		notNilStruct *st = &st{}
	)

	defer contracts.Ensure(
		ctype.IsNil(nilFn),
		ctype.IsNotNil(notNilFn),
		ctype.IsNotNil(""),
		ctype.IsNotNil(0),
		ctype.IsNil(nilStruct),
		ctype.IsNotNil(notNilStruct),
	)
}

func TestEnsurePanic_Type(t *testing.T) {
	type fn func()
	var (
		nilFn    fn = nil
		notNilFn fn = func() {}
	)

	type st struct{}
	var (
		nilStruct    *st = nil
		notNilStruct *st = &st{}
	)

	testEnsurePanic(t, []contracts.Check{
		ctype.IsNil(notNilFn),
		ctype.IsNotNil(nilFn),
		ctype.IsNil(notNilStruct),
		ctype.IsNotNil(nilStruct),
	})
}

func testEnsurePanic(t *testing.T, testCases []contracts.Check) {
	for i, tc := range testCases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			defer func() {
				if r := recover(); r == nil {
					t.Error("should have panic")
				}
			}()
			func() {
				defer contracts.Ensure(tc)
			}()
		})
	}
}
