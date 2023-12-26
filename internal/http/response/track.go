package response

import (
	"github.com/ptaas-tool/base-api/pkg/models/track"

	"time"
)

type TrackResponse struct {
	ID          uint   `json:"id"`
	ProjectID   uint   `json:"project_id"`
	DocumentID  uint   `json:"document_id"`
	Description string `json:"description"`
	Type        string `json:"type"`
	Service     string `json:"service"`
	CreatedAt   string `json:"time"`
}

func (t TrackResponse) DTO(track *track.Track) *TrackResponse {
	t.ID = track.ID
	t.ProjectID = track.ProjectID
	t.DocumentID = track.DocumentID
	t.Description = track.Description
	t.Type = track.Type.ToString()
	t.Service = track.Service
	t.CreatedAt = track.CreatedAt.Format(time.DateTime)

	return &t
}
