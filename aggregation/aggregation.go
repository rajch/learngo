package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/rajch/learngo/aggregation/aggregators"
	"github.com/rajch/learngo/aggregation/aggregators/intaggregators"
)

func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	fmt.Printf("%s took %s\n", name, elapsed)
}

func generateNumbers(size int) (numbers []int) {
	defer timeTrack(time.Now(), "generateNumbers()")

	fmt.Printf("Generating %d numbers..\n.", size)
	numbers = make([]int, size)
	for i := 0; i < size; i++ {
		numbers[i] = rand.Int()
	}
	return
}

func simpleSum(numbers []int, partionnum int, done chan int) {
	defer timeTrack(time.Now(), "simpleSum()")

	fmt.Printf("\nCalculating sum for partition %d...\n", partionnum)
	sumag := new(intaggregators.IntSumAggregator)
	total, cnt := aggregators.SimpleAggregate(sumag, numbers)
	fmt.Printf("Partition %d: Sum of %d numbers is %d\n", partionnum, cnt, total)

	if done != nil {
		done <- partionnum
	}
}

func partitionedSum(numbers []int, partitionsize int) {
	defer timeTrack(time.Now(), "partitionedSum()")

	donechannel := make(chan int)

	fmt.Printf("\nCalculating sum for %d numbers...\n", partitionsize)
	partcnt := len(numbers) / partitionsize
	index := 0

	var i int
	for i = 0; i < partcnt; i++ {
		go simpleSum(numbers[index:index+(partitionsize-1)], i, donechannel)
		index += partitionsize
	}
	if index < len(numbers) {
		partcnt++
		go simpleSum(numbers[index:], partcnt, donechannel)

	}

	fmt.Println("All partitions dispatched. Waiting...")
	for hit := 0; hit < partcnt; hit++ {
		<-donechannel
	}
}

func main() {
	numbers := generateNumbers(1000)
	simpleSum(numbers, 0, nil)
	partitionedSum(numbers, 300)

	//var i string
	//fmt.Scanln(&i)
}
