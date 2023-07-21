package request

type NamespaceRequest struct {
	Name string `json:"name"`
}

type NamespaceUserRequest struct {
	UserID      uint `json:"user_id"`
	NamespaceID uint `json:"namespace_id"`
	Add         bool `json:"add"`
}
