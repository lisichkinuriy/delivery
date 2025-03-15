package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"lisichkinuriy/delivery/cmd"
	"net/http"
	"os"
)

func main() {
	// TODO. context

	app := cmd.NewCompositionRoot()

	port := getEnvVariable("HTTP_PORT", "8081")
	startWebServer(app, port)
}

func startWebServer(compositionRoot cmd.CompositionRoot, port string) {
	e := echo.New()
	e.GET("/health", func(c echo.Context) error {
		e.Logger.Info("Health check")
		return c.String(http.StatusOK, "healthy")
	})

	err := e.Start(fmt.Sprintf("0.0.0.0:%s", port))
	if err != nil {
		e.Logger.Fatal(err)
	}
}

func getEnvVariable(key, fallback string) string {
	// TODO: вынести load файла отдельно?
	err := godotenv.Load(".env")
	if err != nil {
		log.Warn("Error loading .env file")
		return fallback
	}

	return os.Getenv(key)
}
