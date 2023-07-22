package request

import "github.com/automated-pen-testing/api/pkg/models/project"

type ProjectRequest struct {
	Name        string `json:"name"`
	Host        string `json:"host"`
	NamespaceID uint   `json:"namespace_id"`
}

func (p ProjectRequest) ToModel() *project.Project {
	return &project.Project{
		Name:        p.Name,
		Host:        p.Host,
		NamespaceID: p.NamespaceID,
	}
}
