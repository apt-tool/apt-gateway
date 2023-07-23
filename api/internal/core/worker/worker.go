package worker

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"

	"github.com/automated-pen-testing/api/internal/config/ftp"
	"github.com/automated-pen-testing/api/pkg/client"
	"github.com/automated-pen-testing/api/pkg/models"
	"github.com/automated-pen-testing/api/pkg/models/instruction"
)

// worker is the smallest unit of our core
type worker struct {
	channel chan int
	done    chan int
	cfg     ftp.Config
	client  client.HTTPClient
	models  *models.Interface
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

		// todo: choose instructions based on system analysis

		var attacks []*instruction.Instruction

		// perform each attack
		for _, attack := range attacks {
			// todo: create document and give id to tmp

			tmp := executeRequest{
				Param: project.Host,
				Path:  attack.Path,
			}

			var buffer bytes.Buffer
			if err := json.NewEncoder(&buffer).Encode(tmp); err != nil {
				log.Fatal(err)
			}

			_, httpError := w.client.Post(w.cfg.Host, &buffer, fmt.Sprintf("x-token:%s", w.cfg.Secret))
			if httpError != nil {
				log.Println(fmt.Errorf("[worker.work] failed to execute script error=%w", httpError))
			}

			// todo: update database
		}

		w.exit(id)
	}
}

func (w worker) exit(id int) {
	w.done <- id
}
