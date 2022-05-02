package engine

import (
	"001_go_env/crawler/fetcher"
	"log"
)

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
	return r.ParserFunc(body, r.Url), nil
}