package main

import "sync"

var wg sync.WaitGroup

func main() {
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			println(i)
			wg.Done()
		}()
	}
	wg.Wait()
}
