package dev

import (
	"fmt"

	"github.com/sirupsen/logrus"
)

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
