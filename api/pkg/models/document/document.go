package document

import (
	"github.com/automated-pen-testing/api/pkg/enum"
	"github.com/automated-pen-testing/api/pkg/models"
	"github.com/automated-pen-testing/api/pkg/models/instruction"
)

type (
	Document struct {
		models.BaseModel
		ProjectID    uint
		LogFile      string
		Status       enum.Status
		Instructions []*instruction.Instruction `gorm:"many2many:document_instructions;"`
	}
)
