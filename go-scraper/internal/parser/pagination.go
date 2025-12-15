package parser

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func FindNextPage(html string) string {

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		return ""
	}

	next, exists := doc.Find("li.next a").Attr("href")
	if !exists {
		return next
	}

	return ""

}
