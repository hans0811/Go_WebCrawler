package main

import (
	"001_go_env/crawler/engine"
	"001_go_env/crawler/zhenai/parser"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

func main() {

	engine.Run(engine.Request{
		Url: "https://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList,
	})

	resp, err := http.Get("https://www.zhenai.com/zhenghun")

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("Error: status code", resp.StatusCode)
		return
	}

	all, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	//fmt.Printf("%s\n", all)

	printCityList(all)
}

func printCityList(contents []byte) {

	re := regexp.MustCompile(`<a href="(http://www.zhenai.com/zhenghun/[a-z0-9]+)"[^>]*>([^<]+)</a>`)
	//matches := re.FindAll(contents, -1)
	matches := re.FindAllSubmatch(contents, -1) //return [][][]
	for _, m := range matches {

		fmt.Printf("City: %s\n, URL: %s\n", m[2], m[1])
		// show sub match array
		//for _, subMatch := range m {
		//	fmt.Printf("%s\n", subMatch)
		//}
		//fmt.Println()
	}
	fmt.Printf("Matches found: %d\n", len(matches))
}
