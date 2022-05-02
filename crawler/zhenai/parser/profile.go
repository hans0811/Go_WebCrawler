package parser

import (
	"001_go_env/crawler/engine"
	"001_go_env/crawler/model"
	"regexp"
	"strconv"
)


//const ageRe = `<div data-v-8b1eac0c="" class="m-btn pink">([.])</div>`
//const ageRe = `<div data-v-499ba28c="" class="des f-cl">[.] | ([\d]+)岁 | [.] | [.] | [.] | [.]<a data-v-499ba28c="" href="//www.zhenai.com/n/login?channelId=905819&amp;fromurl=https%3A%2F%2Falbum.zhenai.com%2Fu%2F1998372165" target="_self" class="online f-fr">查看最后登录时间</a></div>`
//const marriageRe = `<div data-v-499ba28c="" class="des f-cl">[.] | ([\d]+)岁 | [.] | ([\d]) | [.] | [.]<a data-v-499ba28c="" href="//www.zhenai.com/n/login?channelId=905819&amp;fromurl=https%3A%2F%2Falbum.zhenai.com%2Fu%2F1998372165" target="_self" class="online f-fr">查看最后登录时间</a></div>`
//阿坝 | 27岁 | 大学本科 | 未婚 | 157cm | 3001-5000元
// Compile regex for saving time
//var ageRe = regexp.MustCompile(`<div data-v-499ba28c="" class="des f-cl">[.] | ([\d]+)岁 | [.] | [.] | [.] | [.]<a data-v-499ba28c="" href="//www.zhenai.com/n/login?channelId=905819&amp;fromurl=https%3A%2F%2Falbum.zhenai.com%2Fu%2F1998372165" target="_self" class="online f-fr">查看最后登录时间</a></div>`)
//var marriageRe = regexp.MustCompile(`<div data-v-499ba28c="" class="des f-cl">[.]+ | [.]+ | [.]+ | (未婚) | [.]+ | 3001-5000元[^<]+<a data-v-499ba28c="" href="//www.zhenai.com/n/login?channelId=905819&amp;fromurl=https%3A%2F%2Falbum.zhenai.com%2Fu%2F1998372165" target="_self" class="online f-fr">查看最后登录时间</a></div>`)
//var heightRe = regexp.MustCompile(`<div data-v-499ba28c="" class="des f-cl">[.] | [.] | [.] | [.] | ([\d]+)cm | [.]<a data-v-499ba28c="" href="//www.zhenai.com/n/login?channelId=905819&amp;fromurl=https%3A%2F%2Falbum.zhenai.com%2Fu%2F1998372165" target="_self" class="online f-fr">查看最后登录时间</a></div>`)
//var weightRe = regexp.MustCompile(`<div data-v-8b1eac0c="" class="m-btn purple">([\d]+)kg</div>`)
//var hokouRe = regexp.MustCompile(`<div data-v-499ba28c="" class="des f-cl">([\S]+) | [.] | [.] | [.] | [.] | [.]<a data-v-499ba28c="" href="//www.zhenai.com/n/login?channelId=905819&amp;fromurl=https%3A%2F%2Falbum.zhenai.com%2Fu%2F1998372165" target="_self" class="online f-fr">查看最后登录时间</a></div>`)
//var incomeRe = regexp.MustCompile(`<div data-v-499ba28c="" class="des f-cl">[] | [.] | [.] | [.] | ([^] | [.]+)元<a data-v-499ba28c="" href="//www.zhenai.com/n/login?channelId=905819&amp;fromurl=https%3A%2F%2Falbum.zhenai.com%2Fu%2F1998372165" target="_self" class="online f-fr">查看最后登录时间</a></div>`)
//var incomeRe = regexp.MustCompile(`<div data-v-4c07f04e="" class="des f-cl">[] | [.] | [.] | [.] | ([^] | [.]+)元`)


var ageRe = regexp.MustCompile(
	`<td><span class="label">年龄：</span>(\d+)岁</td>`)
var heightRe = regexp.MustCompile(
	`<td><span class="label">身高：</span>(\d+)CM</td>`)
var incomeRe = regexp.MustCompile(
	`<td><span class="label">月收入：</span>([^<]+)</td>`)
var weightRe = regexp.MustCompile(
	`<td><span class="label">体重：</span><span field="">(\d+)KG</span></td>`)
var genderRe = regexp.MustCompile(
	`<td><span class="label">性别：</span><span field="">([^<]+)</span></td>`)
var xinzuoRe = regexp.MustCompile(
	`<td><span class="label">星座：</span><span field="">([^<]+)</span></td>`)
var marriageRe = regexp.MustCompile(
	`<td><span class="label">婚况：</span>([^<]+)</td>`)
var educationRe = regexp.MustCompile(
	`<td><span class="label">学历：</span>([^<]+)</td>`)
var occupationRe = regexp.MustCompile(
	`<td><span class="label">职业：</span><span field="">([^<]+)</span></td>`)
var hokouRe = regexp.MustCompile(
	`<td><span class="label">籍贯：</span>([^<]+)</td>`)
var houseRe = regexp.MustCompile(
	`<td><span class="label">住房条件：</span><span field="">([^<]+)</span></td>`)
var carRe = regexp.MustCompile(
	`<td><span class="label">是否购车：</span><span field="">([^<]+)</span></td>`)
var guessRe = regexp.MustCompile(
	`<a class="exp-user-name"[^>]*href="(.*album\.zhenai\.com/u/[\d]+)">([^<]+)</a>`)
var idUrlRe = regexp.MustCompile(
	`.*album\.zhenai\.com/u/([\d]+)`)

func ParseProfile(contents []byte, url string, name string) engine.ParseResult{
	profile := model.Profile{}
	profile.Name = name
	//re := regexp.MustCompile(ageRe)
	//match := ageRe.FindSubmatch(contents)

	age, err := strconv.Atoi(extractString(contents, ageRe))

	if err == nil{
		profile.Age = age
	}

	//if match != nil{
	//	age, err := strconv.Atoi(string(match[1]))
	//	if err != nil{
	//		// user age is age
	//		profile.Age = age
	//	}
	//}
	height, err := strconv.Atoi(extractString(contents, heightRe))
	if err == nil{
		profile.Height = height
	}

	weight, err := strconv.Atoi(extractString(contents, weightRe))
	if err == nil{
		profile.Weight = weight
	}

	profile.Height = height
	profile.Marriage = extractString(contents, marriageRe)
	profile.Hokou = extractString(contents, hokouRe)
	profile.Income = extractString(contents, incomeRe)


	//re := regexp.MustCompile(marriageRe)
	//match := marriageRe.FindSubmatch(contents)
	//
	//if match != nil{
	//	profile.Marriage = string(match[1])
	//}

	result := engine.ParseResult{
		Items: []engine.Item{
			{
				Url: url,
				Type: "zhenai",
				Id: extractString([]byte(url), idUrlRe),
				Payload: profile,

			},
		},
	}

	matches := guessRe.FindAllSubmatch(contents, -1)

	for _, m := range matches{
		result.Requests = append(result.Requests,
			engine.Request{
			Url: string(m[1]),
			ParserFunc: ProfileParser(string(m[2])),
			})
	}

	return result
}

func extractString(contents []byte, re *regexp.Regexp) string{
	match := re.FindSubmatch(contents)

	if len(match) >= 2{
		return  string(match[1])
	}else{
		return ""
	}
}

func ProfileParser(name string) engine.ParserFunc{
	return func(c []byte, url string) engine.ParseResult{
		return ParseProfile(c, url, name)
	}

}