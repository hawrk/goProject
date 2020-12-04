package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

const num = 1001002300

func main() {
	simple1()
	//simple2()
	simple3()

}

func simple1() {
	start := time.Now()
	sum := 0
	for i := 0; i < num; i++ {
		sum += i
	}
	lancy := time.Now().Sub(start)
	fmt.Println("sum=", sum, ".lancy:", lancy)
}

func simple2() {
	startTime := time.Now()
	n := runtime.GOMAXPROCS(0)
	wg := sync.WaitGroup{}
	sum := make([]int, n)
	for i := 0; i < n; i++ {
		wg.Add(1)
		go func(i int) {
			start := num / n * i
			end := start + num/n
			for j := start; j < end; j++ {
				sum[i] += j
			}
			wg.Done()
		}(i)
	}
	wg.Done()
	sums := 0
	for _, s := range sum {
		sums += s
	}
	lancy := time.Now().Sub(startTime)
	fmt.Println("sums:", sums, "lancy:", lancy)

}

func simple3() {
	n := runtime.GOMAXPROCS(0)
	startTime := time.Now()
	res := make(chan int)
	for i := 0; i < n; i++ {
		go func(i int, r chan<- int) {
			sum := 0
			start := num / n * i
			var end int
			if i+1 == n {
				end = num
			} else {
				end = start + num/n
			}
			//fmt.Println("in i = :", i , ",start:", start, ",end:",end)
			for j := start; j < end; j++ {
				sum += j
			}
			r <- sum
		}(i, res)
	}
	sums := 0
	for i := 0; i < n; i++ {
		sums += <-res
	}
	lancy := time.Now().Sub(startTime)
	fmt.Println("sum=", sums, ",lancy:", lancy)
}
