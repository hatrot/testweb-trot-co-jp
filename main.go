package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/labstack/echo"
)

var TimeZone = time.FixedZone("Asia/Tokyo", 9*60*60)

func main() {
	e := echo.New()
	http.Handle("/", e)
	//
	e.GET("/_ah/health", healthCheckHandler)
	e.GET("/_ah/warmup", warmupHandler)
	// e.GET("/*", Error503Handler)
	e.GET("/*", IndexHandler)
	//
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}
	log.Printf("Listening on port %s", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}

func healthCheckHandler(c echo.Context) error {
	c.Response().Header().Set("Cache-Control", "no-store")
	c.Response().Header().Set("Content-Type", "text/plain")
	return c.String(http.StatusOK, "ok")
}

func warmupHandler(c echo.Context) error {
	c.Response().Header().Set("Cache-Control", "no-store")
	c.Response().Header().Set("Content-Type", "text/plain")
	return c.String(http.StatusOK, "warmup done")
}

// 503 error function =============================================
func Error503Handler(c echo.Context) error {
	c.Response().Header().Set("Cache-Control", "no-store")
	c.Response().Header().Set("Content-Type", "text/plain")
	return c.String(http.StatusServiceUnavailable, "HTTP/2 503 Bad\n\n 503 error : Service Unavailable.\n\n")
}

// Main function ===============================================...
func IndexHandler(c echo.Context) error {
	c.Response().Header().Set("Cache-Control", "no-store")
	c.Response().Header().Set("Content-Type", "text/plain")
	return c.String(http.StatusOK, "HTTP/2 200 OK\n\nGitHub Actions : Hello world..\n\n")
}
