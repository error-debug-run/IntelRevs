package reddit

import "encoding/json"

type Listing struct {
	Data struct {
		Children []Child `json:"children"`
	} `json:"data"`
}

type Child struct {
	Kind string `json:"kind"`
	Data Data   `json:"data"`
}

type Data struct {
	Body    string          `json:"body"`
	Replies json.RawMessage `json:"replies"`
}
