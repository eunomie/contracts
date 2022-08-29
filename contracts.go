package contracts

import (
	"fmt"
	"log"
)

type (
	Check func() (bool, string)

	Logger interface {
		Panic(...any)
	}

	EmptyLogger struct{}

	StdLogger struct{}
)

var (
	logger Logger = StdLogger{}
)

func (l EmptyLogger) Panic(v ...any) {
	s := fmt.Sprint(v...)
	panic(s)
}

func (l StdLogger) Panic(v ...any) {
	log.Panic(v...)
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

func Requires(checks ...Check) {
	runChecks(checks...)
}

func Ensure(checks ...Check) {
	runChecks(checks...)
}
