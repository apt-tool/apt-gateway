package namespace

import (
	"fmt"

	"github.com/automated-pen-testing/api/pkg/models/user"

	"gorm.io/gorm"
)

// Interface manages the namespace db methods
type Interface interface {
	Create(namespace *Namespace) error
	Delete(namespaceID uint) error
	Get() ([]*Namespace, error)
	GetByID(namespaceID uint, users bool) (*Namespace, error)
	GetUserNamespaces(userID uint) ([]*Namespace, error)
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

func (c core) Get() ([]*Namespace, error) {
	list := make([]*Namespace, 0)

	if err := c.db.Find(&list).Error; err != nil {
		return nil, fmt.Errorf("[db.Namespace.Get] failed to get records error=%w", err)
	}

	return list, nil
}

func (c core) GetByID(namespaceID uint, users bool) (*Namespace, error) {
	namespace := new(Namespace)

	query := c.db

	if users {
		query = query.Preload("Users")
	} else {
		query = query.Preload("Projects")
	}

	if err := query.Where("id = ?", namespaceID).First(&namespace).Error; err != nil {
		return nil, fmt.Errorf("[db.Namespace.GetByID] failed to get record error=%w", err)
	}

	if namespace.ID != namespaceID {
		return nil, ErrRecordNotFound
	}

	return namespace, nil
}

func (c core) GetUserNamespaces(userID uint) ([]*Namespace, error) {
	list := make([]*Namespace, 0)

	query := c.db.Preload("Users").Where("namespace_users.user_id = ?", userID)

	if err := query.Find(&list).Error; err != nil {
		return nil, fmt.Errorf("[db.Namespace.Get] failed to get records error=%w", err)
	}

	return list, nil
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
