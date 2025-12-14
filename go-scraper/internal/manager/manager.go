package manager

import (
	"time"

	"github.com/error-debug-run/go-scraper/internal/detector"
)

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

	site := detector.DetectSite(url)

	var p parser.Parser

	switch site {
	case detector.SiteAmazon:
		p = parser.NewGenericParser() // placeholder
	case detector.SiteReddit:
		p = parser.NewGenericParser() // placeholder
	default:
		p = parser.NewGenericParser()
	}

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
