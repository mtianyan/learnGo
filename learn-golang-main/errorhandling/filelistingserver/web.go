package main

import (
	"GoDemoProj/errorhandling/filelistingserver/consts"
	"GoDemoProj/errorhandling/filelistingserver/filelisting"
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
)

type appHandler func(writer http.ResponseWriter, request *http.Request) error

func errWrapper(handler appHandler) func(http.ResponseWriter, *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		defer func() {
			if r := recover(); r != nil {
				log.Printf("Panic: %v", r)
				http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			}
		}()

		err := handler(writer, request)
		if err != nil {
			log.Printf("Error handling request: %s \n", err.Error())
			if userErr, ok := err.(userError); ok {
				http.Error(writer, userErr.Error(), http.StatusBadRequest)
				return
			}

			code := http.StatusOK
			switch {
			case os.IsNotExist(err):
				code = http.StatusNotFound
			case os.IsPermission(err):
				code = http.StatusForbidden
			default:
				code = http.StatusInternalServerError
			}
			http.Error(writer, http.StatusText(code), code)
		}
	}
}

type userError interface {
	error
	Message() string
}

func main() {
	http.HandleFunc("/", errWrapper(filelisting.HandleFileListing))

	err := http.ListenAndServe(consts.ServerPort, nil)
	if err != nil {
		panic(err)
	}
}
