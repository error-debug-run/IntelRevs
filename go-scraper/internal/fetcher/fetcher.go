package fetcher

import "context"

// Fetcher defines the interface implemented by all site-specific scrapers.
//
// A Fetcher is responsible only for retrieving raw content
// from a target URL.
//
// It MUST:
// - respect the provided context
// - return raw, unprocessed payloads
// - avoid parsing or interpreting content
//
// It MUST NOT:
// - perform NLP or classification
// - normalize or clean data
// - modify payload structure
//
// Payload is intentionally typed as `any` to allow
// JSON, HTML, text, or binary data.
type Fetcher interface {
	Fetch(ctx context.Context, url string) (payload any, meta map[string]string, err error)
}

// registry maps site identifiers to their corresponding fetchers.
//
// This registry is intentionally explicit.
// No reflection, no dynamic loading, no magic.
var registry = map[string]Fetcher{}

// Register associates a site identifier with a Fetcher.
//
// This function should be called from init() functions
// inside site-specific fetcher packages.
//
// Panics if a fetcher is registered twice for the same site.
// This is a startup-time error and must be caught early.
func Register(site string, f Fetcher) {
	if _, exists := registry[site]; exists {
		panic("fetcher already registered for site: " + site)
	}
	registry[site] = f
}

// Get returns the Fetcher associated with the given site.
//
// If no fetcher is registered for the site,
// the generic fetcher is returned as a fallback.
//
// This guarantees that the scraper always has
// a valid fetcher to execute.
func Get(site string) Fetcher {
	if f, ok := registry[site]; ok {
		return f
	}
	return registry["generic"]
}
