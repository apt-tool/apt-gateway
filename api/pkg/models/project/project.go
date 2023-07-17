package project

import "github.com/automated-pen-testing/api/pkg/models"

type Project struct {
	models.BaseModel
	Name        string
	Host        string
	NamespaceID uint
}
