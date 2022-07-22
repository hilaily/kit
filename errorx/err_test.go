package errorx

import (
	"errors"
	"fmt"
	"testing"
)

func TestErr(t *testing.T) {
	e := test()
	ok := errors.Is(e, myErr1)
	fmt.Println("is ==: ", ok)
	ok = errors.As(e, &myErr2)
	fmt.Println("as ==: ", ok)
	v, ok := Unwrap[*MyErr](e)
	fmt.Println("assert ==: ", ok, v)

	e = myErr1
	v, ok = Unwrap[*MyErr](e)
	fmt.Println("assert ==: ", ok, v)

	v2, ok := Unwrap[IInter](e)
	fmt.Println("assert ==: ", ok, v2)
}

func test() error {
	e := myErr1
	e2 := fmt.Errorf("aa %w", e)
	return fmt.Errorf("bb %w", e2)
}

type MyErr struct {
	Code string
}

func (m *MyErr) Error() string {
	return m.Code
}

func (m *MyErr) TestInterface() string {
	return "this is a implement"
}

type IInter interface {
	TestInterface() string
}

var (
	myErr1 = &MyErr{Code: "1"}
	myErr2 = &MyErr{Code: "2"}
)
