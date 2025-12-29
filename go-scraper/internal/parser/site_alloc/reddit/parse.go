package reddit

import "encoding/json"

func ParseJSON(jsonText string) ([]string, error) {
	var listings []Listing

	if err := json.Unmarshal([]byte(jsonText), &listings); err != nil {
		return nil, err
	}

	if len(listings) < 2 {
		return nil, nil
	}

	var results []string
	ExtractComments(listings[1].Data.Children, &results)

	return results, nil
}
