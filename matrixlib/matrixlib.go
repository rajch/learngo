package matrixlib

// IntMatrix is an square integer matrix of the specified dimentions
type IntMatrix struct {
	matrix [][]int
	Size   int
}

// Create creates a square matrix of the specified dimension
func Create(size int) IntMatrix {
	result := make([][]int, size)
	for i := range result {
		result[i] = make([]int, size)
	}
	return IntMatrix{matrix: result, Size: size}
}

func iterate(lhs IntMatrix, operation func(int, int)) {
	for i, onerow := range lhs.matrix {
		for j := range onerow {
			operation(i, j)
		}
	}
}

func iteratewithresult(lhs IntMatrix, rhs IntMatrix, operation func(int, int) int) (IntMatrix, bool) {
	var result IntMatrix

	status := lhs.Size == rhs.Size

	if status {
		result = Create(lhs.Size)
		iterate(lhs, func(i int, j int) {
			result.matrix[i][j] = operation(lhs.matrix[i][j], rhs.matrix[i][j])
		})
	}

	return result, status
}

// Add adds two matrices and returns the results in a new matrix
func Add(lhs IntMatrix, rhs IntMatrix) (IntMatrix, bool) {
	result, status := iteratewithresult(lhs, rhs, func(leftval int, rightval int) int {
		return leftval + rightval
	})
	return result, status
}

// Subtract subtracts two matrices and returns the results in a new matrix
func Subtract(lhs IntMatrix, rhs IntMatrix) (IntMatrix, bool) {
	result, status := iteratewithresult(lhs, rhs, func(leftval int, rightval int) int {
		return leftval - rightval
	})
	return result, status
}

// Set sets the element at co-ordinates x, y to value
func (lhs IntMatrix) Set(x int, y int, value int) bool {
	result := (x < lhs.Size && y < lhs.Size && x >= 0 && y >= 0)
	if result {
		lhs.matrix[x][y] = value
	}
	return result
}

// SetAll sets all elements in the matrix to the given value
func (lhs IntMatrix) SetAll(value int) {
	iterate(lhs, func(x int, y int) {
		lhs.matrix[x][y] = value
	})
}

// SetFrom sets all elements in the matrix from another matrix
func (lhs IntMatrix) SetFrom(rhs IntMatrix) {
	iterate(lhs, func(x int, y int) {
		lhs.matrix[x][y] = rhs.matrix[x][y]
	})
}

// SetValues sets the values passed into this matrix in row-first order
func (lhs IntMatrix) SetValues(values ...int) {
	if len(values) > 0 {
		index := 0
		iterate(lhs, func(x int, y int) {
			if index < len(values) {
				lhs.matrix[x][y] = values[index]
				index++

			}
		})
	}
}

// Add adds elements from another matrix into this one
func (lhs IntMatrix) Add(rhs IntMatrix) {
	iterate(lhs, func(x int, y int) {
		lhs.matrix[x][y] += rhs.matrix[x][y]
	})
}

// Subtract subtracts elements from another matrix in this one
func (lhs IntMatrix) Subtract(rhs IntMatrix) {
	iterate(lhs, func(x int, y int) {
		lhs.matrix[x][y] -= rhs.matrix[x][y]
	})
}
