package enum

type Status int

const (
	StatusInit Status = iota + 1
	StatusPending
	StatusDone
	StatusFailed
)
