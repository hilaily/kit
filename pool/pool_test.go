package pool

import (
	"fmt"
	"testing"
	"time"
)

func TestPool(t *testing.T) {
	p := NewPool(2)
	for i := 0; i < 10; i++ {
		p.Go(func() {
			fmt.Println(time.Now().String())
			time.Sleep(1 * time.Second)
		})
	}
}
