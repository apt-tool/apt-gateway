package document

import (
	"github.com/automated-pen-testing/api/pkg/enum"
	"github.com/automated-pen-testing/api/pkg/models"
)

type (
	Document struct {
		models.BaseModel
		ProjectID   uint
		LogFile     string
		Status      enum.Status
		Instruction []*DocumentInstructions
	}

	DocumentInstructions struct {
		models.BaseModel
		Instruction enum.Instruction
		DocumentID  uint
	}
)
