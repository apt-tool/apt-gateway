package project

import (
	"github.com/automated-pen-testing/api/pkg/models/document"

	"gorm.io/gorm"
)

type Project struct {
	gorm.Model
	Name        string
	Host        string
	NamespaceID uint
	Documents   []*document.Document `gorm:"foreignKey:project_id"`
}
