package dev

import (
	"errors"
	"fmt"
	"log"
	"os"
	"runtime"
)

const k = 1 << 10

// Deprecated: use helper package
type Printf = func(format string, v ...interface{})

// Deprecated: use helper package
func Go(f func()) {
	go func() {
		defer func() {
			_ = Recover(recover())
		}()
		f()
	}()
}

// Deprecated: use helper package
func Go2(f func(), printf Printf) {
	go func() {
		defer func() {
			_ = Recover2(recover(), printf)
		}()
		f()
	}()
}

// Deprecated: use helper package
func Recover(r interface{}) error {
	loggerStderr := log.New(os.Stderr, "", log.LstdFlags)
	return Recover2(r, loggerStderr.Printf)
}

// Deprecated: use helper package
func Recover2(r interface{}, printf Printf) error {
	if r != nil {
		buf := make([]byte, 4*k)
		n := runtime.Stack(buf, false)
		printf("[Recovery] panic recovered:\n%v\n%s\n", r, string(buf[:n]))

		var err error
		switch x := r.(type) {
		case string:
			err = errors.New(x)
		case error:
			err = x
		default:
			err = fmt.Errorf("%v", r)
		}

		return err
	}

	return nil
}
