package server

import (
	"net/http"
	"time"

	"github.com/RishabAkalankan/stringinator/logger"
	iv "github.com/RishabAkalankan/stringinator/validator"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Start() {
	e := echo.New()
	e.Validator = &iv.InputValidator{Validator: validator.New()}

	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		logger.UpdateRequestId()
		return next
	})

	rateLimiterConfig := middleware.RateLimiterConfig{
		Skipper: middleware.DefaultSkipper,
		Store:   middleware.NewRateLimiterMemoryStoreWithConfig(middleware.RateLimiterMemoryStoreConfig{Rate: 10, Burst: 5, ExpiresIn: 3 * time.Minute}),
		IdentifierExtractor: func(context echo.Context) (string, error) {
			ip := context.RealIP()
			return ip, nil
		},
		DenyHandler: func(context echo.Context, identifier string, err error) error {
			logger.Errorf("Rate limit exceeded for ip: %s", identifier)
			return context.JSON(http.StatusTooManyRequests, map[string]string{
				"message": "rate limit exceeded. please try again after sometime",
			})
		},
	}

	e.Use(middleware.RateLimiterWithConfig(rateLimiterConfig))

	//middleware for logging the incoming request
	e.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogURI:    true,
		LogStatus: true,
		LogMethod: true,
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			logger.Infof("Received Request for: [%s] %s", v.Method, v.URI)
			return nil
		},
	}))

	registerRoutes(e)
	err := e.Start(":1323")

	if err != nil {
		logger.Fatalf("error occured while starting the server %+v", err)
	}
}

func registerRoutes(e *echo.Echo) {
	e.GET("/", GetWelcomeMessage)
	e.POST("/stringinate", Stringinate)
	e.GET("/stringinate", Stringinate)
	e.GET("/stats", GetStats)
}
