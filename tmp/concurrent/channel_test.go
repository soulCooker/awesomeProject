package tmp

import (
	"fmt"
	"testing"
	"time"
)

func TestChannel(t *testing.T) {
	ch := make(chan int)
	timeout := time.After(time.Second * 10)
	select {
	case val := <-ch:
		fmt.Println(val)
	case <-timeout:
		fmt.Println("timeout")
	default:
		fmt.Println("empty")
	}
}
