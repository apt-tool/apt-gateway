package response

import (
	"time"

	"github.com/automated-pen-testing/api/pkg/models/project"
)

type ProjectResponse struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	Host      string    `json:"host"`
	CreatedAt time.Time `json:"created_at"`
}

func (p ProjectResponse) DTO(project *project.Project) *ProjectResponse {
	p.ID = project.ID
	p.Name = project.Name
	p.Host = project.Host
	p.CreatedAt = project.CreatedAt

	return &p
}
