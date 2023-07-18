package models

import (
	"github.com/automated-pen-testing/api/pkg/models/document"
	"github.com/automated-pen-testing/api/pkg/models/instruction"
	"github.com/automated-pen-testing/api/pkg/models/namespace"
	"github.com/automated-pen-testing/api/pkg/models/project"
	"github.com/automated-pen-testing/api/pkg/models/user"

	"gorm.io/gorm"
)

// Interface manages the models interfaces
type Interface struct {
	Documents    document.Interface
	Instructions instruction.Interface
	Namespaces   namespace.Interface
	Projects     project.Interface
	Users        user.Interface
}

func New(db *gorm.DB) *Interface {
	return &Interface{
		Documents:    document.New(db),
		Instructions: instruction.New(db),
		Namespaces:   namespace.New(db),
		Projects:     project.New(db),
		Users:        user.New(db),
	}
}
