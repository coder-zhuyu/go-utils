package autoinc


type AutoInc struct {
	start, step, end int
	queue chan int
	running bool
}

func New(start, step, end int) (ai *AutoInc) {
	ai = &AutoInc{
		start: start,
		step: step,
		end: end,
		queue: make(chan int, 100),
		running: true,
	}
	go ai.process()
	return
}

func (ai *AutoInc) process() {
	defer func() {recover()}()
	for i := ai.start; ai.running; i = i + ai.step {
		ai.queue <- i
		if i >= ai.end {
			i = ai.start
		}
	}
}

func (ai *AutoInc) Id() int {
	return <- ai.queue
}

func (ai *AutoInc) Close() {
	ai.running = false
	close(ai.queue)
}
