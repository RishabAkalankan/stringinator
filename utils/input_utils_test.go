package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStripWhiteSpacesAndPunctuationsWithAccentedLetters(t *testing.T) {
	inputString := "Gúnter;-Warner"
	outputString := StripWhiteSpacesAndPunctuations(inputString)
	assert.Equal(t, "GúnterWarner", outputString)
}

func TestStripWhiteSpacesAndPunctuationsWithAccentedLettersAndWhiteSpace(t *testing.T) {
	inputString := "Gúnter;- Warners"
	outputString := StripWhiteSpacesAndPunctuations(inputString)
	assert.Equal(t, "GúnterWarners", outputString)
}
