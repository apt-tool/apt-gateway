package response

import (
	"time"

	"github.com/automated-pen-testing/api/pkg/models/namespace"
)

type NamespaceResponse struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
}

func (n NamespaceResponse) DTO(namespace *namespace.Namespace) *NamespaceResponse {
	n.ID = namespace.ID
	n.Name = namespace.Name
	n.CreatedAt = namespace.CreatedAt

	return &n
}
