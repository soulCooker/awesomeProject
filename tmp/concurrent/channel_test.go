package tmp

import (
	"fmt"
	"testing"
)

func TestChannel(t *testing.T) {
	ch := make(chan int)

	select {
	case val := <-ch:
		fmt.Println(val)
	default:
		fmt.Println("empty")
	}
}
