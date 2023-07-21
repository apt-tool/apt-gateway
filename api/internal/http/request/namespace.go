package request

import "github.com/automated-pen-testing/api/pkg/models/namespace"

type NamespaceRequest struct {
	Name string `json:"name"`
}

type NamespaceUpdateRequest struct {
	UserID      uint `json:"user_id"`
	NamespaceID uint `json:"namespace_id"`
	Add         bool `json:"add"`
}

type NamespaceQueryRequest struct {
	Populate bool `query:"populate"`
}

func (n NamespaceRequest) ToModel() *namespace.Namespace {
	return &namespace.Namespace{
		Name: n.Name,
	}
}
