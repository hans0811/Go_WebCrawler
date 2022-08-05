package main

import (
	"encoding/json"
	"io"
	"os"
	"strings"

	"net/http"
	"time"
)

func main() {
	Printfln("Start Http Server")
	go http.ListenAndServe(":5000", nil)
	time.Sleep(time.Second)
	var builder strings.Builder

	err := json.NewEncoder(&builder).Encode(Products[0])

	if(err != nil) {
		Printfln("Error: %v", err.Error)
	}

	req, err := http.NewRequest(http.MethodPost, 
								"http://localhost:5000/echo",
								io.NopCloser(strings.NewReader(builder.String())))
	if(err != nil){
		Printfln("Error: %v", err.Error)
	}

	req.Header["Content-Type"] = []string{ "application/josn" }
	response, err := http.DefaultClient.Do(req)

	if(err != nil){ Printfln("Error: %v", err.Error) }
	io.Copy(os.Stdout, response.Body)
	defer response.Body.Close()

	// formData := map[string][]string{
	// 	"name": { "Kayak "},
	// 	"category": { "Watersports"},
	// 	"price": { "279"},
		
	// }
	// response, err := http.PostForm("http://localhost:5000/echo", formData)

	// response, err := http.Post("http://localhost:5000/echo",
	// 							"application/json",
	// 						strings.NewReader(builder.String()))

	// if(err == nil && response.StatusCode == http.StatusOK){
	// 	io.Copy(os.Stdout, response.Body)
	// 	defer response.Body.Close()
	// 	//response.Write(os.Stdout)
	// }else{
	// 	Printfln("Error: %v", err.Error())
	// }
}