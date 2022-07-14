package errorx

import "errors"

// Unwrap unwrap error to get custom err object
func Unwrap[T any](err error) (T, bool) {
	var current = err
	var last = err
	for {
		current = errors.Unwrap(current)
		if current == nil {
			err = last
			break
		}
		last = current
	}
	v, ok := err.(T)
	return v, ok
}
