package main

import (
	"fmt"
	"log"
	"math"
	"math/rand"
	"time"
)

func timethis(start time.Time, procname string) {
	end := time.Now()
	log.Printf("Fuction %s took %v", procname, end.Sub(start))
}

func dosumasync(arr []int, result chan<- int) {
	acc := 0
	for _, value := range arr {
		acc += int(math.Sqrt(float64(value)))
	}

	result <- acc
}

func dosum(arr []int) int {
	acc := 0
	for _, value := range arr {
		acc += int(math.Sqrt(float64(value)))
	}

	return acc
}

func fillarray(dataarr []int) {
	timethis(time.Now(), "fillarray")
	for i := range dataarr {
		dataarr[i] = rand.Intn(100)
	}
}

func dototal(dataarr []int, arraysize int, slicesize int) {
	defer timethis(time.Now(), "dototal")

	total := dosum(dataarr)
	log.Printf("Dosum Total: %d\n", total)
}

func dototalasync(dataarr []int, arraysize int, slicesize int) {
	defer timethis(time.Now(), "dototalasync")

	result := make(chan int)

	loopcount := arraysize / slicesize
	remainder := arraysize % slicesize

	for i := 0; i < loopcount; i++ {
		start := slicesize * i
		end := start + slicesize

		go dosumasync(dataarr[start:end], result)
	}

	channelcount := loopcount

	if remainder > 0 {
		start := slicesize * loopcount
		end := start + remainder
		go dosumasync(dataarr[start:end], result)

		channelcount++
	}

	total := 0

	for i := 0; i < channelcount; i++ {
		total += <-result
	}

	log.Printf("Total: %d", total)
}

func doexercise(arraysize int, slicesize int) {
	dataarr := make([]int, arraysize)

	fillarray(dataarr)

	dototal(dataarr, arraysize, slicesize)

	dototalasync(dataarr, arraysize, slicesize)
}

func main() {
	var arraysize int
	var blocksize int

	fmt.Print("Enter total size and block size seperated by a space:")
	fmt.Scanf("%d %d", &arraysize, &blocksize)
	doexercise(arraysize, blocksize)
}
