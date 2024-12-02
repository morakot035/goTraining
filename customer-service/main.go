package main

import (
	"context"
	"goTraining/customer-service/service"
	"goTraining/customer-service/config"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
)

func main() {
	go func() {
		config.DatabaseInit()
		gorm := config.Database()

		dbGorm, err := gorm.DB()
		if err != nil {
			panic(err)
		}
		err = dbGorm.Ping()
		if err != nil {
			panic(err)
		}
	}()

	app := echo.New()
	app.Use(middleware.Recover())
	app.Use(middleware.Logger())

	app.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Success")
	})

	app.POST("/order", service.PlaceOrder)

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	go func() {
		if err := app.Start(":3000"); err != nil && err != http.ErrServerClosed {
			log.Fatal("shutting down the server")
		}
	}()

	<-ctx.Done()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := app.Shutdown(ctx); err != nil {
		log.Fatal(err)
	}
}
