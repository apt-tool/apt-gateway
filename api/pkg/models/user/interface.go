package user

import (
	"github.com/automated-pen-testing/api/internal/utils/crypto"

	"gorm.io/gorm"
)

// Interface manages the user database methods
type Interface interface {
	Create(user *User) error
	Delete(userID uint) error
	Update(userID uint, user *User) error
	Get() ([]*User, error)
	GetByID(userID uint, populate bool) (*User, error)
	GetByName(name string, populate bool) (*User, error)
	Validate(name, pass string) (*User, error)
}

func New(db *gorm.DB) Interface {
	return &core{
		db: db,
	}
}

type core struct {
	db *gorm.DB
}

func (c core) Create(user *User) error {
	user.Password = crypto.GetMD5Hash(user.Password)

	return c.db.Create(user).Error
}

func (c core) Delete(userID uint) error {
	return c.db.Delete(&User{}, "id = ?", userID).Error
}

func (c core) Update(userID uint, user *User) error {
	return c.db.Model(&User{}).Update("role", user.Role).Where("id = ?", userID).Error
}

func (c core) Get() ([]*User, error) {
	list := make([]*User, 0)

	if err := c.db.Find(&list).Error; err != nil {
		return nil, ErrUserNotFound
	}

	return list, nil
}

func (c core) GetByID(userID uint, populate bool) (*User, error) {
	user := new(User)

	query := c.db
	if populate {
		query = query.Preload("Namespaces")
	}

	if err := query.First(&user).Where("id = ?", userID).Error; err != nil {
		return nil, ErrUserNotFound
	}

	return user, nil
}

func (c core) GetByName(name string, populate bool) (*User, error) {
	user := new(User)

	query := c.db
	if populate {
		query = query.Preload("Namespaces")
	}

	if err := query.First(&user).Where("username = ?", name).Error; err != nil {
		return nil, ErrUserNotFound
	}

	return user, nil
}

func (c core) Validate(name, pass string) (*User, error) {
	user := new(User)

	if err := c.db.First(&user).Where("username = ?", name).Error; err != nil {
		return nil, ErrUserNotFound
	}

	if user.Password != crypto.GetMD5Hash(pass) {
		return nil, ErrIncorrectPassword
	}

	return user, nil
}
