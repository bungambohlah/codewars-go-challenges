package main

import (
	build_a_pile_of_cubes "codewars/solutions/build_a_pile_of_cubes"
	decode_the_morse_code "codewars/solutions/decode_the_morse_code"
	equal_sides_of_an_array "codewars/solutions/equal_sides_of_an_array"
	eureka "codewars/solutions/eureka"
	firstnonrepeatingcharacter "codewars/solutions/first_non_repeating_character"
	matrix_determinant "codewars/solutions/matrix_determinant"
	rectintorectangles "codewars/solutions/rect_into_rectangles"
	romansnumeraldecoder "codewars/solutions/romans_numeral_decoder"
	twosum "codewars/solutions/two_sum"
	"fmt"
)

func main() {
	fmt.Println("build_a_pile_of_cubes :", build_a_pile_of_cubes.FindNb(1071225))
	fmt.Println("decode_the_morse_code :", decode_the_morse_code.DecodeMorse(".... . -.--   .--- ..- -.. ."))
	fmt.Println("Equal Sides of an Array :", equal_sides_of_an_array.FindEvenIndex([]int{1, 2, 3, 4, 3, 2, 1}))
	fmt.Println("eurek :", eureka.SumDigPow(1, 100))
	fmt.Println("first_non_repeating_character :", firstnonrepeatingcharacter.FirstNonRepeating("aaabbc"))
	fmt.Println("matrix_determinant :", matrix_determinant.Determinant([][]int{{1, 3}, {2, 5}}))
	fmt.Println("rect_into_rectangles :", rectintorectangles.RectIntoRects(22, 6))
	fmt.Println("romans_numeral_decoder :", romansnumeraldecoder.Decode("IV"))
	fmt.Println("two_sum :", twosum.TwoSum([]int{1, 2, 3}, 4))
}
