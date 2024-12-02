package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"time"

	"goTraining/restaurant-service/client"
	"goTraining/restaurant-service/config"
	"goTraining/restaurant-service/service"

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
		return c.String(http.StatusOK, "Hello, World!")
	})
	app.GET("/restaurant", service.GetRestaurants)
	app.GET("/menu", service.GetMenu)
	app.POST("/restaurant/order/accept", service.AcceptingOrder)

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	go func() {
		if err := app.Start(":8082"); err != nil && err != http.ErrServerClosed {
			log.Fatal("shutting down the server")
		}
	}()

	doneListening := make(chan bool)
	go client.ListenForNotification(doneListening)

	<-ctx.Done()
	close(doneListening) 
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := app.Shutdown(ctx); err != nil {
		log.Fatal(err)
	}
}