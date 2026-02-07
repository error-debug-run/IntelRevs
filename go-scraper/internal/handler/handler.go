package handler

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/error-debug-run/go-scraper/internal/detector"
	"github.com/error-debug-run/go-scraper/internal/envelope"
	"github.com/error-debug-run/go-scraper/internal/fetcher"
	_ "github.com/error-debug-run/go-scraper/internal/fetcher/generic"
	_ "github.com/error-debug-run/go-scraper/internal/fetcher/reddit"
)

// ScraperHandler handles all incoming scrape requests.
//
// Endpoint:
//
//	GET /v1/scraper?url=<target_url>
//
// Behavior:
// - Accepts only GET requests
// - Requires a `url` query parameter
// - Enforces a strict timeout
// - Delegates site detection and fetching
// - Returns a canonical response envelope
//
// This handler never panics and never returns malformed JSON.
func ScraperHandler(w http.ResponseWriter, r *http.Request) {

	// All responses from this handler are JSON.
	w.Header().Set("Content-Type", "application/json")

	// Enforce HTTP method.
	//
	// POST, PUT, DELETE, etc. are intentionally rejected
	// to keep the API surface minimal and predictable.
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		_ = json.NewEncoder(w).Encode(
			envelope.Error("METHOD_NOT_ALLOWED", "only GET is supported"),
		)
		return
	}

	// Extract target URL.
	//
	// The scraper operates on exactly one URL per request.
	targetURL := r.URL.Query().Get("url")
	if targetURL == "" {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(
			envelope.Error("MISSING_URL", "query parameter 'url' is required"),
		)
		return
	}

	// Create a request-scoped context with timeout.
	//
	// This prevents:
	// - hanging connections
	// - slow remote servers exhausting resources
	// - unbounded scraper execution
	ctx, cancel := context.WithTimeout(r.Context(), 8*time.Second)
	defer cancel()

	// Detect the source site.
	//
	// Detection is deterministic and does not perform
	// any network operations.
	source := detector.Detect(targetURL)

	// Select appropriate fetcher for the detected site.
	//
	// Each fetcher encapsulates site-specific behavior.
	f := fetcher.Get(source)

	if f == nil {
		log.Fatal("fetcher not registered")
	}

	// Execute fetch operation.
	//
	// Fetchers return raw payloads only.
	payload, meta, err := f.Fetch(ctx, targetURL)
	if err != nil {
		w.WriteHeader(http.StatusBadGateway)
		_ = json.NewEncoder(w).Encode(
			envelope.FromError(source, targetURL, err),
		)
		return
	}

	// Construct successful response envelope.
	//
	// Payload is passed through untouched.
	response := envelope.Success(source, targetURL, payload, meta)

	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(response)
}
