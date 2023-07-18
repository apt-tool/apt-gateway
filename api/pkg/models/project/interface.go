package project

import "gorm.io/gorm"

type Interface interface {
	Create(project *Project) error
	Delete(projectID uint) error
	GetByID(projectID uint) (*Project, error)
}

type core struct {
	db *gorm.DB
}

func (c *core) Create(project *Project) error {
	return c.db.Create(project).Error
}

func (c *core) Delete(projectID uint) error {
	return c.db.Delete(&Project{}, "id = ?", projectID).Error
}

func (c *core) GetByID(projectID uint) (*Project, error) {
	prj := new(Project)

	query := c.db.
		First(&prj, "id = ?", projectID).
		Preload("Documents").
		Preload("Documents.Instructions")

	if err := query.Error; err != nil {
		return nil, err
	}

	if prj.ID != projectID {
		return nil, nil
	}

	return prj, nil
}
