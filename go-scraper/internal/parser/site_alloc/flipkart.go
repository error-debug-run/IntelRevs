package site_alloc

import "github.com/error-debug-run/go-scraper/internal/parser"

type FlipkartParser struct{}

func NewFlipkartParser() parser.Parser {
	return &FlipkartParser{}
}

func (p *FlipkartParser) Parse(html string) ([]string, error) {

	//TODO:Implement real flipkart reveiw parsing
	return parser.ExtractReviews(html)
}
