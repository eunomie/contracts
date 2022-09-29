// Package contracts allow to define a set of pre- and post-conditions to any methods.
// It will check the requirements of the method, in a way finer grain than the typing system of Go, and will ensure the
// result values also correspond to the defined contracts.
package contracts

import (
	"fmt"
	"log"
)

type (
	// Check returns a boolean indicating if the condition has been verified and an error message if needed.
	Check func() (bool, string)

	// Logger is a minimal logger interface with only a Panic method. You can set the current logger implementation using
	// SetLogger method.
	// Three different implementations are provided to cover the main cases but feel free to define your own:
	// - PanicOnlyLogger: panic but do not write log line
	// - StdPanicLogger: default logger, calling log.Panic from the standard library.
	// - NoPanicLogger: log only, no panic
	Logger interface {
		// Panic will be called when pre- or post-condition are not met.
		Panic(...any)
	}

	// PanicOnlyLogger is an implementation of Logger interface that only panics. No log lines will be sent.
	PanicOnlyLogger struct{}

	// StdPanicLogger is an implementation of Logger interface using the default log.Panic method. It will print a log
	// line then call panic.
	// This is the default logger set.
	StdPanicLogger struct{}

	// NoPanicLogger is an implementation of Logger interface not calling panic. It will print a log line on failure
	// but will continue. It might be controversial, but it's the type of behaviour we often see when using asserts in
	// other languages when run in production. Use it at your own risk.
	NoPanicLogger struct{}
)

var (
	logger Logger = StdPanicLogger{}
)

// Panic will be called when pre- or post-condition are not met.
// This implementation will only panic, no log line will be output.
func (l PanicOnlyLogger) Panic(v ...any) {
	s := fmt.Sprint(v...)
	panic(s)
}

// Panic will be called when pre- or post-condition are not met.
// This implementation will call the default log.Panic method.
func (l StdPanicLogger) Panic(v ...any) {
	log.Panic(v...)
}

// Panic will be called when pre- or post-condition are not met.
// This implementation will only log the details, no panic.
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
