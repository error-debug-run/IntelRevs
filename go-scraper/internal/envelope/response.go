package envelope

import "time"

// Response represents the canonical response envelope
// returned by the Go scraper service.
//
// This structure is intentionally minimal and stable.
// Python code MUST be able to rely on these fields always existing.
//
// Fields:
// - Meta:     minimal metadata for routing and debugging
// - Payload: raw, unprocessed data fetched from the target site
// - Error:   structured error information (null on success)
//
// IMPORTANT:
// - Payload is opaque to Go and Python adapters
// - Go MUST NOT interpret or transform payload content
type Response struct {
	Meta    map[string]string `json:"meta"`
	Payload any               `json:"payload"`
	Error   *ErrorInfo        `json:"error"`
}

// ErrorInfo represents a non-fatal scraper error.
//
// Errors are returned as data, not raised as panics,
// to ensure downstream systems never crash.
//
// An error here indicates that scraping failed or was incomplete,
// not that the API itself is broken.
type ErrorInfo struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

// Success constructs a successful scraper response envelope.
//
// This function MUST be used for all successful fetches.
//
// Arguments:
// - source:     detected site (reddit, amazon, generic, etc.)
// - inputURL:  original user-provided URL
// - payload:   raw fetched data (JSON, HTML, text)
// - meta:      optional fetcher-specific metadata
//
// Guarantees:
// - Error is always nil
// - Meta always includes source and input_url
func Success(
	source string,
	inputURL string,
	payload any,
	meta map[string]string,
) Response {

	if meta == nil {
		meta = make(map[string]string)
	}

	meta["source"] = source
	meta["input_url"] = inputURL
	meta["fetched_at"] = time.Now().UTC().Format(time.RFC3339)

	return Response{
		Meta:    meta,
		Payload: payload,
		Error:   nil,
	}
}

// Error constructs a generic error response envelope.
//
// This is used for API-level errors such as:
// - invalid HTTP method
// - missing query parameters
// - malformed requests
//
// This function does NOT represent site-specific scraping failures.
func Error(code string, message string) Response {
	return Response{
		Meta:    map[string]string{},
		Payload: nil,
		Error: &ErrorInfo{
			Code:    code,
			Message: message,
		},
	}
}

// FromError constructs an error response envelope
// from a site fetcher failure.
//
// This is used when the scraper attempted to fetch data
// from a remote site but failed (403, timeout, network error, etc.).
//
// The error is serialized instead of being raised,
// allowing Python to continue processing with reduced confidence.
func FromError(source string, inputURL string, err error) Response {
	return Response{
		Meta: map[string]string{
			"source":    source,
			"input_url": inputURL,
		},
		Payload: nil,
		Error: &ErrorInfo{
			Code:    "FETCH_FAILED",
			Message: err.Error(),
		},
	}
}
