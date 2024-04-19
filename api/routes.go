package api

import (
	"github.com/RishabAkalankan/stringinator/logger"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func StartServer() {
	e := echo.New()
	e.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogURI:    true,
		LogStatus: true,
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			logger.Infof("Received Request for: %s", v.URI)
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
