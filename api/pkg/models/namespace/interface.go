package namespace

import (
	"errors"

	"github.com/automated-pen-testing/api/pkg/models/user"

	"gorm.io/gorm"
)

// Interface manages the namespace db methods
type Interface interface {
	Create(namespace *Namespace) error
	Delete(namespaceID uint) error
	Get(populate bool) ([]*Namespace, error)
	GetByID(namespaceID uint) (*Namespace, error)
	AddUser(namespaceID uint, user *user.User) error
	RemoveUser(namespaceID uint, user *user.User) error
}

func New(db *gorm.DB) Interface {
	return &core{
		db: db,
	}
}

type core struct {
	db *gorm.DB
}

func (c core) Create(namespace *Namespace) error {
	return c.db.Create(namespace).Error
}

func (c core) Delete(namespaceID uint) error {
	return c.db.Delete(&Namespace{}, "id = ?", namespaceID).Error
}

func (c core) Get(populate bool) ([]*Namespace, error) {
	list := make([]*Namespace, 0)

	query := c.db

	if populate {
		query = query.Preload("Users").Preload("Projects")
	}

	if err := c.db.Find(&list).Error; err != nil {
		return nil, err
	}

	return list, nil
}

func (c core) GetByID(namespaceID uint) (*Namespace, error) {
	namespace := new(Namespace)

	if err := c.db.Preload("Projects").Where("id = ?", namespaceID).First(&namespace).Error; err != nil {
		return nil, err
	}

	if namespace.ID != namespaceID {
		return nil, errors.New("namespace not found")
	}

	return namespace, nil
}

func (c core) AddUser(namespaceID uint, user *user.User) error {
	return c.db.Model(&Namespace{}).
		Where("id = ?", namespaceID).
		Association("Users").
		Append(user)
}

func (c core) RemoveUser(namespaceID uint, user *user.User) error {
	return c.db.Model(&Namespace{}).
		Where("id = ?", namespaceID).
		Association("Users").
		Delete(user)
}
