package response

import (
	"fmt"
	"time"

	"github.com/apt-tool/apt-core/pkg/models/project"
)

type (
	SetResponse struct {
		Key   string `json:"key"`
		Value string `json:"value"`
	}

	ProjectResponse struct {
		ID          uint                `json:"id"`
		Name        string              `json:"name"`
		Description string              `json:"description"`
		Host        string              `json:"host"`
		Creator     string              `json:"creator"`
		CreatedAt   time.Time           `json:"created_at"`
		Endpoints   []string            `json:"endpoints"`
		Labels      []*SetResponse      `json:"labels"`
		Params      []*SetResponse      `json:"params"`
		Documents   []*DocumentResponse `json:"documents"`
	}
)

func (p ProjectResponse) DTO(project *project.Project) *ProjectResponse {
	p.ID = project.ID
	p.Name = project.Name
	p.Description = project.Description
	p.Creator = project.Creator
	p.CreatedAt = project.CreatedAt

	p.Host = p.createHost(project.Host, project.Port, project.HTTPSecure)

	list := make([]*SetResponse, 0)

	for _, item := range project.Labels {
		list = append(list, &SetResponse{
			Key:   item.Key,
			Value: item.Value,
		})
	}

	list2 := make([]*DocumentResponse, 0)

	for _, item := range project.Documents {
		list2 = append(list2, DocumentResponse{}.DTO(item))
	}

	list3 := make([]*SetResponse, 0)

	for _, item := range project.Params {
		list3 = append(list3, &SetResponse{
			Key:   item.Key,
			Value: item.Value,
		})
	}

	list4 := make([]string, 0)

	for _, item := range project.Endpoints {
		list4 = append(list4, item.Endpoint)
	}

	p.Labels = list
	p.Documents = list2
	p.Params = list3
	p.Endpoints = list4

	return &p
}

func (p ProjectResponse) createHost(host string, port int, secure bool) string {
	tmp := "http"
	if secure {
		tmp = "https"
	}

	return fmt.Sprintf("%s://%s:%d", tmp, host, port)
}
