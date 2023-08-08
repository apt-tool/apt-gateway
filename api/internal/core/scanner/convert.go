package scanner

import "encoding/json"

// convert scanner report to our system report
func convert(context []byte, r *report) error {
	// convert our json object into our report
	if er := json.Unmarshal(context, r); er != nil {
		return er
	}

	return nil
}
