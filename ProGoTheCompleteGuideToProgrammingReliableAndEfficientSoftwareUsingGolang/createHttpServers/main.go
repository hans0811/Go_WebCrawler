package main

import (
	"io"
	"net/http"
)

type StringHandler struct {
	message string
}

func (sh StringHandler) ServeHTTP(writer http.ResponseWriter,
									request *http.Request){
										Printfln("Request for %v", request.URL.Path)
										io.WriteString(writer, sh.message)
}

func main() {
	http.Handle("/message", StringHandler{ "Hello, World"})
	http.Handle("/favicon.ico", http.NotFoundHandler())
    http.Handle("/", http.RedirectHandler("/message", http.StatusTemporaryRedirect))

	fsHandler := http.FileServer(http.Dir("./static"))

	http.Handle("/files/", http.StripPrefix("/files", fsHandler))
	err := http.ListenAndServe(":5000", nil)

	if (err != nil){
		Printfln("Error: %v", err.Error())
	}

	// for _, p := range Products {
	// 	Printfln("Product: %v, Category: %v, Price: $%.2f",
	// 		p.Name, p.Category, p.Price)
	// }
}