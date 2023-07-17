package enum

// Role of user controls the access of each user
type Role int

const (
	RoleAdmin Role = iota + 1
	RoleUser
	RoleViewer
)
