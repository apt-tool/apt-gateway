package worker

// worker is the smallest unit of our core
type worker struct {
	channel chan int
	done    chan int
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
