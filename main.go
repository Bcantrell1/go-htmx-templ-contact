package main

import (
	"context"
	"embed"
	"fmt"
	"io/fs"
	"net/http"
	"os"
	"os/signal"
	"time"

	"app/routes"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/joho/godotenv"
)

//go:embed static
var static embed.FS

func Static() http.Handler {
	//Server static files to support TailwindCSS
	dist, err := fs.Sub(static, "static")
	if err != nil {
		panic(err)
	}

	return http.FileServer(http.FS(dist))
}

func main() {
	// Create a new echo instance
	e := echo.New()

	godotenv.Load(".env")

	// Add static files to the server
	e.GET("/static/*", echo.WrapHandler(http.StripPrefix("/static/", Static())))

	// Middleware
	e.Use(middleware.Logger())

	// bind routes to the server
	routes.Routes(e)

	// Start server
	go func() {
		if err := e.Start(fmt.Sprintf(":%s", os.Getenv("PORT"))); err != nil && err != http.ErrServerClosed {
			e.Logger.Fatal("shutting down the server", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with a timeout of 10 seconds.
	// Use a buffered channel to avoid missing signals as recommended for signal.Notify
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}
