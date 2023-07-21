package worker

type Pool struct {
	capacity int
	inuse    int
	channel  chan int
	done     chan int
}

func New(capacity int) *Pool {
	return &Pool{
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
