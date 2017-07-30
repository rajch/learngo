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
	fmt.Printf("%s took %s\n\n\n", name, elapsed)
}

func generateNumbers(size int) (numbers []int) {
	defer timeTrack(time.Now(), "generateNumbers()")

	fmt.Printf("Generating %d numbers..\n.", size)
	numbers = make([]int, size)
	for i := 0; i < size; i++ {
		numbers[i] = rand.Intn(200)
	}
	return
}

func doSum(numbers []int, partionnum int, done chan int) {
	fmt.Printf("Calculating sum for partition %d...\n", partionnum)
	sumag := new(intaggregators.IntSumAggregator)
	total, cnt := aggregators.SimpleAggregate(sumag, numbers)
	fmt.Printf("Partition %d: Sum of %d numbers is %d\n", partionnum, cnt, total)

	if done != nil {
		done <- total
	}
}

func simpleSum(numbers []int) {
	defer timeTrack(time.Now(), "simpleSum()")
	fmt.Println("\nCalculating simple sum...")

	doSum(numbers, 0, nil)
}

func partitionedSum(numbers []int, partitioncount int) {
	defer timeTrack(time.Now(), "partitionedSum()")
	fmt.Println("\nCalculating partitioned sum...")

	donechannel := make(chan int)

	partsize := len(numbers) / partitioncount
	leftover := len(numbers) % partitioncount

	index := 0

	for i := 0; i < partitioncount; i++ {
		offset := partsize
		if i == 0 {
			offset += leftover
		}

		go doSum(numbers[index:index+(offset)], i, donechannel)
		index += offset
	}

	fmt.Println("All partitions dispatched. Waiting...")

	total := 0
	for hit := 0; hit < partitioncount; hit++ {
		partitiontotal := <-donechannel
		total += partitiontotal
	}
	fmt.Printf("Sum of %d partitions is %d\n", partitioncount, total)
}

func main() {
	size, partioncount := 100000, 5

	numbers := generateNumbers(size)
	simpleSum(numbers)
	partitionedSum(numbers, partioncount)
}
