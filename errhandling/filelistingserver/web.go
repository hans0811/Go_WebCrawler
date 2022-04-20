package main

import (
	"001_go_env/errhandling/filelistingserver/filelisting"
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
)

type appHandler func(writer http.ResponseWriter,
	request *http.Request) error

func errWrapper(
	handler appHandler) func(
	http.ResponseWriter, *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {

		defer func() {
			if r := recover(); r != nil {
				log.Printf("Panic: %v", r)
				http.Error(writer, http.StatusText(http.StatusInternalServerError),
					http.StatusInternalServerError)
			}

		}()

		err := handler(writer, request)

		if err != nil {
			//log.Warn("Error handling request: %s", err.Error())
			log.Printf("Error ocurred Error handling request: %s", err.Error())

			if userError, ok := err.(userError); ok {
				http.Error(
					writer,
					userError.Message(),
					http.StatusBadRequest)
				return
			}

			code := http.StatusOK
			switch {
			case os.IsNotExist(err):
				code = http.StatusNotFound
				//http.Error( writer, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			case os.IsPermission(err):
				code = http.StatusForbidden
			default:
				code = http.StatusInternalServerError
			}
			http.Error(writer,
				http.StatusText(code), code)
		}
	}
}

// let user can see
type userError interface {
	error
	Message() string
}

func main() {

	http.HandleFunc("/",
		errWrapper(filelisting.HandleFileList))
	//func(writer http.ResponseWriter, request *http.Request) {
	//	path := request.URL.Path[len("/list/"):] // /list/fib/txt
	//	file, err := os.Open(path)
	//	if err != nil {
	//		http.Error(writer,
	//			err.Error(),
	//			http.StatusInternalServerError,
	//		)
	//		return
	//	}
	//	defer file.Close()
	//
	//	all, err := ioutil.ReadAll(file)
	//	if err != nil {
	//		panic(err)
	//	}
	//
	//	writer.Write(all)

	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		panic(err)
	}
}
