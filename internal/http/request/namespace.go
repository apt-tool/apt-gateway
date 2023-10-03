package request

import "github.com/ptaas-tool/base-api/pkg/models/namespace"

type (
	NamespaceRequest struct {
		Name string `json:"name"`
	}

	NamespaceUpdateRequest struct {
		Name    string `json:"name"`
		UserIDs []uint `json:"user_ids"`
	}

	NamespaceQueryRequest struct {
		Populate bool `query:"populate"`
	}
)

func (n NamespaceRequest) ToModel(creator string) *namespace.Namespace {
	return &namespace.Namespace{
		Name:      n.Name,
		CreatedBy: creator,
	}
}
