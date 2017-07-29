package matrixlib

// Create creates a square matrix of the specified dimension
func Create(size int) [][]int {
	result := make([][]int, size)
	for i := range result {
		result[i] = make([]int, size)
	}
	return result
}

func rowwiseoperate(one [][]int, two [][]int, operation func(int, int) int) ([][]int, bool) {
	var resultmatrix [][]int

	resultstatus := len(one) == len(two)

	if resultstatus {
		resultmatrix = Create(len(one))
		for i, onerow := range one {
			if len(onerow) == len(two[i]) {
				for j, oneval := range onerow {
					resultmatrix[i][j] = operation(oneval, two[i][j])
				}
			} else {
				resultstatus = false
				resultmatrix = nil
				break
			}
		}
	}

	return resultmatrix, resultstatus
}

// Add adds two matrices
func Add(one [][]int, two [][]int) ([][]int, bool) {
	resultmatrix, resultstatus := rowwiseoperate(one, two, func(leftval int, rightval int) int {
		return leftval + rightval
	})
	return resultmatrix, resultstatus
}

// Subtract subtracts two matrices
func Subtract(one [][]int, two [][]int) ([][]int, bool) {
	resultmatrix, resultstatus := rowwiseoperate(one, two, func(leftval int, rightval int) int {
		return leftval - rightval
	})
	return resultmatrix, resultstatus
}
