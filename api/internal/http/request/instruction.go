package request

import "github.com/automated-pen-testing/api/pkg/models/instruction"

type InstructionRequest struct {
	Name string `json:"name"`
	Path string `json:"path"`
}

func (i InstructionRequest) ToModel() *instruction.Instruction {
	return &instruction.Instruction{
		Name: i.Name,
		Path: i.Path,
	}
}
