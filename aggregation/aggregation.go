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

func simpleSum(numbers []int, partionnum int) {
	defer timeTrack(time.Now(), "simpleSum()")

	fmt.Printf("\nCalculating sum for partition %d...\n", partionnum)
	sumag := new(intaggregators.IntSumAggregator)
	total, cnt := aggregators.SimpleAggregate(sumag, numbers)
	fmt.Printf("Partition %d: Sum of %d numbers is %d\n", partionnum, cnt, total)

}

func partitionedSum(numbers []int, partitions int) {
	defer timeTrack(time.Now(), "partitionedSum()")

	fmt.Printf("\nCalculating sum for %d partitions...\n", partitions)
	partsize := len(numbers) / partitions
	index := 0

	var i int
	for i = 0; i < partitions; i++ {
		simpleSum(numbers[index:index+(partsize-1)], i)
		index += partsize
	}
	if index < len(numbers) {
		simpleSum(numbers[index:], i)
	}
}

func main() {
	numbers := generateNumbers(1000)
	simpleSum(numbers, 0)
	partitionedSum(numbers, 3)
}
