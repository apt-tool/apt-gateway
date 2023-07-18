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
