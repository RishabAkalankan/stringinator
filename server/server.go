package server

import (
	"github.com/RishabAkalankan/stringinator/logger"
	iv "github.com/RishabAkalankan/stringinator/validator"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Start() {
	e := echo.New()
	e.Validator = &iv.InputValidator{Validator: validator.New()}

	//middleware for logging the incoming request
	//TODO: Print the request body for debugging purposes
	e.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogURI:    true,
		LogStatus: true,
		LogMethod: true,
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			logger.Infof("Received Request for: [%s] %s", v.Method, v.URI)
			return nil
		},
	}))
	e.GET("/", GetWelcomeMessage)

	e.POST("/stringinate", Stringinate)
	e.GET("/stringinate", Stringinate)
	e.GET("/stats", GetStats)
	err := e.Start(":1323")

	if err != nil {
		logger.Fatalf("error occured while starting the server %+v", err)
	}
}
