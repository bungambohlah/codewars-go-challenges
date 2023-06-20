package main

import (
	build_a_pile_of_cubes "codewars/build_a_pile_of_cubes"
	eureka "codewars/eureka"
	firstnonrepeatingcharacter "codewars/first_non_repeating_character"
	rectintorectangles "codewars/rect_into_rectangles"
	"fmt"
)

func main() {
	fmt.Println("rect_into_rectangles :", rectintorectangles.RectIntoRects(22, 6))
	fmt.Println("first_non_repeating_character :", firstnonrepeatingcharacter.FirstNonRepeating("aaabbc"))
	fmt.Println("build_a_pile_of_cubes :", build_a_pile_of_cubes.FindNb(1071225))
	fmt.Println("eurek :", eureka.SumDigPow(1, 100))
}
