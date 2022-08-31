package contracts

import (
	"fmt"
	"log"
)

type (
	// Check returns a bool indicating if the condition is good and an error message if needed.
	Check func() (bool, string)

	Logger interface {
		Panic(...any)
	}

	PanicOnlyLogger struct{}

	StdPanicLogger struct{}

	NoPanicLogger struct{}
)

var (
	logger Logger = StdPanicLogger{}
)

func (l PanicOnlyLogger) Panic(v ...any) {
	s := fmt.Sprint(v...)
	panic(s)
}

func (l StdPanicLogger) Panic(v ...any) {
	log.Panic(v...)
}

func (l NoPanicLogger) Panic(v ...any) {
	log.Print(v...)
}

func SetLogger(l Logger) {
	logger = l
}

func runChecks(checks ...Check) {
	for _, check := range checks {
		if ok, errMsg := check(); !ok {
			logger.Panic(errMsg)
		}
	}
}

// Requires verify a list of conditions are all good and panic if that's not the case.
// contracts.Requires is intended to be used to verify pre-conditions on input values.
//
//	contracts.Requires(ctype.IsNotNil(inputValue)
func Requires(checks ...Check) {
	runChecks(checks...)
}

// Ensure verify a list of conditions are all good and panic if that's not the case.
// contract.Ensure is intended to be used to verify post-conditions on returning values.
//
//	defer contracts.Ensure(ctype.IsNotNil(returnValue))
func Ensure(checks ...Check) {
	runChecks(checks...)
}
