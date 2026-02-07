package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/error-debug-run/go-scraper/internal/handler"
)

// main is the entry point of the Go scraper service.
//
// It performs only three actions:
// 1. load minimal configuration
// 2. register HTTP routes
// 3. start the HTTP server
//
// No application logic should ever be added here.
func main() {

	// Server address
	//
	// This allows:
	// - local development (:8080)
	// - containerized deployment (0.0.0.0:8080)
	// - environment-based overrides
	addr := setEnv("SCRAPER_ADDR", ":8080")

	// HTTP request multiplexer
	//
	// net/http ServeMux is intentionally used instead of
	// third-party frameworks to ensure:
	// - full control
	// - predictable behavior
	// - minimal dependencies
	mux := http.NewServeMux()

	////Explicit fetcher registration
	//fetcher.Register("reddit", reddit.New())
	//fetcher.Register("generic", generic.New())

	// Route registration
	//
	// All API routes must be registered explicitly here.
	// No dynamic routing or middleware chains are allowed.
	mux.HandleFunc("/v1/scraper", handler.ScraperHandler)

	// HTTP server configuration
	//
	// Timeouts are critical for scraper services to:
	// - prevent resource exhaustion
	// - protect against slow or abusive clients
	// - keep latency predictable
	server := &http.Server{
		Addr:              addr,
		Handler:           mux,
		ReadHeaderTimeout: 5 * time.Second,
		ReadTimeout:       10 * time.Second,
		WriteTimeout:      15 * time.Second,
		IdleTimeout:       60 * time.Second,
	}

	log.Printf("[scraper] server starting on %s", addr)

	// Start HTTP server (blocking call)
	//
	// http.ErrServerClosed is ignored because it is
	// returned during normal shutdown.
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("[scraper] server failed: %v", err)
	}
}

// getEnv retrieves an environment variable value or returns
// a fallback if the variable is not set.
//
// This function exists to keep configuration handling
// explicit and centralized.
func setEnv(key string, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}
