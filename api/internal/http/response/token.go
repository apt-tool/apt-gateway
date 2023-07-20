package response

import "time"

type JToken struct {
	Token   string    `json:"token"`
	Expires time.Time `json:"expires"`
}
