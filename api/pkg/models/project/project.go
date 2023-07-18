package project

import (
	"github.com/automated-pen-testing/api/pkg/models"
	"github.com/automated-pen-testing/api/pkg/models/document"
)

type Project struct {
	models.BaseModel
	Name        string
	Host        string
	NamespaceID uint
	Documents   []*document.Document `gorm:"foreignKey:project_id"`
}
