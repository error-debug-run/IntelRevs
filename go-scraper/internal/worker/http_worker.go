package worker

import (
	"io"
	"net/http"
	"time"
)

type HTTPWorker struct {
	client *http.Client
}

func NewHTTPWorker() *HTTPWorker {
	return &HTTPWorker{
		client: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

func (w *HTTPWorker) FetchHTML(url string) (string, error) {
	resp, err := w.client.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}
