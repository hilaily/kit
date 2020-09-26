package datetime

import (
	"fmt"
	"testing"
)

func TestDateToday(t *testing.T) {
	res := DateToday()
	fmt.Println(Format2006_01_02(res))
}
