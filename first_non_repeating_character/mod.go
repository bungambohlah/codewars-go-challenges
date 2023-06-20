package firstnonrepeatingcharacter

import "strings"

func FirstNonRepeating(str string) string {
	// if str length is 0 then return empty string
	if len(str) == 0 {
		return ""
	}

	// make string lowercased
	strCopy := strings.ToLower(str)
	lengthStr := len(str)
	// return non repeating character in string
	for i := 0; i < lengthStr; i++ {
		if strings.Count(strCopy, string(strCopy[i])) == 1 {
			return string(str[i])
		}
	}

	return ""
}
