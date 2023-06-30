package decodethemorsecode

import "strings"

var MORSE_CODE = map[string]string{
	".-":   "A",
	"-...": "B",
	"-.-.": "C",
	"-..":  "D",
	".":    "E",
	"..-.": "F",
	"--.":  "G",
	"....": "H",
	"..":   "I",
	".---": "J",
	"-.-":  "K",
	".-..": "L",
	"--":   "M",
	"-.":   "N",
	"---":  "O",
	".--.": "P",
	"--.-": "Q",
	".-.":  "R",
	"...":  "S",
	"-":    "T",
	"..-":  "U",
	"...-": "V",
	".--":  "W",
	"-..-": "X",
	"-.--": "Y",
	"--..": "Z",
}

func DecodeMorse(morseCode string) string {
	morseCode = strings.ReplaceAll(morseCode, "   ", " / ")
	words := strings.Split(morseCode, " ")
	decodedText := ""
	for _, word := range words {
		decodedChar, exists := MORSE_CODE[word]
		if exists {
			decodedText += decodedChar
		} else if word == "/" {
			decodedText += " "
		}
	}

	// Trim leading and trailing spaces
	decodedText = strings.TrimSpace(decodedText)

	return decodedText
}
