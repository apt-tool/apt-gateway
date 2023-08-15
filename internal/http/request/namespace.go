package request

import "github.com/apt-tool/apt-core/pkg/models/namespace"

type NamespaceRequest struct {
	Name string `json:"name"`
}

type NamespaceUpdateRequest struct {
	UserIDs     []uint `json:"user_ids"`
	NamespaceID uint   `json:"namespace_id"`
}

type NamespaceQueryRequest struct {
	Populate bool `query:"populate"`
}

func (n NamespaceRequest) ToModel() *namespace.Namespace {
	return &namespace.Namespace{
		Name: n.Name,
	}
}
