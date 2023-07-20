package document

import (
	"github.com/automated-pen-testing/api/pkg/enum"
	"github.com/automated-pen-testing/api/pkg/models/instruction"

	"gorm.io/gorm"
)

type Document struct {
	gorm.Model
	ProjectID     uint
	InstructionID uint
	LogFile       string
	Status        enum.Status
	Instruction   *instruction.Instruction
}
