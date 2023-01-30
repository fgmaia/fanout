package main

import (
	"fmt"
	"sync"
)

func main() {
	ch1, err := read("file1.csv")
	if err != nil {
		panic(fmt.Errorf("could not read file1 %v", err))
	}

	//-
	var wg sync.WaitGroup

	breakup(&wg, "worker 1", ch1)
	breakup(&wg, "worker 2", ch1)
	breakup(&wg, "worker 3", ch1)

	wg.Wait()

	fmt.Println("All completed, exiting")
}

func breakup(wg *sync.WaitGroup, worker string, ch <-chan string) {
	wg.Add(1)
	go func() {
		for v := range ch {
			fmt.Println(worker, v)
		}
		wg.Done()
	}()
}

func read(fileName string) (<-chan string, error) {
	ch := make(chan string, 10)

	go func(ch chan string) {
		for i := 1; i <= 1000; i++ {
			ch <- fmt.Sprintf("fileName %s val: %d", fileName, i)
		}
		close(ch)
	}(ch)

	return ch, nil
}
