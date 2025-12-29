package reddit

import "encoding/json"

func ExtractComments(children []Child, results *[]string) {
	for _, child := range children {
		if child.Kind != "t1" {
			continue
		}

		if len(child.Data.Body) > 40 {
			*results = append(*results, child.Data.Body)
		}

		if len(child.Data.Replies) > 0 && string(child.Data.Replies) != `""` {
			var replyListing Listing
			if err := json.Unmarshal(child.Data.Replies, &replyListing); err == nil {
				ExtractComments(replyListing.Data.Children, results)
			}
		}
	}
}
