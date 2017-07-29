package matrixlib

import (
	"fmt"
	"strings"
)

func (lhs IntMatrix) String() string {
	elements := make([]string, lhs.Size*lhs.Size)
	index := 0
	iterate(lhs, func(x int, y int) {
		element := fmt.Sprintf("%3d", lhs.matrix[x][y])
		if y == lhs.Size-1 {
			element += "\n"
		} else {
			element += "\t"
		}
		elements[index] = element
		index++
	})
	return "\n" + strings.Join(elements, "")
}
