package response

import (
	"time"

	"github.com/automated-pen-testing/api/pkg/models/instruction"
)

type InstructionResponse struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	Path      string    `json:"path"`
	CreatedAt time.Time `json:"created_at"`
}

func (i InstructionResponse) DTO(instruction *instruction.Instruction) *InstructionResponse {
	i.ID = instruction.ID
	i.Name = instruction.Name
	i.Path = instruction.Path
	i.CreatedAt = instruction.CreatedAt

	return &i
}
