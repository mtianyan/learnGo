package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
)

func main() {
	//httpGet()
	httpGetMobile()
}

func httpGetMobile() {
	// manually build request and set header properly
	const googlecom = "https://www.google.com"
	const amazoncom = "https://www.amazon.com"
	const useragent = "Mozilla/5.0 (iPhone; CPU iPhone OS 13_2_3 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/13.0.3 Mobile/15E148 Safari/604.1"

	request, err := http.NewRequest(http.MethodGet, googlecom, nil)
	request.Header.Set("User-Agent", useragent)

	client := http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			fmt.Println("Redirect:", req)
			return nil
		},
	}
	//resp, err := http.DefaultClient.Do(request)
	resp, err := client.Do(request)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	content, err := httputil.DumpResponse(resp, true)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", content)
}

func httpGet() {
	resp, err := http.Get("https://www.google.com")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	content, err := httputil.DumpResponse(resp, true)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", content)
}
