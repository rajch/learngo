package main

import "fmt"
import "github.com/rajch/learngo/matrixlib"

func main() {
	// Case 1: Correct
	fmt.Println("Case 1: Add m1(3x3) to m2(3x3)")
	m1 := matrixlib.Create(3)
	m1[0][0] = 10
	m1[1][1] = 20
	m1[2][2] = 30

	m2 := matrixlib.Create(3)
	m2[0][0] = 40
	m2[1][1] = 50
	m2[2][2] = 60

	m3, ok := matrixlib.Add(m1, m2)

	fmt.Println(m1)
	fmt.Println(m2)
	fmt.Println(m3, ok)

	// Case 2: Wrong
	fmt.Println("\nCase 2: Add m1(3x3) to m4(4x4)")
	m4 := matrixlib.Create(4)
	m4[0][0] = 40
	m4[1][1] = 50
	m4[2][2] = 60

	m3, ok = matrixlib.Add(m1, m4)

	fmt.Println(m1)
	fmt.Println(m4)
	fmt.Println(m3, ok)

	// Case 3: Correct
	fmt.Println("\nCase 3: Subtract m1(3x3) from m2(3x3)")
	m3, ok = matrixlib.Subtract(m2, m1)

	fmt.Println(m2)
	fmt.Println(m1)
	fmt.Println(m3, ok)

	// Case 4: Wrong
	fmt.Println("\nCase 4: Add m1(3x3) to m5(jagged)")
	m5 := matrixlib.Create(3)
	m5[0][0] = 40
	m5[1][1] = 50
	m5[2][2] = 60
	m5[1] = append(m5[1], 75)

	m3, ok = matrixlib.Add(m1, m5)
	fmt.Println(m1)
	fmt.Println(m5)
	fmt.Println(m3, ok)

}
