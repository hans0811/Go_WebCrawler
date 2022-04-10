package main

import (
	"001_go_env/retriever/mock"
	"001_go_env/retriever/real"
	"fmt"
	"time"
)

type Retriever interface {
	Get(url string) string
}

type Poster interface {
	Post(url string, form map[string]string) string
}

const url = "http://www.imooc.com"

func downloader(r Retriever) string {
	return r.Get(url)
}

func post(poster Poster) {
	poster.Post(url,
		map[string]string{
			"name":   "ccnmuse",
			"course": "golang",
		})
}

type RetrieverPoster interface {
	Retriever
	Poster
}

func session(s RetrieverPoster) string {
	s.Post(url, map[string]string{
		"contents": "another fake imooc.com",
	})
	return s.Get(url)
}

func main() {
	var r Retriever
	retriever := mock.Retriever{"This is a fake imooc.com"}
	r = &retriever
	inspect(r)

	r = &real.Retriever{
		UserAgent: "Mozilla/5.0",
		Timeout:   time.Minute,
	}
	//inspect(r)

	// type assertion
	realRetriever := r.(*real.Retriever)
	fmt.Println(realRetriever)

	//fmt.Println(downloader(r))
	var r1 Retriever
	retriever1 := mock.Retriever{
		"this is a fake imooc.com",
	}
	r1 = &retriever1
	inspect(r1)
	fmt.Println("Try a session")
	fmt.Println(session(&retriever1))

}

func inspect(r Retriever) {
	fmt.Printf("Inspecting", r)
	fmt.Printf(" > %T %v\n", r, r)
	fmt.Printf(" > Type switch: ")
	switch v := r.(type) {
	case *mock.Retriever:
		fmt.Println("Contents: ", v.Contents)
	case *real.Retriever:
		fmt.Println("UserAgent:", v.UserAgent)
	}
	fmt.Println()
}
