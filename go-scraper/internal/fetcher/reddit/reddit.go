package reddit

import (
	"context"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/error-debug-run/go-scraper/internal/fetcher"
)

// RedditFetcher implements Fetcher for Reddit URLs.
//
// It fetches Reddit's native JSON representation
// and returns it as raw bytes.
type RedditFetcher struct {
	client *http.Client
}

// New creates a RedditFetcher with conservative defaults.
func New() *RedditFetcher {
	return &RedditFetcher{
		client: &http.Client{
			Timeout: 20 * time.Second,
		},
	}
}

func (r *RedditFetcher) Fetch(ctx context.Context, rawURL string) (any, map[string]string, error) {
	jsonURL, err := normalizedRedditURL(rawURL)
	if err != nil {
		return nil, nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, jsonURL, nil)
	if err != nil {
		return nil, nil, err
	}

	req.Header.Set("User-Agent", "IntelRevsBot/0.1 (by /u/intelrevs)")
	req.Header.Set("Accept", "application/json")

	resp, err := r.client.Do(req)
	if err != nil {
		return nil, nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, nil, fmt.Errorf("unexpected status code non-200: %d", resp.StatusCode)
	}

	log.Println("Content-Encoding:", resp.Header.Get("Content-Encoding"))
	log.Println("Content-Type:", resp.Header.Get("Content-Type"))

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, nil, err
	}

	log.Println("payload preview:", string(body[:300]))

	meta := map[string]string{
		"source":         "reddit",
		"content-type":   resp.Header.Get("Content-Type; charset=utf-8"),
		"status":         resp.Status,
		"normalized_url": jsonURL,
	}

	return string(body), meta, nil
}

// normalizeRedditURL converts a Reddit post URL
// into its JSON endpoint.
//
// Examples:
// https://www.reddit.com/r/golang/comments/abc123/post/
// → https://old.reddit.com/r/golang/comments/abc123/post/.json
func normalizedRedditURL(raw string) (string, error) {
	u, err := url.Parse(raw)
	if err != nil {
		return "", err
	}

	if !strings.Contains(u.Host, "reddit.com") {
		return "", errors.New("invalid reddit url")
	}

	u.Host = "old.reddit.com"

	if !strings.HasSuffix(u.Path, ".json") {
		u.Path = strings.TrimSuffix(u.Path, "/") + ".json"
	}

	return u.String(), nil
}

func init() {
	fetcher.Register("reddit", New())
}
