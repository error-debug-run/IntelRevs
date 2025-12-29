package manager

import (
	"errors"
	"fmt"
	"time"

	"github.com/error-debug-run/go-scraper/internal/detector"
	"github.com/error-debug-run/go-scraper/internal/parser"
	"github.com/error-debug-run/go-scraper/internal/parser/site_alloc"
	"github.com/error-debug-run/go-scraper/internal/parser/site_alloc/reddit"
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

	// 1️⃣ Detect site
	site := detector.DetectSite(url)

	// 2️⃣ Select parser
	var p parser.Parser
	switch site {
	case detector.SiteAmazon:
		p = site_alloc.NewAmazonParser()
	case detector.SiteReddit:
		p = reddit.NewRedditParser()
	case detector.SiteFlipkart:
		p = site_alloc.NewFlipkartParser()
	default:
		p = site_alloc.NewGenericParser()
	}

	var allReviews []string
	currentURL := url

	const maxPages = 10
	pageCount := 0

	for {
		if pageCount >= maxPages {
			break
		}

		var content string
		var err error

		// 3️⃣ Fetch page (DIFFERENT for Reddit)
		if site == detector.SiteReddit {
			content, err = worker.FetchPage(currentURL)
		} else {
			httpWorker := worker.NewHTTPWorker()
			content, err = httpWorker.FetchHTML(currentURL)
		}

		if err != nil {
			return nil, err
		}

		fmt.Println("FETCHED CONTENT LENGTH:", len(content))

		// 4️⃣ Parse content
		reviews, err := p.Parse(content)
		if err != nil {
			return nil, err
		}

		fmt.Println("PARSED REVIEWS COUNT:", len(reviews))

		allReviews = append(allReviews, reviews...)

		// 5️⃣ Pagination (HTML only for now)
		if site == detector.SiteReddit {
			break // pagination handled later via Reddit JSON
		}

		next := parser.FindNextPage(content)
		if next == "" {
			break
		}

		currentURL = resolveURL(currentURL, next)
		pageCount++
	}

	// 6️⃣ Final result
	return &ScrapeResult{
		URL:       url,
		Reviews:   allReviews,
		Source:    string(site),
		Timestamp: time.Now().Unix(),
	}, nil
}
