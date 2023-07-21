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

func (p Pool) Register() {
	for i := 0; i < p.capacity; i++ {
		go worker{
			channel: p.channel,
			done:    p.done,
		}.work()
	}
}
