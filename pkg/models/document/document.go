package document

import (
	"github.com/apt-tool/apt-gateway/pkg/enum"

	"gorm.io/gorm"
)

// Document represents core log files
type Document struct {
	gorm.Model
	ProjectID   uint
	LogFile     string
	Instruction string
	Status      enum.Status
}
