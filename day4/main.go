package main

import (
	"crypto/md5"
	"fmt"
	"runtime"
	"strconv"
	"sync"
)

func calc(s int, e int, done chan int, index int, results []int, mu *sync.Mutex) {
	x := s
	for x <= e {
		input := "yzbqklnj" + strconv.Itoa(x)
		data := []byte(input)
		sum := md5.Sum(data)
		if sum[0] == 0 && sum[1] == 0 && sum[2] == 0 {
			mu.Lock()
			results[index] = x
			mu.Unlock()
			done <- index
			return
		}
		x += 1
	}
}

func main() {
	done := make(chan int)
	var mu sync.Mutex
	var wg sync.WaitGroup

	numWorkers := runtime.NumCPU()
	totalRange := int(^uint(0) >> 1)
	rangeSize := totalRange / numWorkers

	results := make([]int, numWorkers)

	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		start := i * rangeSize
		end := start + rangeSize - 1
		if i == numWorkers-1 {
			end = totalRange
		}
		go func(s, e, idx int) {
			calc(s, e, done, idx, results, &mu)
			wg.Done()
		}(start, end, i)
	}
	<-done
	go func() {
		for i := 0; i < numWorkers; i++ {
			index := <-done
			fmt.Printf("Goroutine %d found the number: %d\n", index, results[index])
		}
	}()

	wg.Wait()
}
