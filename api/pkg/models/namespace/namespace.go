package namespace

import (
	"github.com/automated-pen-testing/api/pkg/models"
	"github.com/automated-pen-testing/api/pkg/models/user"
)

type (
	// Namespace manage projects admin can create namespaces
	Namespace struct {
		models.BaseModel
		Name  string
		Users user.User
	}

	// NamespaceUsers stores users of namespace
	NamespaceUsers struct {
		UserID      uint //todo: ref key to user table
		NamespaceID uint //todo: ref key to namespace
	}
)
