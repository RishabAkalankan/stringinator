package server

import (
	"net/http"
	"unicode/utf8"

	"github.com/RishabAkalankan/stringinator/logger"
	"github.com/RishabAkalankan/stringinator/utils"
	"github.com/labstack/echo/v4"
)

var seen_strings map[string]int = make(map[string]int)

type StringRequest struct {
	Input string `param:"input" query:"input" form:"input" json:"input" xml:"input" validate:"required,min=1"`
}

type StatsData struct {
	Inputs               map[string]int `json:"inputs"`
	MostPopular          string         `json:"most_popular"`
	LongestInputReceived string         `json:"longest_input_received"`
}

type InputStats struct {
	Value    string   `json:"value"`
	Length   int      `json:"length"`
	Insights Insights `json:"insights"`
}
type Insights struct {
	MostOccuringLetters []string `json:"mostOccuringLetters"`
	Occurrences         int      `json:"occurrences"`
}

func Stringinate(c echo.Context) (err error) {
	request_data := new(StringRequest)
	if err = c.Bind(request_data); err != nil {
		logger.Errorf("binding failed. %+v", err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := c.Validate(request_data); err != nil {
		logger.Errorf("validation failed. %+v", err)
		return echo.NewHTTPError(http.StatusBadRequest, map[string]string{
			"message": "validation failed.'input' is a required parameter",
		})
	}
	frequentlyOccurringCharacters, times := getMostOccurringCharacters(request_data.Input)
	saveUserInputs(request_data.Input)

	//considers each character as one rune
	response := InputStats{Value: request_data.Input,
		Length:   utf8.RuneCountInString(request_data.Input),
		Insights: Insights{MostOccuringLetters: frequentlyOccurringCharacters, Occurrences: times}}
	return c.JSON(http.StatusOK, response)
}

func GetStats(c echo.Context) (err error) {
	mostpopular, longestInput := getMostPopularAndLongestInputReceived()
	return c.JSON(http.StatusOK, StatsData{Inputs: seen_strings, MostPopular: mostpopular, LongestInputReceived: longestInput})
}

func GetWelcomeMessage(c echo.Context) (err error) {
	welcomeMessage := `
	<pre>
	Welcome to the Stringinator 3000 for all of your string manipulation needs.
	GET / - You're already here!
	POST /stringinate - Get all of the info you've ever wanted about a string. Takes JSON of the following form: {"input":"your-string-goes-here"}
	GET /stats - Get statistics about all strings the server has seen, including the longest and most popular strings.
	</pre>
`
	return c.HTML(http.StatusOK, welcomeMessage)
}

func getMostPopularAndLongestInputReceived() (string, string) {
	var statistics struct {
		value              string
		maxNoOfTimesCalled int
	} = struct {
		value              string
		maxNoOfTimesCalled int
	}{value: "", maxNoOfTimesCalled: 0}

	var lengthStatistics struct {
		value  string
		length int
	} = struct {
		value  string
		length int
	}{value: "", length: 0}

	for k, v := range seen_strings {
		lengthOfInput := utf8.RuneCountInString(k)
		if lengthOfInput > lengthStatistics.length {
			lengthStatistics.length = lengthOfInput
			lengthStatistics.value = k
		}
		if v > statistics.maxNoOfTimesCalled {
			statistics.maxNoOfTimesCalled = v
			statistics.value = k
		}
	}
	return statistics.value, lengthStatistics.value
}

func getMostOccurringCharacters(input string) ([]string, int) {
	strippedInput := utils.StripWhiteSpacesAndPunctuations(input)
	characterToOccurrenceMap := make(map[string]int)
	for _, c := range strippedInput {
		if _, ok := characterToOccurrenceMap[string(c)]; ok {
			characterToOccurrenceMap[string(c)] = characterToOccurrenceMap[string(c)] + 1
			continue
		}
		characterToOccurrenceMap[string(c)] = 1
	}
	occurrenceTocharactersMap := make(map[int][]string)
	maxLengthArray := 0
	for k, v := range characterToOccurrenceMap {
		if _, ok := occurrenceTocharactersMap[v]; ok {
			occurrenceTocharactersMap[v] = append(occurrenceTocharactersMap[v], k)
			if v > maxLengthArray {
				maxLengthArray = v
			}
			continue
		}
		occurrenceTocharactersMap[v] = []string{k}
		if v > maxLengthArray {
			maxLengthArray = v
		}
	}
	return occurrenceTocharactersMap[maxLengthArray], maxLengthArray
}

func saveUserInputs(input string) {
	if seen_strings[input] == 0 {
		seen_strings[input] = 1
	} else {
		seen_strings[input] += 1
	}
}
