package response

import (
	"time"

	"github.com/automated-pen-testing/api/pkg/models/project"
)

type ProjectResponse struct {
	ID        uint                `json:"id"`
	Name      string              `json:"name"`
	Host      string              `json:"host"`
	CreatedAt time.Time           `json:"created_at"`
	Documents []*DocumentResponse `json:"documents"`
}

func (p ProjectResponse) DTO(project *project.Project) *ProjectResponse {
	p.ID = project.ID
	p.Name = project.Name
	p.Host = project.Host
	p.CreatedAt = project.CreatedAt

	list := make([]*DocumentResponse, 0)

	for _, item := range project.Documents {
		list = append(list, DocumentResponse{}.DTO(item))
	}

	p.Documents = list

	return &p
}
