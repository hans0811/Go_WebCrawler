package parser

import (
	"001_go_env/crawler/engine"
	"regexp"
)

const cityRe=`<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`

func ParseCity(contents []byte) engine.ParseResult{
	re := regexp.MustCompile(cityRe)
	matches := re.FindAllSubmatch(contents, -1) //return [][][]

	result := engine.ParseResult{}
	for _, m := range matches {

		// m[2] is in for block
		name := string(m[2])

		result.Items = append(result.Items, "User " + name)

		result.Requests = append(result.Requests, engine.Request{
			Url:        string(m[1]),
			// using anonymous function, because username can be pass in
			ParserFunc: func(c []byte) engine.ParseResult{
				return ParseProfile(c, name)
			},
		})

		//fmt.Printf("City: %s\n, URL: %s\n", m[2], m[1])
	}

	//fmt.Printf("Matches found: %d\n", len(matches))
	return result
}
