package instruction

import "gorm.io/gorm"

// Instruction represents that core attacks
type Instruction struct {
	gorm.Model
	Name string
	Path string
}
