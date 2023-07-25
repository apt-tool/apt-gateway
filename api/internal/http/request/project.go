package request

import "github.com/automated-pen-testing/api/pkg/models/project"

type ProjectRequest struct {
	Name        string            `json:"name"`
	Description string            `json:"description"`
	Host        string            `json:"host"`
	Port        int               `json:"port"`
	HTTPSecure  bool              `json:"http_secure"`
	Endpoints   []string          `json:"endpoints"`
	Labels      map[string]string `json:"labels"`
	Params      map[string]string `json:"params"`
}

func (p ProjectRequest) ToModel(namespaceID uint, creator string) *project.Project {
	params := make([]*project.ParamSet, 0)
	labels := make([]*project.LabelSet, 0)

	for _, item := range p.Params {
		params = append(params, &project.ParamSet{
			Key:   item,
			Value: p.Params[item],
		})
	}

	for _, item := range p.Labels {
		labels = append(labels, &project.LabelSet{
			Key:   item,
			Value: p.Params[item],
		})
	}

	return &project.Project{
		Name:        p.Name,
		Host:        p.Host,
		Port:        p.Port,
		HTTPSecure:  p.HTTPSecure,
		Endpoints:   p.Endpoints,
		Creator:     creator,
		NamespaceID: namespaceID,
		Params:      params,
		Labels:      labels,
	}
}
