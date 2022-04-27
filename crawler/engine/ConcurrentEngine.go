package engine

import (
	"001_go_env/crawler/model"
	"fmt"
	"log"
)

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
		// URL dedup
		if isDuplicate(r.Url){
			log.Printf("Duplivate request: "+ "%s", r.Url)
			continue
		}
		e.Scheduler.Submit(r)
	}

	// count real
	itemCount := 0
	profileCount := 0

	// received out
	for{
		result := <- out
		for _, item := range result.Items{
			if _, ok := item.(model.Profile); ok{
				fmt.Printf("Got item #%d: %v", profileCount, item)
				profileCount++
			}

			fmt.Printf("Got item #%d: %v", itemCount, item)
			fmt.Println()
			itemCount++
		}

		// URL dedup
		for _, request := range result.Requests{
			if isDuplicate(request.Url){
				log.Printf("Duplivate request: "+ "%s", request.Url)
				continue
			}
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

var visitedUrls = make(map[string]bool)
func isDuplicate(url string) bool{
	if visitedUrls[url]{
		return true
	}
	visitedUrls[url] = true
	return false
}