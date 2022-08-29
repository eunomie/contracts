package cnumber

import (
	"fmt"

	"github.com/eunomie/contracts"
	"golang.org/x/exp/constraints"
)

type (
	Number interface {
		constraints.Integer | constraints.Float
	}
)

func IsPositive[T Number](v T) contracts.Check {
	return IsPositiveFunc[T]()(v)
}

func IsPositiveFunc[T Number]() func(T) contracts.Check {
	return func(v T) contracts.Check {
		return func() (bool, string) {
			return v >= T(0), fmt.Sprintf("number should be positive but is %v", v)
		}
	}
}

func IsStrictlyPositive[T Number](v T) contracts.Check {
	return IsStrictlyPositiveFunc[T]()(v)
}

func IsStrictlyPositiveFunc[T Number]() func(T) contracts.Check {
	return func(v T) contracts.Check {
		return func() (bool, string) {
			return v > T(0), fmt.Sprintf("number should be strictly positive but is %v", v)
		}
	}
}

func IsNegative[T Number](v T) contracts.Check {
	return IsNegativeFunc[T]()(v)
}

func IsNegativeFunc[T Number]() func(T) contracts.Check {
	return func(v T) contracts.Check {
		return func() (bool, string) {
			return v <= T(0), fmt.Sprintf("number should be negative but is %v", v)
		}
	}
}

func IsStrictlyNegative[T Number](v T) contracts.Check {
	return IsStrictlyNegativeFunc[T]()(v)
}

func IsStrictlyNegativeFunc[T Number]() func(T) contracts.Check {
	return func(v T) contracts.Check {
		return func() (bool, string) {
			return v < T(0), fmt.Sprintf("number should be strictly negative but is %v", v)
		}
	}
}

func IsBetween[T Number](v, min, max T) contracts.Check {
	return IsBetweenFunc(min, max)(v)
}

func IsBetweenFunc[T Number](min, max T) func(T) contracts.Check {
	return func(v T) contracts.Check {
		return func() (bool, string) {
			return min <= v && v <= max, fmt.Sprintf("number should inside [%v %v] but is %v", min, max, v)
		}
	}
}

func IsStrictlyBetween[T Number](v, min, max T) contracts.Check {
	return IsStrictlyBetweenFunc(min, max)(v)
}

func IsStrictlyBetweenFunc[T Number](min, max T) func(T) contracts.Check {
	return func(v T) contracts.Check {
		return func() (bool, string) {
			return min < v && v < max, fmt.Sprintf("number should be inside ]%v %v[ but is %v", min, max, v)
		}
	}
}

func IsLessOrEqual[T Number](v, max T) contracts.Check {
	return IsLessOrEqualFunc(max)(v)
}

func IsLessOrEqualFunc[T Number](max T) func(T) contracts.Check {
	return func(v T) contracts.Check {
		return func() (bool, string) {
			return v <= max, fmt.Sprintf("number should be <= %v but is %v", max, v)
		}
	}
}

func IsLess[T Number](v, max T) contracts.Check {
	return IsLessFunc(max)(v)
}

func IsLessFunc[T Number](max T) func(T) contracts.Check {
	return func(v T) contracts.Check {
		return func() (bool, string) {
			return v < max, fmt.Sprintf("number should be < %v but is %v", max, v)
		}
	}
}

func IsGreaterOrEqual[T Number](v, min T) contracts.Check {
	return IsGreaterOrEqualFunc(min)(v)
}

func IsGreaterOrEqualFunc[T Number](min T) func(T) contracts.Check {
	return func(v T) contracts.Check {
		return func() (bool, string) {
			return min <= v, fmt.Sprintf("number should be >= %v but is %v", min, v)
		}
	}
}

func IsGreater[T Number](v, min T) contracts.Check {
	return IsGreaterFunc(min)(v)
}

func IsGreaterFunc[T Number](min T) func(T) contracts.Check {
	return func(v T) contracts.Check {
		return func() (bool, string) {
			return min < v, fmt.Sprintf("number should be > %v but is %v", min, v)
		}
	}
}
