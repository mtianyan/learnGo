package controller

import (
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

type userError string

func (e userError) Error() string {
	return e.Message()
}

func (e userError) Message() string {
	return string(e)
}

const prefix = "/list/"

func ListFile(writer http.ResponseWriter, request *http.Request) error {
	if !strings.HasPrefix(request.URL.Path, prefix) {
		return userError("path must start with " + prefix)
	}
	file, err := os.Open(request.URL.Path[len(prefix):])
	if err != nil {
		return err
	}
	defer file.Close()

	all, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}

	writer.Write(all)
	return nil
}
