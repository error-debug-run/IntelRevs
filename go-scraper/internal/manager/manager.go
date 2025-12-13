package manager

import "time"

import (
	"github.com/error-debug-run/go-scraper/internal/parser"
	"github.com/error-debug-run/go-scraper/internal/worker"
)

type ScrapeResult struct {
	URL       string   `json:"url"`
	Reviews   []string `json:"reviews"`
	Source    string   `json:"source"`
	Timestamp int64    `json:"timestamp"`
}

func RunScrapeJob(url string) (*ScrapeResult, error) {

	rawHTML, err := worker.FetchPage(url)
	if err != nil {
		return nil, err
	}

	reviews, err := parser.ExtractReviews(rawHTML)
	if err != nil {
		return nil, err
	}

	if reviews == nil {
		reviews = []string{}
	}

	return &ScrapeResult{
		URL:       url,
		Reviews:   reviews,
		Source:    "generic",
		Timestamp: time.Now().Unix(),
	}, nil

}
