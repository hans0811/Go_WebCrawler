package scheduler

import "001_go_env/crawler/engine"

type SimpleScheduler struct {
	workerChan chan engine.Request
}

// it will change point
func (s *SimpleScheduler) ConfigureMasterWorkChan(c chan engine.Request) {
	s.workerChan = c
}

func (s *SimpleScheduler) Submit(r engine.Request) {
	// send request down to worker chan
	//s.workerChan <- r

	// Use go routine to submit
	go func() {
		s.workerChan <- r
	}()
}

