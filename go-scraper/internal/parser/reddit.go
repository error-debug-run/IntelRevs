package parser

import "encoding/json"

type RedditParser struct{}

func NewRedditParser() Parser {
	return &RedditParser{}
}

func (p *RedditParser) Parse(html string) ([]string, error) {
	var data []interface{}
	if err := json.Unmarshal([]byte(html), &data); err != nil {
		return nil, err
	}

	comments := []string{}
	listing := data[1].(map[string]interface{})
	children := listing["data"].(map[string]interface{})["children"].([]interface{})

	for _, child := range children {
		c := child.(map[string]interface{})
		if c["kind"] == "t1" {
			body := c["data"].(map[string]interface{})["body"]
			if text, ok := body.(string); ok {
				comments = append(comments, text)
			}
		}
	}
	return comments, nil
}
