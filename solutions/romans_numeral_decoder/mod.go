package romansnumeraldecoder

func Decode(roman string) int {
	romanMap := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}
	romanLength := len(roman)
	romanNumeral := 0
	romanIndex := 0
	romanValue := romanMap[string(roman[romanIndex])]
	for romanIndex < romanLength {
		if romanIndex+1 < romanLength && romanValue < romanMap[string(roman[romanIndex+1])] {
			romanNumeral -= romanValue
		} else {
			romanNumeral += romanValue
		}

		romanIndex++

		if romanIndex < romanLength {
			romanValue = romanMap[string(roman[romanIndex])]
		}
	}

	return romanNumeral
}
