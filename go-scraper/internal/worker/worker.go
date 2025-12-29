package worker

import (
	"bytes"
	"strings"

	"github.com/error-debug-run/go-scraper/internal/detector"
	"github.com/gocolly/colly/v2"
)

// FetchPage fetches raw content (HTML or JSON) as string
func FetchPage(url string) (string, error) {

	c := colly.NewCollector(
		colly.UserAgent("Mozilla/5.0 (compatible; ReviewBot/1.0)"),
	)

	var buffer bytes.Buffer

	c.OnResponse(func(r *colly.Response) {
		buffer.Write(r.Body)
	})

	// Reddit requires JSON
	if detector.DetectSite(url) == "reddit" {
		url = strings.TrimSuffix(url, "/") + ".json"
	}

	if err := c.Visit(url); err != nil {
		return "", err
	}

	return buffer.String(), nil
}
