package response

import (
	"time"

	"github.com/ptaas-tool/base-api/pkg/models/namespace"
)

type NamespaceResponse struct {
	ID        uint               `json:"id"`
	Name      string             `json:"name"`
	CreatedBy string             `json:"created_by"`
	CreatedAt time.Time          `json:"created_at"`
	Users     []*UserResponse    `json:"users"`
	Projects  []*ProjectResponse `json:"projects"`
}

func (n NamespaceResponse) DTO(namespace *namespace.Namespace) *NamespaceResponse {
	n.ID = namespace.ID
	n.Name = namespace.Name
	n.CreatedAt = namespace.CreatedAt
	n.CreatedBy = namespace.CreatedBy

	list := make([]*UserResponse, 0)

	for _, item := range namespace.Users {
		list = append(list, UserResponse{}.DTO(item))
	}

	n.Users = list

	list2 := make([]*ProjectResponse, 0)

	for _, item := range namespace.Projects {
		list2 = append(list2, ProjectResponse{}.DTO(item))
	}

	n.Projects = list2

	return &n
}
