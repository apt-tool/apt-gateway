package namespace

import (
	"github.com/automated-pen-testing/api/pkg/models"
	"github.com/automated-pen-testing/api/pkg/models/project"
	"github.com/automated-pen-testing/api/pkg/models/user"
)

type (
	// Namespace manage projects admin can create namespaces
	Namespace struct {
		models.BaseModel
		Name     string
		Users    []*user.User       `gorm:"many2many:namespace_users;"`
		Projects []*project.Project `gorm:"foreignKey:namespace_id"`
	}
)
