package worker

import (
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
		<-p.done

		p.inuse--
	}
}

func (p *Pool) Register() {
	for i := 0; i < p.capacity; i++ {
		go worker{
			channel: p.channel,
			done:    p.done,
		}.work()
	}
}

func (p *Pool) Do(id int) bool {
	if p.inuse == p.capacity {
		return false
	}

	p.inuse++

	p.channel <- id

	return true
}
