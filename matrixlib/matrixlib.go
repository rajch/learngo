package matrixlib

func Create(size int)([][]int) {
	result := make([][]int, size)
	for i:= range result {
		result[i] = make([]int, size)
	}
	return result
}

func Add(one [][]int, two [][]int) ([][]int, bool) {
	var resultval [][]int
	result := len(one) == len(two) && len(one[0]) == len(two[0])
	if result {
		resultval = Create(len(one))
		for i, onerow := range one {
			for j, oneval := range onerow {
				resultval[i][j] = oneval + two[i][j]
			}
		}
	} else {
		resultval = nil
	}
	return resultval, result
}
