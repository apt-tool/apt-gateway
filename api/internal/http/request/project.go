package request

import "github.com/automated-pen-testing/api/pkg/models/project"

type ProjectRequest struct {
	Name string `json:"name"`
	Host string `json:"host"`
}

func (p ProjectRequest) ToModel(namespaceID uint) *project.Project {
	return &project.Project{
		Name:        p.Name,
		Host:        p.Host,
		NamespaceID: namespaceID,
	}
}
