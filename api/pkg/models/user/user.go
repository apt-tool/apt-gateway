package user

import (
	"github.com/automated-pen-testing/api/pkg/enum"
	"github.com/automated-pen-testing/api/pkg/models/namespace"

	"gorm.io/gorm"
)

// User is the base entity of our clients
type User struct {
	gorm.Model
	Username   string `gorm:"unique"`
	Password   string
	Role       enum.Role
	Namespaces []*namespace.Namespace `gorm:"many2many:namespace_users;"`
}
