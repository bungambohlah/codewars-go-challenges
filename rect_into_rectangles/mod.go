package rectintorectangles

import (
	"fmt"
)

func RectIntoRects(l, w int) []string {
	if l == w || l == 0 || w == 0 {
		return []string{}
	}

	// ensure that length is always greater than width
	if w > l {
		placeholder := l
		l = w
		w = placeholder
	}

	// attach initial rectangle to output list
	var outputRectangles []string
	outputRectangles = append(outputRectangles, fmt.Sprintf("(%d*%d)", l, w))

	// Create list for side lengths and cumulative counts and
	// previous shortest for special case. (When the shortest
	// side remains the shortest)
	sideLengths := []int{l, w}
	shortestCumulativeCount := 1

	// While loop finds the shortest side, and takes the length
	// of this away from the longest, checking for the special condition
	for sideLengths[0] != sideLengths[1] {
		var shortestIndex, longestIndex int

		if sideLengths[0] < sideLengths[1] {
			shortestIndex = 0
			longestIndex = 1
		} else {
			shortestIndex = 1
			longestIndex = 0
		}

		shortest := sideLengths[shortestIndex]

		// Special condition: shortest side remains the same (additional rectangles can be created)
		if Min(sideLengths) == shortest && sideLengths[longestIndex]-shortest > shortest {
			shortestCumulativeCount++
			for i := 2; i <= shortestCumulativeCount; i++ {
				rectSize := i * shortest
				rect := fmt.Sprintf("(%d*%d)", rectSize, shortest)
				outputRectangles = append(outputRectangles, rect)
			}
		} else {
			shortestCumulativeCount = 1
		}

		// Remove the shortest from the longest
		sideLengths[longestIndex] -= shortest

		// Append the new rectangle size
		if sideLengths[0] < sideLengths[1] {
			rect := fmt.Sprintf("(%d*%d)", sideLengths[1], sideLengths[0])
			outputRectangles = append(outputRectangles, rect)
		} else if sideLengths[1] < sideLengths[0] {
			rect := fmt.Sprintf("(%d*%d)", sideLengths[0], sideLengths[1])
			outputRectangles = append(outputRectangles, rect)
		}
	}

	return outputRectangles
}

func Min(values []int) int {
	if len(values) == 0 {
		return 0
	}

	var min = values[0]
	for _, v := range values {
		if v < min {
			min = v
		}
	}

	return min
}
