package response

import (
	"fmt"
	"time"

	"github.com/automated-pen-testing/api/pkg/models/project"
)

type ProjectResponse struct {
	ID          uint                `json:"id"`
	Name        string              `json:"name"`
	Description string              `json:"description"`
	Host        string              `json:"host"`
	Creator     string              `json:"creator"`
	Labels      map[string]string   `json:"labels"`
	CreatedAt   time.Time           `json:"created_at"`
	Documents   []*DocumentResponse `json:"documents"`
}

func (p ProjectResponse) DTO(project *project.Project) *ProjectResponse {
	p.ID = project.ID
	p.Name = project.Name
	p.Description = project.Description
	p.Creator = project.Creator
	p.CreatedAt = project.CreatedAt

	p.Host = p.createHost(project.Host, project.Port, project.HTTPSecure)

	p.Labels = make(map[string]string)

	for _, item := range project.Labels {
		p.Labels[item.Key] = item.Value
	}

	list := make([]*DocumentResponse, 0)

	for _, item := range project.Documents {
		list = append(list, DocumentResponse{}.DTO(item))
	}

	p.Documents = list

	return &p
}

func (p ProjectResponse) createHost(host string, port int, secure bool) string {
	tmp := "http"
	if secure {
		tmp = "https"
	}

	return fmt.Sprintf("%s://%s:%d", tmp, host, port)
}
