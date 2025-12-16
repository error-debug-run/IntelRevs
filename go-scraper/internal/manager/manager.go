package manager

import (
	"time"

	"github.com/error-debug-run/go-scraper/internal/parser"
	"github.com/error-debug-run/go-scraper/internal/worker"
)

type ScrapeResult struct {
	URL     string   `json:"url"`
	Reviews []string `json:"reviews"`
}

func RunScrapeJob(url string) (*ScrapeResult, error) {
	site := detector.DetectSite(url)

	var p parser.Parser
	switch site {
	case detector.SiteReddit:
		p = parser.NewRedditParser()
	default:
		p = parser.NewGenericParser()
	}

	worker := worker.NewHTTPWorker()
	html, err := worker.FetchHTML(url)
	if err != nil {
		return nil, err
	}

	reviews, err := p.Parse(html)
	if err != nil {
		return nil, err
	}

	return &ScrapeResult{
		URL:       url,
		Reviews:   reviews,
		Source:    site.String(),
		Timestamp: time.Now().Unix(),
	}, nil
}
