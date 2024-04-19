package server

import (
	"net/http"

	"github.com/RishabAkalankan/stringinator/logger"
	"github.com/labstack/echo/v4"
)

var seen_strings map[string]int = make(map[string]int)

type StringRequest struct {
	Input string `param:"input" query:"input" form:"input" json:"input" xml:"input" validate:"required,min=1"`
}

type StringData struct {
	Input  string `param:"input" query:"input" form:"input" json:"input" xml:"input" validate:"required,min=1"`
	Length int    `json:"length"`
}

type StatsData struct {
	Inputs map[string]int `json:"inputs"`
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
	remember(request_data.Input)
	response_data := StringData{request_data.Input, len(request_data.Input)}
	return c.JSON(http.StatusOK, response_data)
}

func GetStats(c echo.Context) (err error) {
	return c.JSON(http.StatusOK, StatsData{seen_strings})
}

func GetWelcomeMessage(c echo.Context) (err error) {
	return c.HTML(http.StatusOK, `
		<pre>
		Welcome to the Stringinator 3000 for all of your string manipulation needs.
		GET / - You're already here!
		POST /stringinate - Get all of the info you've ever wanted about a string. Takes JSON of the following form: {"input":"your-string-goes-here"}
		GET /stats - Get statistics about all strings the server has seen, including the longest and most popular strings.
		</pre>
	`)
}

func remember(input string) {
	if seen_strings[input] == 0 {
		seen_strings[input] = 1
	} else {
		seen_strings[input] += 1
	}
}