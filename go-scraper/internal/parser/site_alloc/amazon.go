package site_alloc

import "github.com/error-debug-run/go-scraper/internal/parser"

// AmazonParser handles Amazon product review pages
type AmazonParser struct{}

func NewAmazonParser() parser.Parser {
	return &AmazonParser{}
}
func (p *AmazonParser) Parse(html string) ([]string, error) {

	//TODO:Implement real amazon reveiw parsing
	return parser.ExtractReviews(html)
}
