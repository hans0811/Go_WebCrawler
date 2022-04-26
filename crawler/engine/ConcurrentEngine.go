package engine

import "fmt"

type ConcurrentEngine struct {
	Scheduler Scheduler
	WorkerCount int
}

type Scheduler interface {
	Submit(Request)
	ConfigureMasterWorkChan(chan Request)
	WorkerReady(chan Request)
	Run()
}


func (e *ConcurrentEngine) Run(seeds ...Request){


	// All the workers share on input
	//in := make(chan Request)
	out := make(chan ParseResult)
	//e.Scheduler.ConfigureMasterWorkChan(in)
	e.Scheduler.Run()

	for i :=0; i < e.WorkerCount; i++{
		creatWorker(out, e.Scheduler)
	}

	// Wait for creating workers
	for _, r := range seeds{
		e.Scheduler.Submit(r)
	}

	// count real
	itemCount := 0

	// received out
	for{
		result := <- out
		for _, itme := range result.Items{
			fmt.Printf("Got item #%d: %v", itemCount, itme)
			itemCount++
		}

		for _, request := range result.Requests{
			// the request need to consume
			e.Scheduler.Submit(request)
		}
	}

}

func creatWorker(out chan ParseResult, s Scheduler){
	in := make(chan Request)
	go func() {
		for{
			// tell scheduler I am ready
			s.WorkerReady(in)
			request := <- in
			result, err := worker(request)

			if err != nil{
				continue
			}
			// Work need to send a result
			out <- result
		}
	}()

}