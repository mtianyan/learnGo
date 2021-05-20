package main

import (
	"GoDemoProj/interfaces/retriever/mock"
	"GoDemoProj/interfaces/retriever/real"
	"fmt"
)

func main() {

	//var r Retriever = mock.Retriever{Contents: "Mock retriever"}
	var r Retriever = &real.Retriever{}
	println(download(r))

	// type switch
	switch v := r.(type) {
	case mock.Retriever:
		fmt.Println("mock retriever", v)
	}

	// Type assertion
	if realRetriever, ok := r.(*real.Retriever); ok {
		download(realRetriever)
	} else {
		panic(ok)
	}
	fmt.Printf("%T, %v", r, r)

}

type Retriever interface {
	Get(url string) string
}

type Poster interface {
	Post(url string, from map[string]string) string
}

type RetrieverPoster interface {
	Retriever
	Poster

	Connect(host string)
}

func download(r Retriever) string {
	return r.Get("https://www.google.com")
}
