package eureka

import (
	"fmt"
	"math"
)

func SumDigPow(a, b uint64) []uint64 {
	output := []uint64{}
	for i := a; i <= b; i++ {
		sum := float64(0)
		digits := fmt.Sprintf("%v", i)
		for j := 0; j < len(digits); j++ {
			sum += math.Pow(float64(digits[j]-'0'), float64(j+1))
		}
		if sum == float64(i) {
			output = append(output, i)
		}
	}

	if len(output) == 0 {
		return nil
	}

	return output
}
