package user

import (
	"github.com/automated-pen-testing/api/pkg/enum"
	"github.com/automated-pen-testing/api/pkg/models"
)

// User is the base entity of our clients
type User struct {
	models.BaseModel
	Username string
	Password string
	Role     enum.Role
}
