package main

import "fmt"
import "github.com/rajch/learngo/first/firstlib"

func main() {
	fmt.Println("Hello")
	fmt.Println("2 + 2 is:", add(2, 2))
	fmt.Println("5 - 2 is:", firstlib.Subtract(5, 2))
	fmt.Println("Starting series...")
	fmt.Println(firstlib.Next())
	fmt.Println(firstlib.Next())
	fmt.Println(firstlib.Next())
	fmt.Println(firstlib.Next())
	fmt.Println(firstlib.Next())
	fmt.Println(firstlib.Next())
	fmt.Println("Restarting series...")
	firstlib.Init()
	fmt.Println(firstlib.Next())
	fmt.Println(firstlib.Next())
	fmt.Println(firstlib.Next())
	fmt.Println(firstlib.Next())
	fmt.Println(firstlib.Next())

}
