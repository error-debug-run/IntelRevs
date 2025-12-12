package manager

import (
	"github.com/error-debug-run/go-scraper/internal/parser"
	"github.com/error-debug-run/go-scraper/internal/worker"
)

type ScrapeResult struct {
	URL     string   `json:"url"`
	Reviews []string `json:"reviews"`
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

	return &ScrapeResult{
		URL:     url,
		Reviews: reviews,
	}, nil
}
