package main

import (
	"GoDemoProj/interfaces/infra"
	"fmt"
)

var googleUrl = "https://www.google.com"

type retriever interface {
	Get(string) string
}

func main() {
	var retriever retriever = infra.Retriever{}
	fmt.Println(retriever.Get(googleUrl))
}
