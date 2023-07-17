package instruction

import "github.com/automated-pen-testing/api/pkg/models"

type Instruction struct {
	models.BaseModel
	Name string
	Path string
}
