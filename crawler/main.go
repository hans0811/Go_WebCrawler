package main

import (
	"001_go_env/crawler/engine"
	"001_go_env/crawler/persist"
	"001_go_env/crawler/scheduler"
	"001_go_env/crawler/zhenai/parser"
)

func main() {

	// engine.SimpleEngine
	//engine.SimpleEngine{}.Run(engine.Request{
	//	Url: "https://www.zhenai.com/zhenghun",
	//	ParserFunc: parser.ParseCityList,
	//})

	itemChan, err := persist.ItemSaver("dating_profile")
	if err != nil{
		// itemChan must be started for saving data
		panic(err)
	}
	e := engine.ConcurrentEngine{
		//QueuedScheduler
		Scheduler: &scheduler.QueuedScheduler{},
		WorkerCount: 100,
		ItemChan: itemChan,
	}

	e.Run(engine.Request{
		Url: "http://localhost:8080/mock/www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList,
	})

	//e.Run(engine.Request{
	//	Url: "http://localhost:8080/mock/www.zhenai.com/zhenghun",
	//	ParserFunc: parser.ParseCity,
	//})


	//resp, err := http.Get("https://www.zhenai.com/zhenghun")
	//
	//if err != nil {
	//	panic(err)
	//}
	//
	//defer resp.Body.Close()
	//
	//if resp.StatusCode != http.StatusOK {
	//	fmt.Println("Error: status code", resp.StatusCode)
	//	return
	//}
	//
	//all, err := ioutil.ReadAll(resp.Body)
	//if err != nil {
	//	panic(err)
	//}
	//
	////fmt.Printf("%s\n", all)
	//
	//printCityList(all)
}

//func printCityList(contents []byte) {
//
//	re := regexp.MustCompile(`<a href="(http://www.zhenai.com/zhenghun/[a-z0-9]+)"[^>]*>([^<]+)</a>`)
//	//matches := re.FindAll(contents, -1)
//	matches := re.FindAllSubmatch(contents, -1) //return [][][]
//	for _, m := range matches {
//
//		fmt.Printf("City: %s\n, URL: %s\n", m[2], m[1])
//		// show sub match array
//		//for _, subMatch := range m {
//		//	fmt.Printf("%s\n", subMatch)
//		//}
//		//fmt.Println()
//	}
//	fmt.Printf("Matches found: %d\n", len(matches))
//}
