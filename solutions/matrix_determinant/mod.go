package matrixdeterminant

func Determinant(matrix [][]int) int {
	matrixLen := len(matrix)

	if matrixLen == 1 {
		return matrix[0][0]
	}

	determinant := 0
	sign := 1

	for i := 0; i < matrixLen; i++ {
		subMatrix := make([][]int, matrixLen-1)
		for j := range subMatrix {
			subMatrix[j] = make([]int, matrixLen-1)
		}

		for j := 1; j < matrixLen; j++ {
			for k := 0; k < matrixLen; k++ {
				if k < i {
					subMatrix[j-1][k] = matrix[j][k]
				} else if k > i {
					subMatrix[j-1][k-1] = matrix[j][k]
				}
			}
		}

		determinant += matrix[0][i] * sign * Determinant(subMatrix)
		sign *= -1
	}

	return determinant

	// BEST PRACTICE OF CODE
	// if len(matrix) == 1 {
	// 		return matrix[0][0]
	// }

	// sliced := matrix[1:]
	// res := 0

	// for n := 0; n < len(matrix); n++ {
	// 	arr := [][]int{}

	// 	for _, v := range sliced {
	// 			arr = append(arr, append(append([]int{}, v[:n]...),  v[n+1:]...))
	// 	}

	// 	sign := -1

	// 	if n % 2 == 0 {
	// 			sign = 1
	// 	}

	// 	res += sign * matrix[0][n] * Determinant(arr)
	// }

	// return res
}
