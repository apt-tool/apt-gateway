package document

import "gorm.io/gorm"

// Interface manages the documents methods
type Interface interface {
	Create(document *Document) error
	Delete(projectID uint) error
}

func New(db *gorm.DB) Interface {
	return &core{
		db: db,
	}
}

type core struct {
	db *gorm.DB
}

func (c core) Create(document *Document) error {
	return c.db.Create(document).Error
}

func (c core) Delete(projectID uint) error {
	return c.db.Delete(&Document{}, "project_id = ?", projectID).Error
}
