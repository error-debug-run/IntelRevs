package reddit

import (
	"github.com/error-debug-run/go-scraper/internal/parser"
)

// RedditParser implements parser.Parser
type RedditParser struct{}

func NewRedditParser() parser.Parser {
	return &RedditParser{}
}

func (p *RedditParser) Parse(url string) ([]string, error) {

	raw, err := FetchRedditJSON(url)
	if err != nil {
		return nil, err
	}

	filePath, err := saveRawJSON(raw)
	if err != nil {
		return nil, err
	}

	// Contract: return file path
	return []string{filePath}, nil
}
