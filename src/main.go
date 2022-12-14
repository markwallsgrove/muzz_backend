package main

import (
	"context"
	"os"
	"os/signal"

	"github.com/labstack/echo/v4"
	"github.com/markwallsgrove/muzz_devops/src/database"
	"github.com/markwallsgrove/muzz_devops/src/middleware"
	"github.com/markwallsgrove/muzz_devops/src/routes"
	"go.uber.org/zap"
)

func main() {
	logger, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}

	dsn := os.Getenv("DSN")
	db, err := database.NewMariaDB(dsn, logger)
	if err != nil {
		panic(err)
	}

	secret := os.Getenv("SIGNING_SECRET")

	// Setup the controllers
	ctx, close := context.WithCancel(context.Background())
	u := routes.NewUserController(ctx, db, logger)
	i := routes.IndexController{}
	l := routes.LoginController{
		Ctx:      ctx,
		Database: db,
		Logger:   logger,
		Secret:   secret,
	}

	auth := middleware.JWTAuth{
		DB:     db,
		Secret: secret,
	}

	// Register the handlers with echo
	e := echo.New()
	e.GET("/", i.Index)
	e.POST("/user/create", u.CreateUser)
	e.GET("/profiles", u.Profiles, auth.Process)
	e.POST("/swipe", u.Swipe, auth.Process)
	e.POST("/login", l.Login)

	// Handle any signals to close the application. The connection to the database
	// must be cleanly closed.
	appSignal := make(chan os.Signal, 3)
	signal.Notify(appSignal, os.Interrupt)

	go func() {
		<-appSignal
		close()
		if err := db.Close(); err != nil {
			logger.Error("failed to close databse connection", zap.Error(err))
		}
	}()

	if err := e.Start(":8080"); err != nil {
		logger.Fatal("failed to start the server", zap.Error(err))
	}
}
