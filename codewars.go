package main

import (
	build_a_pile_of_cubes "codewars/solutions/build_a_pile_of_cubes"
	equal_sides_of_an_array "codewars/solutions/equal_sides_of_an_array"
	eureka "codewars/solutions/eureka"
	firstnonrepeatingcharacter "codewars/solutions/first_non_repeating_character"
	rectintorectangles "codewars/solutions/rect_into_rectangles"
	"fmt"
)

func main() {
	fmt.Println("build_a_pile_of_cubes :", build_a_pile_of_cubes.FindNb(1071225))
	fmt.Println("Equal Sides of an Array :", equal_sides_of_an_array.FindEvenIndex([]int{1, 2, 3, 4, 3, 2, 1}))
	fmt.Println("eurek :", eureka.SumDigPow(1, 100))
	fmt.Println("first_non_repeating_character :", firstnonrepeatingcharacter.FirstNonRepeating("aaabbc"))
	fmt.Println("rect_into_rectangles :", rectintorectangles.RectIntoRects(22, 6))
}
