package main

import (
	"fmt"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"net/http"
	"simpleservice/capitals"
	"time"
)

func corsConfig() middleware.CORSConfig {
	headers := []string{echo.HeaderAccept, echo.HeaderAcceptEncoding, echo.HeaderAuthorization, echo.HeaderContentLength, echo.HeaderContentType, echo.HeaderOrigin, echo.HeaderXCSRFToken}

	return middleware.CORSConfig{
		AllowHeaders:     headers,
		AllowCredentials: true,
		ExposeHeaders:    headers,
		AllowOrigins:     []string{"*"},
	}
}

func serverGet(e *echo.Echo) *http.Server {
	webservicePort := 80
	readWriteTimeout := time.Second * 15

	return &http.Server{
		Addr:         fmt.Sprintf(":%d", webservicePort),
		Handler:      e,
		ReadTimeout:  readWriteTimeout,
		WriteTimeout: readWriteTimeout,
		IdleTimeout:  time.Second * 60,
	}
}

func webserviceStart() {
	fmt.Println("starting webservices")

	e := echo.New()

	// CORS is not working if you define it on group or subgroup.
	// So, "echo" instance is the highest level, and is where we should define cors middleware
	e.Use(middleware.CORSWithConfig(corsConfig()))

	r := e.Group("/simple")

	r.Use(middleware.Recover(), middleware.Logger())

	r.GET("/uptime/", func(c echo.Context) error {
		return c.String(http.StatusOK, time.Since(bootTime).String())
	})

	r.GET("/ping/", func(c echo.Context) error {
		return c.String(http.StatusOK, "pong")
	})

	r.GET("/capital/:country/", func(c echo.Context) error {
		return capitals.ByCountry(c)
	})

	r.GET("/capital/", func(c echo.Context) error {
		return capitals.All(c)
	})

	server := serverGet(e)

	if err := e.StartServer(server); err != nil {
		fmt.Println(err)
	}
}
