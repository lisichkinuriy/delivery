package main

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"github.com/robfig/cron/v3"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"lisichkinuriy/delivery/cmd"
	httpin "lisichkinuriy/delivery/internal/adapters/in/http"
	"lisichkinuriy/delivery/internal/adapters/out/postgres/courierrepo"
	"lisichkinuriy/delivery/internal/adapters/out/postgres/orderrepo"
	"lisichkinuriy/delivery/pkg/servers"
	"net/http"
	"os"
)

func main() {
	ctx := context.Background()
	gormDB := mustGormOpen()
	mustAutoMigrate(gormDB)
	app := cmd.NewCompositionRoot(ctx, gormDB)

	port := getEnvVariable("HTTP_PORT", "8081")
	startCron(app)
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

func startCron(compositionRoot cmd.CompositionRoot) {
	c := cron.New()
	_, err := c.AddFunc("@every 1s", compositionRoot.Jobs.AssignOrdersJob.Run)
	if err != nil {
		log.Fatalf("ошибка при добавлении задачи: %v", err)
	}
	_, err = c.AddFunc("@every 2s", compositionRoot.Jobs.MoveCouriersJob.Run)
	if err != nil {
		log.Fatalf("ошибка при добавлении задачи: %v", err)
	}
	c.Start()
}

func startWebServer(compositionRoot cmd.CompositionRoot, port string) {

	serverHandlers, err := httpin.NewServer(compositionRoot.CommandHandlers.CreateOrderHandler,
		compositionRoot.QueryHandlers.GetAllCouriersQueryHandler,
		compositionRoot.QueryHandlers.GetNotCompletedOrdersQueryHandler)
	if err != nil {
		log.Fatalf(err.Error())
	}

	e := echo.New()

	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.DELETE, echo.OPTIONS},
	}))

	servers.RegisterHandlers(e, serverHandlers)

	e.GET("/health", func(c echo.Context) error {
		e.Logger.Info("Health check")
		return c.String(http.StatusOK, "Delivery is healthy")
	})

	e.Logger.Fatal(e.Start(fmt.Sprintf("0.0.0.0:%s", port)))

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
