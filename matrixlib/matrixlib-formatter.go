package matrixlib

import (
	"fmt"
	"strings"
)

func (lhs IntMatrix) String() string {
	elements := make([]string, len(lhs)*len(lhs))
	index := 0
	rowwiseexecute(lhs, nil, func(x int, y int) {
		element := fmt.Sprintf("%3d", lhs[x][y])
		if y == len(lhs)-1 {
			element += "\n"
		} else {
			element += "\t"
		}
		elements[index] = element
		index++
	})
	return "\n" + strings.Join(elements, "")
}
