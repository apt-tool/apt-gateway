package worker

import (
	"fmt"
	"github.com/automated-pen-testing/api/pkg/client"
	"github.com/automated-pen-testing/api/pkg/models"
	"log"
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
		// todo: analyse
		// todo: use model
		// todo: get instructions
		// todo: execute instructions
		// todo: save into log file
		// todo: update database

		w.exit(id)
	}
}

func (w worker) exit(id int) {
	w.done <- id
}
