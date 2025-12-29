package reddit

import (
	"strings"

	"github.com/error-debug-run/go-scraper/internal/worker"
)

// FetchRedditJSON ensures .json is fetched
func FetchRedditJSON(url string) (string, error) {
	url = strings.TrimSuffix(url, "/") + ".json"
	return worker.FetchPage(url)
}
