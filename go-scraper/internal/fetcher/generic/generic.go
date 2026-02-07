package generic

import (
	"context"
	"io"
	"net/http"
	"time"

	"github.com/error-debug-run/go-scraper/internal/fetcher"
)

// GenericFetcher implements the Fetcher interface.
//
// It performs a simple HTTP GET request and returns
// the raw response body as bytes.
//
// No parsing, no transformation, no assumptions.
type GenericFetcher struct {
	client *http.Client
}

// New creates a new GenericFetcher with a sane HTTP client.
func New() *GenericFetcher {
	return &GenericFetcher{
		client: &http.Client{
			Timeout: 20 * time.Second,
		},
	}
}

// Fetch retrieves raw content from the provided URL.
//
// Returns:
// - payload: []byte (raw body)
// - meta: basic HTTP metadata
// - error: network / protocol errors only
func (g *GenericFetcher) Fetch(
	ctx context.Context,
	url string,
) (any, map[string]string, error) {

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, nil, err
	}

	// Minimal but valid headers
	req.Header.Set("User-Agent", "IntelRevsBot/0.1")
	req.Header.Set("Accept", "*/*")

	resp, err := g.client.Do(req)
	if err != nil {
		return nil, nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, nil, err
	}

	meta := map[string]string{
		"status":       resp.Status,
		"status_code":  http.StatusText(resp.StatusCode),
		"content_type": resp.Header.Get("Content-Type"),
		"source":       "generic",
	}

	return body, meta, nil
}

func init() {
	fetcher.Register("generic", New())
}
