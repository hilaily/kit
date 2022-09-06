package dev

import (
	"fmt"

	"github.com/sirupsen/logrus"
)

type devTool struct {
	log logrus.Entry
}

// Deprecated: use helper package
func (d *devTool) Catch() {
	if r := recover(); r != nil {
		d.log.Errorln("got error: ", r)
	}
}

// Deprecated: use helper package
func (d *devTool) Return() (err error) {
	if r := recover(); r != nil {
		err = fmt.Errorf("%v", r)
	}
	return
}

// Deprecated: use helper package
func (d *devTool) Check(err error) {
	if err != nil {
		panic(err)
	}
}
