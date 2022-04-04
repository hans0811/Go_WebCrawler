package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func converToBin(n int) interface{} {

	result := ""
	for ; n > 0; n /= 2 {
		lsb := n % 2
		result = strconv.Itoa(lsb) + result
	}

	return result

}

func printFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	printFileContent(file)
}

func printFileContent(reader io.Reader) {
	scanner := bufio.NewScanner(reader)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func forever() {
	for {
		fmt.Println("/Users/hans/Desktop/GO/GO_tutorial/001_go_env/basic/abc.txt")
	}
}

func main() {

	fmt.Println(
		converToBin(5), // 101
		converToBin(13),
	)

	s := `abd"d"
			kkkk
			123

			p`
	printFileContent(strings.NewReader(s))
	printFile("/Users/hans/Desktop/GO/GO_tutorial/001_go_env/basic/abc.txt")
}
