package response

import (
	"time"

	"github.com/apt-tool/apt-core/pkg/enum"
	"github.com/apt-tool/apt-core/pkg/models/document"
)

type DocumentResponse struct {
	ID            uint          `json:"id"`
	Instruction   string        `json:"instruction"`
	ExecutedBy    string        `json:"executed_by"`
	Result        enum.Result   `json:"result"`
	Status        enum.Status   `json:"status"`
	CreatedAt     time.Time     `json:"created_at"`
	ExecutionTime time.Duration `json:"execution_time"`
}

func (d DocumentResponse) DTO(document *document.Document) *DocumentResponse {
	d.ID = document.ID
	d.Instruction = document.Instruction
	d.ExecutedBy = document.ExecutedBy
	d.Result = document.Result
	d.Status = document.Status
	d.CreatedAt = document.CreatedAt
	d.ExecutionTime = document.ExecutionTime

	return &d
}
