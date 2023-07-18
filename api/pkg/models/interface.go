package models

import (
	"github.com/automated-pen-testing/api/pkg/models/namespace"
	"github.com/automated-pen-testing/api/pkg/models/user"

	"gorm.io/gorm"
)

// Interface manages the models interfaces
type Interface struct {
	Namespaces namespace.Interface
	Users      user.Interface
}

func New(db *gorm.DB) *Interface {
	return &Interface{
		Namespaces: namespace.New(db),
		Users:      user.New(db),
	}
}
