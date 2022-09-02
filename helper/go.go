package helper

import (
	"errors"
	"fmt"
	"log"
	"os"
	"runtime"
)

const k = 1 << 10

type Printf = func(format string, v ...interface{})

func Go(f func()) {
	go func() {
		defer func() {
			Recover(recover())
		}()
		f()
	}()
}

func Go2(f func(), printf Printf) {
	go func() {
		defer func() {
			Recover2(recover(), printf)
		}()
		f()
	}()
}

func Recover(r interface{}) error {
	loggerStderr := log.New(os.Stderr, "", log.LstdFlags)
	return Recover2(r, loggerStderr.Printf)
}

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
