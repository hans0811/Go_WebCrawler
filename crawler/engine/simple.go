package engine

import (
	"001_go_env/crawler/fetcher"
	"log"
)

type SimpleEngine struct {

}

func (e SimpleEngine) Run(seeds ...Request) {
	var requests []Request
	for _, r := range seeds{
		requests = append(requests, r)
	}

	for len(requests) > 0 {
		// queue
		r := requests[0]
		requests = requests[1:]

		//log.Printf("Fetching %s", r.Url)
		//body, err := fetcher.Fetch(r.Url)
		//if err != nil{
		//	log.Printf("Fetcher: error fetching url %s: %v",
		//		r.Url, err)
		//	continue
		//}
		//// Parser part
		//ParseResult := r.ParserFunc(body)

		ParseResult, err := worker(r)
		if err != nil{
			continue
		}

		requests = append(requests,
			// ParseResult.Requests[0]....
			ParseResult.Requests...)

		// print result
		for _, item := range ParseResult.Items {
			log.Printf("Got item %v", item)
		}
	}
}

func worker(r Request) (ParseResult, error){
	log.Printf("Fetching %s", r.Url)
	body, err := fetcher.Fetch(r.Url)
	if err != nil{
		log.Printf("Fetcher: error fetching url %s: %v",
			r.Url, err)
		//Even error, still run it
		//continue
		return ParseResult{}, err
	}
	// Parser part
	return r.ParserFunc(body), nil
}
