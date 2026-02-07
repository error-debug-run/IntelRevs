package detector

import (
	"net/url"
	"strings"
)

const (
	SiteReddit   = "reddit"
	SiteAmazon   = "amazon"
	SiteFlipkart = "flipkart"
	SiteGeneric  = "generic"
)

// Detect determines the source site for a given URL.
//
// The detection logic is intentionally simple and conservative.
// It relies only on hostname inspection and string matching.
//
// Guarantees:
// - No network calls
// - Deterministic output
// - Always returns a value (fallback to generic)
//
// NOTE:
// Incorrect detection does not break the pipeline;
// adapters will handle mismatches downstream.
func Detect(rawURL string) string {

	parsed, err := url.Parse(rawURL)
	if err != nil {
		return SiteGeneric
	}

	host := strings.ToLower(parsed.Host)

	switch {
	case strings.Contains(host, "reddit.com"):
		return SiteReddit

	case strings.Contains(host, "amazon."):
		return SiteAmazon

	case strings.Contains(host, "flipkart."):
		return SiteFlipkart

	default:
		return SiteGeneric
	}
}
