package apiprac

import (
	"fmt"
	"goSnippets/logger"
	"strings"
)

type AddSuffixFunc func(name string) string

func MakeAddSuffix(suffix string) AddSuffixFunc {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}

func init() {
	logger.DefaultLogger.Log()
	addJpg := MakeAddSuffix(".jpg")
	fmt.Println(addJpg("Luffy"))
}
