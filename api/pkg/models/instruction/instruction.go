package instruction

import "gorm.io/gorm"

type Instruction struct {
	gorm.Model
	Name string
	Path string
}
