package scheduler

import "001_go_env/crawler/engine"

type QueuedScheduler struct{
	requestChan chan engine.Request
	// channels in the channel
	workerChan chan chan engine.Request
}

func (s *QueuedScheduler) Submit(r engine.Request) {
	s.requestChan <- r
}

func (s *QueuedScheduler) WorkerReady(w chan engine.Request) {
	s.workerChan <- w
}

func (s *QueuedScheduler) ConfigureMasterWorkChan(r chan engine.Request) {
	panic("implement me")
}

func (s *QueuedScheduler) Run(){
	s.workerChan = make(chan chan engine.Request)
	s.requestChan = make(chan engine.Request)
	go func(){
		// create queue
		var requestQ []engine.Request
		var workerQ []chan engine.Request

		for{
			var activeRequest engine.Request
			var activeWorker chan engine.Request
			// When both queue have data
			if len(requestQ) > 0 && len(workerQ) > 0{
				activeWorker = workerQ[0]
				activeRequest = requestQ[0]
			}

			// use select
			select {
				case r:= <-s.requestChan:
					requestQ = append(requestQ, r)

				//send r to a >worker
				case w:= <-s.workerChan:
					//send ?next_request to w
					workerQ = append(workerQ, w)

				case activeWorker <- activeRequest:
					workerQ = workerQ[1:]
					requestQ = requestQ[1:]
			}

		}
	}()
}


