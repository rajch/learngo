package main

import "fmt"
import "github.com/rajch/learngo/matrixlib"

func main() {
	m1 := matrixlib.Create(3)
	m1[0][0] = 10
	m1[1][1] = 20
	m1[2][2] = 30
	fmt.Println(m1)
	m2 := matrixlib.Create(4)
	m2[0][0] = 40
	m2[1][1] = 50
	m2[2][2] = 60
	fmt.Println(m2)
	m3, ok := matrixlib.Add(m1,m2)
	fmt.Println(m3, ok)
}
