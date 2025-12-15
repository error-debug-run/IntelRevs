package manager

import (
	"errors"
	"fmt"
	"time"

	"github.com/error-debug-run/go-scraper/internal/detector"
	"github.com/error-debug-run/go-scraper/internal/parser"
	"github.com/error-debug-run/go-scraper/internal/parser/site_alloc"
	"github.com/error-debug-run/go-scraper/internal/worker"
)

type ScrapeResult struct {
	URL       string   `json:"url"`
	Reviews   []string `json:"reviews"`
	Source    string   `json:"source"`
	Timestamp int64    `json:"timestamp"`
}

func RunScrapeJob(url string) (*ScrapeResult, error) {
	if url == "" {
		return nil, errors.New("empty url")
	}

	site := detector.DetectSite(url)

	var p parser.Parser
	switch site {
	case detector.SiteAmazon:
		p = site_alloc.NewAmazonParser()
	case detector.SiteReddit:
		p = site_alloc.NewRedditParser()
	case detector.SiteFlipkart:
		p = site_alloc.NewFlipkartParser()
	default:
		p = site_alloc.NewGenericParser()
	}

	w := worker.NewHTTPWorker()

	var allReviews []string
	currentURL := url

	const maxPages = 10
	pageCount := 0

	for {
		if pageCount >= maxPages {
			break
		}

		html, err := w.FetchHTML(currentURL)
		if err != nil {
			return nil, err
		}
		fmt.Println("Fetched HTML length:", len(html))

		reviews, err := p.Parse(html)
		if err != nil {
			return nil, err
		}
		println("PARSED REVIEWS COUNT:", len(reviews))
		println("HTML LENGTH:", len(html))

		allReviews = append(allReviews, reviews...)

		next := parser.FindNextPage(html)
		if next == "" {
			break
		}

		currentURL = resolveURL(currentURL, next)
		pageCount++
	}

	return &ScrapeResult{
		URL:       url,
		Reviews:   allReviews,
		Source:    string(site),
		Timestamp: time.Now().Unix(),
	}, nil
}
