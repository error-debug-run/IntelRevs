package worker

import (
	"bytes"

	"github.com/gocolly/colly/v2"
)

func FetchPage(url string) (string, error) {

	c := colly.NewCollector(
		colly.UserAgent("Mozilla/5.0 (compatible; ReviewBot/1.0)"),
	)

	var buffer bytes.Buffer

	c.OnResponse(func(r *colly.Response) {
		buffer.Write(r.Body)
	})

	if err := c.Visit(url); err != nil {
		return "", err
	}

	return buffer.String(), nil
}
