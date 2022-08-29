# Contracts

Just a simple library to try to have design by contracts in Go.

```go
package main

import (
	"github.com/eunomie/contracts"
	"github.com/eunomie/contracts/carray"
	"github.com/eunomie/contracts/cnumber"
)

func SumAngles(values ...int) (res int) {
	contracts.Requires(
		carray.IsNotEmpty(values),
		carray.WithEachElement(values,
			cnumber.IsPositiveFunc[int](),
			cnumber.IsLessFunc(360)))
	defer contracts.Ensure(
		cnumber.IsBetween(res, 0, 360))

	for _, el := range values {
		res += el
	}

	return res % 360
}
```
