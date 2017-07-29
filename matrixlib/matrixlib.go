package matrixlib

// IntMatrix is an square integer matrix of the specified dimentions
type IntMatrix [][]int

// Create creates a square matrix of the specified dimension
func Create(size int) IntMatrix {
	result := make([][]int, size)
	for i := range result {
		result[i] = make([]int, size)
	}
	return result
}

func rowwiseoperate(one IntMatrix, two IntMatrix, operation func(int, int) int) (IntMatrix, bool) {
	var resultmatrix IntMatrix

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
func Add(one IntMatrix, two IntMatrix) (IntMatrix, bool) {
	resultmatrix, resultstatus := rowwiseoperate(one, two, func(leftval int, rightval int) int {
		return leftval + rightval
	})
	return resultmatrix, resultstatus
}

// Subtract subtracts two matrices
func Subtract(one IntMatrix, two IntMatrix) (IntMatrix, bool) {
	resultmatrix, resultstatus := rowwiseoperate(one, two, func(leftval int, rightval int) int {
		return leftval - rightval
	})
	return resultmatrix, resultstatus
}

func rowwiseexecute(lhs IntMatrix, rhs IntMatrix, operation func(int, int)) {
	for i, onerow := range lhs {
		for j := range onerow {
			operation(i, j)
		}
	}
}

// SetAll sets all elements in the matrix to the given value
func (lhs IntMatrix) SetAll(value int) {
	rowwiseexecute(lhs, nil, func(x int, y int) {
		lhs[x][y] = value
	})
}

// SetFrom sets all elements in the matrix from another matrix
func (lhs IntMatrix) SetFrom(rhs IntMatrix) {
	rowwiseexecute(lhs, rhs, func(x int, y int) {
		lhs[x][y] = rhs[x][y]
	})
}

// Add adds elements from another matrix into this one
func (lhs IntMatrix) Add(rhs IntMatrix) {
	rowwiseexecute(lhs, rhs, func(x int, y int) {
		lhs[x][y] += rhs[x][y]
	})
}

// Subtract subtracts elements from another matrix in this one
func (lhs IntMatrix) Subtract(rhs IntMatrix) {
	rowwiseexecute(lhs, rhs, func(x int, y int) {
		lhs[x][y] -= rhs[x][y]
	})
}
