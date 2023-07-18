package namespace

import "gorm.io/gorm"

type Interface interface {
	Create(namespace *Namespace) error
	Delete(namespaceID uint) error
	AddUser(userID, namespaceID uint) error
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
