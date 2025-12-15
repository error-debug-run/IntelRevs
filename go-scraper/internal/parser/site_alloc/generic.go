package site_alloc

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/error-debug-run/go-scraper/internal/parser"
)

// GenericParser is a fallback parser for unknown websites.
// It implements the Parser interface.
type GenericParser struct{}

// NewGenericParser returns a Parser implementation
// that performs basic, site-agnostic parsing.
func NewGenericParser() parser.Parser {
	return &GenericParser{}
}

// Parse extracts review-like text blocks from raw HTML.
// For now, this is a stub and will be expanded later.
func (p *GenericParser) Parse(html string) ([]string, error) {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		return nil, err
	}

	var reviews []string

	doc.Find("div.quote span.text").Each(func(i int, s *goquery.Selection) {
		text := strings.TrimSpace(s.Text())
		if text != "" {
			reviews = append(reviews, text)
		}
	})

	return reviews, nil
}
