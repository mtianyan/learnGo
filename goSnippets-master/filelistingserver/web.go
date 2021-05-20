package main

import (
	"goSnippets/filelistingserver/controller"
	"log"
	"net/http"
	"os"
)

type appHandler func(http.ResponseWriter, *http.Request) error

func errorWrapper(handler appHandler) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("Error occurred handling request: %s\n", err)
				http.Error(writer,
					http.StatusText(http.StatusInternalServerError),
					http.StatusInternalServerError)
			}
		}()
		err := handler(writer, request)
		if err != nil {
			switch errTyped := err.(type) {
			case userError:
				handleUserError(writer, errTyped)
			default:
				handleSystemError(writer, err)
			}
		}
	}
}

func handleSystemError(writer http.ResponseWriter, err error) {
	log.Printf("Error occurred handling request: %s\n", err)
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

func handleUserError(writer http.ResponseWriter, err userError) {
	log.Printf("Error occurred handling request: %s\n", err.Message())
	http.Error(writer,
		err.Message(),
		http.StatusBadRequest)
}

type userError interface {
	error
	Message() string
}

func main() {
	http.HandleFunc("/", errorWrapper(controller.ListFile))
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		panic(err)
	}
}
