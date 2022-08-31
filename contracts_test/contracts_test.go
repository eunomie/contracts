package contracts_test

import (
	"fmt"
	"testing"

	"github.com/eunomie/contracts"
	"github.com/eunomie/contracts/carray"
	"github.com/eunomie/contracts/cnumber"
	"github.com/eunomie/contracts/cvar"
)

func TestRequires(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Error("should not have panic")
		}
	}()
	contracts.Requires(
		cnumber.IsStrictlyPositive(0.1),
		cnumber.IsPositive(12.3),
		carray.IsNotEmpty([]string{""}),
		carray.WithEachElement([]int{-1, 0, 1, 2}, cnumber.IsBetweenFunc(-1, 2)),
		cvar.IsNotNil(false),
	)
}

func TestRequiresPanic(t *testing.T) {
	testRequiresPanic(t, []contracts.Check{
		cnumber.IsStrictlyPositive(0),
		cvar.IsNotNil((func())(nil)),
		cvar.IsNil(false),
	})
}

func TestEnsure(t *testing.T) {
	func() {
		defer contracts.Ensure(
			cnumber.IsStrictlyPositive(0.1),
			cvar.IsNotNil(""),
			carray.WithEachElement([]int{0, 1, 2, 0}, cnumber.IsPositiveFunc[int]()),
		)
	}()
}

func TestEnsurePanic(t *testing.T) {
	testEnsurePanic(t, []contracts.Check{
		cnumber.IsNegative(0.1),
		cvar.IsNotNil((*struct{})(nil)),
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
