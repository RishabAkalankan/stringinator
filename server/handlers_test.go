package server

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	iv "github.com/RishabAkalankan/stringinator/validator"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func Test_getMostOccurringCharactersWithSingleMostOccurringCharacter(t *testing.T) {
	inputString := "Güüüüünter"
	strArr, times := getMostOccurringCharacters(inputString)
	assert.Equal(t, 5, times)
	assert.Equal(t, []string{"ü"}, strArr)
}

func Test_getMostOccurringCharactersWithMultipleMostOccurringCharacter(t *testing.T) {
	inputString := "Güüüüünter Waaaaarner"
	strArr, times := getMostOccurringCharacters(inputString)
	assert.Equal(t, 5, times)
	assert.Contains(t, strArr, "ü")
	assert.Contains(t, strArr, "a")
}

func Test_StringinateUsingGetMethodThrowsErrorWhenTheInputIsInvalid(t *testing.T) {
	e := echo.New()
	e.Validator = &iv.InputValidator{Validator: validator.New()}
	req := httptest.NewRequest("GET", "/stringinate", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	err := Stringinate(c)
	assert.NotNil(t, err)
	httpError := err.(*echo.HTTPError)
	errMsg := httpError.Message.(map[string]string)["message"]
	assert.Contains(t, errMsg, "validation failed.'input' is a required parameter")
	assert.Equal(t, http.StatusBadRequest, httpError.Code)
}

func Test_StringinateUsingPostThrowsErrorWhenTheInputIsInvalid(t *testing.T) {
	inputJSON := `{"inputs":"David Warner"}`
	e := echo.New()
	e.Validator = &iv.InputValidator{Validator: validator.New()}
	req := httptest.NewRequest("POST", "/stringinate", strings.NewReader(inputJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	err := Stringinate(c)
	assert.NotNil(t, err)
	httpError := err.(*echo.HTTPError)
	errMsg := httpError.Message.(map[string]string)["message"]
	assert.Contains(t, errMsg, "validation failed.'input' is a required parameter")
	assert.Equal(t, http.StatusBadRequest, httpError.Code)
}

func Test_StringinateUsingPostMethodReturnsProperResponse(t *testing.T) {
	inputJSON := `{"input":"David Warner"}`
	e := echo.New()
	e.Validator = &iv.InputValidator{Validator: validator.New()}
	req := httptest.NewRequest("POST", "/stringinate", strings.NewReader(inputJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	err := Stringinate(c)
	assert.Nil(t, err)
	//expected response - `{"value":"David Warner","length":12,"insights":{"mostOccuringLetters":["r","a"],"occurrences":2}}`
	var resp InputStats
	json.Unmarshal(rec.Body.Bytes(), &resp)
	assert.Equal(t, "David Warner", resp.Value)
	assert.Equal(t, 12, resp.Length)
	assert.Contains(t, resp.Insights.MostOccuringLetters, "r")
	assert.Contains(t, resp.Insights.MostOccuringLetters, "a")
	assert.Equal(t, resp.Insights.Occurrences, 2)
}

func Test_StringinateUsingGetMethodReturnsProperResponse(t *testing.T) {
	e := echo.New()
	e.Validator = &iv.InputValidator{Validator: validator.New()}

	q := make(url.Values)
	q.Set("input", "Jos Buttler")
	req := httptest.NewRequest("GET", "/stringinate?"+q.Encode(), nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	err := Stringinate(c)
	assert.Nil(t, err)
	//expected response - `{"value":"Jos Buttler","length":11,"insights":{"mostOccuringLetters":["t"],"occurrences":2}}`
	var resp InputStats
	json.Unmarshal(rec.Body.Bytes(), &resp)
	assert.Equal(t, "Jos Buttler", resp.Value)
	assert.Equal(t, 11, resp.Length)
	assert.Contains(t, resp.Insights.MostOccuringLetters, "t")
	assert.Equal(t, resp.Insights.Occurrences, 2)
}
