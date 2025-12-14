package detector

import "strings"

type SiteType string

const (
	SiteAmazon   SiteType = "amazon"
	SiteFlipkart SiteType = "flipkart"
	SiteReddit   SiteType = "reddit"
	SiteGeneric  SiteType = "generic"
)

func DetectSite(url string) SiteType {
	u := strings.ToLower(url)

	switch {
	case strings.Contains(u, "amazon."):
		return SiteAmazon
	case strings.Contains(u, "flipkart."):
		return SiteFlipkart
	case strings.Contains(u, "reddit.com"):
		return SiteReddit
	default:
		return SiteGeneric
	}
}
