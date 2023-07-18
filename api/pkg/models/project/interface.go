package project

type Interface interface {
	Create(project *Project) error
	Delete(projectID uint) error
	GetByID(projectID uint) error
}
