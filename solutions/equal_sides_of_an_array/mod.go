package equalsidesofanarray

func FindEvenIndex(arr []int) int {
	for idx := range arr {
		leftSum := 0
		rightSum := 0

		for i := 0; i < idx; i++ {
			leftSum += arr[i]
		}

		for i := idx + 1; i < len(arr); i++ {
			rightSum += arr[i]
		}

		if leftSum == rightSum {
			return idx
		}
	}

	return -1
}
