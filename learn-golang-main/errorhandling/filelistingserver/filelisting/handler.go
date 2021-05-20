package filelisting

import (
	"GoDemoProj/errorhandling/filelistingserver/consts"
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

func HandleFileListing(writer http.ResponseWriter, request *http.Request) error {
	if strings.Index(request.URL.Path, consts.ListPath) != 0 {
		return userError("path must start with: " + consts.ListPath)
	}
	path := request.URL.Path[len(consts.ListPath):]
	file, err := os.Open(path)
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
