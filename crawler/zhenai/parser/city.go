package parser

import (
	"001_go_env/crawler/engine"
	"regexp"
)

var (profileRe = regexp.MustCompile(
	`<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`)
cityUrlRe = regexp.MustCompile(`href="(http://www.zhenai.com/zhenghun/[^"]+)"`)

)

func ParseCity(contents []byte) engine.ParseResult{
	matches := profileRe.FindAllSubmatch(contents, -1) //return [][][]

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

	matches = cityUrlRe.FindAllSubmatch(contents, -1)

	for _, m := range matches{
		result.Requests = append(result.Requests,
			engine.Request{
			Url: string(m[1]),
			ParserFunc: ParseCityList,
			})
	}

	//fmt.Printf("Matches found: %d\n", len(matches))
	return result
}
