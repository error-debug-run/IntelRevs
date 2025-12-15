package manager

import "net/url"

func resolveURL(baseURL string, nextPath string) string {
	// If already absolute, return directly
	parsedNext, err := url.Parse(nextPath)
	if err == nil && parsedNext.IsAbs() {
		return nextPath
	}

	base, err := url.Parse(baseURL)
	if err != nil {
		return ""
	}

	resolved := base.ResolveReference(parsedNext)
	return resolved.String()
}
