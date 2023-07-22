package worker

import (
	"log"

	"github.com/automated-pen-testing/api/pkg/client"
	"github.com/automated-pen-testing/api/pkg/models"
)

type Pool struct {
	client *client.Client
	models *models.Interface

	capacity int
	inuse    int
	channel  chan int
	done     chan int
}

func New(client *client.Client, models *models.Interface, capacity int) *Pool {
	return &Pool{
		client:   client,
		models:   models,
		capacity: capacity,
		inuse:    0,
		channel:  make(chan int),
		done:     make(chan int),
	}
}

func (p *Pool) update() {
	for {
		id := <-p.done

		p.inuse--

		log.Printf("[worker.update] finished process for id=%d\n", id)
	}
}

func (p *Pool) Register() {
	for i := 0; i < p.capacity; i++ {
		go worker{
			client:  p.client,
			models:  p.models,
			channel: p.channel,
			done:    p.done,
		}.work()
	}

	log.Printf("[worker.Register] started %d workers\n", p.capacity)
}

func (p *Pool) Do(id int) bool {
	if p.inuse == p.capacity {
		return false
	}

	p.inuse++

	p.channel <- id

	log.Printf("[worker.Do] start process for id=%d\n", id)

	return true
}
