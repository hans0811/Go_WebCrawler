package main

import (
	"fmt"
	"regexp"
)

//const text = "My email is ccmouse1@gmail.com"
//const text = "My email is ccmouse1@gmail.com@abc.com"
const text = `
My email is ccmouse1@gmail.com@abc.com
email1 is abc@deg.org
email2 is     kkk@qq.com
email3 is ddd@abc.com.cn`

func main() {

	//re,err := regexp.Compile("ccmouse@gmail.com")

	// normal way ".+@.+\\..+"
	re := regexp.MustCompile(`([a-zA-Z0-9]+)@([a-zA-Z0-9.]+)(\.[a-zA-Z0-9.]+)`)

	//re.Find()
	match := re.FindString(text)

	// if there are multiply lines
	match1 := re.FindAllString(text, -1)

	// return
	match2 := re.FindAllStringSubmatch(text, -1)

	fmt.Println(match)
	fmt.Println(match1)
	fmt.Println(match2)
	fmt.Println()
	for _, m := range match2 {
		fmt.Println(m)
	}

}
