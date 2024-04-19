package utils

import (
	"regexp"
)

//Matching all the numbers as well as characters except punctuations
func StripWhiteSpacesAndPunctuations(input string) string {
	str := regexp.MustCompile(`[^\p{N}\p{L}]+`).ReplaceAllString(input, "")
	return str
}
