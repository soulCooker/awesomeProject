package main

import "sync"

func main() {
	crossPrint()
}

func crossPrint() {
	ch1 := make(chan interface{}, 1)
	ch2 := make(chan interface{}, 1)

	ch1 <- 1

	arr1 := make([]int, 0, 10)
	arr2 := make([]int, 0, 10)

	for i := 0; i < 10; i += 2 {
		arr1 = append(arr1, i)
	}
	for i := 1; i < 10; i += 2 {
		arr2 = append(arr2, i)
	}

	countDown := sync.WaitGroup{}

	countDown.Add(2)

	go func() {
		for i := 0; i < len(arr1); i++ {
			select {
			case <-ch1:
				println(arr1[i])
				ch2 <- 1
			}
		}
		countDown.Done()
	}()

	go func() {
		for i := 0; i < len(arr2); i++ {
			select {
			case <-ch2:
				println(arr2[i])
				ch1 <- 1
			}
		}
		countDown.Done()
	}()

	countDown.Wait()
}
