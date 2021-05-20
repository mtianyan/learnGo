package real

import (
	"net/http"
	"net/http/httputil"
	"time"
)

type Retriever struct {
	UserAgent string
	Timeout   time.Duration
}

func (*Retriever) Get(url string) string {
	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	result, errResult := httputil.DumpResponse(response, true)
	response.Body.Close()

	if errResult != nil {
		panic(errResult)
	}

	return string(result)
}
