# Contracts

Just a simple library to try to have design by contracts in Go.

```go
package main

import (
	"github.com/eunomie/contracts"
	"github.com/eunomie/contracts/carray"
	"github.com/eunomie/contracts/cnumber"
	"github.com/eunomie/contracts/cvar"
)

func SumAngles(values ...int) (res int) {
	contracts.Requires(
		carray.IsNotEmpty(values),
		carray.WithEachElement(values,
			cnumber.IsPositiveFunc[int](),
			cnumber.IsStrictlyLessFunc(360)))
	defer contracts.Ensure(
		cnumber.IsBetween(res, 0, 359))

	for _, el := range values {
		res += el
	}

	res = res % 360

	return
}

func RotateByQuarter(angle, rotation int) (res int) {
	contracts.Requires(
		cnumber.IsBetween(angle, 0, 359),
		cvar.IsInsideEnum(rotation, 0, 90, 180, 270))
	defer contracts.Ensure(
		cnumber.IsBetween(res, 0, 359))
	
	res = (res + rotation) % 360
	
	return
}
```
