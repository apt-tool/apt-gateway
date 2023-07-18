package namespace

import "gorm.io/gorm"

type Interface interface {
	Create(namespace *Namespace) error
	Delete(namespaceID uint) error
	Get(namespaceIDs []uint, populate bool) ([]*Namespace, error)
}

type core struct {
	db *gorm.DB
}

func (c *core) Create(namespace *Namespace) error {
	return c.db.Create(namespace).Error
}

func (c *core) Delete(namespaceID uint) error {
	return c.db.Delete(&Namespace{}, "id = ?", namespaceID).Error
}

func (c *core) Get(namespaceIDs []uint, populate bool) ([]*Namespace, error) {
	list := make([]*Namespace, 0)

	query := c.db.Where("id in ?", namespaceIDs)
	if populate {
		query = query.Preload("Users").Preload("Projects")
	}

	if err := c.db.Where("id in ?", namespaceIDs).Find(&list).Error; err != nil {
		return nil, err
	}

	return list, nil
}
