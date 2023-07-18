package models

import (
	"github.com/automated-pen-testing/api/pkg/models/user"

	"gorm.io/gorm"
)

// Interface manages the models interfaces
type Interface struct {
	Users user.Interface
}

func New(db *gorm.DB) *Interface {
	return &Interface{
		Users: user.New(db),
	}
}
