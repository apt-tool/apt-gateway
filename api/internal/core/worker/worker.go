package worker

import (
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

		// todo: remove history (if exists)
		// todo: analyse
		// todo: use model
		// todo: get instructions
		// todo: execute instructions
		// todo: save into log file
		// todo: update database

		w.done <- id
	}
}
