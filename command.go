package transmission

import "encoding/json"

type Command struct {
	Method    string      `json:"method"`
	Arguments interface{} `json:"arguments"`
	Tag       int         `json:"tag"`
}

type Response struct {
	Result    string          `json:"result"`
	Arguments json.RawMessage `json:"arguments"`
	Tag       int             `json:"tag"`
}

type WhichTorrents struct {
	IDs    []int
	SHAs   []string
	Recent bool
}

func mergeIDsAndSHAs(which WhichTorrents) []interface{} {

	if len(which.IDs) == 0 && len(which.SHAs) == 0 {
		return []interface{}{}
	}

	partOne, partTwo := len(which.IDs), len(which.SHAs)
	ids := make([]interface{}, partOne+partTwo)
	for i, v := range which.IDs {
		ids[i] = v
	}
	for i, v := range which.SHAs {
		ids[partOne+i] = v
	}
	return ids
}
