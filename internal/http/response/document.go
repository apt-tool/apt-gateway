package response

import (
	"time"

	"github.com/apt-tool/apt-gateway/pkg/enum"
	"github.com/apt-tool/apt-gateway/pkg/models/document"
)

type DocumentResponse struct {
	ID          uint        `json:"id"`
	Instruction string      `json:"instruction"`
	Status      enum.Status `json:"status"`
	CreatedAt   time.Time   `json:"created_at"`
}

func (d DocumentResponse) DTO(document *document.Document) *DocumentResponse {
	d.ID = document.ID
	d.Instruction = document.Instruction
	d.Status = document.Status
	d.CreatedAt = document.CreatedAt

	return &d
}
