package parser

import (
	"errors"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type Parser interface {
	parse(html string) ([]string, error)
}

func ExtractReviews(html string) ([]string, error) {
	if html == "" {
		return nil, errors.New("empty HTML")
	}

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		return nil, err
	}

	var reviews []string

	// placeholder selectors for now
	doc.Find(".review, .customer-review, .a-size-base").Each(func(i int, s *goquery.Selection) {
		text := strings.TrimSpace(s.Text())
		if text != "" {
			reviews = append(reviews, text)
		}
	})

	return reviews, nil
}
