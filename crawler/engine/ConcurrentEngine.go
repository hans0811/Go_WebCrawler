package engine

import "fmt"

type ConcurrentEngine struct {
	Scheduler Scheduler
	WorkerCount int
}

type Scheduler interface {
	ReadyNotifier
	Submit(Request)
	WorkerChan() chan Request
	//ConfigureMasterWorkChan(chan Request)
	//WorkerReady(chan Request)
	Run()
}

type ReadyNotifier interface {
	WorkerReady(chan Request)
}


func (e *ConcurrentEngine) Run(seeds ...Request){


	// All the workers share on input
	//in := make(chan Request)
	out := make(chan ParseResult)
	//e.Scheduler.ConfigureMasterWorkChan(in)
	e.Scheduler.Run()

	for i :=0; i < e.WorkerCount; i++{
		creatWorker(e.Scheduler.WorkerChan(), out, e.Scheduler)
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
			fmt.Println()
			itemCount++
		}

		for _, request := range result.Requests{
			// the request need to consume
			e.Scheduler.Submit(request)
		}
	}

}

func creatWorker(in chan Request, out chan ParseResult, ready ReadyNotifier){
	go func() {
		for{
			// tell scheduler I am ready
			ready.WorkerReady(in)
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