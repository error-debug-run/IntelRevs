package site_alloc

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/error-debug-run/go-scraper/internal/parser"
)

// RedditParser handles Reddit comment threads
type RedditParser struct{}

func NewRedditParser() parser.Parser {
	return &RedditParser{}
}

func (p *RedditParser) Parse(html string) ([]string, error) {

	println("HTML snippet:")
	println(html[:(len(html) - 1)]) // first 1000 chars
	//return []string{}, nil

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		return nil, err
	}

	var comments []string

	doc.Find("div[data-test-id=\"comment\"]").Each(func(i int, s *goquery.Selection) {

		text := strings.TrimSpace(s.Text())

		if len(text) < 50 {
			return
		}

		text = NormalizeWhiteText(text)

		comments = append(comments, text)

	})

	return comments, nil

}

func NormalizeWhiteText(s string) string {

	lines := strings.Fields(s)
	return strings.Join(lines, "")

}
