package main

import "fmt"
import "github.com/rajch/learngo/matrixlib"

func main() {
	// Case 1: Correct
	fmt.Println("Case 1: Add m1(3x3) to m2(3x3)")
	m1 := matrixlib.Create(3)
	m1.Set(0, 0, 10)
	m1.Set(1, 1, 20)
	m1.Set(2, 2, 30)

	m2 := matrixlib.Create(3)
	m2.Set(0, 0, 40)
	m2.Set(1, 1, 50)
	m2.Set(2, 2, 60)

	m3, ok := matrixlib.Add(m1, m2)

	fmt.Println(m1)
	fmt.Println(m2)
	fmt.Println(m3, ok)

	// Case 2: Wrong
	fmt.Println("\nCase 2: Add m1(3x3) to m4(4x4)")
	m4 := matrixlib.Create(4)
	m4.Set(0, 0, 70)
	m4.Set(1, 1, 80)
	m4.Set(2, 2, 90)
	m4.Set(2, 2, 100)

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

	// OBSOLETE Case 4: Wrong
	fmt.Println("\nOBSOLETE Case 4: Add m1(3x3) to m5(jagged)")
	/*
		m5 := matrixlib.Create(3)
		m5[0][0] = 40
		m5[1][1] = 50
		m5[2][2] = 60
		m5[1] = append(m5[1], 75)

		m3, ok = matrixlib.Add(m1, m5)
		fmt.Println(m1)
		fmt.Println(m5)
		fmt.Println(m3, ok)
	*/

	// Case 5: Should not compile, but does. m6 is [][]int, which is an "unnamed" type
	// Unnamed types are interchangable with their redifinition, in this case IntMatrix
	fmt.Println("\nCase 5: Add m1(3x3) to m6(3x3 [][]int) NOW does not compile because IntMatrix is a complex type")
	/*
		m6 := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}

		m3, ok = matrixlib.Add(m1, m6)
		fmt.Println(m1)
		fmt.Println(m6)
		fmt.Println(m3, ok)
	*/

	// Case 6: Casting for type
	fmt.Println("\nCase 6: Add m1(3x3) to m7(m6 cast as IntMatrix) NOW not available any more because IntMatrix is a complex type")
	/*
		m7 := matrixlib.IntMatrix(m6)

		m3, ok = matrixlib.Add(m1, m7)
		fmt.Println(m1)
		fmt.Println(m7)
		fmt.Println(m3, ok)
	*/

	// Case 7: using methods
	fmt.Println("\nCase 7: setting m1 to all 10s, adding m2 to m1 itself")
	m1.SetAll(10)
	fmt.Println("m1 after setting, before adding:", m1)
	fmt.Println(m2)
	m1.Add(m2)
	fmt.Println(m1)

	// Case 8: variadic method
	fmt.Println("\nCase 8: setting m8 using variadic SetValues method")
	m8 := matrixlib.Create(3)
	m8.SetValues(10, 20, 30, 40, 50)
	fmt.Println(m8)

}
