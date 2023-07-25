package worker

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"

	"github.com/automated-pen-testing/api/internal/config/ftp"
	"github.com/automated-pen-testing/api/internal/core/ai"
	"github.com/automated-pen-testing/api/pkg/client"
	"github.com/automated-pen-testing/api/pkg/enum"
	"github.com/automated-pen-testing/api/pkg/models"
	"github.com/automated-pen-testing/api/pkg/models/document"
)

// worker is the smallest unit of our core
type worker struct {
	channel chan int
	done    chan int
	cfg     ftp.Config
	client  client.HTTPClient
	models  *models.Interface
	ai      *ai.AI
}

// executeRequest is used to call ftp system
type executeRequest struct {
	Param      string `json:"param"`
	Path       string `json:"path"`
	DocumentID uint   `json:"document_id"`
}

// work method will do the logic of penetration testing
func (w worker) work() {
	for {
		// get project id from channel
		id := <-w.channel

		projectID := uint(id)

		// remove all used documents
		if err := w.models.Documents.Delete(projectID); err != nil {
			log.Println(fmt.Errorf("[worker.work] failed to remove documents error=%w", err))

			w.exit(id)
		}

		// get project from db
		project, er := w.models.Projects.GetByID(projectID)
		if er != nil {
			log.Println(fmt.Errorf("[worker.work] failed to get project error=%w", er))

			w.exit(id)
		}

		// get attacks from ai module
		attacks := w.ai.GetAttacks(nil)

		docs := make([]*document.Document, 0)

		// create document
		for _, attack := range attacks {
			// create document
			doc := &document.Document{
				ProjectID:   projectID,
				Instruction: attack,
				Status:      enum.StatusInit,
			}

			if err := w.models.Documents.Create(doc); err != nil {
				log.Println(fmt.Errorf("[worker.work] failed to create document error=%w", err))

				continue
			}
		}

		// perform each attack
		for _, doc := range docs {
			// update doc status
			doc.Status = enum.StatusPending
			_ = w.models.Documents.Update(doc)

			// create ftp request
			tmp := executeRequest{
				Param:      project.Host,
				Path:       doc.Instruction,
				DocumentID: doc.ID,
			}

			// send ftp request
			var buffer bytes.Buffer
			if err := json.NewEncoder(&buffer).Encode(tmp); err != nil {
				log.Println(fmt.Errorf("[worker.work] failed to create request error=%w", err))

				continue
			}

			// update document based of response
			if response, httpError := w.client.Post(w.cfg.Host, &buffer, fmt.Sprintf("x-token:%s", w.cfg.Secret)); httpError != nil {
				log.Println(fmt.Errorf("[worker.work] failed to execute script error=%w", httpError))

				doc.Status = enum.StatusFailed
			} else {
				if response.StatusCode == 200 {
					doc.Status = enum.StatusDone
				} else {
					doc.Status = enum.StatusFailed
				}
			}

			_ = w.models.Documents.Update(doc)
		}

		w.exit(id)
	}
}

func (w worker) exit(id int) {
	w.done <- id
}
