package main

import (
	"net/http"

	"github.com/error-debug-run/go-scraper/internal/manager"
	"github.com/labstack/echo/v4"
)

type ScrapeRequest struct {
	URL string `query:"url"`
}

func main() {
	e := echo.New()

	e.GET("/v1/scraper", func(c echo.Context) error {
		req := new(ScrapeRequest)
		if err := c.Bind(req); err != nil || req.URL == "" {
			return c.JSON(http.StatusBadRequest, map[string]string{
				"error": "Missing ?url parameter",
			})
		}

		result, err := manager.RunScrapeJob(req.URL)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{
				"error": err.Error(),
			})
		}

		return c.JSON(http.StatusOK, result)
	})

	e.Logger.Fatal(e.Start(":8080"))
}
