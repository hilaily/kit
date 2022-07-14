package dev

import (
	"errors"
	"fmt"

	"github.com/sirupsen/logrus"
)

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

type devTool struct {
	log logrus.Entry
}

func (d *devTool) Catch() {
	if r := recover(); r != nil {
		d.log.Errorln("got error: ", r)
	}
}

func (d *devTool) Return() (err error) {
	if r := recover(); r != nil {
		err = fmt.Errorf("%v", r)
	}
	return
}

func (d *devTool) Check(err error) {
	if err != nil {
		panic(err)
	}
}
