package main

type semaphore chan any

func main() {

}

func (s semaphore) P(n int) {
	e := new(any)
	for i := 0; i < n; i++ {
		s <- e
	}
}

func (s semaphore) V(n int) {
	for i := 0; i < n; i++ {
		<-s
	}
}

// mutexes
func (s semaphore) Lock() {
	s.P(1)
}

func (s semaphore) Unlock() {
	s.V(1)
}

// signal-wait
func (s semaphore) Wait(n int) {
	s.P(n)
}

func (s semaphore) Signal() {
	s.V(1)
}
