package worker

import (
	"fmt"
	"log"
	"os/exec"

	"github.com/automated-pen-testing/api/pkg/client"
	"github.com/automated-pen-testing/api/pkg/models"
)

// worker is the smallest unit of our core
type worker struct {
	channel chan int
	done    chan int
	client  client.HTTPClient
	models  *models.Interface
}

// work method will do the logic of penetration testing
func (w worker) work() {
	for {
		id := <-w.channel

		if err := w.models.Documents.Delete(uint(id)); err != nil {
			log.Println(fmt.Errorf("[worker.work] failed to remove documents error=%w", err))

			w.exit(id)
		}

		project, err := w.models.Projects.GetByID(uint(id))
		if err != nil {
			log.Println(fmt.Errorf("[worker.work] failed to get project error=%w", err))

			w.exit(id)
		}

		cmd, er := exec.Command("nmap", "-sV", "--script", "nmap-vulners/", project.Host).Output()
		if er != nil {
			log.Println(fmt.Errorf("[worker.work] failed to analyse project error=%w", err))

			w.exit(id)
		}

		_ = string(cmd)

		// todo: use model

		var ids []uint

		for _, instructionID := range ids {
			_, err := w.models.Instructions.GetByID(instructionID)
			if err != nil {
				log.Println(fmt.Errorf("[worker.work] failed to get instruction error=%w", err))

				continue
			}

			// todo: execute instructions
			// todo: save into log file
			// todo: update database
		}

		w.exit(id)
	}
}

func (w worker) exit(id int) {
	w.done <- id
}
