package main

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"lisichkinuriy/delivery/cmd"
	"lisichkinuriy/delivery/internal/adapters/out/postgres/courierrepo"
	"lisichkinuriy/delivery/internal/adapters/out/postgres/orderrepo"
	"net/http"
	"os"
)

func main() {
	ctx := context.Background()
	gormDB := mustGormOpen()
	mustAutoMigrate(gormDB)
	app := cmd.NewCompositionRoot(ctx, gormDB)

	port := getEnvVariable("HTTP_PORT", "8081")
	startWebServer(app, port)
}

func mustGormOpen() *gorm.DB {
	dsn := "host=localhost user=username password=secret dbname=delivery port=5490 sslmode=disable TimeZone=Europe/Moscow"
	pgGorm, err := gorm.Open(postgres.New(
		postgres.Config{
			DSN:                  dsn,
			PreferSimpleProtocol: true,
		},
	), &gorm.Config{})
	if err != nil {
		log.Fatalf("connection to postgres through gorm\n: %s", err)
	}
	return pgGorm
}

func mustAutoMigrate(db *gorm.DB) {
	err := db.AutoMigrate(&courierrepo.CourierDTO{})
	if err != nil {
		log.Fatalf("Ошибка миграции: %v", err)
	}

	err = db.AutoMigrate(&courierrepo.TransportDTO{})
	if err != nil {
		log.Fatalf("Ошибка миграции: %v", err)
	}

	err = db.AutoMigrate(&orderrepo.OrderDTO{})
	if err != nil {
		log.Fatalf("Ошибка миграции: %v", err)
	}
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
