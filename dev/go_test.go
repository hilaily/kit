package dev

import (
	"fmt"
	"testing"
	"time"
)

func TestPanic(t *testing.T) {
	f := func() {
		defer func() {
			Recover(recover())
		}()
		panic("123")
	}
	f()
	time.Sleep(1 * time.Second)
}

func TestGo(t *testing.T) {
	Go(func() {
		panic("haha")
	})
	time.Sleep(1 * time.Second)
	fmt.Println("a")
}
