package main

import (
	"fmt"
	"log"
)

var lastcounted int

func doloop(name string, times int, finished chan bool) {
	for i := 0; i < times; i++ {
		log.Printf("%s: %d out of %d\n", name, i, times)
		lastcounted = i
	}

	if finished != nil {
		finished <- true
	}
}

func main() {

	var done chan bool

	done = make(chan bool)
	go doloop("sub", 10, done)

	doloop("main", 15, nil)

	<-done

	fmt.Printf("Lastcounted: %d\n", lastcounted)

}
