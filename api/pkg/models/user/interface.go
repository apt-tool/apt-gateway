package user

import "gorm.io/gorm"

type Interface interface {
	Create(user *User) error
	Delete(userID uint) error
	Update(userID uint, user *User) error
	Get() ([]*User, error)
	Validate(name, pass string) (*User, error)
}

type core struct {
	db *gorm.DB
}

func (c *core) Create(user *User) error {
	return c.db.Create(user).Error
}

func (c *core) Delete(userID uint) error {
	return c.db.Delete(&User{}, "id = ?", userID).Error
}

func (c *core) Update(userID uint, user *User) error {
	return c.db.Model(&User{}).Update("role", user.Role).Where("id = ?", userID).Error
}

func (c *core) Get() ([]*User, error) {
	list := make([]*User, 0)

	if err := c.db.Find(&list).Error; err != nil {
		return nil, err
	}

	return list, nil
}

func (c *core) Validate(name, pass string) (*User, error) {
	user := new(User)

	if err := c.db.First(&user).Where("username = ? and password = ?", name, pass).Error; err != nil {
		return nil, err
	}

	return user, nil
}
